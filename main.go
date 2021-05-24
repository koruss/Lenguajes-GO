package main

import "fmt"

/*

	REQ 1

*/

func isPrimo(seed int) bool { //
	if seed >= 11 && seed <= 101 { // si esta en el rango permitido
		for num := 2; num < seed/2; num++ { //revisa si es primo a fuerza bruta
			if (seed % num) == 0 {
				return false
			}
		}
		return true
	}
	return false
}

func arreglo(seed int, size int) []int {
	if isPrimo(seed) { // revisa si la semilla ingresada es valida
		var slice = make([]int, size) //asi se crea un arreglo dinamico en go (slice)
		for i := range slice {
			seed = pseudoRand(seed) // restringe a que sean valores menores a 200
			slice[i] = seed
		}
		return slice
	} else {
		return nil
	}

}

func pseudoRand(seed int) int {
	m := 4096             // entre 11 y 101	//averiguar sobre este
	a := 109              // y este numero
	xn1 := (a * seed) % m // algoritmo multiplicativo
	return xn1
}

/*

	Arbol Binario

*/
type binNode struct {
	left  *binNode
	right *binNode
	value int
	cont  int
}

type binTree struct {
	root *binNode
}

func (tree *binTree) insertNode(data int) *binTree {
	if tree.root == nil { // si el arbol esta vacio cree un nodo
		tree.root = &binNode{
			value: data,
			left:  nil,
			right: nil}
	} else { // si no esta vacio llame a la funcion de insertar
		tree.root.insert(data)
	}
	return tree
}

func (node *binNode) insert(data int) {
	if node == nil { // si llego a una rama nula
		return
	} else if data < node.value { // si es menor
		if node.left == nil { //si no hay nodo en el lado izq
			node.left = &binNode{value: data, left: nil, right: nil} // cree un nodo y lo asigna
		} else { // si hay nodo entonces muevase
			node.left.insert(data) // se llama recursivamente moviendose hasta encontrar un espacio vacio
		}
	} else { // si es mayor
		if node.right == nil { // si no hay nodo en el lado derecho
			node.right = &binNode{value: data, left: nil, right: nil}
		} else {
			node.right.insert(data)
		}
	}
}

func inOrder(root *binNode) {
	if root == nil {
		return
	}
	inOrder(root.left)         // izquierda
	fmt.Print(root.value, " ") // centro
	inOrder(root.right)        // derecha
}

func binSearch(binNode *binNode, key int) {
	

}

func binSearchAux(root *binNode, key int, count int) int {
	if root == nil || root.value == key {
		return count
	}
	if root.value < key {
		count++
		return binSearchAux(root.right, key, count)
	}
	count++
	return binSearchAux(root.left, key, count)
}

/*

	Funcion Principal

*/
func main() {
	//fmt.Println( arreglo(37,200) )
	tree := binTree{}
	tree.insertNode(10)
	tree.insertNode(5)
	tree.insertNode(12)
	tree.insertNode(130)
	tree.insertNode(45)
	tree.insertNode(75)
	tree.insertNode(35)
	inOrder(tree.root)
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
