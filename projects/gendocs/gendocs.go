package main

import (
	"flag"

	"github.com/isaaguilar/terraform-operator/pkg/apis/tf/v1alpha2"
)

var (
	templatefile string
	outputfile   string
)

func init() {
	flag.StringVar(&templatefile, "tpl", "", "Path to the template")
	flag.StringVar(&outputfile, "out", "", "Path to save rendered template")
	flag.Parse()
}

func main() {
	v1alpha2.Generate(templatefile, outputfile)
}
