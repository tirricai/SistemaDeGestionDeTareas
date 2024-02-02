package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type TareaAdmin struct {
	Tareas       []Tarea
	TareaCurrent *Tarea
	Scanner      *bufio.Scanner
	Comparator   func(t1, t2 *Tarea) bool
	PilaTareas   []Tarea
}

// Métodos elementales de una pila
func (ta *TareaAdmin) Push(t Tarea) {
	ta.PilaTareas = append(ta.PilaTareas, t)
}

func (ta *TareaAdmin) Pop() *Tarea {
	if ta.IsEmpty() {
		return nil
	}
	index := len(ta.PilaTareas) - 1
	popped := ta.PilaTareas[index]
	ta.PilaTareas = ta.PilaTareas[:index]
	return &popped
}

func (ta *TareaAdmin) IsEmpty() bool {
	return len(ta.PilaTareas) == 0
}

// Función para imprimir la pila de tareas
func (ta *TareaAdmin) ImprimirPilaTareas() {
	fmt.Println("Pila de Tareas:")
	for _, tarea := range ta.PilaTareas {
		fmt.Printf("Prioridad: %s, Duracion: %d\n", tarea.Prioridad, tarea.Duracion)
	}
}

// =================================================================================================================
// ====================================================FUNCIONES===================================================
// =================================================================================================================
func NuevaTareaAdmin() *TareaAdmin {
	scanner := bufio.NewScanner(os.Stdin)
	return &TareaAdmin{
		Tareas:       []Tarea{},
		TareaCurrent: nil,
		Scanner:      scanner,
		Comparator: func(t1, t2 *Tarea) bool {
			if t1.Prioridad == "alta" && t2.Prioridad == "baja" {
				return true
			}
			if t1.Prioridad == "baja" && t2.Prioridad == "alta" {
				return false
			}
			return t1.Duracion < t2.Duracion
		},
	}
}

func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func despedida() {
	fmt.Println("============================================================================================================================")
	fmt.Println("===========================================================================================================================")
	fmt.Println("============================================================================================================================")
	fmt.Println("====================================================ADIOS===================================================================")
	fmt.Println("============================================================================================================================")
	fmt.Println("============================================================================================================================")
	fmt.Println("============================================================================================================================")
}

func BuscarPorClaveTag(buscar int, TareaAdmin *TareaAdmin) {
	switch buscar {
	case 1:
		fmt.Print("Palabra clave: ")
		palabraClave := readLine()
		TareaAdmin.BuscarTareaPorClave(palabraClave)
	case 2:
		fmt.Print("Ingrese Tag: ")
		Tag := readLine()
		TareaAdmin.BuscarTareaPorTag(Tag)
	default:
		fmt.Println("=========================================================")
		fmt.Println("Opción inválida.")
		fmt.Println("=========================================================")
	}
}

func (tm *TareaAdmin) ListasDeTareas() {
	if len(tm.Tareas) == 0 {
		fmt.Println("=========================================================")
		fmt.Println("No hay tareas disponibles.")
		fmt.Println("=========================================================")
	} else {
		fmt.Println("=========================================================")
		fmt.Println("Lista de tareas:")
		for i, tarea := range tm.Tareas {
			fmt.Printf("%d - %s\n", i, tarea.String())
		}
		fmt.Println("=========================================================")
		menuDeListadoDeTareas(tm)
	}
}

func (tm *TareaAdmin) TareaActual() {
	if tm.TareaCurrent != nil {
		fmt.Println("=========================================================")
		fmt.Println("Tarea actual:")
		fmt.Println("   Nombre:", tm.TareaCurrent.Nombre)
		fmt.Printf("   Duración: %.2f horas\n", tm.TareaCurrent.Duracion)
		fmt.Printf("   Prioridad: %s\n", tm.TareaCurrent.Prioridad)
		fmt.Printf("   Estado: %s\n", tm.TareaCurrent.Estado)
		fmt.Println("=========================================================")
	} else {
		fmt.Println("=========================================================")
		fmt.Println("No hay tarea en proceso.")
		fmt.Println("=========================================================")
	}
}

