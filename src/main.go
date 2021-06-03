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
		} else {
			AVL_COMP += tree.insertar(num, false)
		}
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
	cant := 200
	for _, num := range p_arr {
		msg += "Árbol de " + strconv.Itoa(cant) + " elementos: " + strconv.Itoa(num) + "\n"
		cant += 200
	}
	return msg
}

func f64_arr_to_string(p_arr []float64) string {
	var msg string
	cant := 200
	for _, num := range p_arr {
		msg += "Árbol de " + strconv.Itoa(cant) + " elementos: " + strconv.FormatFloat(num, 'g', 3, 64) + "\n"
		cant += 200
	}
	return msg
}

func buscar_avl(p_arr []int, arbol *arbol_avl) int {
	// Reiniciar el valor de las comparaciones del AVL a cero.
	AVL_COMP = 0
	var cmp int
	// var fnd string
	for _, num := range p_arr {
		// En este caso el valor booleano no sirve de nada, solo las comparaciones.
		_, cmp = arbol.buscar(num, false)
		AVL_COMP += cmp
	}
	return AVL_COMP
}

func buscar_bst(p_arr []int, arbol binTree) int {
	// Reiniciar el valor de las comparaciones del BST/DSW a cero.
	BST_COMP = 0
	var resp response
	for _, num := range p_arr {
		resp = arbol.buscar(num)
		BST_COMP += resp.comparisons
	}
	return BST_COMP
}

func div_entero_to_float(int1 int, int2 int) float64 {
	return float64(int1) / float64(int2)
}

/*
	Función Principal
*/

