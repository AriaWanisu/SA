package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ACER/app/ent"
	"github.com/ACER/app/ent/course"
	"github.com/ACER/app/ent/teacher"
	"github.com/gin-gonic/gin"
)

// CourseController defines the struct for the course controller
type CourseController struct {
	client *ent.Client
	router gin.IRouter
}

// Course struct
type Course struct {
	name  string
	Owner int
}

// CreateCourse handles POST requests for adding Course entities
// @Summary Create Course
// @Description Create Course
// @ID create-Course
// @Accept   json
// @Produce  json
// @Param Course body ent.Course true "Course entity"
// @Success 200 {object} ent.Course
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Course [post]
func (ctl *CourseController) CreateCourse(c *gin.Context) {
	obj := Course{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Course binding failed",
		})
		return
	}

	u, err := ctl.client.Teacher.
		Query().
		Where(teacher.IDEQ(int(obj.Owner))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "owner not found",
		})
		return
	}

	p, err := ctl.client.Course.
		Create().
		SetCourseName(obj.name).
		SetOwner(u).
		Save(context.Background())

	fmt.Println(err)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, p)
}

// GetCourse handles GET requests to retrieve a Course entity
// @Summary Get a Course entity by ID
// @Description get Course by ID
// @ID get-Course
// @Produce  json
// @Param id path int true "Course ID"
// @Success 200 {object} ent.Course
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Courses/{id} [get]
func (ctl *CourseController) GetCourse(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	p, err := ctl.client.Course.
		Query().
		Where(course.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

// ListCourse handles request to get a list of Course entities
// @Summary List Course entities
// @Description list Course entities
// @ID list-Course
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Course
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Course [get]
func (ctl *CourseController) ListCourse(c *gin.Context) {
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

	Courses, err := ctl.client.Course.
		Query().
		Where(course.HasOwnerWith(teacher.IDEQ(1))).
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, Courses)
}

// NewCourseControllercreates and registers handles for the Course controller
func NewCourseController(router gin.IRouter, client *ent.Client) *CourseController {
	cc := &CourseController{
		client: client,
		router: router,
	}

	cc.register()

	return cc
}

// InitCourseController registers routes to the main engine
func (ctl *CourseController) register() {
	Courses := ctl.router.Group("/Course")

	Courses.POST("", ctl.CreateCourse)
	Courses.GET(":id", ctl.GetCourse)
	Courses.GET("", ctl.ListCourse)

}
