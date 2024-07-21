package main

import (
	"bufio"
	"fmt"
	"os"
)

type UserScore struct {
	name  string
	score int
}

func NewUserScore(name string, score int) *UserScore {
	user := UserScore{
		name:  name,
		score: score,
	}
	return &user
}

func (u *UserScore) WriteRecord() error {
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

// TODO: возвращать значения структуры user и score
func (u *UserScore) ReadRecord() error {
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

// TODO: Сделать функцию(!),а не метод,которая будет находить по user строку в файле
// и возвращать в виде обьекта (используя конструктор)
func FindUserScore(name string) (*UserScore, error) {
	/* file, err := os.Open(path)
	if err != nil {
		return nil, err
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
	return 0, fmt.Errorf("user not found") */

	return nil, nil
}