func (tm *TareaAdmin) PrepararColaTareasPorTiempo(tiempoDisponible float64) {
	if len(tm.Tareas) == 0 {
		fmt.Println("=========================================================")
		fmt.Println("No hay tareas disponibles para preparar la cola.")
		fmt.Println("=========================================================")
		return
	}

	cola := make([]*Tarea, len(tm.Tareas))
	for i := range tm.Tareas {
		cola[i] = &tm.Tareas[i]
	}

	for i := 0; i < len(cola)-1; i++ {
		for j := i + 1; j < len(cola); j++ {
			if cola[j].Prioridad == "alta" && cola[i].Prioridad != "alta" {
				cola[i], cola[j] = cola[j], cola[i]
			}
		}
	}

	fmt.Println("=========================================================")
	fmt.Println("Tareas que puedes realizar en", tiempoDisponible, "horas:")
	tareasRealizables := []*Tarea{}
	tiempoRestante := tiempoDisponible

	for _, tarea := range cola {
		if tarea.Duracion <= tiempoRestante {
			tareasRealizables = append(tareasRealizables, tarea)
			tiempoRestante -= tarea.Duracion
		}
	}

	for i, tarea := range tareasRealizables {
		fmt.Printf("%d: %s\n", i, tarea.Nombre)
		fmt.Printf("   Duración: %.2f horas\n", tarea.Duracion)
		fmt.Printf("   Prioridad: %s\n", tarea.Prioridad)
		fmt.Printf("   Estado: %s\n", tarea.Estado)
	}
	fmt.Println("=========================================================")
}

func (tm *TareaAdmin) BuscarTareaPorClave(palabraClave string) {
	tareaEncontradas := []Tarea{}

	for _, tarea := range tm.Tareas {
		if strings.Contains(strings.ToLower(tarea.Nombre), strings.ToLower(palabraClave)) {
			tareaEncontradas = append(tareaEncontradas, tarea)
		}
	}

	if len(tareaEncontradas) == 0 {
		fmt.Println("=========================================================")
		fmt.Println("No se encontraron tareas con la palabra clave proporcionada.")
		fmt.Println("=========================================================")
		return
	}
	fmt.Println("=========================================================")
	fmt.Printf("Tareas encontradas con la palabra clave '%s':\n", palabraClave) //o
	for i, tarea := range tareaEncontradas {
		fmt.Printf("%d: %s\n", i, tarea.Nombre)
		fmt.Printf("   Duración: %.2f horas\n", tarea.Duracion)
		fmt.Printf("   Prioridad: %s\n", tarea.Prioridad)
		fmt.Printf("   Estado: %s\n", tarea.Estado)
	}
	fmt.Println("=========================================================")
}

func (tm *TareaAdmin) BuscarTareaPorTag(tag string) {
	tareaEncontradas := []Tarea{}

	for _, tarea := range tm.Tareas {
		for _, tareaTag := range tarea.Etiquetas {
			if strings.Contains(strings.ToLower(tareaTag), strings.ToLower(tag)) {
				tareaEncontradas = append(tareaEncontradas, tarea)
				break
			}
		}
	}

	if len(tareaEncontradas) == 0 {
		fmt.Println("=========================================================")
		fmt.Println("No se encontraron tareas con la etiqueta proporcionada.")
		fmt.Println("=========================================================")
		return
	}
	fmt.Println("=========================================================")
	fmt.Printf("Tareas encontradas con la etiqueta '%s':\n", tag)
	for i, tarea := range tareaEncontradas {
		fmt.Printf("%d: %s\n", i, tarea.Nombre)
		fmt.Printf("   Duración: %.2f horas\n", tarea.Duracion)
		fmt.Printf("   Prioridad: %s\n", tarea.Prioridad)
		fmt.Printf("   Estado: %s\n", tarea.Estado)
	}
	fmt.Println("=========================================================")
}

