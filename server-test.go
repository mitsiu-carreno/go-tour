package main

import (
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
	"fmt"
)

type declarations struct{
	_id		string `bson:"_id,omitempty"`
	ANIO	string `bson:"ANIO"`
	ARCHIVO string `bson:"ARCHIVO"`
}

func check(e error){
	if e != nil{
		panic(e)
	}
}

func main(){
	// GOOS=linux GOARCH=amd64 go build server-test.go
	// scp -i .pem /Path/to/source /Path/to/remote/dest
	// ./server-test
	
	var hosts      = os.Getenv("MAIN_DB_HOST")
    var database   = os.Getenv("MAIN_DB_DB")
    var username   = os.Getenv("MAIN_DB_USER")
    var password   = os.Getenv("MAIN_DB_PASSWORD")
	var collection = os.Getenv("MAIN_DB_COLLECTION")
	

	info := &mgo.DialInfo{
		Addrs 		: []string{hosts},
		Timeout		: 60 * time.Second,
		Database	: database,
		Username	: username,
		Password 	: password,
	}

	session, err := mgo.DialWithInfo(info)
	check(err)
	defer session.Close()

	col := session.DB(database).C(collection)

	count, err := col.Count()
	check(err)
	fmt.Println(count, " documents found")

	var mgoResult []declarations
	err = col.Find(bson.M{"DIA":2017}).All(&mgoResult)
	check(err)
	fmt.Println(mgoResult)

}