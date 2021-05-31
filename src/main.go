package main

import (
	"fmt"
	"strconv"
)

/*
	Funciones auxiliares requeridas por la función principal.
*/

var AVL_COMP int = 0
var BST_COMP int = 0

func insert_arr_avl(p_arr []int) *arbol_avl {
	// Reiniciar el valor de AVL_COMP
	AVL_COMP = 0

	var tree *arbol_avl
	first := true

	for _, num := range p_arr {
		if first {
			tree = crear_avl(num)
			first = false
			AVL_COMP += 1
		}
		AVL_COMP += tree.insertar(num, false)
	}
	return tree
}

func insert_arr_bst(p_arr []int) binTree {
	// Reiniciar el valor de BST_COMP
	BST_COMP = 0

	tree := binTree{}
	for _, num := range p_arr {
		BST_COMP += tree.insertNode(num)
	}
	return tree
}

func int_arr_to_string(p_arr []int) string {
	var msg string
	for _, num := range p_arr {
		msg += strconv.Itoa(num) + "\n"
	}
	return msg
}

func buscar_avl(p_arr []int, arbol *arbol_avl) {
	var cmp int
	var res bool
	var fnd string
	for _, num := range p_arr {
		res, cmp = arbol.buscar(num, false)
		if res {
			fnd = "SÍ"
		} else {
			fnd = "NO"
		}
		println("Busqueda en AVL con llave: " + strconv.Itoa(num) + " ,¿Encontrado?: " + fnd + " ,Cantidad de comparaciones: " + strconv.Itoa(cmp))
	}
}

func buscar_bst(p_arr []int, arbol binTree) {
	var fnd string
	var resp response
	for _, num := range p_arr {
		resp = binSearch(arbol.root, num, 0)
		if resp.state {
			fnd = "SÍ"
		} else {
			fnd = "NO"
		}
		println("Busqueda en BST con llave: " + strconv.Itoa(num) + " ,¿Encontrado?: " + fnd + " ,Cantidad de comparaciones: " + strconv.Itoa(resp.comparisons))
	}
}

func buscar_dsw(p_arr []int, arbol binTree) {
	var fnd string
	var resp response
	for _, num := range p_arr {
		resp = binSearch(arbol.root, num, 0)
		if resp.state {
			fnd = "SÍ"
		} else {
			fnd = "NO"
		}
		println("Busqueda en DSW con llave: " + strconv.Itoa(num) + " ,¿Encontrado?: " + fnd + " ,Cantidad de comparaciones: " + strconv.Itoa(resp.comparisons))
	}
}

/*
	Función Principal
*/

