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

func (s *Server) GetAllAuthors(c *macaron.Context) {
	if s.checkCache(c) {
		return
	}
	var authors []models.Author
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

	authors, count, err = helpers.GetAuthors(s.DB, limitInt, order, pageInt, filters, fields)
	if err != nil {
		c.JSON(404, APIErrors{"ErrDataNotFound", ErrDataNotFound})
		return
	}

	// pagination
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
		"authors": authors,
		"meta":  meta,
	}
	s.setCache(jsonData, c)
	c.JSON(200, jsonData)
}

func (s *Server) GetAuthorById(c *macaron.Context) {
	if s.checkCache(c) {
		return
	}
	id := c.Params("authorId")
	if ok := bson.IsObjectIdHex(id); !ok {
		c.JSON(404, APIErrors{"ErrDataNotFound", ErrDataNotFound})
		return
	}
	author := models.Author{Id: bson.ObjectIdHex(id)}
	err := author.GetAuthor(s.DB, "_id")
	if err != nil {
		c.JSON(404, APIErrors{"ErrDataNotFound", ErrDataNotFound})
		return
	}
	jsonData := gin.H{"users": []models.Author{author}}
	s.setCache(jsonData, c)
	c.JSON(200, jsonData)
}

func (s *Server) GetAuthorBySlug(c *macaron.Context) {
	if s.checkCache(c) {
		return
	}
	slug := c.Params("authorSlug")
	author := models.Author{Slug: slug}
	err := author.GetAuthor(s.DB, "slug")
	if err != nil {
		c.JSON(404, APIErrors{"ErrDataNotFound", ErrDataNotFound})
		return
	}
	jsonData := gin.H{"tags": []models.Author{author}}
	s.setCache(jsonData, c)
	c.JSON(200, jsonData)
}
