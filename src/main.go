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

func f64_arr_to_string(p_arr []float64) string {
	var msg string
	for _, num := range p_arr {
		msg += strconv.FormatFloat(num, 'g', 3, 64) + "\n"
	}
	return msg
}

func buscar_avl(p_arr []int, arbol *arbol_avl) {
	// var cmp int
	// var res bool
	// var fnd string
	for _, num := range p_arr {
		// res, cmp =
		arbol.buscar(num, false)
		// if res {
		// 	fnd = "SÍ"
		// } else {
		// 	fnd = "NO"
		// }
		// println("Busqueda en AVL con llave: " + strconv.Itoa(num) + " ,¿Encontrado?: " + fnd + " ,Cantidad de comparaciones: " + strconv.Itoa(cmp))
	}
}

func buscar_bst(p_arr []int, arbol binTree) {
	// var fnd string
	// var resp response
	for _, num := range p_arr {
		// resp =
		binSearch(arbol.root, num)
		// if resp.state {
		// 	fnd = "SÍ"
		// } else {
		// 	fnd = "NO"
		// }
		// println("Busqueda en BST con llave: " + strconv.Itoa(num) + " ,¿Encontrado?: " + fnd + " ,Cantidad de comparaciones: " + strconv.Itoa(resp.comparisons))
	}
}

func buscar_dsw(p_arr []int, arbol binTree) {
	// var fnd string
	// var resp response
	for _, num := range p_arr {
		// resp =
		binSearch(arbol.root, num)
		// if resp.state {
		// 	fnd = "SÍ"
		// } else {
		// 	fnd = "NO"
		// }
		// println("Busqueda en DSW con llave: " + strconv.Itoa(num) + " ,¿Encontrado?: " + fnd + " ,Cantidad de comparaciones: " + strconv.Itoa(resp.comparisons))
	}
}

