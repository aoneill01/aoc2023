package main

import (
	"strings"
	"testing"
)

const sample1 = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func gotWantHelper(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestPart1(t *testing.T) {
	got := part1(strings.Split(sample1, "\n"))
	gotWantHelper(t, got, 6440)
}

func TestPart2(t *testing.T) {
	got := part2(strings.Split(sample1, "\n"))
	gotWantHelper(t, got, 5905)
}

func TestParseCard(t *testing.T) {
	cardTests := []struct {
		card  string
		value Card
	}{
		{"A", 14},
		{"T", 10},
		{"9", 9},
		{"2", 2},
	}

	for _, cardTest := range cardTests {
		t.Run("test card "+cardTest.card, func(t *testing.T) {
			gotWantHelper(t, int(parseCard(cardTest.card)), int(cardTest.value))
		})
	}
}
