package helper_test

import (
	"testing"

	"github.com/ariefsam/workshop/helper"
)

func TestJoinText(t *testing.T) {
	input := []string{"Halo", "halo", "Bandung"}
	expected := "Halo halo Bandung"

	actual := helper.JoinText(input)
	if expected != actual {
		t.Error("Output "+actual, ". Padahal seharusnya "+expected)
	}

}
