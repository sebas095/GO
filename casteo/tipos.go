package main;

import (
  "fmt"
  "strconv"
);

func main() {
  edad := 21;
  num_str := "22";

  edad_str := strconv.Itoa(edad);
  num_int,_ := strconv.Atoi(num_str);

  fmt.Println("Mi edad es " + edad_str);
  fmt.Println(num_int + 10);
}
