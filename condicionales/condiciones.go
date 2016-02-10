package main;

import "fmt";

/*
 == igual a
 != diferente de
 < menor que
 > mayor que
 >= mayor o igual que
 <= menor o igual que
 && AND
 || OR
*/

func main() {
  x := 10;
  y := 10;

  if (x > y) {
    fmt.Printf("%d es mayor que %d\n", x, y);
  } else if(x < y) {
    fmt.Printf("%d es menor que %d\n", x, y);
  } else {
    fmt.Println("Son iguales");
  }
}
