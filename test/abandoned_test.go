package test

import (
	"testing"
)

func TestAbandoned(t *testing.T) {
	// this just tests the name generation works correctly
	r := Root{
		Name:      "jonson",
		Abandoned: &abandoned.PackageList{},
	}
	// the test is the presence of the Abandoned field
	if r.Abandoned == nil {
		t.Fatal("thats the test")
	}
}
