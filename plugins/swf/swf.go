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

package swf

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"

	"github.com/RangelReale/swfinfo"
	"github.com/gocaio/goca"
	log "github.com/sirupsen/logrus"
)

const plugName = "swf"

func init() {
	goca.RegisterPlugin(plugName, goca.Plugin{
		Type:    plugName,
		Assoc:   []string{"application/x-shockwave-flash"}, // SWF
		Action:  setup,
		Matcher: nil,
	})
}

var swf *swfMetaExtractor

func setup(m goca.Manager) error {
	swf = new(swfMetaExtractor)
	swf.Manager = m
	swf.Subscribe(goca.Topics["NewTarget"], swf.readSWF)
	return nil
}

type swfMetaExtractor struct {
	goca.Manager
}

func (swf *swfMetaExtractor) readSWF(target string, data []byte) {
	log.Debugf("[SWF] Received Data Length: %d - TARGET: %s\n", len(data), target)

	defer goca.PluginRecover("SWF", "readSWF")

	var fileName = strconv.Itoa(rand.Int()) + ".swf"

	ioutil.WriteFile(os.TempDir()+"/"+fileName, data, 0777)

	doc, err := swfinfo.Open(os.TempDir() + "/" + fileName)

	if err == nil && doc != nil {
		out := goca.NewOutput()
		out.MainType = "swf"
		out.Target = target

		out.Duration = doc.Duration().String()
		out.Version = doc.Version
		out.FrameRate = doc.FrameRate
		out.ImageWidth = doc.FrameSize.Xmax.Pixels()
		out.ImageHeight = doc.FrameSize.Ymax.Pixels()
		out.FrameCount = doc.FrameCount
		out.Comment = doc.Compression.String()

		swf.Publish(goca.Topics["NewOutput"], plugName, target, out)
	} else {
		log.Debugf("[SWF] - Err: %s\n", err.Error())
	}
}
