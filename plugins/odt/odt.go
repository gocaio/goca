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

package odt

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"

	"github.com/gocaio/goca"
	"github.com/kevinborras/metagopenoffice"
	log "github.com/sirupsen/logrus"
)

const plugName = "odt"

func init() {
	goca.RegisterPlugin(plugName, goca.Plugin{
		Type:    "odt",
		Assoc:   []string{"application/vnd.oasis.opendocument.text"}, // odt
		Action:  setup,
		Matcher: nil,
	})
}

var odt *odtMetaExtractor

func setup(m goca.Manager) error {
	odt = new(odtMetaExtractor)
	odt.Manager = m
	odt.Subscribe(goca.Topics["NewTarget"], odt.readODT)
	return nil
}

type odtMetaExtractor struct {
	goca.Manager
}

func (odt *odtMetaExtractor) readODT(target string, data []byte) {
	log.Debugf("[ODT] Received Data Length: %d - TARGET: %s\n", len(data), target)

	defer goca.PluginRecover("ODT", "readODT")

	var fileName = strconv.Itoa(rand.Int()) + ".odt"

	ioutil.WriteFile(os.TempDir()+"/"+fileName, data, 0777)
	file, err := os.Open(os.TempDir() + "/" + fileName)

	if err != nil {
		log.Debugf("[ODT] - Open Err: %s\n", err.Error())
	}
	file.Close()

	mdata, err := metagopenoffice.GetMetada(file)

	if err != nil {
		log.Debugf("[ODT] - Metagoffice Err: %s\n", err.Error())
	}

	if err == nil {
		out := goca.NewOutput()
		out.MainType = "ODT"
		out.Target = target
		out.Title = mdata.Meta.Title
		out.Description = mdata.Meta.Description
		out.Comment = mdata.Meta.Subject
		out.Lang = mdata.Meta.Language
		out.Producer = mdata.Meta.InitialCreator
		out.CreatorTool = mdata.Meta.Generator
		out.Keywords = mdata.Meta.Keyword
		out.ModifiedBy = mdata.Meta.Creator
		out.CreateDate = mdata.Meta.CreationDate
		out.ModifyDate = mdata.Meta.Date

		odt.Publish(goca.Topics["NewOutput"], plugName, target, out)
	} else {
		log.Debugf("[ODT] - Err: %s\n", err.Error())
	}
}
