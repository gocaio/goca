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

package dsstore

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/gehaxelt/ds_store"
	"github.com/gocaio/Goca"
	log "github.com/sirupsen/logrus"
)

const plugName = "dsstore"

func init() {
	goca.RegisterPlugin(plugName, goca.Plugin{
		Type:    "dsstore",
		Assoc:   []string{"application/octet-stream"}, //.DS_Store
		Action:  setup,
		Matcher: matcher,
	})
}

var dsstore *dsstoreMetaExtractor

func setup(m goca.Manager) error {
	dsstore = new(dsstoreMetaExtractor)
	dsstore.Manager = m
	dsstore.Subscribe(goca.Topics["NewTarget"], dsstore.readDSSTORE)
	return nil
}

func matcher(buf []byte) bool {
	return len(buf) > 1 && (buf[0] == 0x00 &&
		buf[1] == 0x00 &&
		buf[2] == 0x00 &&
		buf[3] == 0x01 &&
		buf[4] == 0x42 &&
		buf[5] == 0x75 &&
		buf[6] == 0x64 &&
		buf[7] == 0x31 &&
		buf[8] == 0x00)
}

type dsstoreMetaExtractor struct {
	goca.Manager
}

func (dsstore *dsstoreMetaExtractor) readDSSTORE(target string, data []byte) {
	log.Debugf("[DSSTORE] Received Data Length: %d - TARGET: %s\n", len(data), target)

	defer goca.PluginRecover("DSSTORE", "readDSSTORE")

	var dataStore []string
	var fileName = strconv.Itoa(rand.Int()) + ".DS_Store"

	ioutil.WriteFile(os.TempDir()+"/"+fileName, data, 0777)
	file, err := os.Open(os.TempDir() + "/" + fileName)

	if err != nil {
		log.Debugf("[DSSTORE] - Open Err: %s\n", err.Error())
	}
	file.Close()

	dat, err := ioutil.ReadFile(os.TempDir() + "/" + fileName)

	if err != nil {
		log.Debugf("[DSSTORE] - ioutil Err: %s\n", err.Error())
	}

	// mdata, err := ds_store.NewAllocator(dat)

	// println(mdata)

	// if err != nil {
	// 	log.Debugf("[DSSTORE] - DS_Store Err: %s\n", err.Error())
	// }

	a, err := ds_store.NewAllocator(dat)
	filenames, err := a.TraverseFromRootNode()

	if err != nil {
		log.Debugf("[DSSTORE] - DS_Store Err: %s\n", err.Error())
	}
	for _, f := range filenames {
		dataStore = append(dataStore, f)
	}

	dataString := strings.Join(dataStore, ", ")
	if err == nil {
		out := goca.NewOutput()
		out.MainType = "DSSTORE"
		out.Target = target
		out.Keywords = dataString

		dsstore.Publish(goca.Topics["NewOutput"], plugName, target, out)
	} else {
		log.Debugf("[DSTORE] - Err: %s\n", err.Error())
	}
}
