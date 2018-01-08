package main

import (

"ethos/altEthos"
"ethos/syscall"
"ethos/log"
"ethos/kernelTypes"
"strings"
"strconv"
)

var logger = log.Initialize("test/myRpcClient")

func init() {
SetupMyRpcPrintAssignmentReply(printAssignmentReply)
SetupMyRpcBookAssignmentReply(bookAssignmentReply)
}


func printAssignmentReply(test Assignmentlist) (MyRpcProcedure) {
logger.Printf("Step 4")
	 logger.Printf("myRpcClient:ReceivedprintAssignmentReply:%v\n")

        data1 := test.Field1
       	data2 := test.Field2
	data3 := test.Field3
	data4 := test.Field4
	data5 := test.Field5
	data6 := test.Field6
	data7 := test.Field7
        file1 := data1.Field1
        file2 := data2.Field1
	file3 := data3.Field1
	file4 := data4.Field1
	file5 := data5.Field1
	file6 := data6.Field1
	file7 := data7.Field1
        pos1 := data1.Field2
        pos2 := data2.Field2
 	pos3 := data3.Field2
        pos4 := data4.Field2
	pos5 := data5.Field2
        pos6 := data6.Field2
	pos7 := data7.Field2

        logger.Printf("Details for "+file1+"\t \t \t \t Total Seats left:%v\n", pos1 )
        logger.Printf("Details for "+file2+"\t \t \t \t Total Seats left:%v\n", pos2 )
	logger.Printf("Details for "+file3+"\t \t \t \t Total Seats left:%v\n", pos3 )
	logger.Printf("Details for "+file4+"\t \t \t \t Total Seats left:%v\n", pos4 )
	logger.Printf("Details for "+file5+"\t \t \t \t Total Seats left:%v\n", pos5 )
	logger.Printf("Details for "+file6+"\t \t \t \t Total Seats left:%v\n", pos6 )
	logger.Printf("Details for "+file7+"\t \t \t \t Total Seats left:%v\n", pos7 )

	               
        str1 := "\nDetails For: "+file1+"\t \tTotal Seats Left:" + strconv.Itoa(int(pos1)) + "\t (To Select, Enter 1)\n"
        str2 := "Details For: "+file2+"\tTotal Seats Left:" + strconv.Itoa(int(pos2)) + "\t (To Select, Enter 2)\n"
	str3 := "Details For: "+file3+"\tTotal Seats Left:" + strconv.Itoa(int(pos3)) + "\t (To Select, Enter 3)\n"
	str4 := "Details For: "+file4+"\tTotal Seats Left:" + strconv.Itoa(int(pos4)) + "\t (To Select, Enter 4)\n"
	str5 := "Details For: "+file5+"\tTotal Seats Left:" + strconv.Itoa(int(pos5)) + "\t (To Select, Enter 5)\n"
	str6 := "Details For: "+file6+"\tTotal Seats Left:" + strconv.Itoa(int(pos6)) + "\t (To Select, Enter 6)\n"
	str7 := "Details For: "+file7+"\t \tTotal Seats Left:" + strconv.Itoa(int(pos7)) + "\t (To Select, Enter 7)\n"
        
        var strk1 kernelTypes.String
        var strk2 kernelTypes.String
	var strk3 kernelTypes.String
	var strk4 kernelTypes.String
	var strk5 kernelTypes.String
	var strk6 kernelTypes.String
	var strk7 kernelTypes.String
        strk1 = kernelTypes.String(str1)
        strk2 = kernelTypes.String(str2) 
	strk3 = kernelTypes.String(str3) 
	strk4 = kernelTypes.String(str4) 
	strk5 = kernelTypes.String(str5) 
	strk6 = kernelTypes.String(str6) 
	strk7 = kernelTypes.String(str7) 

        statusW := altEthos.WriteStream(syscall.Stdout, &strk1)
        if statusW != syscall.StatusOk {
                logger.Printf("Error while writing syscall.Stdout: %v", statusW)
        }
        statusW = altEthos.WriteStream(syscall.Stdout, &strk2)
        if statusW != syscall.StatusOk {
                logger.Printf("Error while writing syscall.Stdout: %v", statusW)
        }
	statusW = altEthos.WriteStream(syscall.Stdout, &strk3)
		if statusW != syscall.StatusOk {
		        logger.Printf("Error while writing syscall.Stdout: %v", statusW)
		}
	statusW = altEthos.WriteStream(syscall.Stdout, &strk4)
		if statusW != syscall.StatusOk {
		        logger.Printf("Error while writing syscall.Stdout: %v", statusW)
		}
	statusW = altEthos.WriteStream(syscall.Stdout, &strk5)
		if statusW != syscall.StatusOk {
		        logger.Printf("Error while writing syscall.Stdout: %v", statusW)
		}
	statusW = altEthos.WriteStream(syscall.Stdout, &strk6)
		if statusW != syscall.StatusOk {
		        logger.Printf("Error while writing syscall.Stdout: %v", statusW)
		}
	statusW = altEthos.WriteStream(syscall.Stdout, &strk7)
		if statusW != syscall.StatusOk {
		        logger.Printf("Error while writing syscall.Stdout: %v", statusW)
		}

        return nil

}

