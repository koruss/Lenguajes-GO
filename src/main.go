package main

import (
	"fmt"
	"strconv"
)

/*
	Funcion Principal

*/

func pruebas_avl() {
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

func main() {
	//fmt.Println( arreglo(37,200) )
	tree := binTree{}
	tree.insertNode(2)
	tree.insertNode(4)
	tree.insertNode(5)
	tree.insertNode(7)
	tree.insertNode(8)
	tree.insertNode(11)
	tree.insertNode(12)
	tree.insertNode(17)
	tree.insertNode(18)
	String(tree.root)

	tree.createDSW()
	String(tree.root)

	// fmt.Println(tree.insertNode(10))
	// fmt.Println(tree.insertNode(5))
	// fmt.Println(tree.insertNode(12))
	// fmt.Println(tree.insertNode(130))
	// fmt.Println(tree.insertNode(45))
	// fmt.Println(tree.insertNode(75))
	// fmt.Println(tree.insertNode(35))
	// fmt.Println(tree.insertNode(50))
	// createBackBone(tree.root)
	// String(tree.root)

	//tree.root.treeStructure(tree.size)
	//inOrder(tree.root)

	//tree.String()

	//result := binSearch(tree.root,75,0)
	//fmt.Println(tree)
	// fmt.Println("\n")
	// fmt.Println(tree.size)
	pruebas_avl()
}

// func main() {
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
