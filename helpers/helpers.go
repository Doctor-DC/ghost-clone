package helpers

import (
	"fmt"
	"github.com/cyantarek/ghost-clone-project/config"
	"github.com/cyantarek/ghost-clone-project/db"
	"github.com/cyantarek/ghost-clone-project/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

// prepareFilters prepare filters for mongo db to query based on filter data from URL query
func prepareFilters(filters string) map[string]interface{} {
	var filterTypes []string
	var filtersListOfEachType = make(map[string]interface{})

	// splits based on "," or ", " as per user inputs
	filterTypes = strings.Split(filters, ", ")
	if len(filterTypes) == 1 {
		filterTypes = strings.Split(filters, ",")
	}


	for _, v := range filterTypes {
		// separate actions based on "command:action"
		d := strings.Split(v, ":")

		// check for inverse value and make appropriate query for it
		if string(d[1][0]) == "-" {
			filtersListOfEachType[d[0]] = map[string]string{"$ne": d[1][1:]}
		} else {
			filtersListOfEachType[d[0]] = d[1]
		}
	}
	return filtersListOfEachType
}

// GetPosts is a helper function that returns all the posts based on certain logic
func GetPosts(db db.Db, limit int, order string, page int, authors, tags bool, filters string, fields string) ([]models.Post, int, error) {
	var posts []models.Post
	var query *mgo.Query

	// get cloned mongodb session
	sess := db.GetSession()

	// default query initialize with filters available or not
	if filters != "" {
		query = sess.DB(config.DB_NAME).C(config.POST_COLLECTION_NAME).Find(prepareFilters(filters))
	} else {
		query = sess.DB(config.DB_NAME).C(config.POST_COLLECTION_NAME).Find(nil)
	}

	// execution 1: get count of posts
	count, _ := query.Count()

	// check fields and split them apart
	var fieldsP []string
	if fields != "" {
		fieldsP = strings.Split(fields, ",")
	}
	if len(fieldsP) > 0 {
		// make bson query
		var selectors = make(bson.M)
		for _, v := range fieldsP {
			selectors[v] = 1
		}
		// construct Select query from selectors
		query = query.Select(selectors)
	}

	// pagination
	if page > 1 {
		query = query.Skip((page - 1) * limit)
	}

	// sorting and limiting
	query = query.Sort(order).Limit(limit)

	// execution 2: execute the main query
	err := query.All(&posts)
	if err != nil {
		fmt.Println(err.Error())
		return nil, 0, err
	}
	defer sess.Close()

	// attach authors field if needed
	if !authors {
		for i := range posts {
			posts[i].Authors = nil
			posts[i].PrimaryAuthor = nil
		}
	}

	// attach tags field if needed
	if !tags {
		for i := range posts {
			posts[i].Tags = nil
			posts[i].PrimaryTag = nil
		}
	}
	return posts, count, nil
}

// Same as above except for Tags
func GetTags(db db.Db, limit int, order string, page int, filters string, fields string) ([]models.Tag, int, error) {
	var tags []models.Tag
	var query *mgo.Query

	// get cloned mongodb session
	sess := db.GetSession()

	// default query initialize with filters available or not
	if filters != "" {
		query = sess.DB(config.DB_NAME).C(config.TAG_COLLECTION_NAME).Find(prepareFilters(filters))
	} else {
		query = sess.DB(config.DB_NAME).C(config.TAG_COLLECTION_NAME).Find(nil)
	}

	// execution 1: get count of posts
	count, _ := query.Count()

	// check fields and split them apart
	var fieldsP []string
	if fields != "" {
		fieldsP = strings.Split(fields, ",")
	}
	if len(fieldsP) > 0 {
		// make bson query
		var selectors = make(bson.M)
		for _, v := range fieldsP {
			selectors[v] = 1
		}

		// construct Select query from selectors
		query = query.Select(selectors)
	}

	// pagination
	if page > 1 {
		query = query.Skip((page - 1) * limit)
	}

	// sorting and limiting
	query = query.Sort(order).Limit(limit)

	// execution 2: execute the main query
	err := query.All(&tags)
	if err != nil {
		fmt.Println(err.Error())
		return nil, 0, err
	}
	defer sess.Close()

	return tags, count, nil
}

// Same as above except for Authors
func GetAuthors(db db.Db, limit int, order string, page int, filters string, fields string) ([]models.Author, int, error) {
	var authors []models.Author
	var query *mgo.Query

	sess := db.GetSession()

	// default query initialize
	if filters != "" {
		query = sess.DB(config.DB_NAME).C(config.AUTHOR_COLLECTION_NAME).Find(prepareFilters(filters))
	} else {
		query = sess.DB(config.DB_NAME).C(config.AUTHOR_COLLECTION_NAME).Find(nil)
	}
	count, _ := query.Count()

	// fields
	var fieldsP []string
	if fields != "" {
		fieldsP = strings.Split(fields, ",")
	}
	if len(fieldsP) > 0 {
		var selectors = make(bson.M)
		for _, v := range fieldsP {
			selectors[v] = 1
		}
		query = query.Select(selectors)
	}

	// pagination
	if page > 1 {
		query = query.Skip((page - 1) * limit)
	}

	// sorting and limiting
	query = query.Sort(order).Limit(limit)

	// execute query
	err := query.All(&authors)
	if err != nil {
		fmt.Println(err.Error())
		return nil, 0, err
	}
	defer sess.Close()
	return authors, count, nil
}
