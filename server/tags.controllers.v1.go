package server

import (
	"github.com/cyantarek/ghost-clone-project/helpers"
	"github.com/cyantarek/ghost-clone-project/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/macaron.v1"
	"gopkg.in/mgo.v2/bson"
	"math"
	"strconv"
)

func (s *Server) GetAllTags(c *macaron.Context) {
	if s.checkCache(c) {
		return
	}
	var tags []models.Tag
	var err error
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

	filters := c.Query("filter")
	pageInt, _ := strconv.Atoi(c.Query("page"))

	if pageInt == 0 {
		pageInt = 1
	}

	fields := c.Query("fields")

	tags, count, err = helpers.GetTags(s.DB, limitInt, order, pageInt, filters, fields)
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
		"tags": tags,
		"meta":  meta,
	}
	s.setCache(jsonData, c)
	c.JSON(200, jsonData)
}

func (s *Server) GetTagById(c *macaron.Context) {
	if s.checkCache(c) {
		return
	}
	id := c.Params("tagId")
	if ok := bson.IsObjectIdHex(id); !ok {
		c.JSON(404, APIErrors{"ErrDataNotFound", ErrDataNotFound})
		return
	}
	tag := models.Tag{Id: bson.ObjectIdHex(id)}
	err := tag.GetTag(s.DB, "_id")
	if err != nil {
		c.JSON(404, APIErrors{"ErrDataNotFound", ErrDataNotFound})
		return
	}
	jsonData := gin.H{"tags": []models.Tag{tag}}
	s.setCache(jsonData, c)
	c.JSON(200, jsonData)
}

func (s *Server) GetTagBySlug(c *macaron.Context) {
	if s.checkCache(c) {
		return
	}
	slug := c.Params("tagSlug")
	tag := models.Tag{Slug: slug}
	err := tag.GetTag(s.DB, "slug")
	if err != nil {
		c.JSON(404, APIErrors{"ErrDataNotFound", ErrDataNotFound})
		return
	}
	jsonData := gin.H{"tags": []models.Tag{tag}}
	s.setCache(jsonData, c)
	c.JSON(200, jsonData)
}
