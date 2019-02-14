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

package doc

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"

	"github.com/gocaio/goca"
	"github.com/kevinborras/metagoffice"
	log "github.com/sirupsen/logrus"
)

const plugName = "doc"

func init() {
	goca.RegisterPlugin(plugName, goca.Plugin{
		Type: "doc",
		Assoc: []string{"application/vnd.openxmlformats-officedocument.wordprocessingml.document", // docx
			"application/msword", // doc, dot
			"application/zip",    // doc, dot
			"application/vnd.openxmlformats-officedocument.wordprocessingml.template", // dotx
			"application/vnd.ms-word.document.macroEnabled.12",                        // docm, dotm
		},
		Action:  setup,
		Matcher: nil,
	})
}

var doc *docMetaExtractor

func setup(m goca.Manager) error {
	doc = new(docMetaExtractor)
	doc.Manager = m
	doc.Subscribe(goca.Topics["NewTarget"], doc.readDOC)
	return nil
}

type docMetaExtractor struct {
	goca.Manager
}

func (doc *docMetaExtractor) readDOC(target string, data []byte) {
	log.Debugf("[DOC] Received Data Length: %d - TARGET: %s\n", len(data), target)

	defer goca.PluginRecover("DOC", "readDOC")

	var fileName = strconv.Itoa(rand.Int()) + ".docx"

	ioutil.WriteFile(os.TempDir()+"/"+fileName, data, 0777)
	file, err := os.Open(os.TempDir() + "/" + fileName)

	if err != nil {
		log.Debugf("[DOC] - Open Err: %s\n", err.Error())
	}
	file.Close()

	mdata, err := metagoffice.GetContent(file)

	if err != nil {
		log.Debugf("[DOC] - Metagoffice Err: %s\n", err.Error())
	}

	if err == nil {
		out := goca.NewOutput()
		out.MainType = "DOCX"
		out.Target = target
		out.Title = mdata.Title
		out.Comment = mdata.Subject
		out.Producer = mdata.Creator
		out.Keywords = mdata.Keywords
		out.Description = mdata.Description
		out.ModifiedBy = mdata.LastModifiedBy
		out.DocumentID = mdata.Revision
		out.CreateDate = mdata.Created
		out.ModifyDate = mdata.Modified
		out.Category = mdata.Category

		doc.Publish(goca.Topics["NewOutput"], plugName, target, out)
	} else {
		log.Debugf("[DOCX] - Err: %s\n", err.Error())
	}
}
