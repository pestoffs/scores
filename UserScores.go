package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type UserScores struct {
	name  string
	score int
}

func CreateNewUserScores(name string, score int) *UserScores {
	user := UserScores{
		name:  name,
		score: score,
	}
	return &user
}

func (u *UserScores) WriteRecord() error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Print(err.Error())
		return err
	}
	defer file.Close()

	//fmt.Printf("%s;%d\n", "user", 546473)
	_, err = fmt.Fprintf(file, "%s;%d\n", u.name, u.score)
	if err != nil {
		fmt.Print(err.Error())
		return err
	}
	return nil
}

func (u *UserScores) ReadRecord() error {
	file, err := os.Open(path)
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

func (u *UserScores) findUserScore() (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ";")
		if len(parts) == 2 && parts[0] == u.name {
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
