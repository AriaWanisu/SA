package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ACER/app/ent"
	"github.com/ACER/app/ent/teacher"
	"github.com/gin-gonic/gin"
)

// TeacherController defines the struct for the teacher controller
type TeacherController struct {
	client *ent.Client
	router gin.IRouter
}

// Teacher struct
type Teacher struct {
	id    int
	name  string
	email string
}

// CreateTeacher handles POST requests for adding Teacher entities
// @Summary Create Teacher
// @Description Create Teacher
// @ID create-Teacher
// @Accept   json
// @Produce  json
// @Param Teacher body ent.Course true "Teacher entity"
// @Success 200 {object} ent.Teacher
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Teachers [post]
func (ctl *TeacherController) CreateTeacher(c *gin.Context) {
	obj := Teacher{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Teacher binding failed",
		})
		return
	}

	u, err := ctl.client.Teacher.
		Create().
		SetTeacherEmail(obj.email).
		SetTeacherName(obj.name).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetTeacher handles GET requests to retrieve a Teacher entity
// @Summary Get a Teacher entity by ID
// @Description get Teacher by ID
// @ID get-Teacher
// @Produce  json
// @Param id path int true "Teacher ID"
// @Success 200 {object} ent.Teacher
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Teachers/{id} [get]
func (ctl *TeacherController) GetTeacher(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.Teacher.
		Query().
		Where(teacher.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListTeacher handles request to get a list of Teacher entities
// @Summary List Teacher entities
// @Description list Teacher entities
// @ID list-Teacher
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Teacher
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Teachers [get]
func (ctl *TeacherController) ListTeacher(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	Teachers, err := ctl.client.Teacher.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, Teachers)
}

// DeleteTeacher handles DELETE requests to delete a Teacher entity
// @Summary Delete a Teacher entity by ID
// @Description get Teacher by ID
// @ID delete-Teacher
// @Produce  json
// @Param id path int true "Teacher ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Teachers/{id} [delete]
func (ctl *TeacherController) DeleteTeacher(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Teacher.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateTeacher handles PUT requests to update a Teacher entity
// @Summary Update a Teacher entity by ID
// @Description update Teacher by ID
// @ID update-Teacher
// @Accept   json
// @Produce  json
// @Param id path int true "Teacher ID"
// @Param Teacher body ent.Teacher true "Teacher entity"
// @Success 200 {object} ent.Teacher
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Teachers/{id} [put]
func (ctl *TeacherController) UpdateTeacher(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Teacher{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Teacher binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	u, err := ctl.client.Teacher.
		UpdateOneID(int(id)).
		SetTeacherEmail(obj.TeacherEmail).
		SetTeacherName(obj.TeacherName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "update failed",
		})
		return
	}

	c.JSON(200, u)
}

// NewTeacherControllercreates and registers handles for the Teacher controller
func NewTeacherController(router gin.IRouter, client *ent.Client) *TeacherController {
	tc := &TeacherController{
		client: client,
		router: router,
	}

	tc.register()

	return tc

}

// InitTeacherController registers routes to the main engine
func (ctl *TeacherController) register() {
	Teachers := ctl.router.Group("/Teacher")

	Teachers.GET("", ctl.ListTeacher)

	// CRUD
	Teachers.POST("", ctl.CreateTeacher)
	Teachers.GET(":id", ctl.GetTeacher)
	Teachers.PUT(":id", ctl.UpdateTeacher)
	Teachers.DELETE(":id", ctl.DeleteTeacher)
}
