package models

import (
	"github.com/cyantarek/ghost-clone-project/config"
	"github.com/cyantarek/ghost-clone-project/db"
	"gopkg.in/mgo.v2/bson"
)

type Post struct {
	Id                 bson.ObjectId `json:"id,omitempty" bson:"_id"`
	UUID               string        `json:"uuid,omitempty"`
	Title              string        `json:"title,omitempty"`
	Slug               string        `json:"slug,omitempty"`
	MobileDoc          string        `json:"mobiledoc,omitempty" bson:"mobiledoc"`
	Markdown           string        `json:"markdown,omitempty"`
	Html               string        `json:"html,omitempty"`
	PlainText          string        `json:"plain_text,omitempty" bson:"plain_text"`
	FeatureImage       string        `json:"feature_image,omitempty" bson:"feature_image"`
	Status             string        `json:"status,omitempty"`
	MetaTitle          string        `json:"meta_title,omitempty" bson:"meta_title"`
	MetaDescription    string        `json:"meta_description,omitempty" bson:"meta_description"`
	Visibility         string        `json:"visibility,omitempty"`
	Page               bool          `json:"page,omitempty"`
	Featured           bool          `json:"featured,omitempty"`
	Published          bool          `json:"published,omitempty"`
	AuthorID           string        `json:"author_id,omitempty" bson:"author_id"`
	PrimaryTag         interface{}   `json:"primary_tag,omitempty" bson:"primary_tag"`
	PrimaryAuthor      interface{}   `json:"primary_author,omitempty" bson:"primary_author"`
	Comments           []Comment     `json:"comments,omitempty"`
	CreatedAt          string        `json:"created_at,omitempty" bson:"created_at"`
	CreatedBy          string        `json:"created_by,omitempty" bson:"created_by"`
	UpdatedAt          string        `json:"updated_at,omitempty" bson:"updated_at"`
	UpdatedBy          string        `json:"updated_by,omitempty" bson:"updated_by"`
	PublishedAt        string        `json:"published_at,omitempty" bson:"published_at"`
	PublishedBy        string        `json:"published_by,omitempty" bson:"published_by"`
	CustomExcerpt      string        `json:"custom_excerpt,omitempty" bson:"custom_excerpt"`
	CodeinjectionHead  string        `json:"codeinjection_head,omitempty" bson:"codeinjection_head"`
	CodeinjectionFoot  string        `json:"codeinjection_foot,omitempty" bson:"codeinjection_foot"`
	CustomTemplate     string        `json:"custom_template,omitempty" bson:"custom_template"`
	URL                string        `json:"url,omitempty" bson:"url"`
	Excerpt            string        `json:"excerpt,omitempty" bson:"excerpt"`
	OgTitle            string        `json:"og_title,omitempty" bson:"og_title"`
	OgDescription      string        `json:"og_description,omitempty" bson:"og_description"`
	OgImage            string        `json:"og_image,omitempty" bson:"og_image"`
	TwitterTitle       string        `json:"twitter_title,omitempty" bson:"twitter_title"`
	TwitterDescription string        `json:"twitter_description,omitempty" bson:"twitter_description"`
	TwitterImage       string        `json:"twitter_image,omitempty" bson:"twitter_image"`
	Authors            []Author      `json:"authors,omitempty" bson:"authors"`
	Tags               []Tag         `json:"tags,omitempty" bson:"tags"`
}

type PostMeta struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Pages int `json:"pages"`
	Total int `json:"total"`
	Next  int `json:"next"`
	Prev  int `json:"prev"`
}

// GetPost returns single post data
func (p *Post) GetPost(db db.Db, field string) error {
	sess := db.GetSession()
	defer sess.Close()

	if field == "_id" {
		err := sess.DB(config.DB_NAME).C(config.POST_COLLECTION_NAME).Find(bson.M{"_id": bson.M{"$eq": p.Id}}).One(&p)
		if err != nil {
			return err
		}
	} else if field == "slug" {
		err := sess.DB(config.DB_NAME).C(config.POST_COLLECTION_NAME).Find(bson.M{"slug": bson.M{"$eq": p.Slug}}).One(&p)
		if err != nil {
			return err
		}
	}
	return nil
}

// UpdatePost updates data
func (p *Post) UpdatePost(db db.Db) error {
	sess := db.GetSession()
	defer sess.Close()

	err := sess.DB(config.DB_NAME).C(config.POST_COLLECTION_NAME).UpdateId(p.Id, &p)

	if err != nil {
		return err
	}
	return nil
}

func (p *Post) DeletePost(db db.Db) error {
	sess := db.GetSession()
	defer sess.Close()

	err := sess.DB(config.DB_NAME).C(config.POST_COLLECTION_NAME).Remove(bson.M{"slug": p.Slug})
	if err != nil {
		return err
	}
	return nil
}

func (p *Post) CreatePost(db db.Db) error {
	sess := db.GetSession()
	defer sess.Close()

	err := sess.DB(config.DB_NAME).C(config.POST_COLLECTION_NAME).Insert(&p)
	if err != nil {
		return err
	}
	return nil
}

func (t *Tag) GetTag(db db.Db, field string) error {
	sess := db.GetSession()
	defer sess.Close()

	if field == "_id" {
		err := sess.DB(config.DB_NAME).C(config.TAG_COLLECTION_NAME).Find(bson.M{"_id": bson.M{"$eq": t.Id}}).One(&t)
		if err != nil {
			return err
		}
	} else if field == "slug" {
		err := sess.DB(config.DB_NAME).C(config.TAG_COLLECTION_NAME).Find(bson.M{"slug": bson.M{"$eq": t.Slug}}).One(&t)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *Tag) UpdateTag(db *db.Db) error {
	return nil
}

func (t *Tag) DeleteTag(db *db.Db) error {
	return nil
}

func (t *Tag) CreateTag(db *db.Db) error {
	return nil
}

func GetTags(db *db.Db) ([]Tag, error) {
	return nil, nil
}
