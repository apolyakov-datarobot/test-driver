package main

import (
	"github.com/tjarratt/babble"
	"os"
	"strconv"
	"testing"
)

func GetValFromENV(name string, t *testing.T) int {
	r := os.Getenv(name)
	res, err := strconv.Atoi(r)
	if err != nil {
		t.Errorf("can't convert env variable %s with value %s to integer", name, r)
		t.FailNow()
	}
	return res
}

func TestUntilCEqualsD(t *testing.T) {
	f, err := os.Create("TestUntilCEqualsD.out")
	if err != nil {
		t.Errorf("can't create output file: %s", err)
		t.FailNow()
	}

	a, b := GetValFromENV("TEST_INPUT_C", t), GetValFromENV("TEST_INPUT_D", t)
	if b < a {
		t.Errorf("input a = %d, input b = %d, expecting a < b", a, b)
		t.FailNow()
	}
	msg := "Now I'm gonna slowly go from a to b..."
	t.Log(msg)
	f.WriteString(msg)
	babbler := babble.NewBabbler()
	babbler.Separator = " "
	for a < b {
		msg = babbler.Babble()
		t.Log(msg)
		f.WriteString(msg + "\n")
		a = a + 1
	}
}
func TestUntilAEqualsB(t *testing.T) {
	a, b := GetValFromENV("TEST_INPUT_A", t), GetValFromENV("TEST_INPUT_B", t)
	if b < a {
		t.Errorf("input a = %d, input b = %d, expecting a < b", a, b)
		t.FailNow()
	}

	t.Log("Now I'm gonna slowly go from a to b...")
	babbler := babble.NewBabbler()
	babbler.Separator = " "
	for a < b {
		t.Log(babbler.Babble())
		a = a + 1
	}
}
