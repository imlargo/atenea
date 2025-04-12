package services

import (
	"github.com/imlargo/atenea/internal/models"
	"github.com/imlargo/atenea/internal/repositories"
)

type CourseService interface {
	GetAll() ([]*models.Course, error)
	GetByID(id int) (*models.Course, error)
	GetByCode(code string) (*models.Course, error)
	Create(user *models.Course) (*models.Course, error)
	Update(user *models.Course) error
	Delete(id int) error
}

type CourseServiceImpl struct {
	courseRepository repositories.CourseRepository
}

func NewCourseService(userRepository repositories.CourseRepository) CourseService {
	return &CourseServiceImpl{courseRepository: userRepository}
}

func (u *CourseServiceImpl) Create(user *models.Course) (*models.Course, error) {
	return u.courseRepository.Create(user)
}

func (u *CourseServiceImpl) Delete(id int) error {
	return u.courseRepository.Delete(id)
}

func (u *CourseServiceImpl) GetAll() ([]*models.Course, error) {
	return u.courseRepository.GetAll()
}

func (u *CourseServiceImpl) GetByID(id int) (*models.Course, error) {
	return u.courseRepository.GetByID(id)
}

func (u *CourseServiceImpl) GetByCode(code string) (*models.Course, error) {
	return nil, nil
}

func (u *CourseServiceImpl) Update(user *models.Course) error {
	return u.courseRepository.Update(user)
}
