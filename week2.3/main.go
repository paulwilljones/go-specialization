package main

import "fmt"

func main() {
  var a, v string
  cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
  bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
  snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

  fmt.Println("Enter an animal and a verb, eg. cow eat")
  for {
    fmt.Print("> ")
    fmt.Scan(&a, &v)

    var animal *Animal

    switch a {
    case "cow":
      animal = &cow
    case "bird":
      animal = &bird
    case "snake":
      animal = &snake
    default:
      fmt.Println("I don't recognise that animal")
      continue
    }

    switch v {
    case "eat":
      animal.Eat()
    case "move":
      animal.Move()
    case "speak":
      animal.Speak()
    default:
      fmt.Println("I don't recognise that action")
      continue
    }
  }
}

func (a Animal) Eat() {
  fmt.Println(a.food)
}

func (a Animal) Move() {
  fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
  fmt.Println(a.noise)
}

type Animal struct {
  food, locomotion, noise string
}
