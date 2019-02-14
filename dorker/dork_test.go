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
	"strings"
	"testing"
)

func TestDorkGoogle(t *testing.T) {
	d := NewDorker("The_Goca_v0.1", 30, 1)
	d.AddDork(Dork{
		Engine: "google",
		String: "Chema Alonso",
	})

	urls := d.Google()

	if len(urls) == 0 {
		t.Errorf("expected some url but found %d", len(urls))
	}
}

func TestGetType(t *testing.T) {
	dl := dorkLib{
		"pdf": []Dork{
			Dork{"google", "filetype:pdf +\"%s\""},
			Dork{"bing", "filetype:pdf +\"%s\""},
		},
		"mp3": []Dork{
			Dork{"google", "filetype:mp3 +\"%s\""},
		},
	}

	typExpected := len(dl["pdf"])
	typ := dl.GetByType("pdf")

	if len(typ) != typExpected {
		t.Errorf("expected %d Dork but found %d", len(typ), typExpected)
	}

	typExpected = len(dl["mp3"])
	typ = dl.GetByType("mp3")

	if len(typ) != typExpected {
		t.Errorf("expected %d Dork but found %d", len(typ), typExpected)
	}
}

func TestGetByEngine(t *testing.T) {
	dl := dorkLib{
		"pdf": []Dork{
			Dork{"google", "filetype:pdf +\"%s\""},
			Dork{"bing", "filetype:pdf +\"%s\""},
		},
		"mp3": []Dork{
			Dork{"google", "filetype:mp3 +\"%s\""},
		},
	}

	typExpected := 1
	typ := dl.GetByEngine("pdf", "google")

	if len(typ) != typExpected {
		t.Errorf("expected %d Dork but found %d", len(typ), typExpected)
	}

	typExpected = 1
	typ = dl.GetByEngine("pdf", "bing")

	if len(typ) != typExpected {
		t.Errorf("expected %d Dork but found %d", len(typ), typExpected)
	}

	typExpected = 0
	typ = dl.GetByEngine("mp3", "bing")

	if len(typ) != typExpected {
		t.Errorf("expected %d Dork but found %d", len(typ), typExpected)
	}
}

func TestUpdateDork(t *testing.T) {
	d := NewDork("bing", "")

	d.UpdateDork("bing", "", "torm")

	if d.Engine != "bing" {
		t.Errorf("expected %s as engine but found %s", "bing", d.Engine)
	}

	if d.String != "+\"torm\"" {
		t.Errorf("expected +\"torm\" as String but found %s", d.String)
	}

	if strings.HasPrefix(d.String, "site") {
		t.Errorf("expected String without site limit, but found %s", d.String)
	}

	d = NewDork("bing", "term")

	d.UpdateDork("google", "google.com", "torm")

	if d.Engine != "google" {
		t.Errorf("expected %s as engine but found %s", "google", d.String)
	}

	if !strings.HasPrefix(d.String, "site") {
		t.Errorf("expected String to have site as prexif, but sit not found")
	}
}

func TestUpdateEngine(t *testing.T) {
	d := NewDork("google", "")

	d.UpdateEngine("bing")

	if d.Engine != "bing" {
		t.Errorf("expected %s as engine but found %s", "bing", d.Engine)
	}

	if d.String != "" {
		t.Errorf("expected '%s' as String but found %s", "", d.String)
	}
}

func TestUpdateString(t *testing.T) {
	d := NewDork("google", "")

	d.UpdateString("", "term")

	if d.Engine != "google" {
		t.Errorf("expected %s as engine but found %s", "google", d.Engine)
	}

	if d.String != "term" {
		t.Errorf("expected '%s' as String but found %s", "term", d.String)
	}

	if strings.HasPrefix(d.String, "site") {
		t.Errorf("expecting string without site limit, but found site")
	}

	d.UpdateString("google.com", "torm")
	if !strings.HasPrefix(d.String, "site:") {
		t.Errorf("expected String to have site, but found %s", d.String)
	}
}

// func TestNewDorker(t *testing.T)          {}
// func TestDorkBing(t *testing.T)           {}
// func TestDorkDuckDuckGo(t *testing.T)     {}
// func TestDorkYahoo(t *testing.T)          {}
// func TestDorkAsk(t *testing.T)            {}
// func TestDorkTLSCertificate(t *testing.T) {}
