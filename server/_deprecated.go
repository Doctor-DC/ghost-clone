package server


func (s *Server) CreatePost(c *gin.Context) {
	// Get and Parse Data
	var postPayload structs.JsonPostIn
	err := c.BindJSON(&postPayload)
	if err != nil {
		log.Println(err)
	}

	// Core Business Logic
	if postPayload.Slug == "" {
		postPayload.Slug = generateSlug(postPayload.Title)
	}

	//Validate data
	err = s.validatePostPayload(postPayload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create new model with payload data
	tags := generateTagsFromString(postPayload.Tags)
	post := models.Post{
		Id:              bson.NewObjectId(),
		Title:           postPayload.Title,
		Slug:            postPayload.Slug,
		Markdown:        postPayload.Markdown,
		Html:            postPayload.Html,
		Featured:        postPayload.Featured,
		Page:            postPayload.Page,
		Published:       postPayload.Published,
		MetaDescription: postPayload.MetaDescription,
		FeatureImage:    postPayload.Image,
		AuthorID:        "",
		CreatedAt:       time.Now(),
		CreatedBy:       "",
		UpdatedAt:       time.Now(),
		UpdatedBy:       "",
		PublishedAt:     time.Now(),
		PublishedBy:     "",
		Status:          "draft",
		Tags:            tags,
	}
	// Insert data
	err = post.CreatePost(s.DB)
	if err != nil {
		c.JSON(500, gin.H{"status": false, "msg": err.Error()})
		return
	}
	c.JSON(200, post)
}

func (s *Server) UpdatePost(c *gin.Context) {
	slug := c.Param("postSlug")
	var payload structs.JsonPostIn
	_ = c.BindJSON(&payload)

	post, err := s.getPostBySlug(slug)
	if err != nil {
		c.Status(404)
		return
	}

	post.Title = payload.Title
	post.Slug = payload.Slug
	post.Markdown = payload.Markdown
	post.Html = payload.Html
	post.Featured = payload.Featured
	post.Page = payload.Page
	post.Published = payload.Published
	post.FeatureImage = payload.Image
	post.MetaDescription = payload.MetaDescription
	post.AuthorID = payload.AuthorID
	post.UpdatedBy = ""
	post.UpdatedAt = time.Now()

	err = post.UpdatePost(s.DB)
	if err != nil {
		c.JSON(404, gin.H{"status": false, "msg": "not found"})
		return
	}
	c.JSON(200, post)
}

func (s *Server) DeletePost(c *gin.Context) {
	slug := c.Param("postSlug")

	post := models.Post{Slug: slug}
	err := post.DeletePost(s.DB)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (s *Server) getPostBySlug(slug string) (models.Post, error) {
	post := models.Post{Slug: slug}
	err := post.GetPost(s.DB, "slug")
	if err != nil {
		return models.Post{}, nil
	} else {
		return post, nil
	}
}

func (s *Server) insertPost(post *models.Post) {
	if exists := s.DB.CheckExists("title", post.Title, "posts"); exists {
		// Already exists
		fmt.Println("Exists")
	} else {
		s.DB.InsertData(&post)
	}
}

func (s *Server) validatePostPayload(post structs.JsonPostIn) error {
	if post.Title == "" || post.Html == "" || post.Image == "" || post.Markdown == "" || post.MetaDescription == "" || len(post.Tags) == 0 || post.AuthorID == "" {
		return errors.New("empty fields")
	} else {
		sess := s.DB.GetSession()
		defer sess.Close()

		count, _ := sess.DB(config.DB_NAME).C(config.POST_COLLECTION_NAME).Find(bson.M{"title": bson.M{"$eq": post.Title}}).Count()
		if count != 0 {
			return errors.New("data already exists")
		} else {
			return nil
		}
	}
}

func generateSlug(str string) string {
	return strings.ToLower(strings.Replace(str, " ", "-", -1))
}

func generateTagsFromString(data []string) []models.Tag {
	var tags []models.Tag
	for _, v := range data {
		tag := models.Tag{Id: bson.NewObjectId(), Name: v, Slug: generateSlug(v)}
		tags = append(tags, tag)
	}
	return tags
}

