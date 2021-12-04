package main

import (
	"flag"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/mklnz/fyro-templater/dataparser"
)

var (
	fieldDefs    *string
	outputFile   *string
	templateFile string
)

func init() {
	fieldDefs = flag.String("f", "", "field definition")
	outputFile = flag.String("o", "", "path to output file")
}

func stringJoin(s []string, sep string) string {
	return strings.Join(s, sep)
}

func main() {
	flag.Parse()

	fields := dataparser.ParseFields(*fieldDefs)
	dataMap := dataparser.FetchData(fields)

	templateFile = flag.Arg(0)

	t, err := template.New("virtualhost.tmpl").Funcs(
		template.FuncMap{"stringJoin": stringJoin},
	).ParseFiles(templateFile)
	if err != nil {
		panic(err)
	}

	var output io.Writer
	if *outputFile != "" {
		output, err = os.Create(*outputFile)
		if err != nil {
			panic(err)
		}
	} else {
		output = os.Stdout
	}

	err = t.Execute(output, dataMap)
	if err != nil {
		panic(err)
	}
}
