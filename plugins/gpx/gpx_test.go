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

package gpx

import (
	"os"
	"strings"
	"testing"

	"github.com/gocaio/goca"
	gocatesting "github.com/gocaio/goca/gocaTesting"
)

// Test server URL.
var testserver = os.Getenv("GOCA_TEST_SERVER")

// T is a global reference for the test. This allows us to use *testing.T
// methods anywhere
var T *testing.T

// TestReadgpx tests the read on gpx files
func TestReadgpx(t *testing.T) {
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
	case "gpx_1.gpx":
		validateCaseA(out)
	case "gpx_2.gpx":
		validateCaseB(out)
	case "gpx_3.gpx":
		validateCaseC(out)
	}
}

func validateCaseA(out *goca.Output) {
	if out.MainType != "GPX" {
		T.Errorf("expected GPX but found %s", out.MainType)
	}

	if out.DocumentID != "1.1" {
		T.Errorf("expected \"1.1\" but found %s", out.DocumentID)
	}
	if out.Producer != "DDAAXX TCX Extractor and Converter" {
		T.Errorf("expected \"DDAAXX TCX Extractor and Converter\" but found %s", out.Producer)
	}
	if out.Title != "" {
		T.Errorf("expected \"\" but found %s", out.Title)
	}
	if out.Description != "" {
		T.Errorf("expected \"\" but found %s", out.Description)
	}
	if out.ModifiedBy != "" {
		T.Errorf("expected \"\" but found %s", out.ModifiedBy)
	}
	if out.Email != "" {
		T.Errorf("expected \"\" but found %s", out.Email)
	}
	if out.CreatorTool != "" {
		T.Errorf("expected \"\" but found %s", out.CreatorTool)
	}
	if out.Year != "" {
		T.Errorf("expected \"\" but found %s", out.Year)
	}
	if out.Comment != "" {
		T.Errorf("expected \"\" but found %s", out.Comment)
	}
	if out.Keywords != "" {
		T.Errorf("expected \"\" but found %s", out.Keywords)
	}
}

func validateCaseB(out *goca.Output) {
	if out.MainType != "GPX" {
		T.Errorf("expected GPX but found %s", out.MainType)
	}

	if out.DocumentID != "1.1" {
		T.Errorf("expected \"1.1\" but found %s", out.DocumentID)
	}
	if out.Producer != "StravaGPX" {
		T.Errorf("expected \"StravaGPX\" but found %s", out.Producer)
	}
	if out.Title != "" {
		T.Errorf("expected \"\" but found %s", out.Title)
	}
	if out.Description != "" {
		T.Errorf("expected \"\" but found %s", out.Description)
	}
	if out.ModifiedBy != "" {
		T.Errorf("expected \"\" but found %s", out.ModifiedBy)
	}
	if out.Email != "" {
		T.Errorf("expected \"\" but found %s", out.Email)
	}
	if out.CreatorTool != "" {
		T.Errorf("expected \"\" but found %s", out.CreatorTool)
	}
	if out.Year != "" {
		T.Errorf("expected \"\" but found %s", out.Year)
	}
	if out.Comment != "" {
		T.Errorf("expected \"\" but found %s", out.Comment)
	}
	if out.Keywords != "" {
		T.Errorf("expected \"\" but found %s", out.Keywords)
	}
}

func validateCaseC(out *goca.Output) {
	if out.MainType != "GPX" {
		T.Errorf("expected GPX but found %s", out.MainType)
	}

	if out.DocumentID != "1.1" {
		T.Errorf("expected \"1.1\" but found %s", out.DocumentID)
	}
	if out.Producer != "Garmin Desktop App" {
		T.Errorf("expected \"Garmin Desktop App\" but found %s", out.Producer)
	}
	if out.Title != "" {
		T.Errorf("expected \"\" but found %s", out.Title)
	}
	if out.Description != "" {
		T.Errorf("expected \"\" but found %s", out.Description)
	}
	if out.ModifiedBy != "" {
		T.Errorf("expected \"\" but found %s", out.ModifiedBy)
	}
	if out.Email != "" {
		T.Errorf("expected \"\" but found %s", out.Email)
	}
	if out.CreatorTool != "" {
		T.Errorf("expected \"\" but found %s", out.CreatorTool)
	}
	if out.Year != "" {
		T.Errorf("expected \"\" but found %s", out.Year)
	}
	if out.Comment != "http://www.garmin.com" {
		T.Errorf("expected \"http://www.garmin.com\" but found %s", out.Comment)
	}
	if out.Keywords != "" {
		T.Errorf("expected \"\" but found %s", out.Keywords)
	}
}
