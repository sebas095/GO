package main;

import "fmt";

func main() {
  // tipo, len, capacity(?)
  slice := make([]int, 3, 5);
  slice = append(slice, 2);
  fmt.Println(len(slice));
}
