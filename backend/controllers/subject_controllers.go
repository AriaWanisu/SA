package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ACER/app/ent"
	"github.com/ACER/app/ent/subject"
	"github.com/ACER/app/ent/teacher"
	"github.com/gin-gonic/gin"
)

// SubjectController defines the struct for the subject controller
type SubjectController struct {
	client *ent.Client
	router gin.IRouter
}

// Subject struct
type Subject struct {
	name  string
	Owner int
}

// CreateSubject  handles POST requests for adding Subject  entities
// @Summary Create Subject
// @Description Create Subject
// @ID create-Subject
// @Accept   json
// @Produce  json
// @Param Subject  body ent.Subject  true "Subject  entity"
// @Success 200 {object} ent.Subject
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Subjects  [post]
func (ctl *SubjectController) CreateSubject(c *gin.Context) {
	obj := Subject{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Subject binding failed",
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

	p, err := ctl.client.Subject.
		Create().
		SetSubjectName(obj.name).
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

// GetSubject GET requests to retrieve a Subject entity
// @Summary Get a Subject entity by ID
// @Description get Subject by ID
// @ID get-Subject
// @Produce  json
// @Param id path int true "Subject ID"
// @Success 200 {object} ent.Subject
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Subjects/{id} [get]
func (ctl *SubjectController) GetSubject(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	p, err := ctl.client.Subject.
		Query().
		Where(subject.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

// ListSubject handles request to get a list of Subject entities
// @Summary List Subject entities
// @Description list Subject entities
// @ID list-Subject
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Subject
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Subjects [get]
func (ctl *SubjectController) ListSubject(c *gin.Context) {
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

	Subjects, err := ctl.client.Subject.
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

	c.JSON(200, Subjects)
}

// NewSubjectControllercreates and registers handles for the Subject controller
func NewSubjectController(router gin.IRouter, client *ent.Client) *SubjectController {
	sj := &SubjectController{
		client: client,
		router: router,
	}

	sj.register()

	return sj

}

// InitSubjectController registers routes to the main engine
func (ctl *SubjectController) register() {
	Subjects := ctl.router.Group("/Subjects")

	Subjects.POST("", ctl.CreateSubject)
	Subjects.GET(":id", ctl.GetSubject)
	Subjects.GET("", ctl.ListSubject)

}
