package main

import (
	"flag"
	"sorter/readFile"
	"sorter/bubbless"
	"os"
	"log"
	"strconv"
)

func main() {
	test()
}

var inputFile string
var outFile string
var action string

func test() {
	flag.StringVar(&inputFile, "i", "uSortData.data", "input file")
	flag.StringVar(&outFile, "o", "sortData.data", "out file")
	flag.Parse()
	var values bubbless.SorterData
	values, _ = readFile.Sort(inputFile)
	values.Bubbless()
	outToFile(outFile,values)

}

func outToFile(fileName string, data []int) error {
	file, err := os.Create(fileName)
	if err != nil {
		log.Println("创建文件失败")
		return err
	}
	for _, value := range data {
		_, err := file.WriteString(strconv.Itoa(value)+"\n")
		if err != nil {
			log.Println("写入文件失败")
			return err
		}
	}
	return nil
}
