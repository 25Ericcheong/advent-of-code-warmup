package day01_secretentrance

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Dial struct {
	Count           int
	PasswordCounter int
}

const minDialCount = 0
const maxDialCount = 99

const normalizeCount = 100

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
		log.Fatalf("Error opening file - %s", err.Error())
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatalf("Error closing file - %s", err.Error())
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning input - %s", err.Error())
	}
}
