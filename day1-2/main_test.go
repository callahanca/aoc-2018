package main

import "testing"

func TestCase1(t *testing.T) {

	f := NewFrequencyLogged()
	dupe := f.Change(1)
	if dupe && f.Value != 0 {
		t.Error("wrong")
	}
	dupe = f.Change(-1)
	if dupe && f.Value != 0 {
		t.Error("wrong")
	}

}

func TestCase2(t *testing.T) {

	f := NewFrequencyLogged()
	dupe := f.Change(3)
	if dupe && f.Value != 10 {
		t.Error("wrong")
	}
	dupe = f.Change(-3)
	if dupe && f.Value != 10 {
		t.Error("wrong")
	}
	dupe = f.Change(4)
	if dupe && f.Value != 10 {
		t.Error("wrong")
	}
	dupe = f.Change(-2)
	if dupe && f.Value != 10 {
		t.Error("wrong")
	}
	dupe = f.Change(-4)
	if dupe && f.Value != 10 {
		t.Error("wrong")
	}

}

func TestCase3(t *testing.T) {

	f := NewFrequencyLogged()
	dupe := f.Change(-6)
	if dupe && f.Value != 5 {
		t.Error("wrong")
	}
	dupe = f.Change(-3)
	if dupe && f.Value != 5 {
		t.Error("wrong")
	}
	dupe = f.Change(8)
	if dupe && f.Value != 5 {
		t.Error("wrong")
	}
	dupe = f.Change(5)
	if dupe && f.Value != 5 {
		t.Error("wrong")
	}
	dupe = f.Change(-6)
	if dupe && f.Value != 5 {
		t.Error("wrong")
	}

}

func TestCase4(t *testing.T) {

	f := NewFrequencyLogged()
	dupe := f.Change(7)
	if dupe && f.Value != 14 {
		t.Error("wrong")
	}
	dupe = f.Change(7)
	if dupe && f.Value != 14 {
		t.Error("wrong")
	}
	dupe = f.Change(-2)
	if dupe && f.Value != 14 {
		t.Error("wrong")
	}
	dupe = f.Change(-7)
	if dupe && f.Value != 14 {
		t.Error("wrong")
	}
	dupe = f.Change(-4)
	if dupe && f.Value != 14 {
		t.Error("wrong")
	}

}
