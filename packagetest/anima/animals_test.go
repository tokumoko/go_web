package anima

import(
	"testing"
)

func TestEleph(t *testing.T) {
	expect := "A"
	actual := Abc()
	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}
func TestMonk(t *testing.T) {
	expect := "K"
	actual := Monkey()
	if expect != actual {
		t.Errorf("%s != %s", expect, actual
}
func TestRab(t *testing.T) {
	expect := "I"
	actual := Kuruit()
	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}