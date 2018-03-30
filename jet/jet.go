package main

import (
	"os"
	"flag"
	"bytes"
	"errors"

	"github.com/CloudyKit/jet"
)

func render(directory string, templateName string) (string, error) {
	view := jet.NewHTMLSet(directory)
	tpl, err := view.GetTemplate(templateName)
	if err != nil {
		return "", err
	}
	var ret bytes.Buffer
	err = tpl.Execute(&ret, make(jet.VarMap), nil)
	if err != nil {
		return "", err
	}
	return ret.String(), nil
}

func parseArgs() (string, string, error) {
	set := flag.NewFlagSet("", flag.ContinueOnError)
	directory := set.String("directory", "./", "The directory to search for templates in")
	err := set.Parse(os.Args[1:])
	if set.NArg() != 1 {
		return "", "", errors.New("One filename of a template to render required!")
	}
	templateName := set.Arg(0)
	return *directory, templateName, err
}

func main() {
	directory, templateName, err := parseArgs()
	if err != nil {
		os.Stderr.WriteString("Illegal arguments: " + err.Error())
		os.Exit(-1)
	}
	rendered, err := render(directory, templateName)
	if err != nil {
		os.Stderr.WriteString("Jet render error: " + err.Error())
		os.Exit(-1)
	}
	print(rendered)
}
