package models

import (
	"github.com/cyantarek/ghost-clone-project/config"
	"github.com/cyantarek/ghost-clone-project/db"
	"gopkg.in/mgo.v2/bson"
)

type Author struct {
	Id              bson.ObjectId `json:"id,omitempty" bson:"_id"`                      // 1
	Name            string        `json:"name,omitempty"`                               // 1
	Slug            string        `json:"slug,omitempty"`                               // 1
	Email           string        `json:"email,omitempty"`                              // 1
	ProfileImage    string        `json:"profile_image,omitempty" bson:"profile_image"` // by default placeholder
	CoverImage      string        `json:"cover_image,omitempty" bson:"cover_image"`     // by default, nothing
	Bio             string        `json:"bio,omitempty"`
	Website         string        `json:"website,omitempty"`
	Facebook        string        `json:"facebook,omitempty"`
	Twitter         string        `json:"twitter,omitempty"`
	URL             string        `json:"url,omitempty"`
	Location        string        `json:"location,omitempty"` // author location
	Password        string        `json:"-"`
	Accessibility   string        `json:"accessibility,omitempty"` // public, private. by default public
	Visibility      string        `json:"visibility,omitempty"`
	Tour            string        `json:"tour,omitempty"`
	Status          string        `json:"status,omitempty"`
	Language        string        `json:"language,omitempty"`
	MetaTitle       string        `json:"meta_title,omitempty" bson:"meta_title"`             // null
	MetaDescription string        `json:"meta_description,omitempty" bson:"meta_description"` // null
	LastLogin       string        `json:"last_login,omitempty" bson:"last_login"`
	Role            int           `json:"role,omitempty"` //1 = Subscriber, 2 = Author
}

// GetAuthor returns single author based on slug or _id
func (u *Author) GetAuthor(db db.Db, field string) error {
	sess := db.GetSession()
	defer sess.Close()

	// checks if slug or _id is provided
	if field == "_id" {
		err := sess.DB(config.DB_NAME).C(config.AUTHOR_COLLECTION_NAME).Find(bson.M{"_id": bson.M{"$eq": u.Id}}).One(&u)
		if err != nil {
			return err
		}
	} else if field == "slug" {
		err := sess.DB(config.DB_NAME).C(config.AUTHOR_COLLECTION_NAME).Find(bson.M{"slug": bson.M{"$eq": u.Slug}}).One(&u)
		if err != nil {
			return err
		}
	}
	return nil
}
