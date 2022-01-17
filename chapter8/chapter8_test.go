package chapter8

import "testing"

func TestMapDays (t *testing.T) {
	flag := DasyIsIn(Days)
	t.Log(flag)
}

func TestMapDrinks (t *testing.T) {
	SortDrinks(Drinks)
}