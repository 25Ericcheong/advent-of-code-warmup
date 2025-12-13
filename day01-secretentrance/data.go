package day01_secretentrance

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Dial struct {
	Count           int
	PasswordCounter int
}

func CreateDial(startCount int) Dial {
	return Dial{
		Count:           startCount,
		PasswordCounter: 0,
	}
}

const initialDialCount = 50
const minDialCount = 0
const maxDialCount = 99

const normalizeCount = 100

const LEFT = "L"
const RIGHT = "R"

func (d *Dial) TurnLeft(number int) {
	d.Count = -number

	if d.Count < minDialCount {
		d.Count += normalizeCount
	}

	d.plusPasswordCounterIfNeeded()
}

func (d *Dial) TurnRight(number int) {
	d.Count += number

	if d.Count > maxDialCount {
		d.Count %= normalizeCount
	}

	d.plusPasswordCounterIfNeeded()
}

func (d *Dial) plusPasswordCounterIfNeeded() {
	if d.Count == minDialCount {
		d.PasswordCounter += 1
	}
}

func LoadData() {
	file, err := os.Open("./day01-secretentrance/input.txt")
	if err != nil {
		log.Fatalf("error opening file - %s", err.Error())
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatalf("error closing file - %s", err.Error())
		}
	}(file)

	scanner := bufio.NewScanner(file)
	dial := CreateDial(initialDialCount)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			log.Fatal("Unexpected empty line")
		}

		valOfInterest := line[1:]
		val, err := strconv.Atoi(valOfInterest)

		if err != nil {
			log.Fatal("error converting value from file to int")
		}

		fmt.Printf("\n NEW LINE: %s", scanner.Text())
		beforeCount := dial.Count
		if line[0:1] == LEFT {
			dial.TurnLeft(val)
			fmt.Printf("\n LEFT(%v) - Count from (%v to %v)", val, beforeCount, dial.Count)
		} else if line[0:1] == RIGHT {
			dial.TurnRight(val)
			fmt.Printf("\n RIGHT(%v) - Count from (%v to %v)", val, beforeCount, dial.Count)
		} else {
			log.Fatal("unexpected value encountered while dialing")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error scanning input - %s", err.Error())
	}

	log.Println("=== SUCCESS ===")
	log.Printf("Password value found to be - %v", dial.PasswordCounter)
}
