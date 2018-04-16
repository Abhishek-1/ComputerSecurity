package main 

import (
	"ethos/syscall"
	"log"
	"ethos/altEthos"
	"strconv"
)

func main () {

	populate( 5)
	find(2)	
}

func populate( count int){

	me := syscall.GetUser()

	path := "/user/" + me + "/boxes"
	
	for i:= 0; i<count; i++ {

	_, status := altEthos.DirectoryOpen(path)
	if status != syscall.StatusOk {
	   log.Fatalf ("Error opening %v %v\n", path, status)
	}
	
	istr := strconv.Itoa(i)
	name := "/box" + istr

	finaldata := MyType { int32(i), int32(i+5), int32(i+10), int32(i+15)}

	var readData MyType

	status = altEthos.Write(path + name, &finaldata)
	if status != syscall.StatusOk {
	   log.Fatalf ("Error Writing %v/boxes %v\n", path, status)
	}

	status = altEthos.Read(path + name, &readData)
	if status != syscall.StatusOk {
	   log.Fatalf ("Error Reading %v/boxes %v\n", path, status)
	}

	log.Printf("Value at %v/boxes was %v\n", path, readData) 

}

}	

func find( point int){
	me := syscall.GetUser()

	path := "/user/" + me + "/boxes"
	_, status := altEthos.DirectoryOpen(path)
	if status != syscall.StatusOk {
	   log.Fatalf ("Error opening %v /boxes %v\n", path, status)
	}

	_,status = altEthos.SubFiles(path)

	log.Println("Searching boxes for Point: ",point)

		
}
