package controller

import (
	"fmt"
	"net/http"

	"api-final/config"
	"api-final/entity"

	"github.com/gin-gonic/gin"
)

//Handler to GET all concepts in database
func AllConcepts(c *gin.Context) {
	var (
		concept  entity.Concept
		concepts []entity.Concept
	)
	rows, err := config.Db.Query("select * from concept;")
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&concept.ConceptID, &concept.Name, &concept.Description, &concept.SubTopicID)
		concepts = append(concepts, concept)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{
		"result": concepts,
		"count":  len(concepts),
	})
}

//Handler to GET concept by ID
func Concept_by_ID(c *gin.Context) {
	var (
		concept entity.Concept
		result  gin.H
	)
	id := c.Param("id")
	row := config.Db.QueryRow("select ConceptID,Name,Description,SubtopicID from concept where ConceptID = ?", id)
	err := row.Scan(&concept.ConceptID, &concept.Name, &concept.Description, &concept.SubTopicID)
	if err != nil {
		// If no results send null
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": concept,
			"count":  1,
		}
	}
	c.JSON(http.StatusOK, result)
}

//Handler to POST a concept in database
func AddConcept(c *gin.Context) {
	name := c.PostForm("Name")
	des := c.PostForm("Description")
	st := c.PostForm("SubtopicID")
	stmt, err := config.Db.Prepare("insert into concept (Name,Description,SubtopicID) values(?,?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(name, des, st)

	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Concept %s successfully created", name),
	})
}

//Handler to PUT/update a Concept
func UpdateConcept(c *gin.Context) {

	id := c.Query("ConceptID")
	name := c.PostForm("Name")
	des := c.PostForm("Description")
	st := c.PostForm("SubtopicID")
	stmt, err := config.Db.Prepare("update concept set  Name = ?, Description = ?, SubtopicID = ? where  ConceptID= ?;")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(name, des, st, id)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully updated ID %s with name %s", id, name),
	})
}

// Handler to DELETE a Concept
func DelConcept(c *gin.Context) {
	id := c.Query("ConceptID")
	stmt, err := config.Db.Prepare("delete from concept where ConceptID=?;")
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
