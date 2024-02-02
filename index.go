package main

import (
	"fmt"
	"strings"
)

// ==============================================================================================================================
// =============================================================================================================================
// ============================================================[MENU PRINCIPAL]=================================================
// =============================================================================================================================
func main() {
	TareaAdmin := NuevaTareaAdmin()
	fmt.Println("============================================================================================================================")
	fmt.Println("====================================[GRUPO 'PizzA': ALGORITMOS Y PROGRAMACION II]===========================================")
	fmt.Println("============================================[SISTEMA DE GESTION DE TAREAS]==================================================")
	fmt.Println("============================================================================================================================")
	fmt.Println()
	for {
		fmt.Println("================TAREA ACTUAL======================")
		if TareaAdmin.TareaCurrent != nil {
			fmt.Println(TareaAdmin.TareaCurrent)
		} else {
			fmt.Println("No hay tareas que mostrar")
		}
		fmt.Println("==================================================")
		fmt.Println()

		fmt.Println("================TAREAS INTERRUMPIDAS==============")
		MostrarTareasInterrumpidas(TareaAdmin.Tareas)
		fmt.Println("==================================================")
		fmt.Println()

		fmt.Println("=================[MENU]===========================")
		fmt.Println("1. Listado de tareas.")
		fmt.Println("2. Agregar nueva tarea.")
		fmt.Println("3. Ver tarea actual.")
		fmt.Println("4. Optimizar mi tiempo.")
		fmt.Println("5. Buscar tarea.")
		fmt.Println("6. Reordenar tareas.")
		fmt.Println("0. Salir.")
		fmt.Print("Selecciona una opción: ")

		var option int
		fmt.Scanln(&option)

		switch option {
		case 0:
			fmt.Println("=====================================")
			fmt.Println("¡Hasta luego!")
			fmt.Println("=====================================")
			despedida()
			return
		case 1:
			clear()
			TareaAdmin.ListasDeTareas()
		case 2:
			fmt.Println("=========================================================")
			fmt.Print("Nombre de la nueva tarea: ")
			nombreTarea := readLine()
			fmt.Print("Duración de la nueva tarea: ")
			var duracion float64
			fmt.Scanln(&duracion)
			fmt.Print("Etiquetas de la nueva tarea (separadas por comas): ")
			tag := readLine()
			tags := strings.Split(tag, ",")
			fmt.Println("=========================================================")
			tarea := Tarea{
				Nombre:    nombreTarea,
				Duracion:  duracion,
				Prioridad: "baja",
				Subtareas: []Tarea{},
				Estado:    "pendiente",
				Etiquetas: tags,
			}
			TareaAdmin.AñadirTarea(tarea)
			clear()
		case 3:
			clear()
			TareaAdmin.TareaActual()
			TareaAdmin.MiniMenuTareaActual()
		case 4:
			clear()
			fmt.Println("=========================================================")
			fmt.Println("Ingrese la cantidad de tiempo disponible para realizar tareas:")
			var tiempo float64
			fmt.Scanln(&tiempo)
			TareaAdmin.PrepararColaTareasPorTiempo(tiempo)
			fmt.Println()
		case 5:
			clear()
			var buscar int
			fmt.Println("==================[Elije una opcion]=====================")
			fmt.Println("     1: Buscar por palabra clave.")
			fmt.Println("     2: Buscar por palabra tag.")
			fmt.Scanln(&buscar)
			fmt.Println("=========================================================")
			BuscarPorClaveTag(buscar, TareaAdmin)
		case 6:
			clear()
			fmt.Println("=========================================================")
			fmt.Println("Criterio de ordenamiento: 1. Prioridad, 2.Duracion, 3.Subtareas")
			var criterio int
			fmt.Scanln(&criterio)
			fmt.Println("=========================================================")
			TareaAdmin.ReordenarTareas(criterio)
			fmt.Println()
			clear()
		default:
			fmt.Println("=========================================================")
			fmt.Println("Opción inválida.")
			fmt.Println("=========================================================")
			fmt.Println()
		}
	}
}
