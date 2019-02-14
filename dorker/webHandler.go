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
	"io/ioutil"
	"net/http"
)

func (d *Dorker) get(url string) ([]byte, error) {
	var body []byte

	// For POST methods, "nil" is a http form
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return body, err
	}

	// Set headers
	req.Header.Set("User-Agent", d.userAgent)

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
