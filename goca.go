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

package goca

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gocaio/goca/dorker"
	log "github.com/sirupsen/logrus"
)

// UserAgent defines the UserAgent used by goca
// FIXME: This should be possible to be overwritten with a flag
const UserAgent = "The_Goca_v0.1"

var (
	// AppName is the Application name
	AppName string
	// AppVersion is the Application version
	AppVersion string
	// LogLevel sets the log level for Goca
	LogLevel string
)

// StartTerm is the Goca library entrypoint for terms
func StartTerm(input Input) {
	setLogLevel()

	if input.ListURL {
		fmt.Println("Goca has got the following URLs for you")
	}

	for _, plugType := range input.FileType {
		dorks := getDorkers(plugType, input)

		log.Debugf("Dorks for plugin %s: %v\n", plugType, dorks.Dorks)

		// Initialize context or controller per each type
		ctx := NewController(input)

		// Initialize all plugins based on context
		executePlugins(plugType, ctx)

		urls := dorks.Google() // In the future user will be able to choose other search engine

		log.Debugf("Got %d url\n", len(urls))
		log.Debugln(urls)

		if input.ListURL {
			listURL(plugType, urls)
			break
		}

		if len(urls) == 0 {
			log.Warnln("Empty URL from dorker, Engine may have ban you.")
		}

		// TODO: Downloader should just download assets.
		ctx.getData(plugType, urls, false, false)
	}
}

// StartFolder is the Goca library entrypoint for folders
func StartFolder(input Input) {
	setLogLevel()
	if !validFolder(input.Folder) {
		log.Fatalf("The folder %s is not valid\n", input.Folder)
	}
	var files []string
	for _, plugType := range input.FileType {
		files = []string{}
		// Initialize context or controller per each type
		ctx := NewController(input)

		// Initialize all plugins based on context
		executePlugins(plugType, ctx)

		err := filepath.Walk(input.Folder,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() {
					files = append(files, path)
				}
				return nil
			})
		if err != nil {
			log.Errorln(err)
		}
		if err != nil {
			log.Errorln(err)
		}

		ctx.getData(plugType, files, true, false)
	}
}

func executePlugins(typ string, ctx *Controller) {
	if typ == "*" {
		for _, plugs := range plugins {
			for name, plug := range plugs {
				log.Debugf("Executing plugin %s", name)
				go plug.Action(ctx)
			}
		}
	} else {
		for name, plug := range plugins[typ] {
			log.Debugf("Executing plugin %s", name)
			go plug.Action(ctx)
		}
	}
}

func getDorkers(typ string, input Input) *dorker.Dorker {
	var dorkers []dorker.Dork
	var types []string
	d := dorker.NewDorker(input.UA, input.TimeOut, input.PagesDork)

	if typ == "*" {
		types = listPluginTypes()
	} else {
		types = []string{typ}
	}

	for _, t := range types {
		for _, mime := range mimeAssociation.getAssoc(t) {
			dorkers = dorker.DorkLib.GetByType(mime)
			if dorkers != nil {
				for _, dork := range dorkers {
					dork.UpdateString(input.URL, input.Term)
					d.AddDork(dork)
				}
			}
		}
	}
	return d
}

func setLogLevel() {
	switch LogLevel {
	case "info":
	case "1":
		log.SetLevel(log.InfoLevel)
	case "warn":
	case "2":
		log.SetLevel(log.WarnLevel)
	case "error":
	case "3":
		log.SetLevel(log.ErrorLevel)
	case "debug":
	case "4":
		log.SetLevel(log.DebugLevel)
	default:
		log.SetOutput(ioutil.Discard)
		log.SetLevel(log.InfoLevel)
	}
}

func validFolder(folder string) bool {
	_, err := os.Stat(folder)
	return !os.IsNotExist(err)
}

func listURL(plugType string, urls []string) {
	fmt.Printf("URLs for: %s\n", plugType)
	for _, url := range urls {
		fmt.Printf("\t- %s\n", url)
	}
}

// Input represents the command line arguments
type Input struct {
	Term      string
	Domain    string
	URL       string
	ListURL   bool
	Folder    string
	FileType  []string
	PagesDork int
	TimeOut   int
	UA        string
}
