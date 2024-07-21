//Программа(CLI-command line interface),которая умеет:
//1. Записывать рекорд в текстовый файл
//2. Читать все рекорды из файла
//3. Находить рекорд определенного юзера

//func WriteRecord(user string,score int) error

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// TODO: добавить использование FindUserScore
// Добав
func main() {
	venya := NewUserScore("venya", 500)
	fmt.Println(venya.ReadRecord())

	//Define CLI commands and flags
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	var userScore string
	addCmd.StringVar(&userScore, "record", "", "User and score in format user:score")

	//TODO: Parse CLI arguments
	/* if len(os.Args)<2{
		fmt.Println("expected 'add' command")
		os.Exit(1)
	} */

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "add":
			addCmd.Parse(os.Args[2:])
		default:
			flag.PrintDefaults()
			os.Exit(1)
		}
	}

	// Execute add command
	if addCmd.Parsed() {
		if userScore == "" {
			addCmd.PrintDefaults()
			os.Exit(1)
		}
		parts := strings.Split(userScore, ":")
		if len(parts) != 2 {
			fmt.Println("Invalid format. Expected user:score")
			os.Exit(1)
		}
		user := NewUserScore(parts[0], 0)
		_, err := fmt.Sscanf(parts[1], "%d", &user.score)
		if err != nil {
			fmt.Println("Invalid score format")
			os.Exit(1)
		}
		err = user.WriteRecord()
		if err != nil {
			fmt.Println("Error writing record:", err)
		} else {
			fmt.Println("Record added successfully")
		}
	}

}
