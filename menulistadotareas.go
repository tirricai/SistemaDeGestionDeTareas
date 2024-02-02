package main

import (
	"fmt"
	"strings"
)

// =============================================================================================================================
// ====================================================[MENU LISTADO DE TAREAS]=================================================
// =============================================================================================================================
func menuDeListadoDeTareas(tm *TareaAdmin) {
	var indice int
	var option string
	fmt.Println("=========================================================")
	fmt.Println("Elija una Tarea : ")
	fmt.Scanln(&indice)
	tm.TareaCurrent = &tm.Tareas[indice]
	clear()
	fmt.Println("=================[TAREA SELECCIONADA]====================")
	tm.ListarTarea()
	for {
		fmt.Println("=========================================================")
		fmt.Println("     a. Editar nombre")
		fmt.Println("     b. Editar duracion")
		fmt.Println("     c. Subir prioridad")
		fmt.Println("     d. Bajar prioridad")
		fmt.Println("     e. Editar etiquetas")
		fmt.Println("     f. Comenzar tarea")
		fmt.Println("     g. Nueva subtarea")
		fmt.Println("     h. Borrar tarea")
		fmt.Println("     i. Marcar subtarea como completada")
		fmt.Println("     j. Volver al menu principal")
		fmt.Println("Ingrese una opcion correcta: ")

		fmt.Scanln(&option)
		switch option {
		case "a":
			fmt.Println("=========================================================")
			fmt.Print("Nuevo nombre: ")
			nombreTarea := readLine()
			fmt.Println("=========================================================")
			tm.EditarNombreTarea(indice, nombreTarea)
		case "b":
			var duration float64
			fmt.Println("=========================================================")
			fmt.Print("Nueva duracion: ")
			fmt.Scanln(&duration)
			fmt.Println("=========================================================")
			tm.EditarDuracionTarea(indice, duration)
		case "c":
			tm.IncrementarPrioridadTarea(indice)
		case "d":
			tm.DecrementarPrioridadTarea(indice)
		case "e":
			fmt.Println("=========================================================")
			fmt.Print("Nuevas etiquetas (separadas por comas): ")
			tagsNuevos := readLine()
			fmt.Println("=========================================================")
			tags := strings.Split(tagsNuevos, ",")
			tm.EditarTagTarea(indice, tags)
		case "f":
			tm.ComenzarTarea()
			fmt.Println("La tarea ha cambiado su estado a: EN CURSO - Cualquier otra tarea anteriormente en curso, esta interrumpida")
		case "g":
			tm.AnadirSubtarea()
		case "h":
			tm.BorrarTarea(indice)
			tm.TareaCurrent = nil
		case "i":
			seleccionarSubtarea(tm.TareaCurrent)
		case "j":
			clear()
			fmt.Println("=========================================================")
			fmt.Println("Volviendo al menú principal.")
			fmt.Println("=========================================================")
		default:
			continue
		}
		return
	}
}
