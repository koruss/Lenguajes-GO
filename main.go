package main

import "fmt"

/* 

	REQ 1 

*/

func isPrimo(seed int) bool {//
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
	if isPrimo(seed) {// revisa si la semilla ingresada es valida
		var slice = make([]int, size) //asi se crea un arreglo dinamico en go (slice)
		for i := range slice {
			seed = pseudoRand(seed)%200 // restringe a que sean valores menores a 200
			slice[i] = seed
		}
		return slice
	}else{	
		return nil
	}
		

}

func pseudoRand(seed int) int {
	m := 4080        // entre 11 y 101	//averiguar sobre este
	a := 109             // y este numero
	xn1 := (a * seed ) % m // algoritmo multiplicativo
	return xn1
}




/*

	Arbol Binario

*/



/*

	Funcion Principal

*/
func main() {
	fmt.Println( arreglo(53,200) )
	// for i:=0; i<5; i++{
	// 	fmt.Println(arreglo[i])
	// }

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
