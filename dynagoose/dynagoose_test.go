package dynagoose

import (
	"testing"
)

func TestCreate(t *testing.T) {
	d := Dynamo{}

	expected := "create"
	if expected != d.Create() {
		t.Fatalf("error: got: %s, expected: %s", d.Create(), expected)
	}
}

func TestSelect(t *testing.T) {
	d := Dynamo{}

	expected := "select"
	if expected != d.Select() {
		t.Fatalf("error: got: %s, expected: %s", d.Select(), expected)
	}
}

func TestUpdate(t *testing.T) {
	d := Dynamo{}

	expected := "update"
	if expected != d.Update() {
		t.Fatalf("error: got: %s, expected: %s", d.Update(), expected)
	}
}

func TestDelete(t *testing.T) {
	d := Dynamo{}

	expected := "delete"
	if expected != d.Delete() {
		t.Fatalf("error: got: %s, expected: %s", d.Delete(), expected)
	}
}
