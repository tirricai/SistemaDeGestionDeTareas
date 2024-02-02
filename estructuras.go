package main

import (
	"fmt"
	"strings"
)

// ======================================[Struct TAREA]==============================================
type Tarea struct {
	Nombre    string
	Duracion  float64
	Prioridad string
	Subtareas []Tarea
	Estado    string
	Etiquetas []string
}

// =================================== Gestionar Tareas (individuales) ==================================
func NuevaTarea(nombre string, duracion float64, tags []string) Tarea {
	var T Tarea
	T.Nombre = nombre
	T.Duracion = duracion
	T.Prioridad = "Baja"
	T.Subtareas = []Tarea{}
	T.Estado = "Pendiente"
	T.Etiquetas = tags
	return T
}
func (t *Tarea) EditarDuracion(nuevaduracion float64) {
	if len(t.Subtareas) < 1 {
		t.Duracion = nuevaduracion
	} else {
		suma := nuevaduracion
		for _, v := range t.Subtareas {
			suma += float64(v.Duracion)
		}
		t.Duracion = suma
	}
}
func (t *Tarea) AnadirSubtarea(nombre string, duracion float64, tags []string) {
	subtarea := NuevaTarea(nombre, duracion, tags)
	t.Subtareas = append(t.Subtareas, subtarea)
	t.Duracion += subtarea.Duracion
}
func (t *Tarea) AgregarTags(tag ...string) {
	if len(tag) > 1 {
		for _, v := range tag {
			t.Etiquetas = append(t.Etiquetas, v)
		}
	} else {
		etiqueta := tag[0]
		t.Etiquetas = append(t.Etiquetas, etiqueta)
	}
}
func (t *Tarea) EliminarTag(tag string) {
	bool := false
	for i, v := range t.Etiquetas {
		if v == tag || strings.ToLower(v) == tag {
			t.Etiquetas = append(t.Etiquetas[:i], t.Etiquetas[i+1:]...)
			bool = true
		}
	}
	if !bool {
		fmt.Println("\t======= No hay coincidencias =======")
	}
}
func (t *Tarea) SubirPrioridad() {
	if t.Prioridad == "Alta" {
		fmt.Println("La tarea ya posee prioridad alta")
	} else {
		t.Prioridad = "Alta"
	}
}
func (t *Tarea) BajarPrioridad() {
	if t.Prioridad == "Baja" {
		fmt.Println("La tarea ya posee prioridad Baja")
	} else {
		t.Prioridad = "Baja"
	}
}
func (t *Tarea) toString() {
	fmt.Println("================================================================================")
	fmt.Print("\tNombre: " + t.Nombre + " | ")
	fmt.Print("Duracion: " + fmt.Sprintf("%.2f", t.Duracion) + " | ")
	fmt.Print("Prioridad: " + t.Prioridad + " | ")
	fmt.Println("Estado: " + t.Estado + " | ")

	fmt.Print("\tTags: [ ")
	for _, v := range t.Etiquetas {
		fmt.Print(" " + v + " ")
	}
	fmt.Println("]")
	if len(t.Subtareas) == 0 {
		fmt.Println("\t======= No hay subtareas disponibles =======")
	} else {
		fmt.Println("\tSubtareas: ")
		for i, subtarea := range t.Subtareas {
			i++
			fmt.Print("\t\t[ " + fmt.Sprintf("%d", i) + " ] ")
			fmt.Print("Nombre: " + subtarea.Nombre + " | ")
			fmt.Print("Duracion: " + fmt.Sprintf("%.2f", subtarea.Duracion) + " | ")
			fmt.Print("Prioridad: " + subtarea.Prioridad + " | ")
			fmt.Println("Estado: " + subtarea.Estado + " | ")
		}
	}
	fmt.Println("================================================================================")
}
func (t *Tarea) Tstring() {
	t.toString()
}
