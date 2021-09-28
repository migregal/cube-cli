package utils

import (
	"testing"
)

func TestEmptyArgsParse(t *testing.T) {
	var args []string
	_, err := ParseArgs(args)
	if err == nil {
		t.Fatal("Wrong args parsing succeed")
	}
	if err.Error() != "missing args detected" {
		t.Fatal("Wrong error msg")
	}
}

func TestEmptyHostArgParse(t *testing.T) {
	args := []string{"", "", ":0", "token", "scope"}
	_, err := ParseArgs(args)
	if err == nil {
		t.Fatal("Wrong args parsing succeed")
	}
	if err.Error() != "missing host value" {
		t.Fatal("Wrong error msg")
	}
}

func TestEmptyPortArgParse(t *testing.T) {
	args := []string{"", "localhost", "", "token", "scope"}
	_, err := ParseArgs(args)
	if err == nil {
		t.Fatal("Wrong args parsing succeed")
	}
	if err.Error() != "missing port value" {
		t.Fatal("Wrong error msg")
	}
}

func TestEmptyTokenArgParse(t *testing.T) {
	args := []string{"", "localhost", ":0", "", "scope"}
	_, err := ParseArgs(args)
	if err == nil {
		t.Fatal("Wrong args parsing succeed")
	}
	if err.Error() != "missing token value" {
		t.Fatal("Wrong error msg")
	}
}

func TestEmptyScopeArgParse(t *testing.T) {
	args := []string{"", "localhost", ":0", "token", ""}
	_, err := ParseArgs(args)
	if err != nil {
		t.Fatal("Empty scope parsing failed")
	}
}

func TestCorrectArgsParse(t *testing.T) {
	expected := Args{Host: "localhost", Port: ":0", Token: "token", Scope: "scope"}
	args, err := ParseArgs([]string{"", "localhost", ":0", "token", "scope"})
	if err != nil {
		t.Fatal("Correct args parsing failed")
	}
	if *args != expected {
		t.Fatal("Got wrong result on args parsing")
	}
}