func div_entero_to_float(int1 int, int2 int) float64 {
	return float64(int1) / float64(int2)
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
	var alturas_avl = make([]int, 5)
	var alturas_bst = make([]int, 5)
	var alturas_dsw = make([]int, 5)

	// Conseguir las alturas de los AVL.
	alturas_avl[0] = avl200.raiz.altura
	alturas_avl[1] = avl400.raiz.altura
	alturas_avl[2] = avl600.raiz.altura
	alturas_avl[3] = avl800.raiz.altura
	alturas_avl[4] = avl1000.raiz.altura

	fmt.Println("\nAltura del AVL de 200 elementos: " + strconv.Itoa(alturas_avl[0]))
	fmt.Println("Altura del AVL de 400 elementos: " + strconv.Itoa(alturas_avl[1]))
	fmt.Println("Altura del AVL de 600 elementos: " + strconv.Itoa(alturas_avl[2]))
	fmt.Println("Altura del AVL de 800 elementos: " + strconv.Itoa(alturas_avl[3]))
	fmt.Println("Altura del AVL de 1000 elementos: " + strconv.Itoa(alturas_avl[4]))

	// Conseguir las alturas de los BST.
	alturas_bst[0] = calcBSTHeight(bst200.root)
	alturas_bst[1] = calcBSTHeight(bst400.root)
	alturas_bst[2] = calcBSTHeight(bst600.root)
	alturas_bst[3] = calcBSTHeight(bst800.root)
	alturas_bst[4] = calcBSTHeight(bst1000.root)

	fmt.Println("\nAltura del BST de 200 elementos: " + strconv.Itoa(alturas_bst[0]))
	fmt.Println("Altura del BST de 400 elementos: " + strconv.Itoa(alturas_bst[1]))
	fmt.Println("Altura del BST de 600 elementos: " + strconv.Itoa(alturas_bst[2]))
	fmt.Println("Altura del BST de 800 elementos: " + strconv.Itoa(alturas_bst[3]))
	fmt.Println("Altura del BST de 1000 elementos: " + strconv.Itoa(alturas_bst[4]))

	// Conseguir las alturas de los DSW.
	alturas_dsw[0] = calcBSTHeight(dsw200.root)
	alturas_dsw[1] = calcBSTHeight(dsw400.root)
	alturas_dsw[2] = calcBSTHeight(dsw600.root)
	alturas_dsw[3] = calcBSTHeight(dsw800.root)
	alturas_dsw[4] = calcBSTHeight(dsw1000.root)

	fmt.Println("\nAltura del DSW de 200 elementos: " + strconv.Itoa(alturas_dsw[0]))
	fmt.Println("Altura del DSW de 400 elementos: " + strconv.Itoa(alturas_dsw[1]))
	fmt.Println("Altura del DSW de 600 elementos: " + strconv.Itoa(alturas_dsw[2]))
	fmt.Println("Altura del DSW de 800 elementos: " + strconv.Itoa(alturas_dsw[3]))
	fmt.Println("Altura del DSW de 1000 elementos: " + strconv.Itoa(alturas_dsw[4]))

	//Cálculo de profundidad promedio por árbol
	var prof_prom_avl = make([]float64, 5)
	var prof_prom_bst = make([]float64, 5)
	var prof_prom_dsw = make([]float64, 5)

	prof_prom_avl[0] = div_entero_to_float(avl200.get_profundidad_sumada(), avl200.size)
	prof_prom_avl[1] = div_entero_to_float(avl400.get_profundidad_sumada(), avl400.size)
	prof_prom_avl[2] = div_entero_to_float(avl600.get_profundidad_sumada(), avl600.size)
	prof_prom_avl[3] = div_entero_to_float(avl800.get_profundidad_sumada(), avl800.size)
	prof_prom_avl[4] = div_entero_to_float(avl1000.get_profundidad_sumada(), avl1000.size)
	fmt.Println("Profunidad promedio de los AVL:\n" + f64_arr_to_string(prof_prom_avl))

	prof_prom_bst[0] = div_entero_to_float(bst200.get_profundidad_sumada_bst(), bst200.size)
	prof_prom_bst[1] = div_entero_to_float(bst400.get_profundidad_sumada_bst(), bst400.size)
	prof_prom_bst[2] = div_entero_to_float(bst600.get_profundidad_sumada_bst(), bst600.size)
	prof_prom_bst[3] = div_entero_to_float(bst800.get_profundidad_sumada_bst(), bst800.size)
	prof_prom_bst[4] = div_entero_to_float(bst1000.get_profundidad_sumada_bst(), bst1000.size)
	fmt.Println("Profunidad promedio de los BST:\n" + f64_arr_to_string(prof_prom_bst))

	prof_prom_dsw[0] = div_entero_to_float(dsw200.get_profundidad_sumada_bst(), dsw200.size)
	prof_prom_dsw[1] = div_entero_to_float(dsw400.get_profundidad_sumada_bst(), dsw400.size)
	prof_prom_dsw[2] = div_entero_to_float(dsw600.get_profundidad_sumada_bst(), dsw600.size)
	prof_prom_dsw[3] = div_entero_to_float(dsw800.get_profundidad_sumada_bst(), dsw800.size)
	prof_prom_dsw[4] = div_entero_to_float(dsw1000.get_profundidad_sumada_bst(), dsw1000.size)
	fmt.Println("Profunidad promedio de los DSW:\n" + f64_arr_to_string(prof_prom_dsw))

	//Cálculo de densidad por árbol
	var dens_avl = make([]float64, 5)
	var dens_bst = make([]float64, 5)
	var dens_dsw = make([]float64, 5)

	dens_avl[0] = div_entero_to_float(avl200.size, avl200.raiz.altura)
	dens_avl[1] = div_entero_to_float(avl400.size, avl400.raiz.altura)
	dens_avl[2] = div_entero_to_float(avl600.size, avl600.raiz.altura)
	dens_avl[3] = div_entero_to_float(avl800.size, avl800.raiz.altura)
	dens_avl[4] = div_entero_to_float(avl1000.size, avl1000.raiz.altura)
	fmt.Println("Densidad de los AVL:\n" + f64_arr_to_string(dens_avl))

	dens_bst[0] = div_entero_to_float(bst200.size, calcBSTHeight(bst200.root))
	dens_bst[1] = div_entero_to_float(bst400.size, calcBSTHeight(bst400.root))
	dens_bst[2] = div_entero_to_float(bst600.size, calcBSTHeight(bst600.root))
	dens_bst[3] = div_entero_to_float(bst800.size, calcBSTHeight(bst800.root))
	dens_bst[4] = div_entero_to_float(bst1000.size, calcBSTHeight(bst1000.root))
	fmt.Println("Densidad de los BST:\n" + f64_arr_to_string(dens_bst))

	dens_dsw[0] = div_entero_to_float(dsw200.size, calcBSTHeight(dsw200.root))
	dens_dsw[1] = div_entero_to_float(dsw400.size, calcBSTHeight(dsw400.root))
	dens_dsw[2] = div_entero_to_float(dsw600.size, calcBSTHeight(dsw600.root))
	dens_dsw[3] = div_entero_to_float(dsw800.size, calcBSTHeight(dsw800.root))
	dens_dsw[4] = div_entero_to_float(dsw1000.size, calcBSTHeight(dsw1000.root))
	fmt.Println("Densidad de los DSW:\n" + f64_arr_to_string(dens_dsw))

	//Comparaciones realizadas

	// x := []int{10, 9, 11, 16, 18, 18, 18, 19, 6}
	// tree := insert_arr_bst(x)
	// // fmt.Println(tree.root)
	// // fmt.Println(tree.insertNode(14))
	// fmt.Println(tree.buscar(19))

}
