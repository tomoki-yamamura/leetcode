package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
)

var source = `1line
2line
3line`

func main() {
	// 出力先のファイルを作成
	outputFile, err := os.Create("example.zip")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// zip.Writerを作成
	zipWriter := zip.NewWriter(outputFile)
	defer zipWriter.Close()

	// 新しいファイルをzipアーカイブに追加
	filename := "example.txt"
	fileContents := "This is an example text file."
	fileWriter, err := zipWriter.Create(filename)
	if err != nil {
		panic(err)
	}
	fileReader := strings.NewReader(fileContents)
	io.Copy(fileWriter, fileReader)
}