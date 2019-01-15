package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"util"
)

var (
	logF       = flag.String("log", "application.log", "Log file name")
	srcFolder  = flag.String("src", "Z:/Backup", "Source folder")
	tarFolder  = flag.String("target", "Z:/Backup", "Target folder")
	startTimeF = flag.String("st", "2017-11", "Execution start time")
	endTimeF   = flag.String("et", "2017-12", "Execution end time")
)

func main() {

	flag.Parse()
	layout := "2006-01"
	timeStart, _ := time.ParseInLocation(layout, *startTimeF, time.Local)
	timeEnd, _ := time.ParseInLocation(layout, *endTimeF, time.Local)

	initLog()
	initFolder(*tarFolder, timeStart, timeEnd)

	startTime := time.Now()
	log.Println("App start： ", startTime.String())

	// util.CopyFiles("H:/Invoice file/Invoice file/","../../bin/NewDir/")
	// util.CopyFiles("../../bin/NewDir/"，"H:/Invoice file/Invoice file/")
	util.CopyFiles(*srcFolder, *tarFolder, timeStart, timeEnd)
	elapsed := time.Since(startTime)
	log.Println("App elapsed: ", elapsed)
}

func initLog() {
	outfile, err := os.OpenFile(*logF, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(*outfile, "open failed")
		os.Exit(1)
	}
	log.SetOutput(outfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func initFolder(destPath string, timeStart time.Time, timeEnd time.Time) {
	var year, month string
	for {
		year = strconv.Itoa(timeStart.Year())
		month = strconv.Itoa(int(timeStart.Month()))
		destFolderPath := util.FormatPath(destPath) + "/" + year + "/" + month
		util.MkDirAllIfPathNotExit(destFolderPath)
		timeStart = timeStart.AddDate(0, 1, 0)

		if timeStart.After(timeEnd) {
			break
		}
	}
}
