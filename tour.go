package main

import (
	"fmt"
	"time"

	"github.com/mitsiu-carreno/go-tour/gobasics"
	"github.com/mitsiu-carreno/go-tour/stringutil"
)

func main() {
	fmt.Printf("The time is %v \n", time.Now())
	fmt.Println(stringutil.Reverse("evol ym yraK"))
	fmt.Println(gobasics.NakedReturn("kary", "<", 3))
}
