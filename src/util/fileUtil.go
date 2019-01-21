package util

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// Copy use to copy single file
func Copy(filepathOrigin, filepathDest string) error {
	srcFile, _ := os.Open(filepathOrigin)
	defer srcFile.Close()

	destFile, _ := os.Create(filepathDest)
	defer destFile.Close()

	_, err := io.Copy(destFile, srcFile)

	//TO DO: Below code has performance issue, comment fisrt
	// err := destFile.Sync()

	return err
}

//Link to reates newname as a hard link to the oldname file.
//Can copy file without updating create/modified time
func Link(filepathOrigin, filepathDest string) error {
	err := os.Link(filepathOrigin, filepathDest)
	return err
}

//CopyFiles use to copy multiply files in folder
func CopyFiles(filepathOrigin, filepathDest string, timeStart, timeEnd time.Time) error {

	log.Println("CopyFiles() start. ", filepathOrigin, filepathDest, timeStart, timeEnd)
	filepathDest = FormatPath(filepathDest)

	filepath.Walk(filepathOrigin, func(pathName string, info os.FileInfo, err error) error {

		if !info.IsDir() {
			modTime := info.ModTime()
			if modTime.Before(timeStart) || modTime.After(timeEnd.AddDate(0, 1, 0)) {

			} else {
				fileName := filepath.Base(pathName)
				year := strconv.Itoa(modTime.Year())
				month := strconv.Itoa(int(modTime.Month()))
				destDir := filepathDest + "/" + year + "/" + month + "/"
				destFileName := destDir + fileName
				exist, _ := PathExists(destFileName)
				if exist {
					log.Println("File exist, will skip this folder copied " + destDir)
					return filepath.SkipDir
				}

				log.Println("Copy(): ", modTime, fileName, destFileName)
				// err = Copy(pathName,destFileName)
				//Need to keep file createtime and modified time
				err = Link(pathName, destFileName)

				if err != nil {
					log.Println(err)
				} else {
					RemoveFile(pathName)
				}
			}
		}
		return err
	})
	log.Println("CopyFiles() end.")
	return nil
}

//PathExists use to check if file of file folder exist
func PathExists(path string) (bool, error) {
	var err error
	var exist = true
	if _, err = os.Stat(path); os.IsNotExist(err) {
		//not exists
		exist = false
		err = nil
	}
	return exist, err
}

//MkDirAllIfPathNotExit use to create make dir in path not exist
func MkDirAllIfPathNotExit(path string) error {

	exist, err := PathExists(path)

	if exist {
		log.Println(path + " exist")
		return err
	}

	err = os.MkdirAll(path, os.ModePerm)
	return err
}

//RemoveFile use to remove single file
func RemoveFile(filepath string) error {
	exist, err := PathExists(filepath)
	if !exist {
		log.Println(filepath + " not exist")
		return err
	}

	err = os.RemoveAll(filepath)

	return err
}

//FormatPath to format path separator '\\' and '/'
func FormatPath(filepath string) string {

	for len(filepath) > 0 && os.IsPathSeparator(filepath[len(filepath)-1]) {
		filepath = filepath[0 : len(filepath)-1]
	}
	return filepath
}

//ZipFile to zip file to target path
func ZipFile(source, target string) error {

	zipFile, err := os.Create(target)

	if err != nil {
		log.Println(err)
		return err
	}
	defer zipFile.Close()

	archive := zip.NewWriter(zipFile)

	defer archive.Close()

	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)

		if err != nil {
			return err
		}

		header.SetModTime(time.Unix(info.ModTime().Unix(), 0))

		if info.IsDir() {
			return nil
		}

		header.Name = path[len(source):len(path)]
		header.Method = zip.Deflate

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)

		return err
	})
}
