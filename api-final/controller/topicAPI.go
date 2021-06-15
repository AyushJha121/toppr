package controller

import (
	"api-final/cache"
	"api-final/config"
	"api-final/entity"
	"bytes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Handler to GET all topics in database
func AllTopic(c *gin.Context) {
	var (
		topic  entity.Topic
		topics []entity.Topic
	)
	rows, err := config.Db.Query("select * from topics;")
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&topic.TopicID, &topic.TopicName, &topic.SubjectID)
		topics = append(topics, topic)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{
		"result": topics,
		"count":  len(topics),
	})
}

//Handler to GET topic by ID
func Topic_by_ID(c *gin.Context) {
	var (
		topic      entity.Topic
		result     gin.H
		topicCache cache.TopicCache = cache.NewRedisCache("localhost:6379", 0, 1000)
	)
	id := c.Param("id")
	topic = topicCache.Get(string(id))
	if (topic == entity.Topic{}) {
		row := config.Db.QueryRow("select TopicID, TopicName,SubjectID from topics where TopicID = ?", id)
		err := row.Scan(&topic.TopicID, &topic.TopicName, &topic.SubjectID)
		if err != nil {
			// If no results send null
			result = gin.H{
				"result": nil,
				"count":  0,
			}
		} else {
			result = gin.H{
				"result": topic,
				"count":  1,
			}
		}
		topicCache.Set(id, topic)
	} else {
		fmt.Println("From Redis")
		result = gin.H{
			"result":  topic,
			"count":   1,
			"message": "Result from Redis",
		}
	}
	c.JSON(http.StatusOK, result)
}

//Handler to POST a topic in database
func AddTopic(c *gin.Context) {
	var buffer bytes.Buffer
	name := c.PostForm("TopicName")
	sub := c.PostForm("SubjectID")
	stmt, err := config.Db.Prepare("insert into topics (TopicName, SubjectID) values(?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(name, sub)

	if err != nil {
		fmt.Print(err.Error())
	}
	buffer.WriteString(name)
	defer stmt.Close()
	name2 := buffer.String()
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Topic %s successfully created", name2),
	})
}

//Handler to PUT/update a Topic
func UpdateTopic(c *gin.Context) {

	var buffer bytes.Buffer
	id := c.Query("TopicID")
	name := c.PostForm("TopicName")
	sub := c.PostForm("SubjectID")
	stmt, err := config.Db.Prepare("update topics set  TopicName= ?,SubjectID = ? where  TopicID= ?;")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(name, sub, id)
	if err != nil {
		fmt.Print(err.Error())
	}
	buffer.WriteString(name)
	name2 := buffer.String()
	defer stmt.Close()
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully updated ID %s with name %s of subject ID %s", id, name2, sub),
	})
}

// Handler to DELETE a Topic
func DelTopic(c *gin.Context) {
	id := c.Query("TopicID")
	stmt, err := config.Db.Prepare("delete from topics where TopicID=?;")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully deleted id %s", id),
	})
}

//Handler to get all child table details for a topic by ID
func Topic_Childs(c *gin.Context) {
	var (
		child  entity.TopicChild
		childs []entity.TopicChild
	)
	id := c.Param("id")
	myQuery := `Select topics.TopicName,subtopic.ST_Name,concept.Name as ConceptName, 
	videosegment.Name as SegmentName, videosegment.Duration,video.URL
	from topics 
	inner join subtopic on (topics.TopicID=subtopic.TopicID)
	inner join concept on (subtopic.SubtopicID=concept.SubtopicID)
	inner join videosegment on (concept.ConceptID=videosegment.ConceptID)
	inner join video on (video.SegmentID=videosegment.SegmentID)
	where topics.TopicID=?`
	rows, err := config.Db.Query(myQuery, id)
	if err != nil {
		fmt.Println(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&child.TopicName, &child.ST_Name, &child.ConceptName, &child.SegmentName, &child.Duration, &child.URL)
		if err != nil {
			fmt.Println(err.Error())
		}
		childs = append(childs, child)
	}
	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{
		"result": childs,
		"count":  len(childs),
	})
}
