package main

import (
	"flag"
	"bytes"

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

func parseArgs() (string, string) {
	directory := flag.String("directory", "./", "The directory to search for templates in")
	flag.Parse()
	if flag.NArg() != 1 {
		panic("One filename of a template to render required!")
	}
	templateName := flag.Arg(0)
	return *directory, templateName
}

func main() {
	directory, templateName := parseArgs()
	rendered, err := render(directory, templateName)
	if err != nil {
		panic(err)
	}
	print(rendered)
}
