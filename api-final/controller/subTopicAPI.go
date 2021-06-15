package controller

import (
	"fmt"
	"net/http"

	"api-final/config"
	"api-final/entity"

	"github.com/gin-gonic/gin"
)

//Handler to GET all topics in database
func AllSubTopic(c *gin.Context) {
	var (
		subtopic  entity.SubTopic
		subtopics []entity.SubTopic
	)
	rows, err := config.Db.Query("select * from subtopic;")
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&subtopic.SubtopicID, &subtopic.ST_Name, &subtopic.TopicID)
		subtopics = append(subtopics, subtopic)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{
		"result": subtopics,
		"count":  len(subtopics),
	})
}

//Handler to GET topic by ID
func ST_by_ID(c *gin.Context) {
	var (
		subtopic entity.SubTopic
		result   gin.H
	)
	id := c.Param("id")
	row := config.Db.QueryRow("select SubtopicID,ST_Name,TopicID from subtopic where SubTopicID = ?", id)
	err := row.Scan(&subtopic.SubtopicID, &subtopic.ST_Name, &subtopic.TopicID)
	if err != nil {
		// If no results send null
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": subtopic,
			"count":  1,
		}
	}
	c.JSON(http.StatusOK, result)
}

//Handler to POST a topic in database
func AddSubTopic(c *gin.Context) {
	name := c.PostForm("ST_Name")
	topic := c.PostForm("TopicID")
	stmt, err := config.Db.Prepare("insert into subtopic (ST_Name, TopicID) values(?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(name, topic)

	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("SubTopic %s successfully created", name),
	})
}

//Handler to PUT/update a Topic
func UpdateSubTopic(c *gin.Context) {

	id := c.Query("SubtopicID")
	name := c.PostForm("ST_Name")
	topic := c.PostForm("TopicID")
	stmt, err := config.Db.Prepare("update subtopic set  ST_Name= ?,TopicID = ? where  SubtopicID= ?;")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(name, topic, id)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully updated ID %s with name %s of Topic ID %s", id, name, topic),
	})
}

// Handler to DELETE a Topic
func DelSubTopic(c *gin.Context) {
	id := c.Query("SubtopicID")
	stmt, err := config.Db.Prepare("delete from subtopic where SubtopicID=?;")
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
