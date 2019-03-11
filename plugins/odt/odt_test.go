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

// TestReadODT tests the reading over ODT files
func TestReadODT(t *testing.T) {
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
	case "Odt1.docx":
		validateCaseA(out)
	case "Odt2.docx":
		validateCaseB(out)
	case "Odt3.docx":
		validateCaseC(out)
	}
}

func validateCaseA(out *goca.Output) {
	if out.MainType != "ODT" {
		T.Errorf("expected ODT but found %s", out.MainType)
	}
	if out.Title != "Laboratory Manual for DC Electrical Circuits" {
		T.Errorf("expected \"Laboratory Manual for DC Electrical Circuits\" but found %s", out.Title)
	}
	if out.Description != "Creative Commons NC-BY-SA license" {
		T.Errorf("expected \"Creative Commons NC-BY-SA license\" but found %s", out.Description)
	}
	if out.Comment != "DC Circuits" {
		T.Errorf("expected \"DC Circuits\" but found %s", out.Comment)
	}
	if out.Lang != "" {
		T.Errorf("expected \"\" but found %s", out.Lang)
	}
	if out.CreatorTool != "OpenOffice/4.1.2$Win32 OpenOffice.org_project/412m3$Build-9782" {
		T.Errorf("expected \"OpenOffice/4.1.2$Win32 OpenOffice.org_project/412m3$Build-9782\" but found %s", out.CreatorTool)
	}
	if out.Producer != "James M. Fiore" {
		T.Errorf("expected \"James M. Fiore\" but found %s", out.Producer)
	}
	if out.Keywords != "DC circuits" {
		T.Errorf("expected \"DC circuits\" but found %s", out.Keywords)
	}
	if out.ModifiedBy != "J Fiore" {
		T.Errorf("expected \"J Fiore\" but found %s", out.ModifiedBy)
	}
	if out.CreateDate != "2017-09-15T16:45:00" {
		T.Errorf("expected \"2017-09-15T16:45:00\" but found %s", out.CreateDate)
	}
	if out.ModifyDate != "2018-12-08T11:26:21.12" {
		T.Errorf("expected \"2018-12-08T11:26:21.12\" but found %s", out.ModifyDate)
	}
}

func validateCaseB(out *goca.Output) {
	if out.MainType != "ODT" {
		T.Errorf("expected ODT but found %s", out.MainType)
	}
	if out.Title != "Laboratory Manual for Computer Programming with Python and Multisim" {
		T.Errorf("expected \"Laboratory Manual for Computer Programming with Python and Multisim\" but found %s", out.Title)
	}
	if out.Description != "Creative Commons NC-BY-SA license" {
		T.Errorf("expected \"Creative Commons NC-BY-SA license\" but found %s", out.Description)
	}
	if out.Comment != "Computer Programming" {
		T.Errorf("expected \"Computer Programming\" but found %s", out.Comment)
	}
	if out.Lang != "" {
		T.Errorf("expected \"\" but found %s", out.Lang)
	}
	if out.CreatorTool != "OpenOffice/4.1.2$Win32 OpenOffice.org_project/412m3$Build-9782" {
		T.Errorf("expected \"OpenOffice/4.1.2$Win32 OpenOffice.org_project/412m3$Build-9782\" but found %s", out.CreatorTool)
	}
	if out.Producer != "James Fiore" {
		T.Errorf("expected \"James Fiore\" but found %s", out.Producer)
	}
	if out.Keywords != "Python simulation" {
		T.Errorf("expected \"Python simulation\" but found %s", out.Keywords)
	}
	if out.ModifiedBy != "j j" {
		T.Errorf("expected \"j j\" but found %s", out.ModifiedBy)
	}
	if out.CreateDate != "2017-09-19T16:26:00" {
		T.Errorf("expected \"2017-09-19T16:26:00\" but found %s", out.CreateDate)
	}
	if out.ModifyDate != "2017-10-05T09:38:01.28" {
		T.Errorf("expected \"2017-10-05T09:38:01.28\" but found %s", out.ModifyDate)
	}
}

func validateCaseC(out *goca.Output) {
	if out.MainType != "ODT" {
		T.Errorf("expected ODT but found %s", out.MainType)
	}
	if out.Title != "Revised Newswest Publishing Schedule 2018" {
		T.Errorf("expected \"Revised Newswest Publishing Schedule 2018\" but found %s", out.Title)
	}
	if out.Description != "Revised April 2018Dates after May have changed!" {
		T.Errorf("expected \"Revised April 2018Dates after May have changed!\" but found %s", out.Description)
	}
	if out.Comment != "" {
		T.Errorf("expected \"\" but found %s", out.Comment)
	}
	if out.Lang != "" {
		T.Errorf("expected \"\" but found %s", out.Lang)
	}
	if out.CreatorTool != "OpenOffice/4.1.4$Win32 OpenOffice.org_project/414m5$Build-9788" {
		T.Errorf("expected \"OpenOffice/4.1.4$Win32 OpenOffice.org_project/414m5$Build-9788\" but found %s", out.CreatorTool)
	}
	if out.Producer != "" {
		T.Errorf("expected \"\" but found %s", out.Producer)
	}
	if out.Keywords != "" {
		T.Errorf("expected \"\" but found %s", out.Keywords)
	}
	if out.ModifiedBy != "" {
		T.Errorf("expected \"\" but found %s", out.ModifiedBy)
	}
	if out.CreateDate != "2009-11-23T06:31:57" {
		T.Errorf("expected \"2009-11-23T06:31:57\" but found %s", out.CreateDate)
	}
	if out.ModifyDate != "2012-10-02T12:33:04" {
		T.Errorf("expected \"2012-10-02T12:33:04\" but found %s", out.ModifyDate)
	}
}
