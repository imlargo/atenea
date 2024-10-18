package main

import (
	"extractor-contenido/core"
)

func main() {
	codigo := "123"
	asignatura := core.GetContenidoAsignatura(codigo)
	println(asignatura.Nombre)
}