func main() {
	// Experimento a)

	a200 := arreglo(37, 30)
	a400 := arreglo(11, 400)
	a600 := arreglo(97, 600)
	a800 := arreglo(17, 800)
	a1000 := arreglo(83, 1000)

	// Experimento b)
	// AVL
	var comparaciones_avl = make([]int, 5)
	var comparaciones_bst = make([]int, 5)
	var comparaciones_dsw = make([]int, 5)

	avl200 := insert_arr_avl(a200)
	comparaciones_avl[0] = AVL_COMP
	avl200.indorden(false)

	avl400 := insert_arr_avl(a400)
	comparaciones_avl[1] = AVL_COMP
	avl400.indorden(false)

	avl600 := insert_arr_avl(a600)
	comparaciones_avl[2] = AVL_COMP
	avl600.indorden(false)

	avl800 := insert_arr_avl(a800)
	comparaciones_avl[3] = AVL_COMP
	avl800.indorden(false)

	avl1000 := insert_arr_avl(a1000)
	comparaciones_avl[4] = AVL_COMP
	avl1000.indorden(false)

	fmt.Println("\nArreglo de comparaciones avl: \n" + int_arr_to_string(comparaciones_avl))

	// BST
	bst200 := insert_arr_bst(a200)
	comparaciones_dsw[0] = BST_COMP
	inorden_bst(bst200.root)

	bst400 := insert_arr_bst(a400)
	comparaciones_dsw[1] = BST_COMP
	inorden_bst(bst400.root)

	bst600 := insert_arr_bst(a600)
	comparaciones_dsw[2] = BST_COMP
	inorden_bst(bst600.root)

	bst800 := insert_arr_bst(a800)
	comparaciones_dsw[3] = BST_COMP
	inorden_bst(bst800.root)

	bst1000 := insert_arr_bst(a1000)
	comparaciones_dsw[4] = BST_COMP
	inorden_bst(bst1000.root)

	fmt.Println("\nArreglo de comparaciones BST: \n" + int_arr_to_string(comparaciones_bst))

	// DSW
	dsw200 := insert_arr_bst(a200)
	comparaciones_bst[0] = BST_COMP
	dsw200.createDSW()
	inorden_bst(dsw200.root)

	dsw400 := insert_arr_bst(a400)
	comparaciones_bst[1] = BST_COMP
	dsw400.createDSW()
	inorden_bst(dsw400.root)

	dsw600 := insert_arr_bst(a600)
	comparaciones_bst[2] = BST_COMP
	dsw600.createDSW()
	inorden_bst(dsw600.root)

	dsw800 := insert_arr_bst(a800)
	comparaciones_bst[3] = BST_COMP
	dsw800.createDSW()
	inorden_bst(dsw800.root)

	dsw1000 := insert_arr_bst(a1000)
	comparaciones_bst[4] = BST_COMP
	dsw1000.createDSW()
	inorden_bst(dsw1000.root)

	fmt.Println("\nArreglo de comparaciones DSW: \n" + int_arr_to_string(comparaciones_dsw))

	// Experimento c)

	a10000 := arreglo(89, 10000)
	fmt.Println("Búsequedas en árboles AVL")
	fmt.Println("\nAVL de 200 elementos.")
	buscar_avl(a10000, avl200)
	fmt.Println("\nAVL de 400 elementos.")
	buscar_avl(a10000, avl400)
	fmt.Println("\nAVL de 600 elementos.")
	buscar_avl(a10000, avl600)
	fmt.Println("\nAVL de 800 elementos.")
	buscar_avl(a10000, avl800)
	fmt.Println("\nAVL de 1000 elementos.")
	buscar_avl(a10000, avl1000)

	fmt.Println("Búsequedas en árboles BST")
	fmt.Println("\nBST de 200 elementos.")
	buscar_bst(a10000, bst200)
	fmt.Println("\nBST de 400 elementos.")
	buscar_bst(a10000, bst400)
	fmt.Println("\nBST de 600 elementos.")
	buscar_bst(a10000, bst600)
	fmt.Println("\nBST de 800 elementos.")
	buscar_bst(a10000, bst800)
	fmt.Println("\nBST de 1000 elementos.")
	buscar_bst(a10000, bst1000)

	fmt.Println("Búsequedas en árboles DSW")
	fmt.Println("\nDSW de 200 elementos.")
	buscar_dsw(a10000, dsw200)
	fmt.Println("\nDSW de 400 elementos.")
	buscar_dsw(a10000, dsw400)
	fmt.Println("\nDSW de 600 elementos.")
	buscar_dsw(a10000, dsw600)
	fmt.Println("\nDSW de 800 elementos.")
	buscar_dsw(a10000, dsw800)
	fmt.Println("\nDSW de 1000 elementos.")
	buscar_dsw(a10000, dsw1000)

	// Experimento d)
}

// func uwu() {
// 	fmt.Println("⡆⣐⢕⢕⢕⢕⢕⢕⢕⢕⠅⢗⢕⢕⢕⢕⢕⢕⢕⠕⠕⢕⢕⢕⢕⢕⢕⢕⢕⢕")
//  fmt.Println("⢐⢕⢕⢕⢕⢕⣕⢕⢕⠕⠁⢕⢕⢕⢕⢕⢕⢕⢕⠅⡄⢕⢕⢕⢕⢕⢕⢕⢕⢕")
// 	fmt.Println("⢕⢕⢕⢕⠁⢜⠕⢁⣴⣿⡇⢓⢕⢵⢐⢕⢕⠕⢁⣾⢿⣧⠑⢕⢕⠄⢑⢕⠅⢕")
// 	fmt.Println("⢕⢕⠵⢁⠔⢁⣤⣤⣶⣶⣶⡐⣕⢽⠐⢕⠕⣡⣾⣶⣶⣶⣤⡁⢓⢕⠄⢑⢅⢑")
// 	fmt.Println("⠍⣧⠄⣶⣾⣿⣿⣿⣿⣿⣿⣷⣔⢕⢄⢡⣾⣿⣿⣿⣿⣿⣿⣿⣦⡑⢕⢤⠱⢐")
// 	fmt.Println("⢠⢕⠅⣾⣿⠋⢿⣿⣿⣿⠉⣿⣿⣷⣦⣶⣽⣿⣿⠈⣿⣿⣿⣿⠏⢹⣷⣷⡅⢐")
// 	fmt.Println("⣔⢕⢥⢻⣿⡀⠈⠛⠛⠁⢠⣿⣿⣿⣿⣿⣿⣿⣿⡀⠈⠛⠛⠁⠄⣼⣿⣿⡇⢔")
// 	fmt.Println("⢕⢕⢽⢸⢟⢟⢖⢖⢤⣶⡟⢻⣿⡿⠻⣿⣿⡟⢀⣿⣦⢤⢤⢔⢞⢿⢿⣿⠁⢕")
// 	fmt.Println("⢕⢕⠅⣐⢕⢕⢕⢕⢕⣿⣿⡄⠛⢀⣦⠈⠛⢁⣼⣿⢗⢕⢕⢕⢕⢕⢕⡏⣘⢕")
// 	fmt.Println("⢕⢕⠅⢓⣕⣕⣕⣕⣵⣿⣿⣿⣾⣿⣿⣿⣿⣿⣿⣿⣷⣕⢕⢕⢕⢕⡵⢀⢕⢕")
// 	fmt.Println("⢑⢕⠃⡈⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢃⢕⢕⢕")
// 	fmt.Println("⣆⢕⠄⢱⣄⠛⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⢁⢕⢕⠕⢁")
// 	fmt.Println("⣿⣦⡀⣿⣿⣷⣶⣬⣍⣛⣛⣛⡛⠿⠿⠿⠛⠛⢛⣛⣉⣭⣤⣂⢜⠕⢑⣡⣴⣿")
// }

