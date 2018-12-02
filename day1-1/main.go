package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Frequency struct {
	Value int
}

func NewFrequency() *Frequency {
	f := Frequency{Value: 0}

	return &f
}

func (f *Frequency) Change(c int) {
	f.Value += c
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
		return
	}

	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatalf("Error closing file: %s", err)
		}
	}()

	scanner := bufio.NewScanner(f)
	frequency := NewFrequency()

	for scanner.Scan() {
		t := scanner.Text()
		v, err := strconv.Atoi(t)
		if err != nil {
			log.Fatalf("Error converting string to int: %s", err)
		}
		frequency.Change(v)
	}

	fmt.Printf("Answer: %v\n", frequency.Value)
}
