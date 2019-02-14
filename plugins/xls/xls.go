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

package xls

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"

	"github.com/gocaio/goca"
	"github.com/kevinborras/metagoffice"
	log "github.com/sirupsen/logrus"
)

const plugName = "xls"

func init() {
	goca.RegisterPlugin(plugName, goca.Plugin{
		Type: "xls",
		Assoc: []string{"application/vnd.ms-excel", // xls, xlt, xla
			"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",    // xlsx
			"application/vnd.openxmlformats-officedocument.spreadsheetml.template", // xltx
			"application/vnd.ms-excel.sheet.macroEnabled.12",                       // xlsm
			"application/vnd.ms-excel.template.macroEnabled.12",                    // xltm
			"application/vnd.ms-excel.addin.macroEnabled.12",                       // xlam
			"application/vnd.ms-excel.sheet.binary.macroEnabled.12",                // xlsb
		},
		Action:  setup,
		Matcher: nil,
	})
}

var xls *xlsMetaExtractor

func setup(m goca.Manager) error {
	xls = new(xlsMetaExtractor)
	xls.Manager = m
	xls.Subscribe(goca.Topics["NewTarget"], xls.readXLS)
	return nil
}

type xlsMetaExtractor struct {
	goca.Manager
}

func (xls *xlsMetaExtractor) readXLS(target string, data []byte) {
	log.Debugf("[XLS] Received Data Length: %d - TARGET: %s\n", len(data), target)

	defer goca.PluginRecover("XLS", "readXLS")

	var fileName = strconv.Itoa(rand.Int()) + ".xlsx"

	ioutil.WriteFile(os.TempDir()+"/"+fileName, data, 0777)
	file, err := os.Open(os.TempDir() + "/" + fileName)

	if err != nil {
		log.Debugf("[XLS] - Open Err: %s\n", err.Error())
	}
	file.Close()

	doc, err := metagoffice.GetContent(file)

	if err != nil {
		log.Debugf("[XLS] - Metagoffice Err: %s\n", err.Error())
	}

	if err == nil {
		out := goca.NewOutput()
		out.MainType = "XLSX"
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

		xls.Publish(goca.Topics["NewOutput"], plugName, target, out)
	} else {
		log.Debugf("[XLS] - Err: %s\n", err.Error())
	}
}
