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

// core.downloader.go implements the file download functionality and the
// multi-threading control system.

import (
	"crypto/tls"
	"errors"
	"io/ioutil"
	"net/http"
	// "sync"
)

// Downloader is the download handler class
type Downloader struct {
	goca        *Goca
	threads     int
	Links       []string
	resultFiles map[string][][]byte
}

// NewDownloader returns a Downloader object
func NewDownloader(goca *Goca, threads int) *Downloader {
	var files = map[string][][]byte{}

	return &Downloader{
		goca:        goca,
		threads:     threads,
		resultFiles: files,
	}
}

// Run begins with the download
func (d *Downloader) Run() (err error) {
	if len(d.Links) < 1 {
		return errors.New("no links provided to downloader")
	}

	// TODO: Spin-up N download threads
	// wg := &sync.WaitGroup{}
	// for i := 0; i < d.threads; i++ {
	// 	wg.Add(1)
	// 	go func() {
	//		// Download routine
	// 		logDebug("Hi")
	// 		wg.Done()
	// 	}()
	// 	wg.Wait()
	// }

	for _, link := range d.Links {
		body, mime, err := d.get(link)
		if err != nil {
			logError(err)
		}

		d.resultFiles[mime] = append(d.resultFiles[mime], body)
	}

	return err
}

// Files returns a map of scrapped files with the mimetype as key
func (d Downloader) Files() map[string][][]byte { return d.resultFiles }

// =================
// = Class helpers =
// =================

func (d Downloader) get(url string) (body []byte, mime string, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return body, mime, err
	}

	// Set headers
	req.Header.Set("User-Agent", d.goca.UserAgent)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
	}

	res, err := client.Do(req)
	if err != nil {
		return body, mime, err
	}
	defer res.Body.Close()

	// Readout the body
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return body, mime, err
	}

	mime = res.Header.Get("Content-Type")

	return body, mime, err
}
