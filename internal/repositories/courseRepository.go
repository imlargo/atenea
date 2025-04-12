package repositories

import (
	"github.com/imlargo/atenea/internal/models"
)

type CourseRepository interface {
	GetByID(id int) (*models.Course, error)
	GetAll() ([]*models.Course, error)
	Create(user *models.Course) (*models.Course, error)
	Update(user *models.Course) error
	Delete(id int) error
}

type CourseRepositoryImpl struct {
	db int
}

func NewCourseRepository(db int) CourseRepository {
	return &CourseRepositoryImpl{db: db}
}

func (u *CourseRepositoryImpl) Create(user *models.Course) (*models.Course, error) {

	return user, nil
}

func (u *CourseRepositoryImpl) Delete(id int) error {

	return nil
}

func (u *CourseRepositoryImpl) GetAll() (users []*models.Course, err error) {

	return users, nil
}

func (u *CourseRepositoryImpl) GetByID(id int) (user *models.Course, err error) {
	return nil, nil
}

func (u *CourseRepositoryImpl) Update(user *models.Course) error {

	return nil
}
