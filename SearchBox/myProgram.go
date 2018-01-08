package main 

import (
	"ethos/syscall"
	"log"
	"ethos/altEthos"
	"fmt"
	"strconv"
	)

func main () {
	

	//This program has 2 functions, Populate( count int), which will create random number of boxes with lower left point (x,y) 
	//and upper right point (x,y). All the points and file (box) name is generated randomly.
	//Function find(pt box.point) goes through directory /user/nobody/boxes and retreives all boxes in that directory. 
	//Find function searchs for a randomly generated point and print to logs boxes containing that point.
   	
	count,_ := Random(1)
	countval := int(count[0])
	
	//Calling populate function with random count
	populate( countval)
	searchindex,_ := Random(2)
	for i:= 0; i<2; i++ {		
		for j := i+1; j<2; j++ {
			if searchindex[i] > searchindex[j] {
				temp := searchindex[j]
				searchindex[j] = searchindex[i]
				searchindex[i] = temp
			}
		}
	}
	xpt := int32(searchindex[0])
	ypt := int32(searchindex[1])	
	pt := Point {xpt,ypt }
	//Calling find function for a random point	
	find( pt)
	log.Printf("Return from function populate")
	
			
}

func populate( count int){

	me := syscall.GetUser()

	path := "/user/" + me + "/boxes"
	
	var countpop int

	//Creating n number of files with lower left point and upper right point

	for k:= 0; k<count; k++ {
	
	countpop = countpop + 1
	
	_, status := altEthos.DirectoryOpen(path)
	if status != syscall.StatusOk {
	   log.Fatalf ("Error opening %v %v\n", path, status)
	}
	
	filename := "/box"
	//Random function to make filename unique and random
	itert,_ := Random(3)
	itertnm1 := int(itert[0])
	itertnm2 := int(itert[1])
	itertnm3 := int(itert[2])
	//Converting to string for appending in filename	
	ident1 := strconv.Itoa(itertnm1)
	ident2 := strconv.Itoa(itertnm2)
	ident3 := strconv.Itoa(itertnm3)
	filename = filename + ident1 + ident2 + ident3
	//Random function to get {x,y} for 
	numbe,_ := Random(4)
	fmt.Println(numbe[0],numbe[1],numbe[2],numbe[3])

	for i:= 0; i<4; i++ {		
		for j := i+1; j<4; j++ {
			if numbe[i] > numbe[j] {
				temp := numbe[j]
				numbe[j] = numbe[i]
				numbe[i] = temp
			}
		}
	}	
			
	var llx,lly,urx,ury int32
	llx = int32(numbe[0])
	lly = int32(numbe[1])
	urx = int32(numbe[2])
	ury = int32(numbe[3])
	//Creating Point struct
	data := Point {llx,lly }
	data1 := Point {urx,ury }
	//Creating struct with 2 points to store
	finaldata := MyType { data, data1}
	var readData MyType

	
	status = altEthos.Write(path + filename, &finaldata)
	if status != syscall.StatusOk {
	   log.Fatalf ("Error Writing to %v"+filename+" %v\n", path, status)
	}

	status = altEthos.Read(path + filename, &readData)
	if status != syscall.StatusOk {
	   log.Fatalf ("Error Reading %v"+filename+" %v\n", path, status)
	}

	//Printing the written data	
	log.Printf("Typed value written at %v"+filename+" was %v\n", path, readData) 

	}

	log.Println("Total box created from populate function is", countpop) 	
}

func find( pt Point){

	me := syscall.GetUser()

	xfind := pt.xpoint
	yfind := pt.ypoint
	//Cond will be set as 1 when match is found
	var cond int

	path := "/user/" + me + "/boxes"
	_, status := altEthos.DirectoryOpen(path)
	if status != syscall.StatusOk {
	   log.Fatalf ("Error opening %v %v\n", path, status)
	}
	//Fetching all files from subfolder
	files,status := altEthos.SubFiles(path)
	var readData MyType
	var data Point
	var data1 Point
	length := len(files)
	log.Println("Searching boxes for Point: ", pt, " in ", length, "number of files")
	for i:=0 ; i<length; i++ {

		
		filenm := "/" + files[i]
		//Reading files and fetching data
		status = altEthos.Read(path + filenm, &readData)
		if status != syscall.StatusOk {
	   	log.Fatalf ("Error Reading %v"+filenm+" %v\n", path, status)
		}
		
	data = readData.Field1
	data1 = readData.Field2
	
	llx := data.xpoint
	lly := data.ypoint
	urx := data1.xpoint
	ury := data1.ypoint
	//if condition to check if random point matches lower left point and upper right point in iterated boxes
	if int(llx) == int(xfind) && int(lly) == int(yfind) {
		
		log.Println("Match Found for point in "+filenm+" for point", pt )			
		cond = 1
				
	} else if int(urx) == int(xfind) && int(ury) == int(yfind) {
		
		log.Println("Match Found for point in "+filenm+" for point", pt )		
		cond = 1
	}
		
	}

	if cond != 1 {
		log.Println("No match found for point", pt)
	}		
}

func Random(size uint32) (randomString []byte, status syscall.Status){
	eventId, status := syscall.Random(size)
        if status != syscall.StatusOk{
        	return
        }

	randomString, _, status = syscall.BlockAndRetire(eventId)
        return
}
