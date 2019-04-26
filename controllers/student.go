package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/crud_golang/models"
	"github.com/crud_golang/repository"
	"github.com/crud_golang/services"
)

var studentRepo = repository.MySqlStudentRepo{Students: make(map[int64]models.Student)}

/******************************************************************************************************/

func GetAllStudents(c *gin.Context) {

	students, err := studentRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, students)

}

/******************************************************************************************************/

func GetStudent(c *gin.Context) {

	id, err := strconv.Atoi(c.Params.ByName("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID deve ser númerico."})
		return
	}

	student, err := studentRepo.Get(int64(id))

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return

	}

	c.JSON(http.StatusOK, student)
}

/******************************************************************************************************/

func Save(c *gin.Context) {

	student := models.Student{}

	body, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(body, &student)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid json, please try send with raw body"})
		return
	}

	resp, err := studentRepo.Save(&student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	//Send mail with JSON to new student
	services.Send(student.Email, student)

	c.JSON(http.StatusOK, gin.H{"id": resp.ID, "name": resp.Name, "age": resp.Age, "course": resp.Course, "email": resp.Email})

}

/******************************************************************************************************/

func Delete(c *gin.Context) {

	id, err := strconv.Atoi(c.Params.ByName("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Id must be number", "cause": err.Error()})
		return
	}

	student, erro := studentRepo.Delete(int64(id))

	if erro != nil {

		c.JSON(http.StatusBadRequest, gin.H{"message": erro.Error()})
		return

	}
	c.JSON(http.StatusOK, gin.H{"message": "Student with id " + strconv.FormatInt(student.ID, 10) + " was deleted"})
}

/******************************************************************************************************/

func Update(c *gin.Context) {

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Id must be number", "cause": err.Error()})
		return
	}

	student := models.Student{}

	err = c.ShouldBindJSON(&student)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid json, please try send with raw body or verified if the fields are in correct format"})
		return
	}

	student.ID = int64(id)

	resp, err := studentRepo.Update(&student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": resp.ID, "name": resp.Name, "age": resp.Age, "email": resp.Email, "course": resp.Course})

}
