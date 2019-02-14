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

package odp

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"

	"github.com/gocaio/Goca"
	"github.com/kevinborras/metagopenoffice"
	log "github.com/sirupsen/logrus"
)

const plugName = "odp"

func init() {
	goca.RegisterPlugin(plugName, goca.Plugin{
		Type:    "odp",
		Assoc:   []string{"application/vnd.oasis.opendocument.presentation"}, // odp
		Action:  setup,
		Matcher: nil,
	})
}

var odp *odpMetaExtractor

func setup(m goca.Manager) error {
	odp = new(odpMetaExtractor)
	odp.Manager = m
	odp.Subscribe(goca.Topics["NewTarget"], odp.readODP)
	return nil
}

type odpMetaExtractor struct {
	goca.Manager
}

func (odp *odpMetaExtractor) readODP(target string, data []byte) {
	log.Debugf("[ODP] Received Data Length: %d - TARGET: %s\n", len(data), target)

	defer goca.PluginRecover("ODP", "readODP")

	var fileName = strconv.Itoa(rand.Int()) + ".odp"

	ioutil.WriteFile(os.TempDir()+"/"+fileName, data, 0777)
	file, err := os.Open(os.TempDir() + "/" + fileName)

	if err != nil {
		log.Debugf("[ODP] - Open Err: %s\n", err.Error())
	}
	file.Close()

	mdata, err := metagopenoffice.GetMetada(file)

	if err != nil {
		log.Debugf("[ODP] - Metagoffice Err: %s\n", err.Error())
	}

	if err == nil {
		out := goca.NewOutput()
		out.MainType = "ODP"
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

		odp.Publish(goca.Topics["NewOutput"], plugName, target, out)
	} else {
		log.Debugf("[ODP] - Err: %s\n", err.Error())
	}
}