func (tm *TareaAdmin) ReordenarTareas(criterio int) {
	switch criterio {
	case 1:
		tm.Comparator = func(t1, t2 *Tarea) bool {
			if t1.Prioridad == "alta" && t2.Prioridad != "alta" {
				return true
			}
			if t1.Prioridad != "alta" && t2.Prioridad == "alta" {
				return false
			}
			return t1.Duracion < t2.Duracion
		}
	case 2:
		tm.Comparator = func(t1, t2 *Tarea) bool {
			return t1.Duracion > t2.Duracion // Ordenar de mayor a menor duración
		}
	case 3:
		tm.Comparator = func(t1, t2 *Tarea) bool {
			if len(t1.Subtareas) > len(t2.Subtareas) {
				return true
			}
			if len(t1.Subtareas) < len(t2.Subtareas) {
				return false
			}
			return t1.Duracion < t2.Duracion
		}
	default:
		fmt.Println("=========================================================")
		fmt.Println("Criterio de ordenamiento inválido.")
		fmt.Println("=========================================================")
		return
	}

	// Ordenar las tareas utilizando el criterio de comparación
	for i := 0; i < len(tm.Tareas)-1; i++ {
		for j := i + 1; j < len(tm.Tareas); j++ {
			if tm.Comparator(&tm.Tareas[j], &tm.Tareas[i]) {
				tm.Tareas[i], tm.Tareas[j] = tm.Tareas[j], tm.Tareas[i]
			}
		}
	}

	criterioStr := ""
	switch criterio {
	case 1:
		criterioStr = "prioridad"
	case 2:
		criterioStr = "duración"
	case 3:
		criterioStr = "subtareas"
	}

	fmt.Println("=========================================================")
	fmt.Printf("La lista de tareas ha sido reordenada por %s.\n", criterioStr)
	fmt.Println("=========================================================")
}

func (tm *TareaAdmin) AñadirTarea(tarea Tarea) {
	tm.Tareas = append(tm.Tareas, tarea)
	fmt.Println("=========================================================")
	fmt.Println("La tarea se ha añadido con éxito.")
	fmt.Println("=========================================================")
}

func (tm *TareaAdmin) EditarNombreTarea(indice int, nuevoNombre string) {
	clear()
	if indice >= 0 && indice < len(tm.Tareas) {
		tm.Tareas[indice].Nombre = nuevoNombre
		fmt.Println("=========================================================")
		fmt.Println("El nombre de la tarea se ha actualizado.")
		fmt.Println("=========================================================")
	} else {
		fmt.Println("=========================================================")
		fmt.Println("Índice de tarea inválido.")
		fmt.Println("=========================================================")
	}
}

func (tm *TareaAdmin) EditarDuracionTarea(indice int, nuevaDuracion float64) {
	clear()
	if indice >= 0 && indice < len(tm.Tareas) {
		tm.Tareas[indice].Duracion = nuevaDuracion
		fmt.Println("=========================================================")
		fmt.Println("La duración de la tarea se ha actualizado.")
		fmt.Println("=========================================================")
	} else {
		fmt.Println("=========================================================")
		fmt.Println("Índice de tarea inválido.")
		fmt.Println("=========================================================")
	}
}

func (tm *TareaAdmin) IncrementarPrioridadTarea(indice int) {
	if indice >= 0 && indice < len(tm.Tareas) {
		clear()
		if tm.Tareas[indice].Prioridad == "baja" {
			tm.Tareas[indice].Prioridad = "alta"
			fmt.Println("=========================================================")
			fmt.Println("La prioridad de la tarea se ha incrementado.")
			fmt.Println("=========================================================")
		} else {
			fmt.Println("=========================================================")
			fmt.Println("La tarea ya tiene la prioridad más alta.")
			fmt.Println("=========================================================")
		}
	} else {
		fmt.Println("=========================================================")
		fmt.Println("Índice de tarea inválido.")
		fmt.Println("=========================================================")
	}
}

