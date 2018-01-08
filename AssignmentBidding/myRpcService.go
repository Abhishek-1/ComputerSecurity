package main
import (
	"ethos/syscall"
	"ethos/altEthos"
	"ethos/log"
	)

var logger = log.Initialize("test/myRpcService")

func init() {
	SetupMyRpcPrintAssignment(printAssignment)
	SetupMyRpcBookAssignment(bookAssignment)
	}

func printAssignment () (MyRpcProcedure) {

logger.Printf("Step 3")
logger.Printf("myRpcService:calledAssignmentlistPrint\n")
me := syscall.GetUser()
path := "/user/" + me + "/assignment"
filenm := "/assign"
var readData Assignmentlist
_, status := altEthos.DirectoryOpen(path)
if status != syscall.StatusOk {
   logger.Printf ("Error opening %v %v\n", path, status)
}
status = altEthos.Read(path + filenm, &readData)
if status != syscall.StatusOk {
  logger.Printf ("Error Reading %v"+filenm+" %v\n", path, status)
}


return &MyRpcPrintAssignmentReply{readData}
}

func bookAssignment(id uint64) (MyRpcProcedure) {
logger.Printf("myRpcService:calledAssignmentbooking\n")
logger.Printf("Step 7")
me := syscall.GetUser()
path := "/user/" + me + "/assignment"
filenm := "/assign"
var readData Assignmentlist
_, status := altEthos.DirectoryOpen(path)
if status != syscall.StatusOk {
   logger.Printf ("Error opening %v %v\n", path, status)
}
status = altEthos.Read(path + filenm, &readData)
if status != syscall.StatusOk {
  logger.Printf ("Error Reading %v"+filenm+" %v\n", path, status)
}

datain1 := readData.Field1
datain2 := readData.Field2
datain3 := readData.Field3
datain4 := readData.Field4
datain5 := readData.Field5
datain6 := readData.Field6
datain7 := readData.Field7


id1 := datain1.Field0
id2 := datain2.Field0
id3 := datain3.Field0
id4 := datain4.Field0
id5 := datain5.Field0
id6 := datain6.Field0
id7 := datain7.Field0
pos1 := datain1.Field2
pos2 := datain2.Field2
pos3 := datain3.Field2
pos4 := datain4.Field2
pos5 := datain5.Field2
pos6 := datain6.Field2
pos7 := datain7.Field2


var conf uint64
logger.Printf("Input received... %v",id)

if int(id) == int(id1) && int(pos1) != 0 {
        conf = 1
        pos1 = pos1 - 1
} else if int(id) == int(id2) && int(pos2) != 0 {
        conf = 1
        pos2 = pos2 - 1
} else if int(id) == int(id3) && int(pos3) != 0 {
        conf = 1
        pos3 = pos3 - 1
} else if int(id) == int(id4) && int(pos4) != 0 {
        conf = 1
        pos4 = pos4 - 1
} else if int(id) == int(id5) && int(pos5) != 0 {
        conf = 1
        pos5 = pos5 - 1
} else if int(id) == int(id6) && int(pos6) != 0 {
        conf = 1
        pos6 = pos6 - 1
} else if int(id) == int(id7) && int(pos7) != 0 {
        conf = 1
        pos7 = pos7 - 1
} else {
        conf = 99
}
dataout1 := Assignment { 1,"Remote File Copy", pos1}
dataout2 := Assignment { 2,"Appointment Calendar", pos2}
dataout3 := Assignment { 3,"Password Encryption", pos3}
dataout4 := Assignment { 4,"Assignment Submission", pos4}
dataout5 := Assignment { 5,"Authorization System", pos5}
dataout6 := Assignment { 6,"Bell_LaPadula System", pos6}
dataout7 := Assignment { 7,"Approvability", pos7}

finaldata1 := Assignmentlist { dataout1, dataout2, dataout3, dataout4, dataout5, dataout6, dataout7}
status1 := altEthos.Write(path + filenm, &finaldata1)
if status1 != syscall.StatusOk {
	  logger.Printf ("Error Writing to %v"+filenm+" %v\n", path, status1)
}

return &MyRpcBookAssignmentReply{conf}

}

func main (){
	me := syscall.GetUser()
	path := "/user/" + me + "/assignment"
	filename := "/assign"
	_, statusO := altEthos.DirectoryOpen(path)
        if statusO != syscall.StatusOk {
           logger.Printf ("Error opening %v %v\n", path, statusO)
        }
	dataout1 := Assignment { 1,"Remote File Copy", 7}
	dataout2 := Assignment { 2,"Appointment Calendar", 5}
	dataout3 := Assignment { 3,"Password Encryption", 2}
	dataout4 := Assignment { 4,"Assignment Submission", 5}
	dataout5 := Assignment { 5,"Authorization System", 7}
	dataout6 := Assignment { 6,"Bell_LaPadula System", 3}
	dataout7 := Assignment { 7,"Approvability", 5}
	finaldata := Assignmentlist { dataout1, dataout2, dataout3, dataout4, dataout5, dataout6, dataout7}
	status1 := altEthos.Write(path + filename, &finaldata)
	if status1 != syscall.StatusOk {
	   logger.Printf ("Error Writing to %v"+filename+" %v\n", path, status1)
	}
	listeningFd, status := altEthos.Advertise("myRpc")
	if status != syscall.StatusOk {
	logger.Printf("Advertisingservicefailed:%s\n",status)
	altEthos.Exit(status)
}

for {
	_,fd,status := altEthos.Import(listeningFd)
	if status != syscall.StatusOk {
	logger.Printf("ErrorcallingImport:%v\n",status)
	altEthos.Exit(status)
	}
	logger.Printf("myRpcService:newconnectionaccepted\n")
	t:=MyRpc{}
	altEthos.Handle(fd,&t)
	}
}
