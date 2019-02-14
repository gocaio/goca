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
	"fmt"
	"net/url"

	log "github.com/sirupsen/logrus"
)

// Google dorks Google engine
func (d *Dorker) Google() (results []string) {
	for _, dork := range d.Dorks {
		if dork.Engine != "google" {
			continue
		}

		re := `"><a href="/url\?q=(.*?)&amp;sa=U&amp;`
		escapedDork := url.QueryEscape(dork.String)

		for i := 0; i < d.depth; i++ {
			u := fmt.Sprintf("https://www.google.com/search?q=%s&start=%d0", escapedDork, i)

			wbuf, err := d.get(u)
			if err != nil {
				log.Printf("[!] (%s): %s", u, err.Error())
				continue
			}

			parsedData := parser(wbuf, re)
			for j := range parsedData {
				urlUnescaped, err := url.QueryUnescape(parsedData[j][1])
				if err != nil {
					log.Printf("[!] %s", err.Error())
					continue
				}

				results = append(results, urlUnescaped)
			}
		}
	}

	return
}