func (tm *TareaAdmin) DecrementarPrioridadTarea(indice int) {
	clear()
	if indice >= 0 && indice < len(tm.Tareas) {
		if tm.Tareas[indice].Prioridad == "alta" {
			tm.Tareas[indice].Prioridad = "baja"
			fmt.Println("=========================================================")
			fmt.Println("La prioridad de la tarea se ha disminuido.")
			fmt.Println("=========================================================")
		} else {
			fmt.Println("=========================================================")
			fmt.Println("La tarea ya tiene la prioridad más baja.")
			fmt.Println("=========================================================")
		}
	} else {
		fmt.Println("=========================================================")
		fmt.Println("Índice de tarea inválido.")
		fmt.Println("=========================================================")
	}
}

func (tm *TareaAdmin) EditarTagTarea(indice int, nuevasEtiquetas []string) {
	clear()
	if indice >= 0 && indice < len(tm.Tareas) {
		tm.Tareas[indice].Etiquetas = append(tm.Tareas[indice].Etiquetas, nuevasEtiquetas...)
		fmt.Println("=========================================================")
		fmt.Println("Las etiquetas de la tarea se han actualizado.")
		fmt.Println("=========================================================")
	} else {
		fmt.Println("=========================================================")
		fmt.Println("Índice de tarea inválido.")
		fmt.Println("=========================================================")
	}
}

func (tm *TareaAdmin) ComenzarTarea() {
	tm.PilaTareas.Push(tm.TareaCurrent)
}

func (tm *TareaAdmin) BorrarTarea(indice int) {
	clear()
	if indice >= 0 && indice < len(tm.Tareas) {
		tm.Tareas = append(tm.Tareas[:indice], tm.Tareas[indice+1:]...)
		fmt.Println("=========================================================")
		fmt.Println("La tarea se ha eliminado.")
		fmt.Println("=========================================================")
	} else {
		fmt.Println("=========================================================")
		fmt.Println("Índice de tarea inválido.")
		fmt.Println("=========================================================")
	}
}

func (tm *TareaAdmin) AnadirSubtarea() {
	clear()
	fmt.Println("=========================================================")
	fmt.Print("Nombre de la nueva subtarea: ")
	nombreTarea := readLine()
	fmt.Print("Duración de la nueva subtarea: ")
	var duracion float64
	fmt.Scanln(&duracion)
	fmt.Print("Etiquetas de la nueva subtarea (separadas por comas): ")
	tag := readLine()
	tags := strings.Split(tag, ",")

	fmt.Println("=========================================================")
	subtarea := Tarea{
		Nombre:    nombreTarea,
		Duracion:  duracion,
		Prioridad: "baja",
		Subtareas: []Tarea{},
		Estado:    "pendiente",
		Etiquetas: tags,
	}
	tm.TareaCurrent.Subtareas = append(tm.TareaCurrent.Subtareas, subtarea)
	tm.calcularDuracionTotal(&subtarea)

}

func (tm *TareaAdmin) calcularDuracionTotal(subtarea *Tarea) {
	tm.TareaCurrent.Duracion += subtarea.Duracion
}

func (tm *TareaAdmin) MiniMenuTareaActual() {
	if tm.TareaCurrent == nil {
		return
	}

	var option int
	for {
		fmt.Println("=========================================================")
		fmt.Println("     1. Marcar como completada.")
		fmt.Println("     2. Abandonar (vuelve a pendiente)")
		fmt.Println("     0. Volver al menú principal")
		fmt.Println("Ingrese una opción correcta: ")
		fmt.Scanln(&option)

		switch option {
		case 0:
			return
		case 1:
			tm.TareaCurrent.Estado = "completada"
			CompletarSubtareas(tm)
			tm.TareaCurrent = nil
			fmt.Println("=========================================================")
			fmt.Println("¡Tarea marcada como completada!")
			fmt.Println("=========================================================")
		case 2:
			tm.TareaCurrent.Estado = "pendiente"
			tm.TareaCurrent = nil
			fmt.Println("=========================================================")
			fmt.Println("¡El estado de la tarea actual volvió a pendiente!")
			fmt.Println("=========================================================")
		default:
			fmt.Println("=========================================================")
			fmt.Println("Opción inválida.")
			fmt.Println("=========================================================")
			continue
		}
		return
	}
}

