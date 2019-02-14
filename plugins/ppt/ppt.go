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

package ppt

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"

	"github.com/gocaio/goca"
	"github.com/kevinborras/metagoffice"
	log "github.com/sirupsen/logrus"
)

const plugName = "ppt"

func init() {
	goca.RegisterPlugin(plugName, goca.Plugin{
		Type: "ppt",
		Assoc: []string{"application/vnd.ms-powerpoint", // ppt, pot, pps, ppa
			"application/vnd.openxmlformats-officedocument.presentationml.presentation", // pptx
			"application/vnd.openxmlformats-officedocument.presentationml.template",     // potx
			"application/vnd.openxmlformats-officedocument.presentationml.slideshow",    // ppsx
			"application/vnd.ms-powerpoint.addin.macroEnabled.12",                       // ppam
			"application/vnd.ms-powerpoint.presentation.macroEnabled.12",                // pptm, potm
			"application/vnd.ms-powerpoint.slideshow.macroEnabled.12",                   // ppsm
		},
		Action:  setup,
		Matcher: nil,
	})
}

var ppt *pptMetaExtractor

func setup(m goca.Manager) error {
	ppt = new(pptMetaExtractor)
	ppt.Manager = m
	ppt.Subscribe(goca.Topics["NewTarget"], ppt.readPPT)
	return nil
}

type pptMetaExtractor struct {
	goca.Manager
}

func (ppt *pptMetaExtractor) readPPT(target string, data []byte) {
	log.Debugf("[PPT] Received Data Length: %d - TARGET: %s\n", len(data), target)

	defer goca.PluginRecover("PPT", "readPPT")

	var fileName = strconv.Itoa(rand.Int()) + ".pptx"

	ioutil.WriteFile(os.TempDir()+"/"+fileName, data, 0777)
	file, err := os.Open(os.TempDir() + "/" + fileName)

	if err != nil {
		log.Debugf("[PPT] - Open Err: %s\n", err.Error())
	}
	file.Close()

	doc, err := metagoffice.GetContent(file)

	if err != nil {
		log.Debugf("[PPT] - Metagoffice Err: %s\n", err.Error())
	}

	if err == nil {
		out := goca.NewOutput()
		out.MainType = "PPTX"
		out.Target = target
		out.Title = doc.Title
		out.Comment = doc.Subject
		out.Producer = doc.Creator
		out.Keywords = doc.Keywords
		out.Description = doc.Description
		out.ModifiedBy = doc.LastModifiedBy
		out.DocumentID = doc.Revision
		out.CreateDate = doc.Created
		out.ModifyDate = doc.Modified
		out.Category = doc.Category

		ppt.Publish(goca.Topics["NewOutput"], plugName, target, out)
	} else {
		log.Debugf("[PPTX] - Err: %s\n", err.Error())
	}
}
