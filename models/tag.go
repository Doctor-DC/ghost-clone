package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Tag struct {
	Id              bson.ObjectId `json:"id,omitempty" bson:"_id"`
	Name            string        `json:"name,omitempty"`
	URL             string        `json:"url,omitempty"`
	Description     string        `json:"description,omitempty"`
	Slug            string        `json:"slug,omitempty"`
	Parent          string        `json:"parent,omitempty"`
	PostCount       int           `json:"post_count,omitempty"`
	FeatureImage    string        `json:"feature_image,omitempty" bson:"feature_image"`
	Visibility      string        `json:"visibility,omitempty"`
	MetaTitle       string        `json:"meta_title,omitempty" bson:"meta_title"`
	MetaDescription string        `json:"meta_description,omitempty" bson:"meta_description"`
	CreatedAt       string        `json:"created_at,omitempty" bson:"created_at"`
	CreatedBy       string        `json:"created_by,omitempty" bson:"created_by"`
	UpdatedAt       string        `json:"updated_at,omitempty" bson:"updated_at"`
	UpdatedBy       string        `json:"updated_by,omitempty" bson:"updated_by"`
}
