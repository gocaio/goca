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

package pdf

import (
	"bytes"
	"strings"

	"github.com/gocaio/goca"
	log "github.com/sirupsen/logrus"
	"trimmer.io/go-xmp/xmp"
)

const plugName = "pdf"

func init() {
	goca.RegisterPlugin(plugName, goca.Plugin{
		Type:    "pdf",
		Assoc:   []string{"application/pdf"},
		Action:  setup,
		Matcher: nil,
	})
}

var pdf *pdfMetaExtractor

func setup(m goca.Manager) error {
	pdf = new(pdfMetaExtractor)
	pdf.Manager = m
	pdf.Subscribe(goca.Topics["NewTarget"], pdf.readPDF)
	return nil
}

type pdfMetaExtractor struct {
	goca.Manager
}

func (pdf *pdfMetaExtractor) readPDF(target string, data []byte) {
	log.Debugf("[PDF] Received Data Length: %d - TARGET: %s\n", len(data), target)

	defer goca.PluginRecover("PDF", "readPDF")

	reader := bytes.NewReader(data)

	doc, err := xmp.Scan(reader)
	if err == nil && doc != nil {
		out := goca.NewOutput()
		out.MainType = "PDF"
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
		pdf.Publish(goca.Topics["NewOutput"], plugName, target, out)
	} else {
		log.Debugf("[PDF] - Err: %s\n", err.Error())
	}
}
