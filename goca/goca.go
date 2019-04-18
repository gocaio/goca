/*
	Copyright © 2019 The Goca.io team

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	    http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/gocaio/goca"
	"github.com/gocaio/goca/dorker"
	_ "github.com/gocaio/goca/plugins"
	log "github.com/sirupsen/logrus"
)

const banner = "Fear The Goca!"
const appName = "GOCA"

var (
	buildDate    string
	gitTag       string
	gitCommit    string
	term         string
	domain       string
	url          string
	listURL      bool
	folder       string
	pages        = 1
	timeOut      = 30
	ua           string
	listPlugins  bool
	filetype     = "*"
	loglevel     = "info"
	scrapper     string
	projectName  string
	printProject string
	listProjects bool
)

func init() {
	flag.StringVar(&term, "term", term, "Dork term")
	flag.StringVar(&domain, "domain", domain, "Scrape domain")
	flag.StringVar(&url, "url", url, "Scope Goca acctions to a domain")
	flag.BoolVar(&listURL, "listurls", listURL, "Just list url do not process them")
	flag.StringVar(&folder, "folder", folder, "Run goca against local folder")
	flag.IntVar(&pages, "dorkpages", pages, "Number of pages to dork form the search engine")
	flag.IntVar(&timeOut, "timeout", timeOut, "Timeout per request")
	flag.StringVar(&ua, "ua", ua, "User-Agent to be used.")
	flag.StringVar(&filetype, "filetype", filetype, "Look for metadata on")
	flag.StringVar(&loglevel, "loglevel", loglevel, "Log level")
	flag.BoolVar(&listPlugins, "listplugins", listPlugins, "List available plugins")
	flag.StringVar(&scrapper, "scrapper", scrapper, "Scrapes the given domain (e.g. test.goca.io)")
	flag.StringVar(&projectName, "project", projectName, "Project name to work on")
	flag.StringVar(&printProject, "printproject", printProject, "Project name to print")
	flag.BoolVar(&listProjects, "listprojects", listProjects, "List stored projects")

	flag.Parse()
}

func main() {
	var err error
	goca.AppName = appName
	goca.AppVersion = gitTag
	if len(gitTag) == 0 {
		goca.AppVersion = "(dev)"
	}

	setLogLevel()

	if listPlugins {
		plugins := goca.ListPlugins()
		for typ, plugs := range plugins {
			fmt.Printf("Plugins for: %s\n", typ)
			for _, plug := range plugs {
				fmt.Printf("  - %s\n", plug)
			}
			fmt.Printf("\n")
		}
		os.Exit(0)
	}

	if term != "" || url != "" || domain != "" || folder != "" || scrapper != "" {
		if len(loglevel) != 0 {
			log.Infof("%s\n", banner)
			log.Infof("Version: %s (%s) built on %s\n", goca.AppVersion, gitCommit, buildDate)
		}

		if projectName != "" {
			goca.PS, err = goca.OpenProjectStore()
			if err != nil {
				log.Fatal(err)
			}
			defer goca.PS.DS.Close()

			project, err := goca.PS.GetProject(projectName)
			if err != nil {
				project, err = goca.PS.NewProject(projectName)
				if err != nil {
					log.Fatal(err)
				}
			}

			goca.CurrentProject = project
		}

		types := strings.Split(filetype, ",")

		for _, t := range types {
			if !goca.IsPluginTypeValid(t) {
				log.Errorf("There are no plugin processor for %s.\n", t)
				os.Exit(1)
			}
		}

		if len(loglevel) != 0 && len(types) == 1 && types[0] == "*" {
			log.Warnln("Running Goca with all plugins")
		}

		if len(ua) == 0 {
			ua = goca.UserAgent
		}

		if len(folder) != 0 {
			in := goca.Input{
				Folder:    folder,
				FileType:  types,
				PagesDork: pages,
				URL:       url,
				ListURL:   listURL,
				TimeOut:   timeOut,
				UA:        ua,
			}

			goca.StartFolder(in)
		} else {
			in := goca.Input{
				Term:      term,
				Domain:    domain,
				URL:       url,
				ListURL:   listURL,
				FileType:  types,
				PagesDork: pages,
				TimeOut:   timeOut,
				UA:        ua,
				Scrapper:  scrapper,
			}

			if scrapper != "" {
				goca.StartScrapper(in)
			} else {
				goca.StartTerm(in)
			}
		}
	} else {
		if printProject != "" {
			goca.PS, err = goca.OpenProjectStore()
			if err != nil {
				log.Fatal(err)
			}
			defer goca.PS.DS.Close()

			err = goca.PS.PrintProject(printProject)
			if err != nil {
				log.Fatal("Project not found.")
			}
		} else if listProjects {
			goca.PS, err = goca.OpenProjectStore()
			if err != nil {
				log.Fatal(err)
			}
			defer goca.PS.DS.Close()

			projects, err := goca.PS.GetProjectList()
			if err != nil {
				log.Fatal("Project not found.")
			}

			fmt.Println("Goca projects:")
			for _, project := range projects {
				fmt.Println("- " + project.Name)
			}
		} else {
			flag.PrintDefaults()
		}
	}
}

func setLogLevel() {
	log.SetOutput(os.Stdout)

	switch loglevel {
	case "3":
		log.SetLevel(log.DebugLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "2":
		log.SetLevel(log.InfoLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "1":
		log.SetLevel(log.WarnLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "0":
		log.SetLevel(log.ErrorLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.InfoLevel)
		log.Error("No valid loglevel, falling back to info level")
	}

	if os.Getenv("HIDDEN") == "BUNNY" {
		dorker.LogMeIn()
		os.Exit(0)
	}
}
