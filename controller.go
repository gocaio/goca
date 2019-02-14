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
	"net/http"
	"os"
	"strings"

	"github.com/asaskevich/EventBus"
	"github.com/gocolly/colly"
	"github.com/h2non/filetype"
	log "github.com/sirupsen/logrus"
)

var (
	// Topics is a list available event topics
	Topics = map[string]string{
		"NewTarget": "topic:new:url",
		"NewOutput": "topic:new:output",
	}
)

// Manager represents what interface should satisfies a Goca controller
type Manager interface {
	Publish(string, ...interface{})
	Subscribe(string, interface{}) error
	SubscribeOnce(string, interface{}) error
	Unsubscribe(string, interface{}) error
}

// Controller satisfy the Manager interface
type Controller struct {
	eventBus EventBus.Bus
	input    Input
}

func NewController(input Input) *Controller {
	ctrl := &Controller{
		eventBus: EventBus.New(),
		input:    input,
	}
	// This could be done outside NewController
	ctrl.Subscribe(Topics["NewOutput"], processOutput) // processOutput could be also a controller method
	return ctrl
}

func NewControllerTest() *Controller {
	ctrl := &Controller{
		eventBus: EventBus.New(),
	}
	return ctrl
}

func (c *Controller) Subscribe(topic string, fn interface{}) error {
	return c.eventBus.SubscribeAsync(topic, fn, false)
}

func (c *Controller) SubscribeOnce(topic string, fn interface{}) error {
	return c.eventBus.SubscribeOnceAsync(topic, fn)
}

func (c *Controller) Unsubscribe(topic string, fn interface{}) error {
	return c.eventBus.Unsubscribe(topic, fn)
}

func (c *Controller) Publish(topic string, args ...interface{}) {
	c.eventBus.Publish(topic, args...)
}

func (c *Controller) getData(plugType string, targets []string, files, recursive bool) {
	var data []byte
	var err error
	
	for _, url := range targets {
		if files {
			data, err = c.getFiles(url)
		} else {
			data, err = c.getURL(url)
		}

		if err != nil {
			log.Errorf("Unable to download %s", url)
		} else {
			kind, unknown := filetype.Match(data)
			if unknown != nil {
				log.Debugf("Unknown: %s", unknown)
			} else {
				if mimeAssociation.existMime(plugType, kind.MIME.Value) {
					// Publishing here may not be the best option
					log.Debugf("Data Length: %d\n", len(data))
					c.Publish(Topics["NewTarget"], url, data)
				} else {
					log.Debugf("Trying to download html for %s (%s)\n", kind.Extension, kind.MIME.Value)
					var newTargets []string
					// TODO: limit scraper to the target domain.
					scrap := colly.NewCollector(
					//colly.AllowedDomains("hackerspaces.org")
					)
					scrap.OnHTML("a[href]", func(e *colly.HTMLElement) {
						// TODO: Some URL has query strings or fragments that must be
						// removed in terms to build a propper url pointing to the file
						link := e.Attr("href")
						if strings.HasSuffix(url, "/") {
							newTargets = append(newTargets, fmt.Sprintf("%s%s", url, link))
						} else {
							newTargets = append(newTargets, fmt.Sprintf("%s/%s", url, link))
						}
					})
					scrap.Visit(url)
					if !recursive {
						c.getData(plugType, newTargets, false, true)
					} else {
						// Only addmited one level of recursion
						log.Errorf("There is no plugin processor for type %s: %s (%s)\n", plugType, kind.Extension, kind.MIME.Value)
					}
				}
			}
		}
	}
}

// get helps downloadURL to download the url content
func (c *Controller) getURL(url string) ([]byte, error) {
	var body []byte

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return body, err
	}

	// Set headers
	req.Header.Set("User-Agent", c.input.UA)

	// Create a new client
	client := &http.Client{} // This struct accepts config params
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

func (c *Controller) getFiles(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return data, nil
}
