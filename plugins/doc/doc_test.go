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

package doc

import (
	"os"
	"strings"
	"testing"

	"github.com/gocaio/goca"
	"github.com/gocaio/goca/gocaTesting"
)

// Test server URL.
var testserver = os.Getenv("GOCA_TEST_SERVER")

// T is a global reference for the test. This allows us to use *testing.T
// methods anywhere
var T *testing.T

// TestReadDOC tests the read on DOC files
func TestReadDOC(t *testing.T) {
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
	case "Doc1.docx":
		validateCaseA(out)
	case "Doc2.docx":
		validateCaseB(out)
	case "Doc3.docx":
		validateCaseC(out)
	}
}

func validateCaseA(out *goca.Output) {
	if out.MainType != "DOCX" {
		T.Errorf("expected DOCX but found %s", out.MainType)
	}
	if out.Title != "2018–2019 Statewide Testing Schedule and Administration Deadlines, January 18, 2019" {
		T.Errorf("expected \"2018–2019 Statewide Testing Schedule and Administration Deadlines, January 18, 2019\" but found %s", out.Title)
	}
	if out.Comment != "" {
		T.Errorf("expected \"\" but found %s", out.Comment)
	}
	if out.Producer != "DESE" {
		T.Errorf("expected \"DESE\" but found %s", out.Producer)
	}
	if out.Keywords != "" {
		T.Errorf("expected \"\" but found %s", out.Keywords)
	}
	if out.Description != "" {
		T.Errorf("expected \"\" but found %s", out.MainType)
	}
	if out.ModifiedBy != "Zou, Dong (EOE)" {
		T.Errorf("expected \"Zou, Dong (EOE)\" but found %s", out.ModifiedBy)
	}
	if out.DocumentID != "16" {
		T.Errorf("expected \"16\" but found %s", out.DocumentID)
	}
	if out.CreateDate != "2018-10-24T14:04:00Z" {
		T.Errorf("expected \"2018-10-24T14:04:00Z\" but found %s", out.CreateDate)
	}
	if out.ModifyDate != "2019-01-18T20:59:00Z" {
		T.Errorf("expected \"2019-01-18T20:59:00Z\" but found %s", out.ModifyDate)
	}
	if out.Category != "" {
		T.Errorf("expected \"\" but found %s", out.Category)
	}
}

func validateCaseB(out *goca.Output) {
	if out.MainType != "DOCX" {
		T.Errorf("expected DOCX but found %s", out.MainType)
	}
	if out.Title != "Arizona’s Instrument to Measure Standards" {
		T.Errorf("expected \"Arizona’s Instrument to Measure Standards\" but found %s", out.Title)
	}
	if out.Comment != "" {
		T.Errorf("expected \"\" but found %s", out.Comment)
	}
	if out.Producer != "Network Services" {
		T.Errorf("expected \"Network Services\" but found %s", out.Producer)
	}
	if out.Keywords != "" {
		T.Errorf("expected \"\" but found %s", out.Keywords)
	}
	if out.Description != "" {
		T.Errorf("expected \"\" but found %s", out.MainType)
	}
	if out.ModifiedBy != "" {
		T.Errorf("expected \"\" but found %s", out.ModifiedBy)
	}
	if out.DocumentID != "" {
		T.Errorf("expected \"16\" but found %s", out.DocumentID)
	}
	if out.CreateDate != "2018-04-30T17:15:26Z" {
		T.Errorf("expected \"2018-04-30T17:15:26Z\" but found %s", out.CreateDate)
	}
	if out.ModifyDate != "2018-04-30T17:15:26Z" {
		T.Errorf("expected \"2018-04-30T17:15:26Z\" but found %s", out.ModifyDate)
	}
	if out.Category != "" {
		T.Errorf("expected \"\" but found %s", out.Category)
	}
}

func validateCaseC(out *goca.Output) {
	if out.MainType != "DOCX" {
		T.Errorf("expected DOCX but found %s", out.MainType)
	}
	if out.Title != "MCAS Permission Request to Test in Alternate Setting Form: 2018" {
		T.Errorf("expected \"MCAS Permission Request to Test in Alternate Setting Form: 2018\" but found %s", out.Title)
	}
	if out.Comment != "" {
		T.Errorf("expected \"\" but found %s", out.Comment)
	}
	if out.Producer != "" {
		T.Errorf("expected \"\" but found %s", out.Producer)
	}
	if out.Keywords != "" {
		T.Errorf("expected \"\" but found %s", out.Keywords)
	}
	if out.Description != "" {
		T.Errorf("expected \"\" but found %s", out.MainType)
	}
	if out.ModifiedBy != "" {
		T.Errorf("expected \"\" but found %s", out.ModifiedBy)
	}
	if out.DocumentID != "1" {
		T.Errorf("expected \"1\" but found %s", out.DocumentID)
	}
	if out.CreateDate != "2018-03-20T20:23:00Z" {
		T.Errorf("expected \"2018-03-20T20:23:00Z\" but found %s", out.CreateDate)
	}
	if out.ModifyDate != "2018-03-20T20:23:00Z" {
		T.Errorf("expected \"2018-03-20T20:23:00Z\" but found %s", out.ModifyDate)
	}
	if out.Category != "" {
		T.Errorf("expected \"\" but found %s", out.Category)
	}
}
