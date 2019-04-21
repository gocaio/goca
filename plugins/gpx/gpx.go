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

package gpx

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"

	gpxParser "github.com/tkrajina/gpxgo/gpx"
	"github.com/gocaio/goca"
	log "github.com/sirupsen/logrus"
)

const plugName = "gpx"

func init() {
	goca.RegisterPlugin(plugName, goca.Plugin{
		Type:    plugName,
		Assoc:   []string{"application/gpx",      // GPX
				  "application/gpx+xml",  // GPX+XML
			 },
		Action:  setup,
		Matcher: nil,
	})
}

var gpx *gpxMetaExtractor

func setup(m goca.Manager) error {
	gpx = new(gpxMetaExtractor)
	gpx.Manager = m
	gpx.Subscribe(goca.Topics["NewTarget"], gpx.readGPX)
	return nil
}

type gpxMetaExtractor struct {
	goca.Manager
}

func (gpx *gpxMetaExtractor) readGPX(target string, data []byte) {
	log.Debugf("[GPX] Received Data Length: %d - TARGET: %s\n", len(data), target)

	defer goca.PluginRecover("GPX", "readGPX")

	var fileName = strconv.Itoa(rand.Int()) + ".gpx"

	ioutil.WriteFile(os.TempDir()+"/"+fileName, data, 0777)
	
	doc, err := gpxParser.ParseFile(os.TempDir() + "/" + fileName)

	if err == nil && doc != nil {
		out := goca.NewOutput()
		out.MainType = "GPX"
		out.Target = target
		out.DocumentID = doc.Version
		out.Producer = doc.Creator
		out.Title = doc.Name
		out.Description = doc.Description
		out.ModifiedBy = doc.AuthorName
		out.Email = doc.AuthorEmail
		out.CreatorTool = doc.Copyright
		out.Year = doc.CopyrightYear
		out.Comment = doc.Link
		out.Keywords = doc.Keywords

		gpx.Publish(goca.Topics["NewOutput"], plugName, target, out)
	} else {
		log.Debugf("[GPX] - Err: %s\n", err.Error())
	}
}
