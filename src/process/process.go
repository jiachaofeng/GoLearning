package process

import (
	"util"
	"fmt"
	"strconv"
)

//ChanProcess files
func ChanProcess() error{

	c := make(chan int)
	defer close(c)

	var filepathOrigin = "../../bin/testCopy.txt"
	var filepathDest = "../../bin/NewDir/testCopied"

	for i := 0; i < 20; i++{
		fmt.Println("Chan ", strconv.Itoa(i) ," start:")
		go ChanCopy(filepathOrigin,filepathDest,c,i)
	}

	for i:=0; i < 20; i++ {
		v := <-c
		fmt.Println("Chan task ", v," done")
	}

	return nil
}

//ChanCopy files
func ChanCopy(filepathOrigin string,filepathDest string,c chan int,intVal int){
	filepathDest = filepathDest + strconv.Itoa(intVal) + ".txt"
	util.Copy(filepathOrigin,filepathDest)
	c <- intVal
}