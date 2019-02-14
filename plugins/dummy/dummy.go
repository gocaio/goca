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

package dummy

import (
	"fmt"

	goca "github.com/gocaio/Goca"
)

// plugName is the plugin name, this constant variable is usefull
// to refer the plugin name anywhere
const plugName = "dummy"

// dummy is a global variable that represents the plugin it self
var dummy *dummyMetaExtractor

func init() {
	// debug message
	fmt.Println("Initiating dummy plugin")

	// Each plugin needs to call this function in terms to get
	// registered in Goca
	goca.RegisterPlugin(plugName, goca.Plugin{
		Type:    "goca",
		Assoc:   []string{"text/plain"},
		Action:  setup,
		Matcher: nil,
	})
}

// this is the entrypoint. Goca will call this function once
func setup(m goca.Manager) error {
	// keep the context in a dummyMetaExtractor struct
	dummy = new(dummyMetaExtractor)
	dummy.Manager = m
	// Each plugin is responsible to subscribe to topics.
	// For now, there is only one topic.
	dummy.Subscribe(goca.Topics["NewTarget"], dummy.handler)
	return nil
}

// dummyMetaExtractor will keep the goca.Manager. This allows the plugin to
// perform publish/subscribe at any time
type dummyMetaExtractor struct {
	goca.Manager
}

// handler is the handler that will process the event from the previous subscription
func (dummy *dummyMetaExtractor) handler(url string, msg []byte) {
	var res string
	// Each plugin needs an goca.Output object to emit back the resutls
	out := goca.NewOutput()

	// Do the plugin job
	if len(msg) >= 100 {
		res = fmt.Sprintf("Got: %s\n", string(msg)[:100])
	} else {
		res = fmt.Sprintf("Got: %s\n", string(msg))
	}

	out.MainType = "dummyType"
	out.Title = res

	// Emit back the results using a NewOutput event
	dummy.Publish(goca.Topics["NewOutput"], plugName, url, out)
}
