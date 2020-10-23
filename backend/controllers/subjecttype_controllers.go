package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ACER/app/ent"
	"github.com/ACER/app/ent/subjecttype"
	"github.com/gin-gonic/gin"
)

// SubjectTypeController defines the struct for the subject type controller
type SubjectTypeController struct {
	client *ent.Client
	router gin.IRouter
}

// SubjectType struct
type SubjectType struct {
	id   int
	name string
}

// CreateSubjectType handles POST requests for adding SubjectType entities
// @Summary Create SubjectType
// @Description Create SubjectType
// @ID create-SubjectType
// @Accept   json
// @Produce  json
// @Param SubjectType body ent.Course true "SubjectType entity"
// @Success 200 {object} ent.SubjectType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /SubjectTypes [post]
func (ctl *SubjectTypeController) CreateSubjectType(c *gin.Context) {
	obj := SubjectType{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Subject Type binding failed",
		})
		return
	}

	p, err := ctl.client.SubjectType.
		Create().
		SetTypeName(obj.name).
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

// GetSubjectType handles GET requests to retrieve a SubjectType entity
// @Summary Get a SubjectType entity by ID
// @Description get SubjectType by ID
// @ID get-SubjectType
// @Produce  json
// @Param id path int true "SubjectType ID"
// @Success 200 {object} ent.SubjectType
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /SubjectTypes/{id} [get]
func (ctl *SubjectTypeController) GetSubjectType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	p, err := ctl.client.SubjectType.
		Query().
		Where(subjecttype.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

// ListSubjectType handles request to get a list of SubjectType entities
// @Summary List SubjectType entities
// @Description list SubjectType entities
// @ID list-SubjectType
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.SubjectType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /SubjectTypes [get]
func (ctl *SubjectTypeController) ListSubjectType(c *gin.Context) {
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

	SubjectTypes, err := ctl.client.SubjectType.
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

	c.JSON(200, SubjectTypes)
}

// NewSubjectTypeControllercreates and registers handles for the SubjectType controller
func NewSubjectTypeController(router gin.IRouter, client *ent.Client) *SubjectTypeController {
	st := &SubjectTypeController{
		client: client,
		router: router,
	}

	st.register()

	return st

}


// InitSubjectTypeController registers routes to the main engine
func (ctl *SubjectTypeController) register() {
	SubjectTypes := ctl.router.Group("/SubjectTypes")

	SubjectTypes.POST("", ctl.CreateSubjectType)
	SubjectTypes.GET(":id", ctl.GetSubjectType)
	SubjectTypes.GET("", ctl.ListSubjectType)

}
