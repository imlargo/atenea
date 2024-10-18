package core

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getDataByTitle(selection *goquery.Selection) string {
	return strings.TrimSpace(selection.Nodes[0].NextSibling.NextSibling.NextSibling.Data)
}
func getDescripcion(selection *goquery.Selection) string {
	return strings.TrimSpace(selection.Nodes[0].NextSibling.NextSibling.NextSibling.Data)
}

func getMetadatos(contenedor *goquery.Selection) *Metadatos {

	metadatos := Metadatos{}

	contenedor.Find("h3").Each(func(i int, s *goquery.Selection) {

		title := normalizeString(s.Text())

		switch title {
		case "asignatura vigente":
			metadatos.Vigente = getDataByTitle(s) == "Si"

		case "nombre asignatura":
			metadatos.Nombre = getDataByTitle(s)
		case "unidad academica basica":
			metadatos.Uab = getDataByTitle(s)
		case "horas presenciales":
			n, _ := strconv.Atoi(getDataByTitle(s))
			metadatos.HorasPresenciales = n
		case "horas no presenciales":
			n, _ := strconv.Atoi(getDataByTitle(s))
			metadatos.HorasNoPresenciales = n
		case "creditos":
			n, _ := strconv.Atoi(getDataByTitle(s))
			metadatos.Creditos = n
		case "validable":
			metadatos.Validable = getDataByTitle(s) == "Si"
		case "libre eleccion":
			metadatos.Electiva = getDataByTitle(s) == "Si"
		case "descripcion":
			metadatos.Descripcion = getDescripcion(s)

		}
	})

	return &metadatos
}

func getPlanes(contenedor *goquery.Selection) []*Plan {

	rows := contenedor.Find("tr").Slice(1, goquery.ToEnd)
	planes := make([]*Plan, rows.Length())

	rows.Each(func(i int, s *goquery.Selection) {
		tds := s.Find("td")
		codigo := tds.Eq(0).Text()
		plan := tds.Eq(1).Text()

		planes[i] = &Plan{
			Codigo: codigo,
			Nombre: plan,
		}
	})

	return planes
}

func getContenido(contenedor *goquery.Selection) string {

	// Regex para eliminar los saltos de linea y dejarlos solo con un salto de linea
	raw := regexp.MustCompile(`\n\n+`).ReplaceAllString(contenedor.Text(), "\n\n")

	return strings.TrimSpace(raw)
}
