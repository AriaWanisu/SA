package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ACER/app/controllers"
	_ "github.com/ACER/app/docs"
	"github.com/ACER/app/ent"
	"github.com/ACER/app/ent/teacher"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Teachers struct {
	Teacher []Teacher
}

type Teacher struct {
	Name  string
	Email string
}

type Courses struct {
	Course []Course
}

type Course struct {
	Name  string
	Owner int
}

type Subjects struct {
	Subject []Subject
}

type Subject struct {
	Name  string
	Owner int
}

type SubjectTypes struct {
	SubjectType []SubjectType
}

type SubjectType struct {
	Name string
}

// @title SUT SA Project API Course Item
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewCourseController(v1, client)
	controllers.NewCourseItemController(v1, client)
	controllers.NewSubjectController(v1, client)
	controllers.NewSubjectTypeController(v1, client)
	controllers.NewTeacherController(v1, client)

	// Set Teacher Data
	teachers := Teachers{
		Teacher: []Teacher{
			Teacher{"Chanwit Kaewkasi", "chanwit@gmail.com"},
			Teacher{"Name Surname", "me@example.com"},
		},
	}

	for _, u := range teachers.Teacher {
		client.Teacher.
			Create().
			SetTeacherName(u.Name).
			SetTeacherEmail(u.Email).
			Save(context.Background())
	}

	// Set Subject Types Data
	SubjectTypes := SubjectTypes{
		SubjectType: []SubjectType{
			SubjectType{"General Course"},
			SubjectType{"Science Course"},
		},
	}
	for _, t := range SubjectTypes.SubjectType {
		client.SubjectType.
			Create().
			SetTypeName(t.Name).
			Save(context.Background())
	}

	// Set Subjects Data
	Subjects := Subjects{
		Subject: []Subject{
			Subject{"SA", 1},
			Subject{"SE", 1},
			Subject{"Micro", 2},
		},
	}

	for _, s := range Subjects.Subject {

		u, err := client.Teacher.
			Query().
			Where(teacher.IDEQ(int(s.Owner))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.Subject.
			Create().
			SetSubjectName(s.Name).
			SetOwner(u).
			Save(context.Background())
	}

	// Set Courses Data
	Courses := Courses{
		Course: []Course{
			Course{"Computer Engineering A", 1},
			Course{"Computer Engineering B", 1},
			Course{"Computer Engineering C", 2},
		},
	}

	for _, co := range Courses.Course {

		u, err := client.Teacher.
			Query().
			Where(teacher.IDEQ(int(co.Owner))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.Course.
			Create().
			SetCourseName(co.Name).
			SetOwner(u).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
