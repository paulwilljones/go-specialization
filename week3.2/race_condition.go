/*
	Race conditions arise in software when an application depends on the sequence or timing of processes or threads for it to operate properly.
	Race conditions have a reputation of being difficult to reproduce and debug, since the end result is nondeterministic and depends on the relative timing between interfering threads.
	Problems occurring in production systems can therefore disappear when running in debug mode, when additional logging is added, or when attaching a debugger, often referred to as a "Heisenbug".
	It is therefore better to avoid race conditions by careful software design rather than attempting to fix them afterwards.

	A race condition occurrs in this instance, when the goroutines changeX and changeY are called but the second calculation of x+y executes before the completion of both goroutines, thus x+y yields a nondeterministic value
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	x, y := 0, 0

	go changeX(&x)
	go changeY(&y)

	fmt.Printf("x+y = %d\n", x + y)
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	fmt.Printf("x+y = %d\n", x + y)
}

func changeX(x *int) {
  fmt.Println("Changing x")
  time.Sleep(1 * time.Second)
  *x++
}

func changeY(y *int) {
  fmt.Println("Changing y")
  time.Sleep(2 * time.Second)
  *y--
}
