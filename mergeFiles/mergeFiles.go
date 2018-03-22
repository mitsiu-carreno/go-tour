package mergeFiles

import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
	"io"
	"encoding/csv"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	inputPath = "mergeFiles/assets/"
	outPath = "mergeFiles/output/"
)

type Pdfs struct{
	_id			string `bson:"_id,omitempty"`
	ANIO		string `bson:"ANIO"`
	ARCHIVO		string `bson:"ARCHIVO"`
}

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

	var hosts      = os.Getenv("MAIN_DB_HOST")
    var database   = os.Getenv("MAIN_DB_DB")
    var username   = os.Getenv("MAIN_DB_USER")
    var password   = os.Getenv("MAIN_DB_PASSWORD")
	var collection = os.Getenv("MAIN_DB_COLLECTION")
	
	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
        Timeout:  60 * time.Second,
        Database: database,
        Username: username,
        Password: password,
	}

	session, err := mgo.DialWithInfo(info)
	check(err)
	defer session.Close()
	 
	col := session.DB(database).C(collection)

	/*
	count, err := col.Count()
	check(err)
	fmt.Println(count)
	*/
	var mgoResult []Pdfs
	err = col.Find(bson.M{"ANIO":2017}).All(&mgoResult)
	check(err)
	fmt.Println(mgoResult)


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
