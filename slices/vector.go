package main;

import "fmt";

func main() {
  // slice simple
  // slice := []int {1, 2, 3, 4};
  // fmt.Println(slice);

  // var matriz []int;
  // if matriz == nil {
  //  fmt.Println("Esta vacio");
  // }

  // Puntero al arreglo
  // Longitud del arreglo al que apunta
  // Capacidad

  arreglo := [3]int {1, 2, 3};
  // slicing
  slice := arreglo[ : 2];
  fmt.Println(slice);
}
