package mergeFiles

import (
	"fmt"
	"os"
)

func check(e error){
	if e != nil{
		panic(e)
	}
}
// MergeUtil read a slice of files and joins in content on a new file
func MergeUtil(){
//func MergeUtil(_filename string, _files[]string){
	f, err := os.Open("mergeFiles/assets/file1.csv")
	check(err)

	defer f.Close()

	_, err = f.Seek(83 ,0)
	check(err)
	buff := make([]byte, 20)
	_, err = f.Read(buff)
	check(err)
	//fmt.Printf("%d bytes @ %d: %s\n", n2, readPoint, string(buff))
	fmt.Println(string(buff))

}