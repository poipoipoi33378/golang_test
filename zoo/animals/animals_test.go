package animals

import (
	"testing"
)

func TestElepahntFeed(t *testing.T) {
	expect := "Grass"
	actual := ElephantFeed()

	if expect != actual {
		t.Errorf("%s != %s", actual, expect)
	}
}

func TestRabbitFeed(t *testing.T) {
	expect := "Carrot"
	actual := RabbitFeed()

	if expect != actual {
		t.Errorf("%s != %s", actual, expect)
	}
}
