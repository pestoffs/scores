package main

import (
	"fmt"
	"os"
)

func WriteRecord(user string, score int) {
	file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	defer file.Close()

	//fmt.Printf("%s;%d\n", "user", 546473)
	_, err = fmt.Fprintf(file, "%s;%d\n", user, score)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

}
