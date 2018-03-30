package main

import (
	"os"
	"fmt"
	"flag"
	"bytes"
	"errors"

	"github.com/CloudyKit/jet"
)

// exists returns whether the given file or directory exists or not
// https://stackoverflow.com/a/10510783/5719760
// returns true, nil if the file exists and false, nil if the file doesn't
// exist and a non-nil error for other Stat errors like a permissions error
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// render renders a template
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
	set.Usage = func() {
		argspec := "[options] TEMPLATE_NAME"
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s %s\n",
			os.Args[0], argspec)
		set.PrintDefaults()
	}
	directory := set.String("dir", "./", "The directory to search for templates in")
	const templateArg = "template"
	templateName := set.String(templateArg, "", "The filename of the template to render")
	err := set.Parse(os.Args[1:])
	if err != nil {
		return "", "", err
	}
	if set.Lookup(templateArg).Value.String() == "" {
		// blank template
		if set.NArg() != 1 {
			return "", "", errors.New("Exactly one filename of a template to render required")
		} else {
			// template name given positionally
			// hack to use string as *string
			tmp := set.Arg(0)
			templateName = &tmp
		}
	}
	return *directory, *templateName, err
}

func main() {
	directory, templateName, err := parseArgs()
	if err == flag.ErrHelp {
		// just printed the help msg, not an error
		os.Exit(-1)
	}
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
