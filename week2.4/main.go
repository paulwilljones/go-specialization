package main

import "fmt"

func main() {
  var cmd, name, attribute string
  animals := make(map[string]Animal)

  fmt.Println("Enter a command [newanimal|query], name and animal|action [cow|bird|snake][eat|move|speak]")

  for {
    fmt.Print("> ")
    fmt.Scan(&cmd, &name, &attribute)

    var animal Animal

    switch cmd {
    case "newanimal":
      switch attribute {
      case "cow":
        animals[name] = Cow{food: "grass", locomotion: "walk", noise: "moo"}
      case "bird":
        animals[name] = Bird{food: "worms", locomotion: "fly", noise: "peep"}
      case "snake":
        animals[name] = Snake{food: "mice", locomotion: "slither", noise: "hsss"}
      default:
        fmt.Println("I don't recognise that animal")
      }
    case "query":
      animal = animals[name]
      switch attribute {
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
    default:
      fmt.Println("I don't recognise that command")
      continue
    }
  }
}

func (c Cow) Eat() {
  fmt.Println(c.food)
}

func (c Cow) Move() {
  fmt.Println(c.locomotion)
}

func (c Cow) Speak() {
  fmt.Println(c.noise)
}

func (b Bird) Eat() {
  fmt.Println(b.food)
}

func (b Bird) Move() {
  fmt.Println(b.locomotion)
}

func (b Bird) Speak() {
  fmt.Println(b.noise)
}

func (s Snake) Eat() {
  fmt.Println(s.food)
}

func (s Snake) Move() {
  fmt.Println(s.locomotion)
}

func (s Snake) Speak() {
  fmt.Println(s.noise)
}

type Animal interface {
  Eat()
  Move()
  Speak()
}

type Cow struct {
  food, locomotion, noise string
}

type Bird struct {
  food, locomotion, noise string
}

type Snake struct {
  food, locomotion, noise string
}
