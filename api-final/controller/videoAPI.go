package controller

import (
	"bytes"
	"fmt"
	"net/http"

	"api-final/config"
	"api-final/entity"

	"github.com/gin-gonic/gin"
)

//Handler to GET all videos in database
func AllVideo(c *gin.Context) {
	var (
		vid    entity.Video
		videos []entity.Video
	)
	rows, err := config.Db.Query("select * from video;")
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&vid.SegmentID, &vid.URL)
		videos = append(videos, vid)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{
		"result": videos,
		"count":  len(videos),
	})
}

//Handler to GET videos by ID
func Video_by_ID(c *gin.Context) {
	var (
		vid    entity.Video
		result gin.H
	)
	id := c.Param("id")
	row := config.Db.QueryRow("select SegmentID, URL from video where SegmentID = ?", id)
	err := row.Scan(&vid.SegmentID, &vid.URL)
	if err != nil {
		// If no results send null
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": vid,
			"count":  1,
		}
	}
	c.JSON(http.StatusOK, result)
}

//Handler to POST a video in database
func AddVideo(c *gin.Context) {
	var buffer bytes.Buffer
	url := c.PostForm("URL")
	stmt, err := config.Db.Prepare("insert into video URL value ?;")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(url)

	if err != nil {
		fmt.Print(err.Error())
	}
	buffer.WriteString(url)
	defer stmt.Close()
	name := buffer.String()
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("URL %s successfully created", name),
	})
}

//Handler to PUT/update a Video
func UpdateVideo(c *gin.Context) {

	var buffer bytes.Buffer
	url := c.PostForm("URL")
	id := c.Query("SegmentID")
	stmt, err := config.Db.Prepare("update video set  URL= ? where SegmentID = ?;")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(url, id)
	if err != nil {
		fmt.Print(err.Error())
	}
	buffer.WriteString(url)
	link := buffer.String()
	defer stmt.Close()
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully updated Segment ID %s with url %s", id, link),
	})
}

// Handler to DELETE a Video
func DelVideo(c *gin.Context) {
	id := c.Query("SegmentID")
	stmt, err := config.Db.Prepare("delete from video where SegmentID=?;")
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
