package atenea

import "github.com/imlargo/atenea/internal/models"

type AteneaService struct {
	driver driver
}

func NewAteneaService(siaUrl string) *AteneaService {
	d := NewDriver(siaUrl)

	return &AteneaService{
		driver: d,
	}
}

func (atenea *AteneaService) GetCourse(code string) (*models.Course, error) {

	course, err := atenea.driver.GetCourse(code)
	if err != nil {
		return nil, err
	}

	return course, nil
}

func (atenea *AteneaService) IsCodeValid(code string) bool {
	length := len(code)
	return length > 3 && length < 15
}
