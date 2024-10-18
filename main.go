package main

import (
	"encoding/json"
	"extractor-contenido/core"
	"os"
	"sync"
)

func loadCodigosFacultad() []string {
	var codigos []string

	bytesCodigos, _ := os.ReadFile("codigos.json")
	json.Unmarshal(bytesCodigos, &codigos)

	return codigos
}

func main() {
	totalCodigos := loadCodigosFacultad()

	sizeGrupo := 20
	gruposCodigos := make([][]string, 0)

	for i := 0; i < len(totalCodigos); i++ {
		if i%sizeGrupo == 0 {
			gruposCodigos = append(gruposCodigos, make([]string, 0))
		}
		gruposCodigos[len(gruposCodigos)-1] = append(gruposCodigos[len(gruposCodigos)-1], totalCodigos[i])
	}

	println(len(gruposCodigos))

	var wg sync.WaitGroup
	chanel := make(chan *core.Asignatura)

	for _, codigos := range gruposCodigos {
		wg.Add(1)
		go func(codigos []string) {
			defer wg.Done()
			for _, codigo := range codigos {
				asignatura := core.GetContenidoAsignatura(codigo)
				if asignatura.Nombre == "" {
					println("Err: ", codigo)
					continue
				}
				chanel <- asignatura
			}
		}(codigos)
	}

	go func() {
		wg.Wait()
		close(chanel)
	}()

	asignaturas := make([]*core.Asignatura, 0)
	for asignatura := range chanel {
		asignaturas = append(asignaturas, asignatura)
	}

	// Marshall json
	json, _ := json.Marshal(asignaturas)
	os.WriteFile("data.json", json, 0644)

}
