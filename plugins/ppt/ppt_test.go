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

package ppt

import (
	"strings"
	"testing"

	"github.com/gocaio/goca"
	"github.com/gocaio/goca/testData"
)

// Test server URL.
var testserver = "https://test.goca.io"

// T is a global reference for the test. This allows us to use *testing.T
// methods anywhere
var T *testing.T

// TestReadPPT tests the read on PPT files
func TestReadPPT(t *testing.T) {
	T = t // Assignment t (*testing.T to a global T variable)
	// Get a controller
	ctrl := goca.NewControllerTest()
	// Subscribe a processOutput. The propper test will be placed in proccessOutput
	ctrl.Subscribe(goca.Topics["NewOutput"], processOutput)

	// Call the plugin entrypoint
	setup(ctrl)

	testData.GetAssets(t, ctrl, testserver, plugName)
}

func processOutput(module, url string, out *goca.Output) {
	// We have to validate goca.Output according to the resource
	parts := strings.Split(out.Target, "/")

	switch parts[len(parts)-1] {
	case "Ppt1.pptx":
		validateCaseA(out)
	case "Ppt2.pptx":
		validateCaseB(out)
	case "Ppt3.pptx":
		validateCaseC(out)
	}
}

func validateCaseA(out *goca.Output) {
	if out.MainType != "PPTX" {
		T.Errorf("expected PPTX but found %s", out.MainType)
	}
	if out.Title != "Testing and assessment" {
		T.Errorf("expected \"Testing and assessment\" but found %s", out.Title)
	}
	if out.Comment != "ELT Methods and Practices" {
		T.Errorf("expected \"ELT Methods and Practices\" but found %s", out.Comment)
	}
	if out.Producer != "Bessie Dendrinos" {
		T.Errorf("expected \"Bessie Dendrinos\" but found %s", out.Producer)
	}
	if out.Keywords != "" {
		T.Errorf("expected \"\" but found %s", out.Keywords)
	}
	if out.Description != "" {
		T.Errorf("expected \"\" but found %s", out.MainType)
	}
	if out.ModifiedBy != "Smaragda Papadopoulou" {
		T.Errorf("expected \"Smaragda Papadopoulou\" but found %s", out.ModifiedBy)
	}
	if out.DocumentID != "100" {
		T.Errorf("expected \"100\" but found %s", out.DocumentID)
	}
	if out.CreateDate != "2015-08-10T14:47:42Z" {
		T.Errorf("expected \"2018-10-24T14:04:00Z\" but found %s", out.CreateDate)
	}
	if out.ModifyDate != "2015-11-25T23:53:26Z" {
		T.Errorf("expected \"2015-11-25T23:53:26Z\" but found %s", out.ModifyDate)
	}
	if out.Category != "Foreign Language Teaching and Learning" {
		T.Errorf("expected \"Foreign Language Teaching and Learning\" but found %s", out.Category)
	}
}

func validateCaseB(out *goca.Output) {
	if out.MainType != "PPTX" {
		T.Errorf("expected PPTX but found %s", out.MainType)
	}
	if out.Title != "Test Planning" {
		T.Errorf("expected \"Test Planning\" but found %s", out.CreateDate)
	}
	if out.Comment != "" {
		T.Errorf("expected \"\" but found %s", out.Comment)
	}
	if out.Producer != "TJ Probert" {
		T.Errorf("expected \"TJ Probert\" but found %s", out.Producer)
	}
	if out.Keywords != "" {
		T.Errorf("expected \"\" but found %s", out.Keywords)
	}
	if out.Description != "" {
		T.Errorf("expected \"\" but found %s", out.MainType)
	}
	if out.ModifiedBy != "TJ Probert" {
		T.Errorf("expected \"TJ Probert\" but found %s", out.ModifiedBy)
	}
	if out.DocumentID != "33" {
		T.Errorf("expected \"33\" but found %s", out.DocumentID)
	}
	if out.CreateDate != "2009-11-18T00:59:14Z" {
		T.Errorf("expected \"2009-11-18T00:59:14Z\" but found %s", out.CreateDate)
	}
	if out.ModifyDate != "2009-11-23T14:36:12Z" {
		T.Errorf("expected \"2009-11-23T14:36:12Z\" but found %s", out.ModifyDate)
	}
	if out.Category != "" {
		T.Errorf("expected \"\" but found %s", out.Category)
	}
}

func validateCaseC(out *goca.Output) {
	if out.MainType != "PPTX" {
		T.Errorf("expected PPTX but found %s", out.MainType)
	}
	if out.Title != "TEST CONTENT" {
		T.Errorf("expected \"TEST CONTENT\" but found %s", out.Title)
	}
	if out.Comment != "" {
		T.Errorf("expected \"\" but found %s", out.Comment)
	}
	if out.Producer != "Nick Fuller" {
		T.Errorf("expected \"Nick Fuller\" but found %s", out.Producer)
	}
	if out.Keywords != "" {
		T.Errorf("expected \"\" but found %s", out.Keywords)
	}
	if out.Description != "" {
		T.Errorf("expected \"\" but found %s", out.MainType)
	}
	if out.ModifiedBy != "Paul Shope" {
		T.Errorf("expected \"Paul Shope\" but found %s", out.ModifiedBy)
	}
	if out.DocumentID != "395" {
		T.Errorf("expected \"395\" but found %s", out.DocumentID)
	}
	if out.CreateDate != "2015-04-03T19:17:38Z" {
		T.Errorf("expected \"2015-04-03T19:17:38Z\" but found %s", out.CreateDate)
	}
	if out.ModifyDate != "2016-04-12T00:07:44Z" {
		T.Errorf("expected \"2016-04-12T00:07:44Z\" but found %s", out.ModifyDate)
	}
	if out.Category != "" {
		T.Errorf("expected \"\" but found %s", out.Category)
	}
}
