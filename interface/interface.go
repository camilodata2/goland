package main

import (
	"fmt"
)

type Nota struct {
	score float64
}

type Estudiante struct {
	Nombre   string
	Apellido string
	Email    string
	Nota     Nota
}

func (e Estudiante) InformacionEstudiante() {
	fmt.Printf("Nombre: %s\n", e.Nombre)
	fmt.Printf("Apellido: %s\n", e.Apellido)
	fmt.Printf("Email: %s\n", e.Email)
	fmt.Printf("Nota: %d\n", e.Nota.score)

	if e.Nota.score > 3.5 {
		fmt.Println("Estado: Aprobado")
	} else {
		fmt.Println("Estado: No aprobado")
	}
	fmt.Println("--------------")
}

func NuevoEstudiante(nombre, apellido, email string, score int) *Estudiante {
	return &Estudiante{
		Nombre:   nombre,
		Apellido: apellido,
		Email:    email,
		Nota:     Nota{score: 3.8},
	}
}

type InfoEstudiante interface {
	InformacionEstudiante()
}

func ImprimirInfo(estudiante InfoEstudiante) {
	estudiante.InformacionEstudiante()
}

func main() {
	estudianteEjemplo := NuevoEstudiante("Juan", "Trujillo", "juank-1521-121@hotmail.com", 3.5)
	ImprimirInfo(estudianteEjemplo)
}
