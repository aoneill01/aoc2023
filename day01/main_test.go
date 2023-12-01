package main

import (
	"strings"
	"testing"
)

const sampleInput1 = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

const sampleInput2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func TestCalibrationValue(t *testing.T) {
	got := CalibrationValue("a1b2c3d4e5f")
	want := 15

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart1(t *testing.T) {
	got := Part1(strings.Split(sampleInput1, "\n"))
	want := 142

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestCalibrationValue2(t *testing.T) {
	got := CalibrationValue2("eightwo")
	want := 82

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := Part2(strings.Split(sampleInput2, "\n"))
	want := 281

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
