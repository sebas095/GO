package main;

import (
  "fmt"
  "bufio"
  "os"
);

func main() {
  /*
  Usando printf y scanf
  var edad int;
  var nombre string;
  %d enteros
  %v tipo de la variable (en scanf lo lee como string)
  %t booleanos
  %e, %f flotantes

  fmt.Println("Coloca tú nombre: ");
  fmt.Scanf("%v\n", &nombre);
  fmt.Printf("Hola %s\n", nombre);
  */

  reader := bufio.NewReader(os.Stdin);
  fmt.Println("Ingresa tú nombre: ");
  nombre, err := reader.ReadString('\n');

  if err != nil {
    fmt.Println(err);
  }else {
    fmt.Println("Hola " + nombre);
  }

}
