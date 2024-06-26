//Программа(CLI-command line interface),которая умеет:
//1. Записывать рекорд в текстовый файл
//2. Читать все рекорды из файла
//3. Находить рекорд определенного юзера

//func WriteRecord(user string,score int) error

package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	defer file.Close()

	//fmt.Printf("%s;%d\n", "user", 546473)
	_, err = fmt.Fprintf(file, "%s;%d\n", "user", 546473)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	fmt.Println("End" + " " + "User")
}
