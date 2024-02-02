package main

import (
	"fmt"
)

func (t Tarea) Equals(other Tarea) bool {
	return true
}

type ListTareas struct {
	ListaTareas []Tarea
	head        *NodoTarea
	tail        *NodoTarea
	size        int
}

type NodoTarea struct {
	next *NodoTarea
	data Tarea
}

func NewTareaList() *ListTareas {
	return &ListTareas{ListaTareas: make([]Tarea, 0), head: nil, tail: nil, size: 0}
}
func newNode(data Tarea) *NodoTarea {
	return &NodoTarea{data: data, next: nil}
}

func (l *ListTareas) InsertAt(value Tarea, position int) {
	if position < 0 || position > l.size {
		return
	}
	newNode := newNode(value)
	// Insertar al principio
	// O(1)
	if position == 0 {
		newNode.next = l.head
		l.head = newNode
		if l.tail == nil {
			l.tail = newNode
		}
		l.size++
		return
	}
	// Insertar al final
	// O(1)
	if position == l.size {
		l.tail.next = newNode
		l.tail = newNode
		l.size++
		return
	}
	// Insertar en position
	current := l.head
	for current != nil && position > 1 {
		current = current.next
		position--
	}

	newNode.next = current.next
	current.next = newNode
	l.size++
}

// RemoveAt elimina el nodo en la posición recibida
// Si la posición es inválida, no hace nada
// La posición 0 elimina el primer nodo de la lista O(1)
// Eliminar en otra posición. O(n)
func (l *ListTareas) RemoveAt(position int) {
	if position < 0 || position >= l.size {
		return
	}
	// Eliminar el primer nodo
	// O(1)
	if position == 0 {
		l.head = l.head.next
		l.size--
		return
	}

	current := l.head
	for current != nil && position > 1 {
		current = current.next
		position--
	}

	current.next = current.next.next
	l.size--
}

// Search busca el primer nodo que contenga el valor recibido
// y devuelve su posición en la lista o -1 si no lo encuentra
// O(n)
func (l *ListTareas) Search(value Tarea) int {
	if l.head == nil {
		return -1
	}
	current := l.head
	position := 0
	for current != nil {
		if current.data.Nombre == value.Nombre {
			return position
		}
		current = current.next
		position++
	}
	return -1
}

// Get devuelve el valor del nodo en la posición recibida
// o un valor nulo si la posición es inválida
// O(n)
func (l *ListTareas) Get(position int) Tarea {
	if position < 0 || position >= l.size {
		var t Tarea
		return t
	}

	current := l.head
	for current != nil && position > 0 {
		current = current.next
		position--
	}

	return current.data
}

// Size devuelve la cantidad de nodos en la lista
// O(n)
func (l *ListTareas) Size() int {
	return l.size
}

// Contains verifica si el valor recibido está en la lista
// O(n)
func (l *ListTareas) Contains(value Tarea) bool {
	return l.Search(value) != -1
}

// IsEmpty verifica si la lista está vacía
// O(1)
func (l *ListTareas) IsEmpty() bool {
	return l.head == nil
}

// String devuelve una representación en cadena de la lista
// en el formato [1 2 3].
// Se puede usar para imprimir la lista con fmt.Println
// O(n)
func (l *ListTareas) String() string {
	if l.head == nil {
		return "[]"
	}
	current := l.head
	result := "lista: ["
	for current != nil {
		result += fmt.Sprintf("%v", current.data)
		if current.next != nil {
			result += " "
		}
		current = current.next
	}
	result += "]"
	return result
}

func (l *ListTareas) Insertar(posicion int, valor Tarea) {
	nodo := newNode(valor)
	if posicion < 0 || posicion > l.size {
		return
	}

	//posicion 0
	if posicion == 0 {
		nodo.next = l.head
		l.head = nodo
		if l.tail == nil {
			l.tail = nodo
		}
	}

	//posicion tail
	if posicion == l.size {
		l.tail.next = nodo
		l.tail = nodo
		l.size++
		return
	}

	//posicion x
	actual := l.head
	for actual != nil && posicion > 1 {
		actual = actual.next
		posicion--
	}
	nodo.next = actual.next
	actual.next = nodo
	l.size++
}