func bookAssignmentReply(cnf uint64) (MyRpcProcedure) {
	logger.Printf("Step 8")
	 var str kernelTypes.String
        logger.Printf("myRpcClient:ReceivedConfirmationReply:%v\n",cnf)
        if cnf == 1 {
                logger.Printf("\nAssignment Registered\n");
                str = "\nAssignment Registered\n"
        } else if cnf == 99 {
                logger.Printf("\nSorry, Not able to Register the Assignment\n");
                str = "\nSorry, Not able to Register the Assignment\n"
        }
        statusW := altEthos.WriteStream(syscall.Stdout, &str)
        if statusW != syscall.StatusOk {
                logger.Printf("Error while writing syscall.Stdout: %v", statusW)
        }
        

        return nil

}

func main() {
logger.Printf("myRpcClient:beforecall\n")

	//var myReader kernelTypes.String
	//statusR := altEthos.ReadStream(syscall.Stdin, &myReader)
		//if statusR != syscall.StatusOk {
		        //logger.Printf("Error while reading syscall.Stdin: %v", statusR)
		//}


		//statusW := altEthos.WriteStream(syscall.Stdout, &myReader)
		//if statusW != syscall.StatusOk {
		        //logger.Printf("Error while writing syscall.Stdout: %v", statusW)
		//}
	fd,status := altEthos.IpcRepeat("myRpc","",nil)
	if status != syscall.StatusOk {
	logger.Printf("Ipcfailed:%v\n",status)
	altEthos.Exit(status)
	}                                                                                                                                       

	call1 := MyRpcPrintAssignment{}
	status = altEthos.ClientCall(fd,&call1)
	if status != syscall.StatusOk {
	logger.Printf("clientCallfailed:%v\n",status)
	}

	var strnew kernelTypes.String
	var strnew1 kernelTypes.String 
		strnew = "Select choice\n"       
		statusW := altEthos.WriteStream(syscall.Stdout, &strnew)
		if statusW != syscall.StatusOk {
		        logger.Printf("Error while writing syscall.Stdout: %v", statusW)
		}


	statusR := altEthos.ReadStream(syscall.Stdin, &strnew1)
		if statusR != syscall.StatusOk {
		        logger.Printf("Error while reading syscall.Stdin: %v", statusR)
		}

	//statusW = altEthos.WriteStream(syscall.Stdout, &strnew1)
		//if statusW != syscall.StatusOk {
		        //logger.Printf("Error while writing syscall.Stdout: %v", statusW)
		//}

	var callvar uint64
	var recvdch string

	recvdch = string(strnew1)

	if strings.TrimRight(recvdch, "\n") == "1" {
		callvar = 1
	} else if strings.TrimRight(recvdch, "\n") == "2" {
		callvar = 2
	} else if strings.TrimRight(recvdch, "\n") == "3" {
		callvar = 3
	}else if strings.TrimRight(recvdch, "\n") == "4" {
		callvar = 4
	}else if strings.TrimRight(recvdch, "\n") == "5" {
		callvar = 5
	}else if strings.TrimRight(recvdch, "\n") == "6" {
		callvar = 6
	}else if strings.TrimRight(recvdch, "\n") == "7" {
		callvar = 7
	}else {
		callvar = 99
	}
		 

	fd,status = altEthos.IpcRepeat("myRpc","",nil)
	if status != syscall.StatusOk {
	logger.Printf("Ipcfailed:%v\n",status)
	altEthos.Exit(status)
	} 

	call2 := MyRpcBookAssignment{callvar}
		status = altEthos.ClientCall(fd,&call2)
		if status != syscall.StatusOk {
		        logger.Printf("clientCallfailed:%v\n",status)

		}

	fd,status = altEthos.IpcRepeat("myRpc","",nil)
	if status != syscall.StatusOk {
	logger.Printf("Ipcfailed:%v\n",status)
	altEthos.Exit(status)
	}

	call3 := MyRpcPrintAssignment{}
	status = altEthos.ClientCall(fd,&call3)
	if status != syscall.StatusOk {
	logger.Printf("clientCallfailed:%v\n",status)
	} 

		
	
altEthos.Exit(status)
logger.Printf("myRpcClient:done\n")

}
