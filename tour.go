package main

import (
	"log"
	"fmt"
	"time"

	"github.com/mitsiu-carreno/go-tour/gobasics/basics"
	"github.com/mitsiu-carreno/go-tour/gobasics/flowControl"
	"github.com/mitsiu-carreno/go-tour/gobasics/moreTypes"

	"github.com/mitsiu-carreno/go-tour/stringutil"
	"github.com/mitsiu-carreno/go-tour/zipfiles"
)

const (
	// Big Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Small Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func main() {
	fmt.Printf("The time is %v \n", time.Now())
	fmt.Println(stringutil.Reverse("evol ym yraK"))

	// ZipFiles
	files := []string{"file1.csv","file2.csv"}
	output := "go-zip.zip"

	log.Println("Start zipping: ", output)
	err := zipfiles.ZipFiles(output, files)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Finish zipping: ", output)

	// Basics
	fmt.Printf("\n----BASICS----\n")
	fmt.Println(gobasics.NakedReturn("kary", "<", 3))
	f, u := gobasics.TypeConversions(1210)
	fmt.Printf("1210 int -> %v of type %T ->  %v of type %T \n", f, f, u, u)
	gobasics.TypeInterface()
	fmt.Println("Constant Small ", Small)
	fmt.Println("Constant Big ", complex64(Big))

	// Flow Control
	
	fmt.Printf("\n----FLOW CONTROL----\n")
	flowControl.ForStruct(3)

	flowControl.IfStruct(3)

	flowControl.SwitchStruct()

	flowControl.DeferStruct()

	// More Types
	fmt.Printf("\n----MORE TYPES----\n")
	moreTypes.Pointers()
	moreTypes.Structs()
	moreTypes.ArrayStruct()


}
