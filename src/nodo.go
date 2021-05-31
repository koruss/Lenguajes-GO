package main

import (
	"strconv"
)

type nodo struct {
	altura, llave, valor int
	izq, der             *nodo
}

func crear_nodo(p_llave int, p_valor int) *nodo {
	return &nodo{
		altura: 1,
		llave:  p_llave,
		valor:  p_valor,
		izq:    nil,
		der:    nil,
	}
}

func get_altura(n *nodo) int {
	if n == nil {
		return 0
	}
	return n.altura
}

//Obtener la altura máxima entre dos nodos.
func max(n_izq, n_der int) int {
	if n_izq > n_der {
		return n_izq
	}
	return n_der
}

func get_factor_balance(n *nodo) int {
	//Retorn 0 en caso de no tener nodo adyacente, para poder realizar la resta
	//sin error de tipos.
	if n == nil {
		return 0
	}
	return get_altura(n.izq) - get_altura(n.der)
}

func (n *nodo) to_string() string {
	msg := "Nodo con llave: " + strconv.Itoa(n.llave) + "\n"
	msg += "\tValor: " + strconv.Itoa(n.valor) + "\n"
	msg += "\tAltura: " + strconv.Itoa(n.altura) + "\n"
	msg += "\tLlave del nodo izquierdo:"
	if n.izq != nil {
		msg += strconv.Itoa(n.izq.llave) + "\n"
	} else {
		msg += "-NULO-" + "\n"
	}
	msg += "\tLlave del nodo derecho:"
	if n.der != nil {
		msg += strconv.Itoa(n.der.llave) + "\n"
	} else {
		msg += "-NULO-" + "\n"
	}
	return msg
}
