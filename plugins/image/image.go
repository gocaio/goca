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

package image

import (
	"bytes"
	"strings"

	"github.com/gocaio/goca"
	log "github.com/sirupsen/logrus"
	"trimmer.io/go-xmp/xmp"
)

const plugName = "image"

func init() {
	goca.RegisterPlugin(plugName, goca.Plugin{
		Type: plugName,
		Assoc: []string{"image/jpeg", // JPEG
			"image/png",                 // PNG
			"image/gif",                 // GIF
			"image/webp",                // WEBP
			"image/vnd.adobe.photoshop", // PSD
			"application/postscript",    // EPS, AI, PS
		},
		Action:  setup,
		Matcher: nil,
	})
}

var img *imgMetaExtractor

func setup(m goca.Manager) error {
	img = new(imgMetaExtractor)
	img.Manager = m
	img.Subscribe(goca.Topics["NewTarget"], img.readIMG)
	return nil
}

type imgMetaExtractor struct {
	goca.Manager
}

func (img *imgMetaExtractor) readIMG(target string, data []byte) {
	log.Debugf("[IMG] Received Data Length: %d - TARGET: %s\n", len(data), target)

	defer goca.PluginRecover("IMG", "readIMG")

	reader := bytes.NewReader(data)

	doc, err := xmp.Scan(reader)
	if err == nil && doc != nil {
		out := goca.NewOutput()
		out.MainType = "img"
		out.Target = target
		p, _ := doc.ListPaths()
		for _, v := range p {
			key, _ := v.Path.Pop()
			switch key {
			case "format":
				out.ContentType = v.Value
			case "Producer":
				out.Producer = v.Value
			case "CreateDate":
				out.CreateDate = v.Value
			case "CreatorTool":
				out.CreatorTool = v.Value
			case "MetadataDate":
				out.MetadataDate = v.Value
			case "ModifyDate":
				out.ModifyDate = v.Value
			case "DocumentID":
				out.DocumentID = v.Value
			case "InstanceID":
				out.InstanceID = v.Value
			default:
				if strings.HasPrefix(key, "title") {
					out.Title = v.Value
				}
			}
		}
		img.Publish(goca.Topics["NewOutput"], plugName, target, out)
	} else {
		log.Debugf("[IMG] - Err: %s\n", err.Error())
	}
}
