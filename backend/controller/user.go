package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanapon395/sa-67-example/config"
	"github.com/tanapon395/sa-67-example/entity"
)

// POST /students
func CreateStudent(c *gin.Context) {
	var student entity.Student

	// Bind JSON to student struct
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB()

	// Find gender by ID
	var gender entity.Gender
	if err := db.First(&gender, student.GenderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "gender not found"})
		return
	}

	// Hash password and handle errors
	hashedPassword, err := config.HashPassword(student.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	// Create Student
	u := entity.Student{
		SID:       student.SID,
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Password:  hashedPassword,
		Birthday:  student.Birthday, // Correct field name
		Year:      student.Year,
		Major:     student.Major,
		GenderID:  student.GenderID,
		Gender:    gender, // Set the Gender relationship
	}

	// Save student to database
	if err := db.Create(&u).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Student created successfully", "data": u})
}

// GET /students/:id
func GetStudent(c *gin.Context) {
	ID := c.Param("id")
	var student entity.Student

	db := config.DB()
	if err := db.Preload("Gender").First(&student, ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}

// GET /students
func ListStudents(c *gin.Context) {
	var students []entity.Student

	db := config.DB()
	if err := db.Preload("Gender").Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, students)
}

// DELETE /students/:id
func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	db := config.DB()

	if err := db.Delete(&entity.Student{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
}

// PATCH /students/:id
func UpdateStudent(c *gin.Context) {
	var student entity.Student
	ID := c.Param("id")

	db := config.DB()
	if err := db.First(&student, ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
	}

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request, unable to map payload"})
		return
	}

	if err := db.Save(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to update student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student updated successfully"})
}
