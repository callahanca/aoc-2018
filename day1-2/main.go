package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type FrequencyLogged struct {
	Value int
	Log   map[int]*interface{}
}

func NewFrequencyLogged() *FrequencyLogged {
	f := FrequencyLogged{Value: 0,
		Log: make(map[int]*interface{})}

	return &f
}

func (f *FrequencyLogged) Change(c int) bool {
	f.Value += c
	_, ok := f.Log[f.Value]
	if !ok {
		f.Log[f.Value] = nil
	}

	return ok

}

func loop(l []int, frequency *FrequencyLogged) {
	for _, v := range l {

		dupe := frequency.Change(v)
		if dupe {
			fmt.Printf("First dupe is: %v\n", frequency.Value)
			return
		}
	}
	loop(l, frequency)
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
	frequency := NewFrequencyLogged()
	l := make([]int, 0)
	found := false
	for scanner.Scan() {
		t := scanner.Text()
		v, err := strconv.Atoi(t)
		if err != nil {
			log.Fatalf("Error converting string to int: %s", err)
		}
		dupe := frequency.Change(v)
		l = append(l, v)
		if dupe {
			found = true
			fmt.Printf("First dupe is: %v\n", frequency.Value)
			return
		}
	}
	if !found {
		loop(l, frequency)

	}

}