func main() {
	// Experimento a)

	a200 := arreglo(37, 200)
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

	avl400 := insert_arr_avl(a400)
	comparaciones_avl[1] = AVL_COMP

	avl600 := insert_arr_avl(a600)
	comparaciones_avl[2] = AVL_COMP

	avl800 := insert_arr_avl(a800)
	comparaciones_avl[3] = AVL_COMP

	avl1000 := insert_arr_avl(a1000)
	comparaciones_avl[4] = AVL_COMP

	// BST
	bst200 := insert_arr_bst(a200)
	comparaciones_dsw[0] = BST_COMP

	bst400 := insert_arr_bst(a400)
	comparaciones_dsw[1] = BST_COMP

	bst600 := insert_arr_bst(a600)
	comparaciones_dsw[2] = BST_COMP

	bst800 := insert_arr_bst(a800)
	comparaciones_dsw[3] = BST_COMP

	bst1000 := insert_arr_bst(a1000)
	comparaciones_dsw[4] = BST_COMP

	// DSW
	dsw200 := insert_arr_bst(a200)
	comparaciones_bst[0] = BST_COMP
	dsw200.createDSW()

	dsw400 := insert_arr_bst(a400)
	comparaciones_bst[1] = BST_COMP
	dsw400.createDSW()

	dsw600 := insert_arr_bst(a600)
	comparaciones_bst[2] = BST_COMP
	dsw600.createDSW()

	dsw800 := insert_arr_bst(a800)
	comparaciones_bst[3] = BST_COMP
	dsw800.createDSW()

	dsw1000 := insert_arr_bst(a1000)
	comparaciones_bst[4] = BST_COMP
	dsw1000.createDSW()

	// Experimento c)

	// Arreglo de 10000 elementos pseudo aleatorios.
	a10000 := arreglo(89, 10000)

	var comp_busq_avl = make([]int, 5)
	var comp_busq_bst = make([]int, 5)
	var comp_busq_dsw = make([]int, 5)

	comp_busq_avl[0] = buscar_avl(a10000, avl200)
	comp_busq_avl[1] = buscar_avl(a10000, avl400)
	comp_busq_avl[2] = buscar_avl(a10000, avl600)
	comp_busq_avl[3] = buscar_avl(a10000, avl800)
	comp_busq_avl[4] = buscar_avl(a10000, avl1000)

	comp_busq_bst[0] = buscar_bst(a10000, bst200)
	comp_busq_bst[1] = buscar_bst(a10000, bst400)
	comp_busq_bst[2] = buscar_bst(a10000, bst600)
	comp_busq_bst[3] = buscar_bst(a10000, bst800)
	comp_busq_bst[4] = buscar_bst(a10000, bst1000)

	comp_busq_dsw[0] = buscar_bst(a10000, dsw200)
	comp_busq_dsw[1] = buscar_bst(a10000, dsw400)
	comp_busq_dsw[2] = buscar_bst(a10000, dsw600)
	comp_busq_dsw[3] = buscar_bst(a10000, dsw800)
	comp_busq_dsw[4] = buscar_bst(a10000, dsw1000)

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

	fmt.Println("\nComparaciones realizadas al insertar.")
	fmt.Println("Comparaciones de los AVL:\n" + int_arr_to_string(comparaciones_avl))
	fmt.Println("Comparaciones de los BST:\n" + int_arr_to_string(comparaciones_bst))
	fmt.Println("Comparaciones de los DSW:\n" + int_arr_to_string(comparaciones_dsw))

	fmt.Println("\nComparaciónes realizadas al buscar.")
	fmt.Println("Comparaciones de los AVL:\n" + int_arr_to_string(comp_busq_avl))
	fmt.Println("Comparaciones de los BST:\n" + int_arr_to_string(comp_busq_bst))
	fmt.Println("Comparaciones de los DSW:\n" + int_arr_to_string(comp_busq_dsw))

	//Cantidad promedio de comparaciones realizadas para las inserciones
	//realizadas, ponderadas por la altura donde se encuentra cada nodo.

	var calc_prom_avl = make([]float64, 5)
	var calc_prom_bst = make([]float64, 5)
	var calc_prom_dsw = make([]float64, 5)

	calc_prom_avl[0] = div_entero_to_float(avl200.calculo_ponderado(), avl200.size)
	calc_prom_avl[1] = div_entero_to_float(avl400.calculo_ponderado(), avl400.size)
	calc_prom_avl[2] = div_entero_to_float(avl600.calculo_ponderado(), avl600.size)
	calc_prom_avl[3] = div_entero_to_float(avl800.calculo_ponderado(), avl800.size)
	calc_prom_avl[4] = div_entero_to_float(avl1000.calculo_ponderado(), avl1000.size)
	fmt.Println("Cantidad promedio de comparaciones realizadas para las inserciones, \nponderadas por la altura donde se encuentra cada nodo de los árboles AVL:\n" + f64_arr_to_string(calc_prom_avl))

	calc_prom_bst[0] = div_entero_to_float(bst200.calculo_ponderado_bst(), bst200.size)
	calc_prom_bst[1] = div_entero_to_float(bst400.calculo_ponderado_bst(), bst400.size)
	calc_prom_bst[2] = div_entero_to_float(bst600.calculo_ponderado_bst(), bst600.size)
	calc_prom_bst[3] = div_entero_to_float(bst800.calculo_ponderado_bst(), bst800.size)
	calc_prom_bst[4] = div_entero_to_float(bst1000.calculo_ponderado_bst(), bst1000.size)
	fmt.Println("Cantidad promedio de comparaciones realizadas para las inserciones, \nponderadas por la altura donde se encuentra cada nodo de los árboles BST:\n" + f64_arr_to_string(calc_prom_bst))

	calc_prom_dsw[0] = div_entero_to_float(dsw200.calculo_ponderado_bst(), dsw200.size)
	calc_prom_dsw[1] = div_entero_to_float(dsw400.calculo_ponderado_bst(), dsw400.size)
	calc_prom_dsw[2] = div_entero_to_float(dsw600.calculo_ponderado_bst(), dsw600.size)
	calc_prom_dsw[3] = div_entero_to_float(dsw800.calculo_ponderado_bst(), dsw800.size)
	calc_prom_dsw[4] = div_entero_to_float(dsw1000.calculo_ponderado_bst(), dsw1000.size)
	fmt.Println("Cantidad promedio de comparaciones realizadas para las inserciones, \nponderadas por la altura donde se encuentra cada nodo de los árboles DSW:\n" + f64_arr_to_string(calc_prom_dsw))

	// x := []int{2, 4, 5, 7, 8, 11, 12, 17, 18, 18, 18, 18}
	// tree := insert_arr_bst(x)
	// tree.createDSW()
	// String(tree.root)
	// // fmt.Println(tree.root)
	// // fmt.Println(tree.insertNode(14))
	// fmt.Println(tree.calculo_ponderado_bst())

}