func (tm *TareaAdmin) ListarTarea() {
	fmt.Println("   Nombre: ", tm.TareaCurrent.Nombre)
	fmt.Printf("   Duración: %.2f horas\n", tm.TareaCurrent.Duracion)
	fmt.Printf("   Prioridad: %s\n", tm.TareaCurrent.Prioridad)
	fmt.Printf("   Estado: %s\n", tm.TareaCurrent.Estado)
	if len(tm.TareaCurrent.Subtareas) != 0 {
		fmt.Println("Subtareas:")
		ListarSubTareas(tm.TareaCurrent.Subtareas)
	}
}
func ListarSubTareas(subTasks []Tarea) {
	for i := 0; i < len(subTasks); i++ {
		fmt.Print("\n")
		fmt.Println("	   Nombre: ", subTasks[i].Nombre)
		fmt.Printf("	   Duración: %.2f horas\n", subTasks[i].Duracion)
		fmt.Printf("	   Prioridad: %s\n", subTasks[i].Prioridad)
		fmt.Printf("	   Estado: %s\n", subTasks[i].Estado)
	}
}

func CompletarSubtareas(tm *TareaAdmin) {
	if len(tm.TareaCurrent.Subtareas) > 0 {
		for i := 0; i < len(tm.TareaCurrent.Subtareas); i++ {
			tm.TareaCurrent.Subtareas[i].Estado = "completada"
		}
		tm.TareaCurrent = nil
	}
}

func (t *Tarea) String() string {
	return fmt.Sprintf("Nombre: %s, Duración: %.2f, Prioridad: %s, Estado: %s, Etiquetas: %s",
		t.Nombre, t.Duracion, t.Prioridad, t.Estado, strings.Join(t.Etiquetas, ", "))
}
func seleccionarSubtarea(t *Tarea) {
	clear()
	if len(t.Subtareas) == 0 {
		fmt.Println("La tarea seleccionada no tiene subtareas disponibles.")
		return
	}

	fmt.Println("Subtareas disponibles:")
	for i, subtarea := range t.Subtareas {
		fmt.Printf("%d. %s\n", i+1, subtarea.Nombre)
	}

	var opcion int
	fmt.Print("Seleccione la subtarea que desea marcar como completada:")
	fmt.Scanln(&opcion)

	if opcion < 1 || opcion > len(t.Subtareas) {
		fmt.Println("Opción inválida.")
		return
	}

	subtareaSeleccionada := &t.Subtareas[opcion-1]

	subtareaSeleccionada.Estado = "completada"

	fmt.Println("La subtarea", subtareaSeleccionada.Nombre, "ha sido marcada como completada.")
	println()
}

func MostrarTareasInterrumpidas(tareas []Tarea) {
	interrumpidas := 0

	fmt.Println("Tareas Interrumpidas:")
	for _, tarea := range tareas {
		if tarea.Estado == "interrumpida" {
			interrumpidas++
			fmt.Printf("- Nombre: %s\n", tarea.Nombre)
			fmt.Printf("  Duración: %.2f\n", tarea.Duracion)
			fmt.Printf("  Prioridad: %s\n", tarea.Prioridad)
			fmt.Printf("  Etiquetas: %s\n", strings.Join(tarea.Etiquetas, ", "))
			fmt.Println()
		}
	}

	fmt.Printf("(Número de tareas interrumpidas: %d)\n", interrumpidas)
}
