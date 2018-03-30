package main

import (
	"os"
	"testing"
)

type TestCase struct {
	Args []string
	Res string
	Err error
}

func mockedArgs(t *testing.T, test TestCase) {
	_Args := os.Args
	os.Args = append([]string{"jetcli"}, test.Args...)
	rendered, err := render(parseArgs())
	if err != test.Err {
		t.Errorf("Incorrect error for %v; Expected: %v but got %v", test.Args, test.Err, err)
	}
	if rendered != test.Res {
		t.Errorf("Failed for %v; Expected: `%v` but got `%v`", test.Args, test.Res, rendered)
	}
	os.Args = _Args
}

func TestBasic(t *testing.T) {
	mockedArgs(t, TestCase{
		Args: []string{"-directory", "testdata", "test.html"},
		Res:  "title: hello from the jet CLI\nbody:  default body\n",
		Err:  nil,
	})
	mockedArgs(t, TestCase{
		Args: []string{"./testdata/test.html"},
		Res:  "title: hello from the jet CLI\nbody:  default body\n",
		Err:  nil,
	})
	mockedArgs(t, TestCase{
		Args: []string{"./testdata/nonexistent"},
		Res:  "",
		Err:  nil,
	})
}
