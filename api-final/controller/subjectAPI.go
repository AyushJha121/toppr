package controller

import (
	"fmt"
	"net/http"

	"api-final/config"
	"api-final/entity"

	"github.com/gin-gonic/gin"
)

//Handler to GET all subjects in database
func AllSub(c *gin.Context) {
	var (
		subject  entity.Subject
		subjects []entity.Subject
	)
	rows, err := config.Db.Query("select * from subject;")
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&subject.SubjectID, &subject.SubName, &subject.Grade)
		subjects = append(subjects, subject)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	defer rows.Close()
	c.JSON(http.StatusOK, gin.H{
		"result": subjects,
		"count":  len(subjects),
	})
}

//Handler to GET Subject by ID
func Sub_by_ID(c *gin.Context) {
	var (
		subject entity.Subject
		result  gin.H
	)
	id := c.Param("id")
	row := config.Db.QueryRow("select SubjectID,SubName, Grade from subject where SubjectID = ?", id)
	err := row.Scan(&subject.SubjectID, &subject.SubName, &subject.Grade)
	if err != nil {
		// If no results send null
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": subject,
			"count":  1,
		}
	}
	c.JSON(http.StatusOK, result)
}

//Handler to POST a Subject in database
func AddSubject(c *gin.Context) {
	name := c.PostForm("SubName")
	grade := c.PostForm("Grade")
	stmt, err := config.Db.Prepare("insert into subject (SubName,Grade) values(?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(name, grade)

	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Subject %s successfully created", name),
	})
}

//Handler to PUT/update a Subject
func UpdateSub(c *gin.Context) {
	name := c.PostForm("SubName")
	grade := c.PostForm("Grade")
	id := c.Query("SubjectID")
	stmt, err := config.Db.Prepare("update subject set SubName = ?, Grade = ? where SubjectID = ?;")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(name, grade, id)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer stmt.Close()
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully updated Subject ID %s with Name %s", id, name),
	})
}

// Handler to DELETE a Subject
func DelSub(c *gin.Context) {
	id := c.Query("SubjectID")
	stmt, err := config.Db.Prepare("delete from subject where SubjectID=?;")
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
