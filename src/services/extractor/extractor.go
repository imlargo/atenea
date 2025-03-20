package extractor

import "atenea/src/models"

const (
	BaseUrl       = "https://siamed.unal.edu.co/academia/apoyo-administrativo/ConsultaContenidos.do?action=Info&idAsignatura="
	NO_ENCONTRADO = ""
)

func GetContenidoAsignatura(codigo string) *models.Asignatura {

	document := getDocumentFromUrl(getUrl(codigo))

	contendores := document.Find(".zona-dato-caja")
	contenedorMetadatos := contendores.Eq(1)
	contenedorPlanes := contendores.Eq(2)
	contenedorContenido := contendores.Eq(3)

	metadatos := getMetadatos(contenedorMetadatos)
	if metadatos.Nombre == "" {
		return nil
	}

	planes := getPlanes(contenedorPlanes)
	contenido := getContenido(contenedorContenido)

	return &models.Asignatura{
		Codigo:              codigo,
		Nombre:              metadatos.Nombre,
		Uab:                 metadatos.Uab,
		Vigente:             metadatos.Vigente,
		HorasPresenciales:   metadatos.HorasPresenciales,
		HorasNoPresenciales: metadatos.HorasNoPresenciales,
		Creditos:            metadatos.Creditos,
		Validable:           metadatos.Validable,
		Electiva:            metadatos.Electiva,
		Descripcion:         metadatos.Descripcion,
		Contenido:           contenido,
		PlanesRelacionados:  planes,
	}

}
