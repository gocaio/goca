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

// TestReadSWF tests the read on SWF files
func TestReadSWF(t *testing.T) {
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
	case "swf_1.swf":
		validateCaseA(out)
	case "swf_2.swf":
		validateCaseB(out)
	case "swf_3.swf":
		validateCaseC(out)
	}
}

func validateCaseA(out *goca.Output) {
	if out.MainType != "swf" {
		T.Errorf("expected swf but found %s", out.MainType)
	}

	if out.Duration != "4m16.739s" {
		T.Errorf("expected \"m16.739s\" but found %s", out.Duration)
	}
	if out.Version != 5 {
		T.Errorf("expected \"5\" but found %d", out.Version)
	}
	if out.FrameRate != 255.199997 {
		T.Errorf("expected \"255.199997\" but found %f", out.FrameRate)
	}
	if out.FrameCount != 65520 {
		T.Errorf("expected \"65520\" but found %d", out.FrameCount)
	}
	if out.Comment != "No compression" {
		T.Errorf("expected \"No compression\" but found %s", out.Comment)
	}
}

func validateCaseB(out *goca.Output) {
	if out.MainType != "swf" {
		T.Errorf("expected swf but found %s", out.MainType)
	}

	if out.Duration != "33ms" {
		T.Errorf("expected \"33ms\" but found %s", out.Duration)
	}
	if out.Version != 8 {
		T.Errorf("expected \"8\" but found %d", out.Version)
	}
	if out.FrameRate != 30.000000 {
		T.Errorf("expected \"30.000000\" but found %f", out.FrameRate)
	}
	if out.FrameCount != 1 {
		T.Errorf("expected \"1\" but found %d", out.FrameCount)
	}
	if out.Comment != "ZLIB compression" {
		T.Errorf("expected \"ZLIB compression\" but found %s", out.Comment)
	}
}

func validateCaseC(out *goca.Output) {
	if out.MainType != "swf" {
		T.Errorf("expected swf but found %s", out.MainType)
	}

	if out.Duration != "0s" {
		T.Errorf("expected \"0s\" but found %s", out.Duration)
	}
	if out.Version != 8 {
		T.Errorf("expected \"8\" but found %d", out.Version)
	}
	if out.FrameRate != 0.170000 {
		T.Errorf("expected \"0.170000\" but found %f", out.FrameRate)
	}
	if out.FrameCount != 0 {
		T.Errorf("expected \"0\" but found %d", out.FrameCount)
	}
	if out.Comment != "ZLIB compression" {
		T.Errorf("expected \"ZLIB compression\" but found %s", out.Comment)
	}
}
