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
	"time"

	"github.com/gocaio/goca/dork"
)

// core.go implements all the core behavior stuff related to Goca, such as
// initialization stuff and type definitions.

// Goca is the core structure from where all the modes inhere
type Goca struct {
	UserAgent  string        // User-Agent to send on the requests
	BaseFolder string        // Base Goca folder where the configuration files/project dbs/tmp files are stored
	DB         *struct{}     // FIXME: This is only a dev placeholder (this should be a pointer to the initialized database)
	Domain     string        // The domain to scrap for available documents
	Threshold  int           // Threshold for limiting the scrap speed
	Depth      int           // Depth of the links to follow while scrapping
	Term       string        // Dorking term
	Pages      int           // Dorking result pages to scrap
	Engines    []dork.Engine // Allowed engines to dork
	Stats      Stats         // Global stats about the current scan
}

// Stats defines a current scann statistics object
type Stats struct {
	Start         time.Time
	Stop          time.Time
	MimeTypeCount map[string]int
}

var pluginHub PluginHub
