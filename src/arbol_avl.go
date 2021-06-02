package main

import (
	"fmt"
	"strconv"
)

// COMP es una variable global para preservar la cantidad de comparaciones a
// través de las llamadas recursivas realizadas a la hora de insertar o buscar
// nodos dentro del árbol AVL.
var COMP int = 0
var PROF_SUM int = 0

type arbol_avl struct {
	raiz *nodo
	size int
}

// Constructor del árbol AVL.
func crear_avl(p_llave int) *arbol_avl {
	fmt.Println("Se ha creado un nuevo árbol AVL, tiene una raíz con llave: " + strconv.Itoa(p_llave))
	return &arbol_avl{
		raiz: crear_nodo(p_llave, 1),
		size: 1,
	}
}

func (arbol *arbol_avl) insertar(p_llave int, verbose bool) int {
	// Reiniciar el valor de COMP.
	COMP = 0

	arbol.size += 1
	arbol.raiz = insertar_aux(arbol.raiz, p_llave, 1)

	return COMP
}

func (arbol *arbol_avl) buscar(p_llave int, verbose bool) (bool, int) {
	// Reiniciar el valor de COMP.
	COMP = 0
	result := buscar_aux(arbol.raiz, p_llave)
	if result == nil {
		if verbose {
			println("Nodo no encontrado")
		}
		return false, COMP
	}
	if verbose {
		println("Nodo encontrado:\n" + result.to_string())
	}
	return true, COMP
}

func (arbol *arbol_avl) remover() {

}

func (arbol *arbol_avl) indorden(verbose bool) {
	if verbose {
		fmt.Println("Imprimiendo el árbol AVL en preorden: ")
	}
	inorden_aux(arbol.raiz, verbose)
}

func (arbol *arbol_avl) preorden(verbose bool) {
	if verbose {
		fmt.Println("Imprimiendo el árbol AVL en preorden: ")
	}
	preorden_aux(arbol.raiz, verbose)
}

func (arbol *arbol_avl) get_profundidad_sumada() int {
	PROF_SUM = 0
	get_profundidad_sumada_aux(arbol.raiz, 0)
	return PROF_SUM
}

func rotacion_derecha(cab *nodo) *nodo {
	cab_nueva := cab.izq
	cab.izq = cab_nueva.der
	cab_nueva.der = cab
	cab.altura = 1 + max(get_altura(cab.izq), get_altura(cab.der))
	cab_nueva.altura = 1 + max(get_altura(cab_nueva.izq), get_altura(cab_nueva.der))
	return cab_nueva
}

func rotacion_izquierda(cab *nodo) *nodo {
	cab_nueva := cab.der
	cab.der = cab_nueva.izq
	cab_nueva.izq = cab
	cab.altura = 1 + max(get_altura(cab.izq), get_altura(cab.der))
	cab_nueva.altura = 1 + max(get_altura(cab_nueva.izq), get_altura(cab_nueva.der))
	return cab_nueva
}

func get_profundidad_sumada_aux(cabeza *nodo, prof int) {
	if cabeza == nil {
		return
	}

	PROF_SUM += prof
	prof += 1

	get_profundidad_sumada_aux(cabeza.izq, prof)
	get_profundidad_sumada_aux(cabeza.der, prof)
}

func inorden_aux(cabeza *nodo, verbose bool) {
	if cabeza == nil {
		return
	}

	inorden_aux(cabeza.izq, verbose)
	if verbose {
		fmt.Println(cabeza.to_string())
	} else {
		fmt.Println(cabeza.llave)
	}
	inorden_aux(cabeza.der, verbose)
}

func preorden_aux(cabeza *nodo, verbose bool) {
	if cabeza == nil {
		return
	}
	if verbose {
		fmt.Println(cabeza.to_string())
	} else {
		fmt.Println(cabeza.llave)
	}
	preorden_aux(cabeza.izq, verbose)
	preorden_aux(cabeza.der, verbose)
}

func insertar_aux(cabeza *nodo, p_llave int, p_valor int) *nodo {
	if cabeza == nil {
		COMP += 1
		temp := crear_nodo(p_llave, p_valor)
		return temp
	}

	COMP += 1
	// Una comparación de llaves.
	if p_llave < cabeza.llave {
		cabeza.izq = insertar_aux(cabeza.izq, p_llave, p_valor)
	} else if p_llave > cabeza.llave {
		cabeza.der = insertar_aux(cabeza.der, p_llave, p_valor)
	} else {
		COMP += 1
		cabeza.valor += 1
	}

	cabeza.altura = 1 + max(get_altura(cabeza.izq), get_altura(cabeza.der))
	fac_bal := get_factor_balance(cabeza)

	if fac_bal > 1 {
		if p_llave < cabeza.izq.llave {
			return rotacion_derecha(cabeza)
		} else {
			cabeza.izq = rotacion_izquierda(cabeza.izq)
			return rotacion_derecha(cabeza)
		}
	} else if fac_bal < -1 {
		if p_llave > cabeza.der.llave {
			return rotacion_izquierda(cabeza)
		} else {
			cabeza.der = rotacion_derecha(cabeza.der)
			return rotacion_izquierda(cabeza)
		}
	}
	return cabeza
}

func buscar_aux(cabeza *nodo, p_llave int) *nodo {
	//No se encuentra el nodo.
	if cabeza == nil {
		COMP += 1
		return nil
	}
	llave := cabeza.llave

	//Se encuentra el nodo.
	if llave == p_llave {
		COMP += 1
		return cabeza
	}
	COMP += 1
	//Comparación de llaves.
	if llave > p_llave {
		return buscar_aux(cabeza.izq, p_llave)
	} else {
		return buscar_aux(cabeza.der, p_llave)
	}
}
