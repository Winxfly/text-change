package edit

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func listPathsAllFilesDir(path string) []string {

	list := make([]string, 0)
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal("[ERROR OPENING DIRECTORY] ", err)
	}

	for _, file := range files {
		ext := filepath.Ext(file.Name())
		if ext == ".txt" {
			list = append(list, path+file.Name())
		}
	}

	return list
}

func replaceTextInFile(path, oldText, newText string, logDate time.Time) string {
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal("[ERROR OPENING TXT FILE] ", err)
	}
	defer file.Close()

	fileLine := make([]byte, 64)
	var textFromFile string

	for {
		readLine, err := file.Read(fileLine)
		if err == io.EOF {
			break
		}
		textFromFile += string(fileLine[:readLine])
	}

	// prepare old and new text for comparison
	refText := strings.Split(textFromFile, "\n")
	textFromFile = strings.ReplaceAll(textFromFile, oldText, newText)
	replaceText := strings.Split(textFromFile, "\n")

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatal("[ERROR SEEK IN replaceTextInFile()] ", err)
	}

	err = file.Truncate(0)
	if err != nil {
		log.Fatal("[ERROR TRUNCATE IN replaceTextInFile()] ", err)
	}

	_, err = file.WriteString(textFromFile)
	if err != nil {
		log.Fatal("[ERROR WRITING TO TXT FILE] ", err)
	}

	// comparison and logging
	changeLines := ""
	for i, _ := range refText {
		if refText[i] != replaceText[i] {
			changeLines += strconv.Itoa(i+1) + " "
		}
	}

	if changeLines != "" {
		finalLog := "[FILE] " + path + "\n[DATE] " + logDate.Format("02 Jan 2006 15:04:05") +
			"\n[LINES CHANGED] " + changeLines +
			"\n[OLD TEXT] " + oldText + "\n[NEW TEXT] " + newText + "\n\n"

		return finalLog
	}

	return ""
}

func ReplaceTextInAllFilesDir(pathDir, oldText, newText string) {
	getwd, _ := os.Getwd()
	logDate := time.Now()
	fileName := getwd + "/log/" + logDate.Format("02-Jan-2006_15:04:05") + ".log"

	cFile, err := os.Create(fileName)
	if err != nil {
		log.Fatal("[ERROR CREATION LOG FILE] ", err)
	}
	defer cFile.Close()

	file, err := os.OpenFile(fileName, os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("[ERROR OPENING LOG FILE] ", err)
	}
	defer file.Close()

	for _, v := range listPathsAllFilesDir(pathDir) {
		_, err := file.WriteString(replaceTextInFile(v, oldText, newText, logDate))
		if err != nil {
			log.Fatal("[ERROR WRITE LOG FILE] ", err)
		}
	}
}
