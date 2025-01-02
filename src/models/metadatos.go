package models

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