// PRUEBAS VIEJAS
// func pruebas_avl() {
// 	tree := crear_avl(1)

// 	// Código para demostrar los requerimientos 6 y 8 de la especificación.
// 	fmt.Printf("\nInserciones al árbol AVL: \n")

// 	var comp int
// 	comp = tree.insertar(7)
// 	println("Cantidad de comparaciones realizadas: " + strconv.Itoa(comp))

// 	comp = tree.insertar(4)
// 	println("Cantidad de comparaciones realizadas: " + strconv.Itoa(comp))

// 	comp = tree.insertar(8)
// 	println("Cantidad de comparaciones realizadas: " + strconv.Itoa(comp))

// 	comp = tree.insertar(3)
// 	println("Cantidad de comparaciones realizadas: " + strconv.Itoa(comp))

// 	comp = tree.insertar(7)
// 	println("Cantidad de comparaciones realizadas: " + strconv.Itoa(comp))

// 	// Código para demostrar el requerimiento 7 de la especificación.
// 	fmt.Printf("\nBúsqueda de una llave: \n")

// 	var llave_busq = 2
// 	var result, cmp = tree.buscar(llave_busq, false)

// 	if result {
// 		fmt.Println("Se logró encontrar el nodo con llave: " + strconv.Itoa(llave_busq))
// 	} else {
// 		fmt.Println("No se logró encontrar el nodo con llave: " + strconv.Itoa(llave_busq))
// 	}
// 	fmt.Println("Número de comparaciones: " + strconv.Itoa(cmp))

// 	// Imprimir el árbol en preorden para verificar su estructura.

// 	fmt.Printf("\nEstructura del árbol: \n")
// 	tree.preorden(true)

// 	fmt.Println(tree.raiz.altura)
// 	fmt.Println(tree.size)
// }

// func main() {
// 	//fmt.Println( arreglo(37,200) )
// 	tree := binTree{}
// 	tree.insertNode(2)
// 	tree.insertNode(4)
// 	tree.insertNode(5)
// 	tree.insertNode(7)
// 	tree.insertNode(8)
// 	tree.insertNode(11)
// 	tree.insertNode(12)
// 	tree.insertNode(17)
// 	tree.insertNode(18)
// 	String(tree.root)

// 	tree.createDSW()
// 	String(tree.root)

// 	// fmt.Println(tree.insertNode(10))
// 	// fmt.Println(tree.insertNode(5))
// 	// fmt.Println(tree.insertNode(12))
// 	// fmt.Println(tree.insertNode(130))
// 	// fmt.Println(tree.insertNode(45))
// 	// fmt.Println(tree.insertNode(75))
// 	// fmt.Println(tree.insertNode(35))
// 	// fmt.Println(tree.insertNode(50))
// 	// createBackBone(tree.root)
// 	// String(tree.root)

// 	//tree.root.treeStructure(tree.size)
// 	//inOrder(tree.root)

// 	//tree.String()

// 	//result := binSearch(tree.root,75,0)
// 	//fmt.Println(tree)
// 	// fmt.Println("\n")
// 	// fmt.Println(tree.size)
// 	pruebas_avl()
// }
