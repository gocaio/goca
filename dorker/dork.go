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
)

const (
	DEFAULT_DEPTH   = 1  // recurse into
	DEFAULT_TIMEOUT = 30 // TODO: Set req timeout
)

type dorkLib map[string][]Dork

var DorkLib = dorkLib{
	"application/pdf": []Dork{
		Dork{"google", "filetype:pdf"},
	},
	"audio/mpeg": []Dork{
		Dork{"google", "intitle:index.of +?last modified? +?parent directory? +(mp3|wma|ogg) -htm -html -php -asp"},
		Dork{"google", "filetype:mp3"},
		Dork{"google", "filetype:mp4"},
		Dork{"google", "filetype:m4a"},
		Dork{"google", "filetype:m4p"},
		Dork{"google", "filetype:m4v"},
		Dork{"google", "filetype:3gp"},
		Dork{"google", "filetype:3g2"},
	},
	"audio/MPA": []Dork{
		Dork{"google", "filetype:mp3"},
		Dork{"google", "filetype:mp4"},
		Dork{"google", "filetype:m4a"},
		Dork{"google", "filetype:m4p"},
		Dork{"google", "filetype:m4v"},
		Dork{"google", "filetype:3gp"},
		Dork{"google", "filetype:3g2"},
	},
	"audio/mpa-robust": []Dork{
		Dork{"google", "filetype:mp3"},
		Dork{"google", "filetype:mp4"},
		Dork{"google", "filetype:m4a"},
		Dork{"google", "filetype:m4p"},
		Dork{"google", "filetype:m4v"},
		Dork{"google", "filetype:3gp"},
		Dork{"google", "filetype:3g2"},
	},
	"application/ogg": []Dork{
		Dork{"google", "filetype:ogg"},
		Dork{"google", "filetype:ogv"},
		Dork{"google", "filetype:oga"},
		Dork{"google", "filetype:ogx"},
		Dork{"google", "filetype:ogm"},
		Dork{"google", "filetype:spx"},
		Dork{"google", "filetype:opus"},
	},
	"video/ogg": []Dork{
		Dork{"google", "filetype:ogg"},
		Dork{"google", "filetype:ogv"},
		Dork{"google", "filetype:oga"},
		Dork{"google", "filetype:ogx"},
		Dork{"google", "filetype:ogm"},
		Dork{"google", "filetype:spx"},
		Dork{"google", "filetype:opus"},
	},
	"audio/ogg": []Dork{
		Dork{"google", "filetype:ogg"},
		Dork{"google", "filetype:ogv"},
		Dork{"google", "filetype:oga"},
		Dork{"google", "filetype:ogx"},
		Dork{"google", "filetype:ogm"},
		Dork{"google", "filetype:spx"},
		Dork{"google", "filetype:opus"},
	},
	"image/jpeg": {
		Dork{"google", "filetype:jpeg"},
		Dork{"google", "filetype:jpg"},
		Dork{"google", "filetype:jpe"},
		Dork{"google", "filetype:jfif"},
		Dork{"google", "filetype:jfi"},
		Dork{"google", "filetype:jif"},
	},
	"image/png": []Dork{
		Dork{"google", "filetype:png"},
	},
	"image/gif": []Dork{
		Dork{"google", "filetype:gif"},
	},
	"image/webp": []Dork{
		Dork{"google", "filetype:webp"},
	},
	"audio/x-flac": []Dork{
		Dork{"google", "filetype:flac"},
	},
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document": []Dork{
		Dork{"google", "filetype:docx"},
	},
	"application/vnd.openxmlformats-officedocument.wordprocessingml.template": []Dork{
		Dork{"google", "filetype:dotx"},
	},
	"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet": []Dork{
		Dork{"google", "filetype:xlsx"},
	},
	"application/vnd.openxmlformats-officedocument.presentationml.presentation": []Dork{
		Dork{"google", "filetype:pptx"},
	},
	"application/vnd.openxmlformats-officedocument.presentationml.template": []Dork{
		Dork{"google", "filetype:potx"},
	},
	"application/vnd.openxmlformats-officedocument.presentationml.slideshow": []Dork{
		Dork{"google", "filetype:ppsx"},
	},
	"application/octet-stream": []Dork{
		Dork{"google", "inurl:.DS_Store intitle:index.of"},
		Dork{"google", "intext:.DS_Store & intitle:index -github"},
	},
	"application/vnd.oasis.opendocument.text": []Dork{
		Dork{"google", "filetype:odt"},
	},
	"application/vnd.oasis.opendocument.spreadsheet": []Dork{
		Dork{"google", "filetype:ods"},
	},
	"application/vnd.oasis.opendocument.presentation": []Dork{
		Dork{"google", "filetype:odp"},
	},
	"application/x-shockwave-flash": []Dork{
		Dork{"google", "filetype:swf"},
	},
	"application/gpx": []Dork{
		Dork{"google", "intitle:index.of +?last modified? +?parent directory? +(gpx) -htm -html -php -asp"},
		Dork{"google", "filetype:gpx"},
	},
	"application/gpx+xml": []Dork{
		Dork{"google", "intitle:index.of +?last modified? +?parent directory? +(gpx) -htm -html -php -asp"},
		Dork{"google", "filetype:gpx"},
	},
	"application/postscript": []Dork{
		Dork{"google", "filetype:ps"},
	},
}

