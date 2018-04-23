package readFile

import (
	"os"
	"log"
	"bufio"
	"io"
	"strconv"
)
func Sort(infile string)(values []int,err error) {
	file, err := os.Open(infile)
	if err != nil {
		log.Printf("打开文件失败%s%s", infile,err)
	}
	defer file.Close()
	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			log.Println(" too long")
			return
		}
		str := string(line)
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}
		values = append(values, value)
	}

	return
}
