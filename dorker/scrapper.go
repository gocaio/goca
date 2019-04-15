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

package dorker

import (
	"net/url"

	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

// Scrapper is the main scrapping class
type Scrapper struct {
	domain  string
	depthN  int
	files   []string
	targets map[string]struct {
		parent    string
		linkDepth int
		scrapped  bool
	}
}

// NewScrapper initializes the scrapper
func NewScrapper(domain string, depth int) *Scrapper {
	return &Scrapper{
		domain: domain,
		depthN: depth,
	}
}

// Run begins the scrapping task
func (s *Scrapper) Run() error {
	log.Infof("Scrapping %s", s.domain)

	u, err := url.Parse(s.domain)
	if err != nil {
		return err
	}

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains(u.Host),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		// log.Infof("Link found: %s", link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Debugf("Found: %s", r.URL.String())
		s.files = append(s.files, r.URL.String())
	})

	// Start scraping on the domain
	c.Visit(s.domain)

	return nil
}

// Links is a helper method to return all the recovered file URLs
func (s *Scrapper) Links() []string {
	log.Debug("Ret scrapped file links")
	return s.files
}

// =================
// = Class helpers =
// =================
func (s *Scrapper) foo() {}
