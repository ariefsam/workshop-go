package helper_test

import (
	"testing"

	"github.com/ariefsam/workshop/helper"
)

func TestHello(t *testing.T) {
	var expected, actual string

	actual = helper.Hello()
	expected = "Hello World!"

	if actual != expected {
		t.Error("Fungsi Hello tidak sesuai!")
	}
}
