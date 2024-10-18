package core

type Metadatos struct {
	Codigo              string
	Nombre              string
	Uab                 string
	Vigente             bool
	HorasPresenciales   int
	HorasNoPresenciales int
	Creditos            int
	Validable           bool
	Electiva            bool
	Descripcion         string
}

type Asignatura struct {
	Codigo              string  `json:"codigo"`
	Nombre              string  `json:"nombre"`
	Uab                 string  `json:"uab"`
	Vigente             bool    `json:"vigente"`
	HorasPresenciales   int     `json:"horasPresenciales"`
	HorasNoPresenciales int     `json:"horasNoPresenciales"`
	Creditos            int     `json:"creditos"`
	Validable           bool    `json:"validable"`
	Electiva            bool    `json:"electiva"`
	Descripcion         string  `json:"descripcion"`
	Contenido           string  `json:"contenido"`
	PlanesRelacionados  []*Plan `json:"planes_relacionados"`
}

type Plan struct {
	Codigo string `json:"codigo"`
	Nombre string `json:"nombre"`
}
