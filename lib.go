package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const names string = "test.txt"

func WriteRecord(user string, score int) error {
	file, err := os.OpenFile(names, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Print(err.Error())
		return err
	}
	defer file.Close()

	//fmt.Printf("%s;%d\n", "user", 546473)
	_, err = fmt.Fprintf(file, "%s;%d\n", user, score)
	if err != nil {
		fmt.Print(err.Error())
		return err
	}
	return nil
}

func ReadRecord(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return scanner.Err()
}

func findUserScore(filename, user string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ";")
		if len(parts) == 2 && parts[0] == user {
			var score int
			fmt.Sscanf(parts[1], "%d", &score)
			return score, nil
		}
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return 0, fmt.Errorf("user not found")
}
