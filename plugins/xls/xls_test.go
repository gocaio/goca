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

package xls

import (
	"strings"
	"testing"

	"github.com/gocaio/goca"
	"github.com/gocaio/goca/testData"
)

// Test server URL.
// For testing locally you need Python 3.5.X and Flask
var testserver = "https://test.goca.io"

// T is a global reference for the test. This allows us to use *testing.T
// methods anywhere
var T *testing.T

func TestReadXLS(t *testing.T) {
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
	case "Xls1.xlsx":
		validateCaseA(out)
	case "Xls2.xlsx":
		validateCaseB(out)
	case "Xls3.xlsx":
		validateCaseC(out)
	}
}

func validateCaseA(out *goca.Output) {
	if out.MainType != "XLSX" {
		T.Errorf("expected XLSX but found %s", out.MainType)
	}
	if out.Title != "Test Item Analysis Calculator" {
		T.Errorf("expected \"Test Item Analysis Calculator\" but found %s", out.Title)
	}
	if out.Comment != "" {
		T.Errorf("expected \"\" but found %s", out.Comment)
	}
	if out.Producer != "Art Tweedie - School Distirct of Osceola Co., Florida" {
		T.Errorf("expected \"Art Tweedie - School Distirct of Osceola Co., Florida\" but found %s", out.Producer)
	}
	if out.Keywords != "" {
		T.Errorf("expected \"\" but found %s", out.Keywords)
	}
	if out.Description != "This spreadsheet is used to analyze test items.  Tabbed graphs report students scores and item analysis data." {
		T.Errorf("expected \"This spreadsheet is used to analyze test items.  Tabbed graphs report students scores and item analysis data.\" but found %s", out.MainType)
	}
	if out.ModifiedBy != "lisa.huddleston" {
		T.Errorf("expected \"lisa.huddleston\" but found %s", out.ModifiedBy)
	}
	if out.DocumentID != "" {
		T.Errorf("expected \"\" but found %s", out.DocumentID)
	}
	if out.CreateDate != "2010-01-26T20:36:20Z" {
		T.Errorf("expected \"2010-01-26T20:36:20Z\" but found %s", out.CreateDate)
	}
	if out.ModifyDate != "2014-11-13T20:24:49Z" {
		T.Errorf("expected \"2014-11-13T20:24:49Z\" but found %s", out.ModifyDate)
	}
	if out.Category != "" {
		T.Errorf("expected \"\" but found %s", out.Category)
	}
}

func validateCaseB(out *goca.Output) {
	if out.MainType != "XLSX" {
		T.Errorf("expected XLSX but found %s", out.MainType)
	}
	if out.Title != "" {
		T.Errorf("expected \"\" but found %s", out.CreateDate)
	}
	if out.Comment != "" {
		T.Errorf("expected \"\" but found %s", out.Comment)
	}
	if out.Producer != "janap" {
		T.Errorf("expected \"janap\" but found %s", out.Producer)
	}
	if out.Keywords != "" {
		T.Errorf("expected \"\" but found %s", out.Keywords)
	}
	if out.Description != "" {
		T.Errorf("expected \"\" but found %s", out.MainType)
	}
	if out.ModifiedBy != "janap" {
		T.Errorf("expected \"janap\" but found %s", out.ModifiedBy)
	}
	if out.DocumentID != "" {
		T.Errorf("expected \"\" but found %s", out.DocumentID)
	}
	if out.CreateDate != "2012-02-16T08:37:28Z" {
		T.Errorf("expected \"2012-02-16T08:37:28Z\" but found %s", out.CreateDate)
	}
	if out.ModifyDate != "2012-10-23T11:20:14Z" {
		T.Errorf("expected \"2012-10-23T11:20:14Z\" but found %s", out.ModifyDate)
	}
	if out.Category != "" {
		T.Errorf("expected \"\" but found %s", out.Category)
	}
}

func validateCaseC(out *goca.Output) {
	if out.MainType != "XLSX" {
		T.Errorf("expected XLSX but found %s", out.MainType)
	}
	if out.Title != "Test Script template" {
		T.Errorf("expected \"Test Script template\" but found %s", out.Title)
	}
	if out.Comment != "Project Management" {
		T.Errorf("expected \"Project Management\" but found %s", out.Comment)
	}
	if out.Producer != "UServices, Program Management Office" {
		T.Errorf("expected \"UServices, Program Management Office\" but found %s", out.Producer)
	}
	if out.Keywords != "test script, test steps, defect verification" {
		T.Errorf("expected \"test script, test steps, defect verification\" but found %s", out.Keywords)
	}
	if out.Description != "" {
		T.Errorf("expected \"\" but found %s", out.MainType)
	}
	if out.ModifiedBy != "Rob Jackson" {
		T.Errorf("expected \"Rob Jackson\" but found %s", out.ModifiedBy)
	}
	if out.DocumentID != "" {
		T.Errorf("expected \"\" but found %s", out.DocumentID)
	}
	if out.CreateDate != "2001-07-19T17:13:32Z" {
		T.Errorf("expected \"2001-07-19T17:13:32Z\" but found %s", out.CreateDate)
	}
	if out.ModifyDate != "2016-09-09T04:31:01Z" {
		T.Errorf("expected \"2016-09-09T04:31:01Z\" but found %s", out.ModifyDate)
	}
	if out.Category != "Test Phase" {
		T.Errorf("expected \"Test Phase\" but found %s", out.Category)
	}
}
