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
	srcFolder  = flag.String("src", "C:/Users/c.jia/InvoiceFiles", "Source folder")
	tarFolder  = flag.String("target", "C:/Users/c.jia/InvoiceFiles", "Target folder")
	startTimeF = flag.String("st", "System Date - 1 month", "Execution start time (YYYY-MM)")
	endTimeF   = flag.String("et", "System Date", "Execution end time (YYYY-MM)")
)

var timeStart time.Time
var timeEnd time.Time

func main() {

	flag.Parse()
	layout := "2006-01"
	if *startTimeF == "System Date - 1 month" {
		curTime := time.Now()
		year, month, _ := curTime.Date()
		timeStart = time.Date(year, month+time.Month(-1), 1, 0, 0, 0, 0, curTime.Location())
	} else {
		timeStart, _ = time.ParseInLocation(layout, *startTimeF, time.Local)
	}

	if *endTimeF == "System Date" {
		timeEnd = time.Now()
	} else {
		timeEnd, _ = time.ParseInLocation(layout, *endTimeF, time.Local)
	}

	initLog()
	initFolder(*tarFolder, timeStart, timeEnd)

	startTime := time.Now()
	log.Println("Copy Files start： ", startTime.String())
	util.CopyFiles(*srcFolder, *tarFolder, timeStart, timeEnd)
	elapsed := time.Since(startTime)
	log.Println("Copy Files elapsed: ", elapsed)

	startTime = time.Now()
	log.Println("Zipped Files start : ", startTime.String())
	zippedFiles(*tarFolder, timeStart, timeEnd)
	elapsed = time.Since(startTime)
	log.Println("Zipped Files end : ", startTime.String())
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

func initFolder(destPath string, timeStart, timeEnd time.Time) {
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

func zippedFiles(destPath string, timeStart, timeEnd time.Time) {
	var year, month string
	for {
		timeEnd = timeEnd.AddDate(0, -1, 0)

		year = strconv.Itoa(timeEnd.Year())
		month = strconv.Itoa(int(timeEnd.Month()))
		source := util.FormatPath(destPath) + "/" + year + "/" + month
		dest := util.FormatPath(destPath) + "/" + year + "/" + month + ".zip"
		log.Println("ZipFile() ", source, dest)
		util.ZipFile(source, dest)
		util.RemoveFile(source)

		if timeStart.Before(timeEnd) {
			break
		}
	}
}
