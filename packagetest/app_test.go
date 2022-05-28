package main

import (
	"testing"
)

func TestApp(t *testing.T) {
	expect := "Application"
	actual := AppName()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}