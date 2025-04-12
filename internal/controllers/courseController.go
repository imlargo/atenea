package controllers

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/imlargo/atenea/internal/atenea"
	"github.com/imlargo/atenea/internal/env"
	"github.com/imlargo/atenea/internal/models"
	"github.com/imlargo/atenea/internal/responses"
	"github.com/imlargo/atenea/internal/services"
)

type CourseController interface {
	GetAll(c *gin.Context)
	GetById(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type CourseControllerImpl struct {
	courseService services.CourseService
}

func NewCourseController(CourseService services.CourseService) CourseController {
	return &CourseControllerImpl{courseService: CourseService}
}

// @Summary		Search All Courses
// @Router			/courses [get]
// @Description	Search All Courses
// @Tags			courses
// @Accept			json
// @Produce		json
// @Success		200	{object}    models.SuccessList[models.Course] "OK"
// @Failure		500	{object}	models.Error	"Internal Server Error"
func (u *CourseControllerImpl) GetAll(c *gin.Context) {

	var courses []*models.Course

	courses, errGet := u.courseService.GetAll()

	if errGet != nil {
		responses.ErrorInternalServer(c)
		return
	}

	responses.List(c, courses)

}

// @Summary		Search Course By ID
// @Router			/courses/{id} [get]
// @Description	Get Course By ID
// @Tags			courses
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"Course ID"
// @Success		200	{object}	models.SuccessData[models.Course] "OK"
// @Failure		400	{object}	models.Error	"Bad Request"
// @Failure		404	{object}	models.Error	"Not Found"
// @Failure		500	{object}	models.Error	"Internal Server Error"
func (u *CourseControllerImpl) GetById(c *gin.Context) {

	courseID := c.Param("id")

	if courseID == "" {
		responses.ErrorBadRequest(c, "Invalid id")
		return
	}

	driver := atenea.NewAteneaService(os.Getenv(env.SIA_URL))
	course, err := driver.GetCourse(courseID)

	if err != nil {
		switch err {
		case atenea.ErrCourseNotFound:
			responses.ErrorNotFound(c, err.Error())
		case atenea.ErrDataExtractionFailed:
			responses.ErrorBadRequest(c, err.Error())
		case atenea.ErrInternal:
			responses.ErrorInternalServer(c)
		default:
			responses.ErrorInternalServer(c)
		}
		return
	}

	if !c.IsAborted() {
		responses.Ok(c, course)
		return
	}
}

// @Summary		Create Course
// @Router			/courses [post]
// @Description	Create Course With The Given Input Data
// @Tags			courses
// @Accept			json
// @Produce		json
// @Param			Input	body		models.Course		true	"Create course object"
// @Success		201	{object}	models.SuccessData[models.Course] "Created"
// @Failure		400	{object}	models.Error	"Bad Request"
// @Failure		409 {object}	models.Error	"Error Conflict"
// @Failure		500	{object}	models.Error	"Internal Server Error"
func (u *CourseControllerImpl) Create(c *gin.Context) {

	var newCourse *models.Course

	if errBind := c.ShouldBindJSON(&newCourse); errBind != nil {
		responses.ErrorBindJson(c, errBind)
		return
	}

	course, errCreate := u.courseService.Create(newCourse)

	if errCreate != nil {
		responses.ErrorInternalServer(c)
		return
	}

	responses.Create(c, course)
}

// @Summary		Update Course By ID
// @Router			/courses [put]
// @Description	Update Data Course By ID With The Given Input Data
// @Tags			courses
// @Accept			json
// @Produce		json
// @Param			Input	body		models.Course		true	"Update course object"
// @Success		200	{object}	models.SuccessData[models.Course] "Updated"
// @Failure		400	{object}	models.Error	"Bad Request"
// @Failure		404		{object}	models.Error	"Not Found"
// @Failure		409 {object}	models.Error	"Error Conflict"
// @Failure		500		{object}	models.Error	"Internal Server Error"
func (u *CourseControllerImpl) Update(c *gin.Context) {

	var course *models.Course

	if errBind := c.ShouldBindJSON(&course); errBind != nil {
		responses.ErrorBindJson(c, errBind)
		return
	}

	if _, errGet := u.courseService.GetByID(course.ID); errGet != nil {
		responses.ErrorInternalServer(c)
		return
	}

	if errUpdate := u.courseService.Update(course); errUpdate != nil {
		responses.ErrorInternalServer(c)
		return
	}

	responses.Update(c, course)
}

// @Summary		Delete Course By ID
// @Router			/courses/{id} [delete]
// @Description	Delete Course By ID
// @Tags			courses
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"Course ID"
// @Success		200	{object}	models.Success[models.Course] "Deleted"
// @Failure		400	{object}	models.Error	"Bed Request"
// @Failure		404	{object}	models.Error	"Not Found"
// @Failure		500	{object}	models.Error	"Internal Server Error"
func (u *CourseControllerImpl) Delete(c *gin.Context) {

	courseID, errParse := strconv.Atoi(c.Param("id"))

	if errParse != nil {
		responses.ErrorBadRequest(c, "UUID invalid")
		return
	}

	if _, errGet := u.courseService.GetByID(courseID); errGet != nil {
		responses.ErrorInternalServer(c)
		return

	}

	if errDelete := u.courseService.Delete(courseID); errDelete != nil {
		responses.ErrorInternalServer(c)
		return
	}

	responses.Delete(c)

}
