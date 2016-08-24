package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/tudurom/orator/config"
	"github.com/tudurom/orator/gen"
	"github.com/tudurom/orator/util"
)

var Fm *util.FrontMatter
var Layouts map[string]gen.Layout
var RootTemplate *template.Template
var SiteConfig *config.SiteConfig

const (
	configFilePath = "config.yaml"
	layoutDir      = "layouts"
	contentDir     = "content"
	outputDir      = "gen"
	staticDir      = "static"
)

func Init() {
	Fm = util.NewFrontMatter("---")
	Layouts = make(map[string]gen.Layout)
	RootTemplate = template.New("root")
}

func usage() {
	fmt.Print(
		`Usage: orator [-h]

Options:
	-h - print this message

Usage:
	Invoke orator to generate the site in the gen directory int the current working directory.
`,
	)
}

func main() {
	var showUsage bool
	flag.BoolVar(&showUsage, "h", false, "Show help")
	flag.Parse()
	if showUsage {
		usage()
		os.Exit(0)
	}

	Init()

	SiteConfig = new(config.SiteConfig)
	SiteConfig.ReadConfig(configFilePath)
	gen.LoadLayouts(layoutDir, Layouts, RootTemplate, Fm, SiteConfig)
	err := gen.GenerateSite(contentDir, outputDir, staticDir, Fm, Layouts, RootTemplate, SiteConfig)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Job's done.")
}
