//Программа(CLI-command line interface),которая умеет:
//1. Записывать рекорд в текстовый файл
//2. Читать все рекорды из файла
//3. Находить рекорд определенного юзера

//func WriteRecord(user string,score int) error

package main

import (
	"fmt"
)

func main() {

	WriteRecord("venya", 50)
	WriteRecord("venya", 504444)

	fmt.Println("End of adding records")

	err := ReadRecord("test.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("End of printing file")
}
