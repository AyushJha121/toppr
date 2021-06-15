package handlers

import (
	"api-final/controller"

	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()

	//Hanle Queries for subjects
	router.GET("/subject/:id", controller.Sub_by_ID)
	router.GET("/subjects", controller.AllSub)
	router.POST("/subject", controller.AddSubject)
	router.PUT("/subject", controller.UpdateSub)
	router.DELETE("/subject", controller.DelSub)

	//Handle Queries for Topics
	router.GET("/topic/:id", controller.Topic_by_ID)
	router.GET("/topics", controller.AllTopic)
	router.POST("/topic", controller.AddTopic)
	router.PUT("/topic", controller.UpdateTopic)
	router.DELETE("/topic", controller.DelTopic)
	router.GET("/topic/:id/all", controller.Topic_Childs)

	//Handle Queries for SubTopics
	router.GET("/subtopic/:id", controller.ST_by_ID)
	router.GET("/subtopics", controller.AllSubTopic)
	router.POST("/subtopic", controller.AddSubTopic)
	router.PUT("/subtopic", controller.UpdateSubTopic)
	router.DELETE("/subtopic", controller.DelSubTopic)

	//Handle Queries for Concepts
	router.GET("/concept/:id", controller.Concept_by_ID)
	router.GET("/concepts", controller.AllConcepts)
	router.POST("/concept", controller.AddConcept)
	router.PUT("/concept", controller.UpdateConcept)
	router.DELETE("/concept", controller.DelConcept)

	//Handle Queries for videos
	router.GET("/video/:id", controller.Video_by_ID)
	router.GET("/videos", controller.AllVideo)
	router.POST("/video", controller.AddVideo)
	router.PUT("/video", controller.UpdateVideo)
	router.DELETE("/video", controller.DelVideo)

	router.Run(":3000")
}
