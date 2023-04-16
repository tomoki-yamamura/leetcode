package main

import (
	"archive/zip"
)

func main() {
	出力先のファイルを作成
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

	// reader, err := zip.OpenReader("example.zip")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer reader.Close()

	// for _, file := range reader.File {
	// 	fmt.Println(file.Name)
	// 	openedFile, err := file.Open()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer openedFile.Close()
	// 	fileContents, err := ioutil.ReadAll(openedFile)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(string(fileContents))
	// }

}