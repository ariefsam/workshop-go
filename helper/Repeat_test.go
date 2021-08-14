package helper_test

import (
	"testing"

	"github.com/ariefsam/workshop/helper"
)

func TestRepeatString(t *testing.T) {
	var expected, actual, actual2, expectedOutput2 string

	input1 := "a"
	input2 := 3
	expected = "aaa"
	expectedOutput2 = "aaaaaa"
	actual, actual2 = helper.RepeatString(input1, input2)

	if expected != actual {
		t.Error("fungsi tidak mengembalikan aaa")
	}

	if expectedOutput2 != actual2 {
		t.Error("output2 tidak sesuai")
	}
}

func TestRepeatString2(t *testing.T) {
	var expected, actual, actual2, expectedOutput2 string

	input1 := "a"
	input2 := 4
	expected = "aaaa"
	expectedOutput2 = "aaaaaaaa"
	actual, actual2 = helper.RepeatString(input1, input2)

	if expected != actual {
		t.Error("fungsi tidak mengembalikan aaa")
	}
	if expectedOutput2 != actual2 {
		t.Error("output2 tidak sesuai")
	}
}
