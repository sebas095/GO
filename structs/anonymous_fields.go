package main;

import "fmt";

type Human struct {
  name string;
}

type Dummy struct {
  name string;
}

type Tutor struct {
  Human;
  Dummy;
}

func (this Human) hablar() string {
  return "Bla bla bla";
}

func (this Tutor) hablar() string {
  return this.Human.hablar() + " Bienvenidos a CodigoFacilito";
}

func main() {
  tutor := Tutor{Human{"Sebastian"}, Dummy{"Duque"}};
  // ex := tutor.hablar;
  fmt.Println(tutor.hablar());
}