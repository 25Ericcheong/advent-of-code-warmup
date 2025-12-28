package day01_secretentrance

import (
	"log"
	"testing"
)

func TestSimpleExpectedCountCaseOne(t *testing.T) {
	var tests = []struct {
		leftOrRight   string
		val           int
		expectedCount int
	}{
		{"L", 50, 1},
		{"L", 150, 2},
		{"R", 50, 1},
		{"R", 100, 1},
		{"R", 150, 2},
		{"R", 200, 2},
	}

	for _, test := range tests {
		dial := CreateDial(initialDialCount)

		input := test.leftOrRight
		if input == LEFT {
			dial.TurnLeft(test.val)
		} else if input == RIGHT {
			dial.TurnRight(test.val)
		} else {
			log.Fatal("unexpected value encountered while dialing")
		}

		got := dial.NewPasswordCounter
		want := test.expectedCount
		if got != want {
			t.Errorf("got %d, wanted %d for %s with %d", got, want, input, test.val)
		}
	}
}

func TestSimpleExpectedCountCaseTwo(t *testing.T) {
	var tests = []struct {
		leftOrRight   string
		val           int
		expectedCount int
	}{
		{"L", 500, 5},
	}

	for _, test := range tests {
		dial := CreateDial(0)

		input := test.leftOrRight
		if input == LEFT {
			dial.TurnLeft(test.val)
		} else if input == RIGHT {
			dial.TurnRight(test.val)
		} else {
			log.Fatal("unexpected value encountered while dialing")
		}

		got := dial.NewPasswordCounter
		want := test.expectedCount
		if got != want {
			t.Errorf("got %d, wanted %d for %s with %d", got, want, input, test.val)
		}
	}
}
