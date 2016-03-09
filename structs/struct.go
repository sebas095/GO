package main;

import "fmt";

type User struct {
  edad int;
  nombre, apellido string
}

func (this User) nombre_completo() string {
  return this.nombre + " " + this.apellido;
}

func (this *User) set_name(n string) {
  this.nombre = n;
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
  sebas.apellido = "Duque";
  sebas.set_name("lol");

  fmt.Println(sebas.nombre_completo());
}
