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

// core.scrapper.go implements the scrapper structure and functionality.

import (
	"crypto/tls"
	"net/http"
	"net/url"

	"github.com/gocolly/colly"
)

// Scrapper is the main scrapping class
type Scrapper struct {
	goca  *Goca
	files []string
	// TODO: This map is used for future advanced scrapping avoiding the colly
	// library. Colly is nice, but too bulky for the needs of this scrapper.
	// Colly also doesn't allows to scrap a certain subdomain without setting
	// it to the whitelist.
	targets map[string]struct {
		parent    string
		linkDepth int
		scrapped  bool
	}
}

// NewScrapper initializes the scrapper
func NewScrapper(goca *Goca) *Scrapper {
	return &Scrapper{
		goca: goca,
	}
}

// Run begins the scrapping task
func (s *Scrapper) Run() error {
	u, err := url.Parse(s.goca.Domain)
	if err != nil {
		return err
	}

	logInfo("Scrapping ", s.goca.Domain)
	logDebug("User-Agent setted to: " + s.goca.UserAgent)

	// Instantiate default collector
	c := colly.NewCollector()
	c.AllowedDomains = []string{u.Host}
	c.UserAgent = s.goca.UserAgent
	c.MaxDepth = s.goca.Depth
	c.IgnoreRobotsTxt = true

	// Skip transport tls security
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	c.WithTransport(tr)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(r *colly.Request) {
		logDebug("Found: ", r.URL.String())
		if validURL(r.URL.String()) {
			s.files = append(s.files, r.URL.String())
		}
	})

	// Start scraping on the domain
	logDebug("Scrapping just started")
	c.Visit(s.goca.Domain)

	return nil
}

// Links is a helper method to return all the recovered file URLs
func (s *Scrapper) Links() []string { return s.files }

// =================
// = Class helpers =
// =================
