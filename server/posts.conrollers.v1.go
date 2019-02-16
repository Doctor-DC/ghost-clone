package server

import (
	"encoding/json"
	"github.com/cyantarek/ghost-clone-project/helpers"
	"github.com/cyantarek/ghost-clone-project/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gopkg.in/macaron.v1"
	"gopkg.in/mgo.v2/bson"
	"math"
	"strconv"
	"strings"
)

func (s *Server) checkCache(c *macaron.Context) bool {
	dataP, err := s.Cache.Get(c.Req.Request.RequestURI)
	if err == nil {
		c.Resp.Header().Set("Content-Type", "application/json")
		c.Resp.WriteHeader(200)
		_, _ = c.Resp.Write(dataP)
		return true
	}
	return false
}

func (s *Server) setCache(data interface{}, c *macaron.Context) {
	dataS, _ := json.Marshal(data)

	s.Cache.Set(c.Req.Request.RequestURI, dataS, "300s")
}

func (s *Server) GetAllPosts(c *macaron.Context) {
	if viper.GetBool("cache.enabled") {
		if s.checkCache(c) {
			return
		}
	}
	var err error
	var posts []models.Post
	var count int

	limit := c.Query("limit")
	if limit == "" {
		limit = "15"
	}
	limitInt, _ := strconv.Atoi(limit)

	order := c.Query("order")
	if order == "" {
		order = "-created_at"
	}

	includes := c.Query("includes")
	var authors, tags bool

	if includes != "" {
		includesD := strings.Split(includes, ",")
		for _, v := range includesD {
			if v == "authors" {
				authors = true
			}
			if v == "tags" {
				tags = true
			}
		}
	}

	pageInt, _ := strconv.Atoi(c.Query("page"))
	if pageInt == 0 {
		pageInt = 1
	}

	// fields
	fields := c.Query("fields")

	// filters
	filters := c.Query("filter")

	posts, count, err = helpers.GetPosts(s.DB, limitInt, order, pageInt, authors, tags, filters, fields)
	if err != nil {
		c.JSON(404, APIErrors{"ErrDataNotFound", ErrDataNotFound})
		return
	}

	// pagination

	//formats := c.Param("formats")
	var pages = float64(count) / float64(limitInt) + 0.4
	var next int
	var pagesInt = int(math.RoundToEven(pages))
	if pagesInt == 0 {
		pagesInt = 1
	}
	if pageInt+1 > pagesInt {
		next = -1
	} else {
		next = pageInt + 1
	}

	meta := models.PostMeta{
		Page:  pageInt,
		Limit: limitInt,
		Pages: pagesInt,
		Total: count,
		Next:  next,
		Prev:  pageInt - 1,
	}
	jsonData := gin.H{
		"posts": posts,
		"meta":  meta,
	}
	if viper.GetBool("cache.enabled") {
		s.setCache(jsonData, c)
	}
	c.JSON(200, jsonData)
}

func (s *Server) GetPostById(c *macaron.Context) {
	if s.checkCache(c) {
		return
	}
	id := c.Params("postId")
	if ok := bson.IsObjectIdHex(id); !ok {
		c.JSON(404, APIErrors{"ErrDataNotFound", ErrDataNotFound})
		return
	}
	post := models.Post{Id: bson.ObjectIdHex(id)}
	err := post.GetPost(s.DB, "_id")
	if err != nil {
		c.JSON(404, APIErrors{"ErrDataNotFound", ErrDataNotFound})
		return
	}
	jsonData :=gin.H{"posts": []models.Post{post}}
	s.setCache(jsonData, c)
	c.JSON(200, jsonData)
}

func (s *Server) GetPostBySlug(c *macaron.Context) {
	if s.checkCache(c) {
		return
	}
	slug := c.Params("postSlug")
	post := models.Post{Slug: slug}
	err := post.GetPost(s.DB, "slug")
	if err != nil {
		c.JSON(404, APIErrors{"ErrDataNotFound", ErrDataNotFound})
		return
	}
	jsonData := gin.H{"posts": []models.Post{post}}
	s.setCache(jsonData, c)
	c.JSON(200, jsonData)
}
