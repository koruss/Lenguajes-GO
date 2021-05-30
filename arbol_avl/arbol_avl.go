package main

import (
	"fmt"
	"strconv"
)

// COMP es una variable global para preservar la cantidad de comparaciones a
// través de las llamadas recursivas realizadas a la hora de insertar o buscar
// nodos dentro del árbol AVL.
var COMP int = 0

type arbol_avl struct {
	raiz *nodo
}

// Constructor del árbol AVL.
func crear_avl(p_llave int) *arbol_avl {
	fmt.Println("Se ha creado un nuevo árbol AVL, tiene una raíz con llave: " + strconv.Itoa(p_llave))
	return &arbol_avl{
		raiz: crear_nodo(p_llave, 1),
	}
}

func (arbol *arbol_avl) insertar(p_llave int) int {
	// Reiniciar el valor de COMP.
	COMP = 0

	// Por la naturaleza de la función de inserción, si se inserta una llave
	// repetida, la función la omite. Por esto es que primero se realiza una
	// búsqueda para ver si el nodo ya existe en el árbol. Si sí existe,
	// entonces se le incrementa el valor al nodo encontrado y no se inserta
	// nada.

	result := buscar_aux(arbol.raiz, p_llave)

	if result == nil {
		arbol.raiz = insertar_aux(arbol.raiz, p_llave, 1)
		fmt.Println("Se ha ingresado un nodo nuevo con llave: " + strconv.Itoa(p_llave))
	} else {
		result.valor += 1
		fmt.Println("Se ha ingresado una llave que ya existe, se ha incrementado el valor del nodo.")
	}
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

func (arbol *arbol_avl) indorden() {
	inorden_aux(arbol.raiz)
}

func (arbol *arbol_avl) preorden(verbose bool) {
	if verbose {
		fmt.Println("Imprimiendo el árbol AVL en preorden: ")
	}
	preorden_aux(arbol.raiz, verbose)
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

func inorden_aux(cabeza *nodo) {
	if cabeza == nil {
		return
	}

	inorden_aux(cabeza.izq)
	fmt.Println(cabeza.to_string())
	inorden_aux(cabeza.der)
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
	COMP += 1
	// Una comparación por el primer if.
	if cabeza == nil {
		temp := crear_nodo(p_llave, p_valor)
		return temp
	}
	COMP += 1
	// Una comparación por el segundo if.
	if p_llave < cabeza.llave {
		cabeza.izq = insertar_aux(cabeza.izq, p_llave, p_valor)
	} else if p_llave > cabeza.llave {
		COMP += 1
		// Una comparación por el else if.
		cabeza.der = insertar_aux(cabeza.der, p_llave, p_valor)
	}

	cabeza.altura = 1 + max(get_altura(cabeza.izq), get_altura(cabeza.der))
	fac_bal := get_factor_balance(cabeza)

	COMP += 1
	// Una comparación por el tercer if.
	if fac_bal > 1 {
		COMP += 1
		// Una comparación por el if anidado.
		if p_llave < cabeza.izq.llave {
			return rotacion_derecha(cabeza)
		} else {
			cabeza.izq = rotacion_izquierda(cabeza.izq)
			return rotacion_derecha(cabeza)
		}
	} else if fac_bal < -1 {
		COMP += 1
		// Una comparación por el segundo else if.
		if p_llave > cabeza.der.llave {
			COMP += 1
			// Una comparación por el if anidado.
			return rotacion_izquierda(cabeza)
		} else {
			cabeza.der = rotacion_derecha(cabeza.der)
			return rotacion_izquierda(cabeza)
		}
	}
	return cabeza
}

func buscar_aux(cabeza *nodo, p_llave int) *nodo {
	COMP += 1
	//Comparación por el primer if
	if cabeza == nil {
		return nil
	}
	llave := cabeza.llave
	COMP += 1
	//Comparación por el segundo if
	if llave == p_llave {
		return cabeza
	}
	COMP += 1
	//Comparación por el tercer if
	if llave > p_llave {
		return buscar_aux(cabeza.izq, p_llave)
	} else if llave < p_llave {
		COMP += 1
		//Comparación por el else if anidado.
		return buscar_aux(cabeza.der, p_llave)
	} else {
		COMP += 1
		//Comparación por el else if anidado.
		return nil
	}
}

func main() {
	tree := crear_avl(1)

	// Código para demostrar los requerimientos 6 y 8 de la especificación.
	fmt.Printf("\nInserciones al árbol AVL: \n")

	var comp int
	comp = tree.insertar(7)
	println("Cantidad de comparaciones realizadas: " + strconv.Itoa(comp))

	comp = tree.insertar(4)
	println("Cantidad de comparaciones realizadas: " + strconv.Itoa(comp))

	comp = tree.insertar(8)
	println("Cantidad de comparaciones realizadas: " + strconv.Itoa(comp))

	comp = tree.insertar(3)
	println("Cantidad de comparaciones realizadas: " + strconv.Itoa(comp))

	comp = tree.insertar(7)
	println("Cantidad de comparaciones realizadas: " + strconv.Itoa(comp))

	// Código para demostrar el requerimiento 7 de la especificación.
	fmt.Printf("\nBúsqueda de una llave: \n")

	var llave_busq = 2
	var result, cmp = tree.buscar(llave_busq, false)

	if result {
		fmt.Println("Se logró encontrar el nodo con llave: " + strconv.Itoa(llave_busq))
	} else {
		fmt.Println("No se logró encontrar el nodo con llave: " + strconv.Itoa(llave_busq))
	}
	fmt.Println("Número de comparaciones: " + strconv.Itoa(cmp))

	// Imprimir el árbol en preorden para verificar su estructura.

	fmt.Printf("\nEstructura del árbol: \n")
	tree.preorden(true)
}
