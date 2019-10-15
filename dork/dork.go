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

package dork

// dork.dork.go defines the public dork interface

// Engine defines the engine type
type Engine int

const (
	// Google engine
	Google Engine = iota + 1
	// Bing engine
	Bing
	// DDG engine
	DDG
)

// EngineList has the engine names
var EngineList = map[Engine]string{
	Google: "Google",
	Bing:   "Bing",
	DDG:    "DDG",
}

// Dork defines a dork for a particular engine
type Dork struct {
	Engine Engine
	Query  string
}

// NewDork returns a new dork
func NewDork(engine Engine, query string) (d *Dork) {
	d = &Dork{}
	if engine > 0 && query != "" {
		d.Engine = engine
		d.Query = query
	}
	return
}
