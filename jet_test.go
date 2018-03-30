package main

import (
	"os"
	"testing"
)

type TestCase struct {
	Args []string
	Expected string
	ArgErr bool
	RenderErr bool
}

func mockedArgs(t *testing.T, test TestCase) {
	_Args := os.Args
	os.Args = append([]string{"jetcli"}, test.Args...)
	tpl, dir, err := parseArgs()
	if err != nil {
		if !test.ArgErr {
			// unexpected error
			t.Errorf("Incorrect argument error for %v; Expected: %v but got %v", test.Args, test.ArgErr, err)
		} else {
			// expected error
			return
		}
	}
	rendered, err := render(tpl, dir)
	if err != nil {
		if !test.RenderErr {
			t.Errorf("Incorrect render error for %v; Expected: %v but got %v", test.Args, test.RenderErr, err)
		} else {
			return
		}
	}
	if rendered != test.Expected {
		t.Errorf("Failed for %v; Expected: `%v` but got `%v`", test.Args, test.Expected, rendered)
	}
	os.Args = _Args
}

func TestBasic(t *testing.T) {
	mockedArgs(t, TestCase{
		Args: []string{"-directory", "testdata", "test.html"},
		Expected: "title: hello from the jet CLI\nbody:  default body\n",
		ArgErr: false,
		RenderErr: false,
	})
	mockedArgs(t, TestCase{
		Args: []string{"./testdata/test.html"},
		Expected: "title: hello from the jet CLI\nbody:  default body\n",
		ArgErr: false,
		RenderErr: false,
	})
	mockedArgs(t, TestCase{
		Args: []string{"./testdata/nonexistent"},
		Expected:  "",
		ArgErr: false,
		RenderErr: true,
	})
	mockedArgs(t, TestCase{
		Args: []string{},
		Expected:  "",
		ArgErr: true,
		RenderErr: false,
	})
	mockedArgs(t, TestCase{
		Args: []string{"x", "y"},
		Expected:  "",
		ArgErr: true,
		RenderErr: false,
	})
}
