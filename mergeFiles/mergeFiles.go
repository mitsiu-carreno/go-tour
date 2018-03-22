package mergeFiles

import (
	"io/ioutil"
	"os"
	"log"
	"io"
	"encoding/csv"
	_"bytes"
	_"fmt"
)

func check(e error){
	if e != nil{
		panic(e)
	}
}

func processCSV(record io.Reader) (ch chan []string){
	ch = make(chan []string, 10)
	go func(){
		reader := csv.NewReader(record)
		_, err := reader.Read(); 
		check(err)

		defer close(ch)
		for{
			record, err := reader.Read()
			if err == io.EOF{
				break
			}
			check(err)
			ch <- record
		}
	}()
	return
}

// MergeUtil read a slice of files and joins in content on a new file
func MergeUtil(){
//func MergeUtil(_filename string, _files[]string){

	outfile, err := os.Create(outPath + "merge.csv")
	check(err)
	defer outfile.Close()

	writter := csv.NewWriter(outfile)

	err = writter.Write([]string{"OBSERVACIONES","INDICE","NOMBRE","DEPENDENCIA","DECLARACION","FECHA","ACUSE","TEMA","SUBTEMA","VALOR"})
	check(err)

	fileList, err := ioutil.ReadDir(inputPath)
	check(err)

	for _, allFiles := range fileList{
		file, err := os.Open(inputPath + allFiles.Name())
		check(err)
		defer file.Close()


		for rec := range processCSV(file){
			if err := writter.Write(rec); err != nil{
				log.Fatal(err)
			}
		}
		writter.Flush()
		if err := writter.Error(); err != nil{
			log.Fatal(err)
		}
	}
}

const inputPath = "mergeFiles/assets/"
const outPath = "mergeFiles/output/"