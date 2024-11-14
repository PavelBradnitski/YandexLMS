package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func ReadContent(filename string) string {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return ""
	}
	data := make([]byte, 1024)
	n, err := file.Read(data)
	if err != nil {
		return ""
	}
	defer file.Close()
	return string(data[:n])
}

func LineByNum(inputFilename string, lineNum int) string {
	f, err := os.Open(inputFilename)
	if err != nil {
		return ""
	}
	fileScanner := bufio.NewScanner(f)
	var i int
	for fileScanner.Scan() {
		if i == lineNum {
			return fileScanner.Text()
		}
		i++
	}
	return ""
}
func CopyFilePart(inputFilename, outFileName string, startpos int) error {
	ifile, err := os.OpenFile(inputFilename, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer ifile.Close()
	data := make([]byte, 1024)
	n, err := ifile.Read(data)
	if err != io.EOF && err != nil {
		return err
	}
	oFile, err := os.Create(outFileName)
	if err != nil {
		return err
	}
	defer oFile.Close()
	_, err = oFile.Write(data[startpos:n])
	if err != nil {
		return err
	}
	return nil
}

func ModifyFile(filename string, pos int, val string) {
	f, _ := os.OpenFile(filename, os.O_WRONLY, 0600)
	f.Seek(int64(pos), 0) // Сместимся на 1024
	f.WriteString(val)
}

func ExtractLog(inputFileName string, start, end time.Time) ([]string, error) {
	var out []string
	layout := "02.01.2006"
	ifile, err := os.OpenFile(inputFileName, os.O_RDONLY, 0666)
	if err != nil {
		return []string{}, err
	}
	fileScanner := bufio.NewScanner(ifile)
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		index := strings.Index(txt, " ")
		if index != 0 {
			t, err := time.Parse(layout, txt[:index])
			if err != nil {
				return []string{}, err
			}
			if t.Compare(start) >= 0 && t.Compare(end) <= 0 {
				out = append(out, txt)
			}
		}
	}
	if len(out) == 0 {
		return []string{}, fmt.Errorf("zero values from log")
	}
	return out, nil
}
func main() {
	layout := "02.01.2006"
	start, _ := time.Parse(layout, "13.12.2022")
	end, _ := time.Parse(layout, "15.12.2022")
	fmt.Println(ExtractLog("file.txt", start, end))
}
