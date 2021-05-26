package main

import "fmt"

/*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-

	REQ 1

*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*/

/*
Recibe un numero entre 11 y 101 y revisa si es primo
Devuelve TRUE si es primo FALSE si no
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


/*
Funcion para crear un arreglo de tamanio  definido por el usuario
Recibe un int la semilla para la generacion aleatoria y un int el tamanio del arreglo 
*/
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
/*
Generadora de numeros aleatorios, basado en la congruencia lineal
*/
func pseudoRand(seed int) int {
	m := 4096             // entre 11 y 101	//averiguar sobre este
	a := 109              // y este numero
	b := 853
	xn1 := (a * seed +b) % m // algoritmo multiplicativo
	return xn1
}

/*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-

								Arbol Binario

*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*/
type binNode struct {
	left  *binNode
	right *binNode
	value int
	cont  int
}

type binTree struct {
	root *binNode
	size int // medir la cantidad
}

type response struct { // la tupla que pide el req 3 devolver
	state       bool
	comparisons int
}

//----------------------------------------------------------------------//

func (tree *binTree) insertNode(data int) int {
	var comps = 0
	if tree.root == nil { // si el arbol esta vacio cree un nodo
		tree.root = &binNode{
			value: data,
			left:  nil,
			right: nil}
		tree.size++
	} else { // si no esta vacio llame a la funcion de insertar
		comps = tree.root.insert(data, 0)
		tree.size++
	}
	return comps
}

func (node *binNode) insert(data int, cont int) int {
	if node == nil { // si llego a una rama nula
		return 1
	} else if data < node.value { // si es menor
		if node.left == nil { //si no hay nodo en el lado izq
			cont++
			node.left = &binNode{value: data, left: nil, right: nil, cont: 1} // cree un nodo y lo asigna
			return cont

			//fmt.Println(cont)
		} else { // si hay nodo entonces muevase
			cont++
			return node.left.insert(data, cont) // se llama recursivamente moviendose hasta encontrar un espacio vacio
		}
	} else if data > node.value { // si es mayor
		if node.right == nil { // si no hay nodo en el lado derecho
			cont++
			node.right = &binNode{value: data, left: nil, right: nil, cont: 1}
			//fmt.Println(cont)
			return cont
		} else {
			cont++
			return node.right.insert(data, cont)

		}
	} else { // si no es mayor y no es menor el unico caso es que sea igual
		node.cont++
	}
	return 0
}

func binSearch(root *binNode, key int, count int) response {
	if root == nil {
		return response{state: false, comparisons: count}
	}
	if root.value == key {
		count++
		return response{state: true, comparisons: count}
	}
	if root.value < key {
		count++
		return binSearch(root.right, key, count)
	}
	count++
	return binSearch(root.left, key, count)
}

//-----------------------------------------------------------------------

func createDSW(tree *binTree){
	if (tree.root != nil){
		createBackBone(tree.root)
		balancedTree(tree.root)
	}
}

func createBackBone(root *binNode){
	grandParent:= &binNode{}
	parent := root
	leftChild:= &binNode{}

	for (parent != nil) {
		leftChild = parent.left;
		if (leftChild != nil) {
		  grandParent = rotateRight(grandParent, parent, leftChild,root);
		  parent = leftChild;
		} else {
		  grandParent = parent;
		  parent = parent.right;
		}
	  }

}

func rotateRight(grandParent *binNode,parent *binNode, leftChild *binNode,root *binNode) *binNode{
	if (grandParent!= nil) {
		grandParent.right = leftChild;
	  } else {
		root = leftChild;
	  }
	  parent.left = leftChild.right;
	  leftChild.right = parent;
	  return grandParent;
}
func rotateLeft(grandParent *binNode,parent *binNode,rightChild *binNode,root *binNode) {
	if (grandParent!=nil) {
	  grandParent.right = rightChild;
	} else {
	  root = rightChild;
	}
	parent.right = rightChild.left;
	rightChild.left = parent;
  }

func balancedTree(root *binNode){
	n:=0
	for temp := root; temp!=nil; temp = temp.right{
		n++
	}
	m:= funcAux(n+1)-1
	m/=2
	makeRotations(m, root)
}

func funcAux(n int) int {
	x := MSB(n);//MSB
	return (1 << x);//2^x
}

func MSB(n int) int{
	ndx := 0
	for (1 < n) {
	  n = (n >> 1);
	  ndx++;
	}
	return ndx;
  }

func makeRotations(bound int,root *binNode)  {
	grandParent:= &binNode{}
	parent := root;
	child := root.right;
	for i:=bound; i > 0; i-- {
		if (child!=nil) {
			rotateLeft(grandParent, parent, child,root);
			grandParent = child;
			parent = grandParent.right;
			child = parent.right;
		} else {
			break;
		}
	} 
}
  
/*

	Funciones Extra

*/

/*
Recibe la raiz de un arbol binario y lo imprime
*/
func inOrder(root *binNode) {
	if root == nil {
		return
	}
	inOrder(root.left)         // izquierda
	fmt.Print(root.value, " ") // centro
	inOrder(root.right)        // derecha
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

	// fmt.Println(tree.insertNode(10))
	// fmt.Println(tree.insertNode(5))
	// fmt.Println(tree.insertNode(12))
	// fmt.Println(tree.insertNode(130))
	// fmt.Println(tree.insertNode(45))
	// fmt.Println(tree.insertNode(75))
	// fmt.Println(tree.insertNode(35))
	// fmt.Println(tree.insertNode(50))

	//tree.root.treeStructure(tree.size)
	inOrder(tree.root)
	
	//result := binSearch(tree.root,75,0)
	//fmt.Println(tree)
	// fmt.Println("\n")
	// fmt.Println(tree.size)
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
