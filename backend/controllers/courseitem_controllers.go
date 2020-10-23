package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ACER/app/ent"
	"github.com/ACER/app/ent/course"
	"github.com/ACER/app/ent/courseitem"
	"github.com/ACER/app/ent/subject"
	"github.com/ACER/app/ent/subjecttype"
	"github.com/gin-gonic/gin"
)

// CourseItemController defines the struct for the course item controller
type CourseItemController struct {
	client *ent.Client
	router gin.IRouter
}

// Course_Item struct
type Course_Item struct {
	Courses      int
	Subjects     int
	SubjectTypes int
}

// CreateCourseItem handles POST requests for adding courseitem entities
// @Summary Create courseitem
// @Description Create courseitem
// @ID create-courseitem
// @Accept   json
// @Produce  json
// @Param courseitem body ent.CourseItem true "courseitem entity"
// @Success 200 {object} ent.CourseItem
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /CourseItems [post]
func (ctl *CourseItemController) CreateCourseItem(c *gin.Context) {
	obj := Course_Item{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Course Item binding failed",
		})
		return
	}

	co, err := ctl.client.Course.
		Query().
		Where(course.IDEQ(int(obj.Courses))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Course not found",
		})
		return
	}

	s, err := ctl.client.Subject.
		Query().
		Where(subject.IDEQ(int(obj.Subjects))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Subjects not found",
		})
		return
	}

	t, err := ctl.client.SubjectType.
		Query().
		Where(subjecttype.IDEQ(int(obj.SubjectTypes))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Subject Types not found",
		})
		return
	}

	ci, err := ctl.client.CourseItem.
		Create().
		SetCourses(co).
		SetSubjects(s).
		SetTypes(t).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": true,
		"data":   ci,
	})
}

// GetCourseItem handles GET requests to retrieve a courseitem entity
// @Summary Get a courseitem entity by ID
// @Description get courseitem by ID
// @ID get-courseitem
// @Produce  json
// @Param id path int true "courseitem ID"
// @Success 200 {object} ent.CourseItem
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /CourseItems/{id} [get]
func (ctl *CourseItemController) GetCourseItem(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	i, err := ctl.client.CourseItem.
		Query().
		Where(courseitem.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, i)
}

// ListCourseItem handles request to get a list of courseitem entities
// @Summary List courseitem entities
// @Description list courseitem entities
// @ID list-courseitem
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.CourseItem
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /CourseItems [get]
func (ctl *CourseItemController) ListCourseItem(c *gin.Context) {
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

	courseitems, err := ctl.client.CourseItem.
		Query().
		WithCourses().
		WithSubjects().
		WithTypes().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, courseitems)
}

// DeleteCourseItem handles DELETE requests to delete a courseitem entity
// @Summary Delete a courseitem entity by ID
// @Description get courseitem by ID
// @ID delete-courseitem
// @Produce  json
// @Param id path int true "courseitems ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /CourseItems/{id} [delete]
func (ctl *CourseItemController) DeleteCourseItem(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.CourseItem.
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

// UpdateCourseItem handles PUT requests to update a courseitem entity
// @Summary Update a courseitem entity by ID
// @Description update courseitem by ID
// @ID update-courseitem
// @Accept   json
// @Produce  json
// @Param id path int true "courseitem ID"
// @Param courseitem body ent.CourseItem true "courseitem entity"
// @Success 200 {object} ent.CourseItem
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /CourseItems/{id} [put]
func (ctl *CourseItemController) UpdateCourseItem(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.CourseItem{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Course Item binding failed",
		})
		return
	}
	obj.ID = int(id)
	ci, err := ctl.client.CourseItem.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, ci)
}

// NewCourseItemControllercreates and registers handles for the CourseItem controller
func NewCourseItemController(router gin.IRouter, client *ent.Client) *CourseItemController {
	ci := &CourseItemController{
		client: client,
		router: router,
	}

	ci.register()

	return ci

}

// InitCourseItemController registers routes to the main engine
func (ctl *CourseItemController) register() {
	CourseItems := ctl.router.Group("/CourseItems")

	CourseItems.GET("", ctl.ListCourseItem)

	// CRUD
	CourseItems.POST("", ctl.CreateCourseItem)
	CourseItems.GET(":id", ctl.GetCourseItem)
	CourseItems.PUT(":id", ctl.UpdateCourseItem)
	CourseItems.DELETE(":id", ctl.DeleteCourseItem)
}
