package atenea

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/imlargo/atenea/internal/models"
)

type driver interface {
	GetCourse(code string) (*models.Course, error)
}

type DriverImpl struct {
	driver
	SiaUrl string
}

func NewDriver(siaUrl string) *DriverImpl {
	return &DriverImpl{
		SiaUrl: siaUrl,
	}
}

func (d *DriverImpl) GetCourse(code string) (*models.Course, error) {

	doc, err := d.getCourseDocument(code)
	if err != nil {
		return nil, err
	}

	course, err := d.getInfoFromDoc(doc)
	if err != nil {
		return nil, errDataExtractionFailed
	}

	if course.Nombre == "" {
		return nil, errCourseNotFound
	}

	course.Codigo = code

	return course, nil
}

func (d *DriverImpl) getInfoFromDoc(document *goquery.Document) (*models.Course, error) {
	containers := document.Find(".zona-dato-caja")
	metadataSection := containers.Eq(1)
	careersSection := containers.Eq(2)
	contentSection := containers.Eq(3)

	metadata, err := d.extractMetadata(metadataSection)
	if err != nil {
		return nil, errDataExtractionFailed
	}

	careers := d.extractCareers(careersSection)
	content := d.extractContent(contentSection)

	course := &models.Course{
		Codigo:              metadata.Codigo,
		Nombre:              metadata.Nombre,
		Uab:                 metadata.Uab,
		Vigente:             metadata.Vigente,
		HorasPresenciales:   metadata.HorasPresenciales,
		HorasNoPresenciales: metadata.HorasNoPresenciales,
		Creditos:            metadata.Creditos,
		Validable:           metadata.Validable,
		Electiva:            metadata.Electiva,
		Descripcion:         metadata.Descripcion,
		Contenido:           content,
		PlanesRelacionados:  careers,
	}

	return course, nil

}

func (d *DriverImpl) extractMetadata(section *goquery.Selection) (*CourseMetadata, error) {
	metadata := CourseMetadata{}

	section.Find("h3").Each(func(i int, s *goquery.Selection) {
		title, _ := normalizeString(s.Text())

		switch title {
		case "asignatura vigente":
			metadata.Vigente = strings.TrimSpace(s.Nodes[0].NextSibling.NextSibling.NextSibling.Data) == "Si"
		case "nombre asignatura":
			metadata.Nombre = strings.TrimSpace(s.Nodes[0].NextSibling.NextSibling.NextSibling.Data)
		case "unidad academica basica":
			metadata.Uab = strings.TrimSpace(s.Nodes[0].NextSibling.NextSibling.NextSibling.Data)
		case "horas presenciales":
			n, _ := strconv.Atoi(strings.TrimSpace(s.Nodes[0].NextSibling.NextSibling.NextSibling.Data))
			metadata.HorasPresenciales = n
		case "horas no presenciales":
			n, _ := strconv.Atoi(strings.TrimSpace(s.Nodes[0].NextSibling.NextSibling.NextSibling.Data))
			metadata.HorasNoPresenciales = n
		case "creditos":
			n, _ := strconv.Atoi(strings.TrimSpace(s.Nodes[0].NextSibling.NextSibling.NextSibling.Data))
			metadata.Creditos = n
		case "validable":
			metadata.Validable = strings.TrimSpace(s.Nodes[0].NextSibling.NextSibling.NextSibling.Data) == "Si"
		case "libre eleccion":
			metadata.Electiva = strings.TrimSpace(s.Nodes[0].NextSibling.NextSibling.NextSibling.Data) == "Si"
		case "descripcion":
			metadata.Descripcion = strings.TrimSpace(s.Nodes[0].NextSibling.NextSibling.NextSibling.Data)
		}
	})

	return &metadata, nil
}

func (d *DriverImpl) extractCareers(section *goquery.Selection) []models.Plan {

	rows := section.Find("tr")
	if rows.Length() <= 1 {
		return nil // No careers found, return an empty slice
	}

	rows = rows.Slice(1, goquery.ToEnd)
	careers := make([]models.Plan, rows.Length())

	rows.Each(func(i int, s *goquery.Selection) {
		tds := s.Find("td")
		codigo := tds.Eq(0).Text()
		plan := tds.Eq(1).Text()

		careers[i] = models.Plan{
			ID:     0,
			Codigo: codigo,
			Nombre: plan,
		}
	})

	return careers
}

func (d *DriverImpl) extractContent(section *goquery.Selection) string {
	raw := regexp.MustCompile(`\n\n+`).ReplaceAllString(section.Text(), "\n\n")
	return strings.TrimSpace(raw)
}

func (d *DriverImpl) getCourseDocument(code string) (*goquery.Document, error) {

	url := d.SiaUrl + "/academia/apoyo-administrativo/ConsultaContenidos.do?action=Info&idAsignatura=" + code

	res, err := http.Get(url)
	if err != nil {
		return nil, errCourseNotFound
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errCourseNotFound
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, errInternal
	}

	return doc, nil
}
