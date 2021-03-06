package main

import (
	_"log"
	"fmt"
	"time"
	_"io/ioutil"

	"github.com/mitsiu-carreno/go-tour/gobasics/basics"
	"github.com/mitsiu-carreno/go-tour/gobasics/flowControl"
	"github.com/mitsiu-carreno/go-tour/gobasics/moreTypes"
	"github.com/mitsiu-carreno/go-tour/gobasics/methodsInterfaces/methods"
	"github.com/mitsiu-carreno/go-tour/gobasics/methodsInterfaces/interfaces"
	"github.com/mitsiu-carreno/go-tour/gobasics/goroutines"

	"github.com/mitsiu-carreno/go-tour/concurrencyPatterns"

	"github.com/mitsiu-carreno/go-tour/stringutil"
	_"github.com/mitsiu-carreno/go-tour/ziputil"
	"github.com/mitsiu-carreno/go-tour/mergeFiles"
)

const (
	// Big Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Small Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func goBasics() {
	fmt.Printf("The time is %v \n", time.Now())
	fmt.Println(stringutil.Reverse("evol ym yraK"))

	// ZipFiles
	/*
	files := []string{}
	filesList, err := ioutil.ReadDir("/Users/jorandradefig/Desktop/Projects/DataCivica/1560000.org/parser/output/csv/")
	if err != nil {
		log.Fatal(err);
	}

	for _, f := range filesList{
		files = append(files, "/Users/jorandradefig/Desktop/Projects/DataCivica/1560000.org/parser/output/csv/"+ f.Name())
	}
	output := "ziputil/go-zip.zip"

	log.Println("Start zipping: ", output)
	err = ziputil.ZipFiles(output, files)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Finish zipping: ", output)
	*/

	// Basics
	fmt.Printf("\n----BASICS----\n")
	fmt.Println(gobasics.NakedReturn("kary", "<", 3))
	fmt.Println()	
	f, u := gobasics.TypeConversions(1210)
	fmt.Printf("1210 int -> %v of type %T ->  %v of type %T \n", f, f, u, u)
	fmt.Println()	
	gobasics.TypeInterface()
	fmt.Println("Constant Small ", Small)
	fmt.Println("Constant Big ", complex64(Big))

	// Flow Control
	
	fmt.Printf("\n----FLOW CONTROL----\n")
	flowControl.ForStruct(3)
	fmt.Println()
	flowControl.IfStruct(3)
	fmt.Println()
	flowControl.SwitchStruct()
	fmt.Println()
	flowControl.DeferStruct()

	// More Types
	fmt.Printf("\n----MORE TYPES----\n")
	moreTypes.Pointers()
	fmt.Println()
	moreTypes.Structs()
	fmt.Println()
	moreTypes.ArrayStruct()
	fmt.Println()	
	moreTypes.SliceStruct()
	fmt.Println()
	moreTypes.LoopingArray()
	fmt.Println()
	moreTypes.MapsStruct()
	fmt.Println()
	moreTypes.FunctionValues()
	fmt.Println()
	moreTypes.Closures()
	fmt.Println()
	moreTypes.Fi()

	// Methods and interfaces
	fmt.Printf("\n----Methods & Interfaces----\n")
	methods.BasicMethod()
	fmt.Println()
	methods.NonStruct()
	fmt.Println()
	methods.PointerReceiver()
	fmt.Println()
	interfaces.BasicInterface()
	fmt.Println()
	interfaces.ImplementImplicit()
	fmt.Println()
	interfaces.InterfaceValues()
	fmt.Println()
	interfaces.InterfaceNil()
	fmt.Println()
	interfaces.EmptyInterface()
	fmt.Println()
	interfaces.TypeAssertions()
	fmt.Println()
	interfaces.TypeSwitch()
	fmt.Println()
	interfaces.Stringers()
	fmt.Println()
	interfaces.Errors()
	fmt.Println()
	interfaces.Readers()

	// Goroutines
	fmt.Printf("\n----Goroutines----\n")
	goroutines.BasicGoroutine()
	fmt.Println()
	goroutines.BasicChannels()
}

func goConcurrency(){
	//concurrencyPatterns.Concurrency()
	//concurrencyPatterns.Channel()
	concurrencyPatterns.FunctionReturnChannel()
	concurrencyPatterns.Multiplexing()
}

func main(){
	//goBasics()
	//goConcurrency()
	mergeFiles.MergeUtil()
}