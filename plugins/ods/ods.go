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

package ods

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"

	"github.com/gocaio/Goca"
	"github.com/kevinborras/metagopenoffice"
	log "github.com/sirupsen/logrus"
)

const plugName = "ods"

func init() {
	goca.RegisterPlugin(plugName, goca.Plugin{
		Type:    "ods",
		Assoc:   []string{"application/vnd.oasis.opendocument.spreadsheet"}, //ods
		Action:  setup,
		Matcher: nil,
	})
}

var ods *odsMetaExtractor

func setup(m goca.Manager) error {
	ods = new(odsMetaExtractor)
	ods.Manager = m
	ods.Subscribe(goca.Topics["NewTarget"], ods.readODS)
	return nil
}

type odsMetaExtractor struct {
	goca.Manager
}

func (ods *odsMetaExtractor) readODS(target string, data []byte) {
	log.Debugf("[ODS] Received Data Length: %d - TARGET: %s\n", len(data), target)

	defer goca.PluginRecover("ODS", "readODS")

	var fileName = strconv.Itoa(rand.Int()) + ".ods"

	ioutil.WriteFile(os.TempDir()+"/"+fileName, data, 0777)
	file, err := os.Open(os.TempDir() + "/" + fileName)

	if err != nil {
		log.Debugf("[ODS] - Open Err: %s\n", err.Error())
	}
	file.Close()

	mdata, err := metagopenoffice.GetMetada(file)

	if err != nil {
		log.Debugf("[ODS] - Metagoffice Err: %s\n", err.Error())
	}

	if err == nil {
		out := goca.NewOutput()
		out.MainType = "ODS"
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

		ods.Publish(goca.Topics["NewOutput"], plugName, target, out)
	} else {
		log.Debugf("[ODS] - Err: %s\n", err.Error())
	}
}
