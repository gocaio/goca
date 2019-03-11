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
	"strings"
	"testing"

	"github.com/gocaio/goca"
	"github.com/gocaio/goca/gocaTesting"
)

// Test server URL.
var testserver = "https://test.goca.io"

// T is a global reference for the test. This allows us to use *testing.T
// methods anywhere
var T *testing.T

// TestReadDSSTORE tests the read on DS_Store files
func TestReadDSSTORE(t *testing.T) {
	T = t // Assignment t (*testing.T to a global T variable)
	// Get a controller
	ctrl := goca.NewControllerTest()
	// Subscribe a processOutput. The propper test will be placed in proccessOutput
	ctrl.Subscribe(goca.Topics["NewOutput"], processOutput)

	// Call the plugin entrypoint
	setup(ctrl)

	gocatesting.GetAssets(t, ctrl, testserver, plugName)
}

func processOutput(module, url string, out *goca.Output) {
	// We have to validate goca.Output according to the resource
	parts := strings.Split(out.Target, "/")

	switch parts[len(parts)-1] {
	case "DS_Store_1":
		validateCaseA(out)
	case "DS_Store_2":
		validateCaseB(out)
	case "DS_Store_3":
		validateCaseC(out)
	case "DS_Store_bad":
		validateCaseD(out)
	}
}

func validateCaseA(out *goca.Output) {
	if out.MainType != "DSSTORE" {
		T.Errorf("expected DSSTORE but found %s", out.MainType)
	}
	if out.Keywords != "Classification Template, Classification Template, Classification Template, Classification Template, Classification Template, Data Preprocessing Template, Data Preprocessing Template, Data Preprocessing Template, Data Preprocessing Template, Data Preprocessing Template" {
		T.Errorf("expected other value but found %s", out.CreateDate)
	}
}

func validateCaseB(out *goca.Output) {
	if out.MainType != "DSSTORE" {
		T.Errorf("expected DSSTORE but found %s", out.MainType)
	}
	if out.Keywords != "bam, bar, baz" {
		T.Errorf("expected \"bam, bar, baz\" but found %s", out.CreateDate)
	}
}

func validateCaseC(out *goca.Output) {
	if out.MainType != "DSSTORE" {
		T.Errorf("expected DSSTORE but found %s", out.MainType)
	}
	if out.Keywords != "., ., ., ., ., ., Pelis, Pelis, Pelis, Pelis, Pelis, Pelis, Pelis, Series, Series, Series, Series, Series, Series, Series, Series, Series, Series, Software, Software, Software, Software, Software, Software, Software, Software, Software" {
		T.Errorf("expected other value but found %s", out.CreateDate)
	}
}

func validateCaseD(out *goca.Output) {
	if out.MainType != "DSSTORE" {
		T.Errorf("expected DSSTORE but found %s", out.MainType)
	}
	if out.Keywords != "" {
		T.Errorf("expected \"\" but found %s", out.CreateDate)
	}
}
