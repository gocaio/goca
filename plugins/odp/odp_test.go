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

package odp

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

// TestReadODP tests the reading over ODP files
func TestReadODP(t *testing.T) {
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
	case "Odp1.docx":
		validateCaseA(out)
	case "Odp2.docx":
		validateCaseB(out)
	case "Odp3.docx":
		validateCaseC(out)
	}
}

func validateCaseA(out *goca.Output) {
	if out.MainType != "ODP" {
		T.Errorf("expected ODP but found %s", out.MainType)
	}
	if out.Title != "Apresentando o CEJUG e o Poder do Java" {
		T.Errorf("expected \"Apresentando o CEJUG e o Poder do Java\" but found %s", out.Title)
	}
	if out.Description != "Palestra usada no 1º encontro CEPUG+CEJUG em Iguatu em 1º de Novembro de 2008." {
		T.Errorf("expected \"Palestra usada no 1º encontro CEPUG+CEJUG em Iguatu em 1º de Novembro de 2008.\" but found %s", out.Description)
	}
	if out.Comment != "Apresentação da plataforma Java e da comunidade CEJUG" {
		T.Errorf("expected \"Apresentação da plataforma Java e da comunidade CEJUG\" but found %s", out.Comment)
	}
	if out.Lang != "" {
		T.Errorf("expected \"\" but found %s", out.Lang)
	}
	if out.CreatorTool != "OpenOffice.org/3.0$Linux OpenOffice.org_project/300m9$Build-9358" {
		T.Errorf("expected \"OpenOffice.org/3.0$Linux OpenOffice.org_project/300m9$Build-9358\" but found %s", out.CreatorTool)
	}
	if out.Producer != "Jose Maria Silveira Neto" {
		T.Errorf("expected \"Jose Maria Silveira Neto\" but found %s", out.Producer)
	}
	if out.Keywords != "Ceará" {
		T.Errorf("expected \"Ceará\" but found %s", out.Keywords)
	}
	if out.ModifiedBy != "Jose Maria Silveira Neto" {
		T.Errorf("expected \"Jose Maria Silveira Neto\" but found %s", out.ModifiedBy)
	}
	if out.CreateDate != "2008-11-01T07:03:41" {
		T.Errorf("expected \"2008-11-01T07:03:41\" but found %s", out.CreateDate)
	}
	if out.ModifyDate != "2008-11-02T09:45:18" {
		T.Errorf("expected \"2008-11-02T09:45:18\" but found %s", out.ModifyDate)
	}
}

func validateCaseB(out *goca.Output) {
	if out.MainType != "ODP" {
		T.Errorf("expected ODP but found %s", out.MainType)
	}
	if out.Title != "VM and SSL" {
		T.Errorf("expected \"VM and SSL\" but found %s", out.Title)
	}
	if out.Description != "VM Workshop 2012, University of Kentuckyhttp://www.vmworkshop.org/node/91/submission/3" {
		T.Errorf("expected \"VM Workshop 2012, University of Kentuckyhttp://www.vmworkshop.org/node/91/submission/3\" but found %s", out.Description)
	}
	if out.Comment != "VM and SSL -or- the Public Life of a Private Key" {
		T.Errorf("expected \"VM and SSL -or- the Public Life of a Private Key\" but found %s", out.Comment)
	}
	if out.Lang != "" {
		T.Errorf("expected \"\" but found %s", out.Lang)
	}
	if out.CreatorTool != "LibreOffice/3.5$Linux_X86_64 LibreOffice_project/350m1$Build-413" {
		T.Errorf("expected \"LibreOffice/3.5$Linux_X86_64 LibreOffice_project/350m1$Build-413\" but found %s", out.CreatorTool)
	}
	if out.Producer != "" {
		T.Errorf("expected \"\" but found %s", out.Producer)
	}
	if out.Keywords != "" {
		T.Errorf("expected \"\" but found %s", out.Keywords)
	}
	if out.ModifiedBy != "Rick Troth" {
		T.Errorf("expected \"Rick Troth\" but found %s", out.ModifiedBy)
	}
	if out.CreateDate != "2012-05-17T12:01:44" {
		T.Errorf("expected \"2012-05-17T12:01:44\" but found %s", out.CreateDate)
	}
	if out.ModifyDate != "2013-06-16T23:21:25" {
		T.Errorf("expected \"2013-06-16T23:21:25\" but found %s", out.ModifyDate)
	}
}

func validateCaseC(out *goca.Output) {
	if out.MainType != "ODP" {
		T.Errorf("expected ODP but found %s", out.MainType)
	}
	if out.Title != "Páros gráfok párosítása" {
		T.Errorf("expected \"Páros gráfok párosítása\" but found %s", out.Title)
	}
	if out.Description != "Készítették:Szabó TiborTábori ZsoltJuhász LászlóSzalóki Gábor" {
		T.Errorf("expected \"Készítették:Szabó TiborTábori ZsoltJuhász LászlóSzalóki Gábor\" but found %s", out.Description)
	}
	if out.Comment != "Gráfelmélet csoportmunka" {
		T.Errorf("expected \"Gráfelmélet csoportmunka\" but found %s", out.Comment)
	}
	if out.Lang != "" {
		T.Errorf("expected \"\" but found %s", out.Lang)
	}
	if out.CreatorTool != "OpenOffice.org/3.0$Win32 OpenOffice.org_project/300m9$Build-9358" {
		T.Errorf("expected \"OpenOffice.org/3.0$Win32 OpenOffice.org_project/300m9$Build-9358\" but found %s", out.CreatorTool)
	}
	if out.Producer != "" {
		T.Errorf("expected \"\" but found %s", out.Producer)
	}
	if out.Keywords != "Szegedi Egyetem Természettudományi Kar" {
		T.Errorf("expected \"Szegedi Egyetem Természettudományi Kar\" but found %s", out.Keywords)
	}
	if out.ModifiedBy != "" {
		T.Errorf("expected \"\" but found %s", out.ModifiedBy)
	}
	if out.CreateDate != "2008-12-07T09:23:18" {
		T.Errorf("expected \"2008-12-07T09:23:18\" but found %s", out.CreateDate)
	}
	if out.ModifyDate != "2009-01-15T21:20:23.66" {
		T.Errorf("expected \"2009-01-15T21:20:23.66\" but found %s", out.ModifyDate)
	}
}
