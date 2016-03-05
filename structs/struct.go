package main;

import "fmt";

type User struct {
  edad int;
  nombre, apellido string
}

func main() {
  /*
    Forma 1
    var sebas User;
  */
  // Forma 2
  // sebas := User{21, "Sebastian", "Duque"}; // User{edad: 21, nombre: "Sebastian", apellido: "Duque"};

  // Forma 3
  sebas := new(User);
  sebas.nombre = "Sebas";
  
  fmt.Println(sebas.nombre);
}
