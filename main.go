package main

import (
	"simple-api-gorm/models"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var errDb error

func index(context *gin.Context) {
	student := []models.Student{}

	db.Order("name asc").Find(&student)

	context.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Data retrivied successfully !",
		"data":    student,
	})

	return
}

func show(context *gin.Context) {
	student := models.Student{}

	var id, _ = strconv.ParseInt(context.Query("student_id"), 10, 64)

	db.Where("id = ?", id).First(&student)

	context.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Data retrivied successfully !",
		"data":    student,
	})

	return
}

func store(context *gin.Context) {
	student := models.Student{}

	student.Name = context.PostForm("name")
	student.Age, _ = strconv.ParseInt(context.PostForm("age"), 10, 64)
	student.Address = context.PostForm("address")
	student.PhoneNumber = context.PostForm("phone_number")

	db.Save(&student)

	context.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Data create successfully !",
		"data":    student,
	})

	return
}

func update(context *gin.Context) {
	student := models.Student{}

	var id, _ = strconv.ParseInt(context.PostForm("student_id"), 10, 64)

	db.Where("id = ?", id).First(&student)

	student.Name = context.PostForm("name")
	student.Age, _ = strconv.ParseInt(context.PostForm("age"), 10, 64)
	student.Address = context.PostForm("address")
	student.PhoneNumber = context.PostForm("phone_number")

	db.Save(student)

	context.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Data updated successfully !",
		"data":    student,
	})
}

func destroy(context *gin.Context) {
	student := models.Student{}

	var id, _ = strconv.ParseInt(context.PostForm("student_id"), 10, 64)

	db.Delete(&student, id)

	context.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Data deleted successfully !",
	})
}

func setupRouter() *gin.Engine {
	db, errDb = gorm.Open("mysql", "fikrisabriansyah:fikrisabriansyah@tcp(127.0.0.1:3306)/test_golang?charset=utf8&parseTime=True")

	if errDb != nil {
		panic("Connection failed !")
	}

	migrate()

	route := gin.Default()

	route.GET("/student", func(context *gin.Context) {
		index(context)
	})

	route.GET("/student/show", func(context *gin.Context) {
		show(context)
	})

	route.POST("/student", func(context *gin.Context) {
		store(context)
	})

	route.PUT("/student", func(context *gin.Context) {
		update(context)
	})

	route.DELETE("/student", func(context *gin.Context) {
		destroy(context)
	})

	return route
}

func migrate() {
	db.DropTableIfExists(&models.Student{})
	db.AutoMigrate(&models.Student{})
	seederUser()
}

func seederUser() {
	var student = new(models.Student)

	student.Name = "Fikri"
	student.Age = 23
	student.Address = "Bandung"
	student.PhoneNumber = "081123456789"

	db.Create(student)
}

func main() {
	r := setupRouter()

	r.Run(":8080")
}
