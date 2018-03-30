package main

import (
	"os"
	"testing"
)

type RenderTestCase struct {
	Args []string
	Expected string
	ArgErr bool
	RenderErr bool
}

// tests rendering
func testRender(t *testing.T, test RenderTestCase) {
	_Args := os.Args
	os.Args = append([]string{"jet"}, test.Args...)
	tpl, dir, err := parseArgs()
	if err != nil {
		if !test.ArgErr {
			// unexpected error
			t.Errorf("Incorrect argument error for %v; Expected no error but got %v",
				test.ArgErr, err)
		} else {
			// expected error
			return
		}
	}
	rendered, err := render(tpl, dir)
	if err != nil {
		if !test.RenderErr {
			t.Errorf("Incorrect render error for %v; Expected no error but got %v",
				test.RenderErr, err)
		} else {
			return
		}
	}
	if rendered != test.Expected {
		t.Errorf("Failed for %v; Expected: `%v` but got `%v`",
			test.Args, test.Expected, rendered)
	}
	os.Args = _Args
}

func TestBasic(t *testing.T) {
	testRender(t, RenderTestCase{
		Args: []string{"-dir", "testdata", "test.html"},
		Expected: "title: hello from the jet CLI\nbody:  default body\n",
		ArgErr: false,
		RenderErr: false,
	})
	testRender(t, RenderTestCase{
		Args: []string{"./testdata/test.html"},
		Expected: "title: hello from the jet CLI\nbody:  default body\n",
		ArgErr: false,
		RenderErr: false,
	})
	testRender(t, RenderTestCase{
		Args: []string{"./testdata/nonexistent"},
		Expected:  "",
		ArgErr: false,
		RenderErr: true,
	})
	testRender(t, RenderTestCase{
		Args: []string{},
		Expected:  "",
		ArgErr: true,
		RenderErr: false,
	})
	testRender(t, RenderTestCase{
		Args: []string{"x", "y"},
		Expected:  "",
		ArgErr: true,
		RenderErr: false,
	})
}

type ArgTestCase struct {
	Args []string
	Template string
	Directory string
	Err bool
}

func testArgs(t *testing.T, test ArgTestCase) {
	_Args := os.Args
	os.Args = append([]string{"jet"}, test.Args...)
	dir, tpl, err := parseArgs()
	if err != nil {
		if !test.Err {
			// unexpected error
			t.Errorf("Unexpected error: Expected none but got %s", err)
		} else {
			// expected error
			return
		}
	}
	if dir != test.Directory {
		t.Errorf("Incorrect directory: expected %s but got %s",
			test.Directory, dir)
	}
	if tpl != test.Template {
		t.Errorf("Incorrect template: expected %s but got %s",
			test.Template, tpl)
	}
	os.Args = _Args
}

func TestArgs(t *testing.T) {
	testArgs(t, ArgTestCase{
		Args: []string{"-dir", "testdata", "test.html"},
		Template: "test.html",
		Directory: "testdata",
		Err: false,
	})
	testArgs(t, ArgTestCase{
		Args: []string{"./testdata/test.html"},
		Template: "./testdata/test.html",
		Directory: "./",
		Err: false,
	})
	testArgs(t, ArgTestCase{
		Args: []string{"./testdata/nonexistent"},
		Template: "./testdata/nonexistent",
		Directory: "./",
		Err: false,
	})
	testArgs(t, ArgTestCase{
		Args: []string{},
		Template: "",
		Directory: "",
		Err: true,
	})
	testArgs(t, ArgTestCase{
		Args: []string{"x", "y"},
		Template: "",
		Directory: "",
		Err: true,
	})
}