func (dl dorkLib) GetByType(typ string) []Dork {
	if _, ok := dl[typ]; ok {
		return dl[typ]
	}
	return nil
}

func (dl dorkLib) GetByEngine(typ, engine string) []Dork {
	dorks := dl.GetByType(typ)
	if dorks != nil {
		var d []Dork
		for _, dork := range dorks {
			if dork.Engine == engine {
				d = append(d, dork)
			}
		}
		return d
	}
	return dorks
}

// Dork defines a dork for a particular engine
// TODO: Load dorks from a custom json file
type Dork struct {
	Engine string `json:"engine"`
	String string `json:"string"`
}

// GetEngine returns dork engine
func (d *Dork) GetEngine() string {
	return d.Engine
}

// GetSrting returns query string for dorking
func (d *Dork) GetSrting() string {
	return d.String
}

// UpdateDork updates the dork values
func (d *Dork) UpdateDork(engine, site, term string) {
	if d.Engine != engine {
		d.Engine = engine
	}
	if site != "" {
		d.String = fmt.Sprintf("site:%s %s", site, term)
	}
	if term != "" {
		if d.String != "" {
			d.String = fmt.Sprintf("%s +\"%s\"", d.String, term)
		} else {
			d.String = fmt.Sprintf("+\"%s\"", term)
		}
	} else {
		d.String = term
	}
}

// UpdateEngine unpdates dork's engine
func (d *Dork) UpdateEngine(engine string) {
	if d.Engine != engine {
		d.Engine = engine
	}
}

// UpdateString unpdates dork's search string
func (d *Dork) UpdateString(site, term string) {
	if d.String == "" {
		d.String = term
	} else {
		if term != "" {
			if d.String != "" {
				d.String = fmt.Sprintf("%s +\"%s\"", d.String, term)
			} else {
				d.String = fmt.Sprintf("+\"%s\"", term)
			}
		}
	}
	if site != "" {
		d.String = fmt.Sprintf("site:%s %s", site, d.String)
	}
}

// NewDork returns a new dork
func NewDork(engine, term string) Dork {
	return Dork{Engine: engine, String: term}
}

// Dorker defines a dorker structure. Used to store dorks.
type Dorker struct {
	userAgent string
	timeout   int
	depth     int
	Dorks     []Dork
}

// NewDorker returns a new dorker
func NewDorker(userAgent string, timeout, depth int) *Dorker {
	if depth < 1 {
		depth = DEFAULT_DEPTH
	}

	if timeout < 1 {
		timeout = DEFAULT_TIMEOUT
	}

	return &Dorker{
		userAgent: userAgent,
		timeout:   timeout,
		depth:     depth,
	}
}

// AddDork adds a dork in the dorker object
func (d *Dorker) AddDork(dork Dork) {
	d.Dorks = append(d.Dorks, dork)
}
