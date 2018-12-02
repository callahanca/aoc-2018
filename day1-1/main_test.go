package main

import "testing"

func TestCase1(t *testing.T) {

	f := NewFrequency()
	f.Change(1)
	f.Change(-2)
	f.Change(3)
	f.Change(1)

	if f.Value != 3 {
		t.Errorf("Value was %v", f.Value)
	}

}

func TestCase2(t *testing.T) {

	f := NewFrequency()
	f.Change(1)
	f.Change(1)
	f.Change(1)

	if f.Value != 3 {
		t.Errorf("Value was %v", f.Value)
	}

}

func TestCase3(t *testing.T) {

	f := NewFrequency()
	f.Change(1)
	f.Change(1)
	f.Change(-2)

	if f.Value != 0 {
		t.Errorf("Value was %v", f.Value)
	}

}

func TestCase4(t *testing.T) {

	f := NewFrequency()
	f.Change(-1)
	f.Change(-2)
	f.Change(-3)

	if f.Value != -6 {
		t.Errorf("Value was %v", f.Value)
	}

}
