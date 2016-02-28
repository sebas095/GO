package main;

import "fmt";

func main() {
  /*
    Copia el minimo de elementosen ambos arreglos
  */
  slice := []int {1, 2, 3, 4, 6};
  copia := make([]int, len(slice), cap(slice) * 2);

  // copy(dest, fuente);
  copy(copia, slice);

  fmt.Println(slice);
  fmt.Println(copia);
}
