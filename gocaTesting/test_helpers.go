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

package gocatesting

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/gocaio/goca"
)

type Assets struct {
	Name      string
	Size      int
	URL       string
	ModTime   time.Time
	Mode      int64
	IsDir     bool
	IsSymlink bool
}

func GetAssets(t *testing.T, ctrl goca.Manager, testserver, plugName string) {
	// Simulate goca search and download system
	server := fmt.Sprintf("%s/%s/", testserver, plugName)

	client := &http.Client{}
	req, err := http.NewRequest("GET", server, nil)
	if err != nil {
		t.Errorf("unable to create a request for server %s. Error: %s", server, err.Error())
	}
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("unable to download assets from test server (%s) %s", server, err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		//t.Errorf("there are no files for %s", plugName)
		t.Skipf("skipping test for %s, because there is no data in test server", plugName)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("unable to read data from server (%s) %s", server, err.Error())
	}

	// Unmarshal json returned by the test server
	// This is like a list of URL
	var assetList []Assets
	err = json.Unmarshal(body, &assetList)
	if err != nil {
		t.Errorf("unable to unmarshal resources. Err: %s", err.Error())
	}

	// Download those URL and emit a NewURL event
	for _, res := range assetList {
		resourceURL := server + res.Name
		req, err = http.NewRequest("GET", resourceURL, nil)
		if err != nil {
			t.Errorf("unable to create a request for server %s. Error: %s", server, err.Error())
		}
		req.Header.Add("Accept", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			t.Errorf("unable to download assets from test server (%s) %s", server, err.Error())
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("unable to read data from server (%s) %s", server, err.Error())
		}
		ctrl.Publish(goca.Topics["NewURL"], resourceURL, body)
	}
}
