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
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/gocaio/goca/dork"
)

// core.dorker.go implements core dorker structure and functionality.

// Dorker defines a dorker structure. Used to store dorks.
type Dorker struct {
	allowedEngines []dork.Engine
	userAgent      string
	pages          int
	term           string
	threshold      int
	dorks          []*dork.Dork
	links          []string
}

// NewDorker returns a new dorker
func NewDorker(g *Goca) *Dorker {
	return &Dorker{
		allowedEngines: g.Engines,
		userAgent:      g.UserAgent,
		pages:          g.Pages,
		term:           g.Term,
		threshold:      g.Threshold,
		dorks:          []*dork.Dork{},
	}
}

// AddDorks adds dorks to dorker
func (d *Dorker) AddDorks(dorks []*dork.Dork) {
	if dorks == nil {
		dorks = []*dork.Dork{}
	}
	anyEngineValid := false
	for _, dd := range dorks {
		if d.addDork(dd) {
			anyEngineValid = true
		}
	}
	if !anyEngineValid {
		logWarning("Dorks with not allowed engines")
	}
}

// Run execute dorks according its engine
func (d *Dorker) Run() error {
	for _, dd := range d.dorks {
		links := []string{}
		switch dd.Engine {
		case dork.Google:
			links = d.Google(dd)
		case dork.Bing:
			links = d.Bing(dd)
		case dork.DDG:
			links = d.DDG(dd)
		}
		d.links = append(d.links, links...)
		logDebug(fmt.Sprintf("%s dork got %d links", dork.EngineList[dd.Engine], len(links)))
	}
	if len(d.links) == 0 {
		return errors.New("no links dorked")
	}
	return nil
}

// Links returns the list of dorked links
func (d *Dorker) Links() []string { return d.links }

// AddDorksFromPluginHub adds dorks to dorker based on plugins already loaded
func (d *Dorker) AddDorksFromPluginHub(plugHub *PluginHub) {
	for _, plug := range plugHub.plugins {
		d.AddDorks(plug.Dorks)

	}
}

func (d *Dorker) addDork(dk *dork.Dork) bool {
	if d.validEngine(dk) {
		dk.Query = fmt.Sprintf("%s %s", d.term, dk.Query)
		d.dorks = append(d.dorks, dk)
		return true
	}
	return false
}

func (d *Dorker) validEngine(de *dork.Dork) bool {
	for _, e := range d.allowedEngines {
		if de.Engine == e {
			return true
		}
	}
	return false
}

func (d *Dorker) get(url string) ([]byte, error) {
	var body []byte

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return body, err
	}

	// Set headers
	req.Header.Set("User-Agent", d.userAgent)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	// Create a new client
	client := &http.Client{
		Transport: tr,
	} // This struct accepts config params
	res, err := client.Do(req)
	if err != nil {
		return body, err
	}
	defer res.Body.Close()

	// Readout the body
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return body, err
	}

	return body, nil
}

// ===========
// = Helpers =
// ===========

func parser(buf []byte, re string) (parsed [][]string) {
	rex := regexp.MustCompile(re)
	parsed = rex.FindAllStringSubmatch(string(buf), -1)[0:]
	return
}
