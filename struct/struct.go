package main

import "fmt"

type Nota struct {
	score float64
}

type Estudiante struct {
	Nombre   string
	Apellido string
	Email    string
	Nota     Nota
}

// Función que imprime información del estudiante
func InformacionEstudiante(alumno Estudiante) {
	fmt.Printf("Nombre: %s\n", alumno.Nombre)
	fmt.Printf("Apellido: %s\n", alumno.Apellido)
	fmt.Printf("Email: %s\n", alumno.Email)
	fmt.Printf("Nota: %d\n", alumno.Nota.score) //  %d para imprimir el entero

	if alumno.Nota.score > 3.5 {
		fmt.Println("Estado: Aprobado")
	} else {
		fmt.Println("Estado: No aprobado")
	}
	fmt.Println("--------------")
}

func main() {
	estudianteEjemplo := Estudiante{
		Nombre:   "Juan",
		Apellido: "Trujillo",
		Email:    "juank-1521-121@hotmail.com",
		Nota:     Nota{score: 3.6},
	}

	// Imprimir información
	InformacionEstudiante(estudianteEjemplo)
}
