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

// TestReadPDF tests the read on PDF files
func Test_readPDF(t *testing.T) {
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
	case "sample_1.pdf":
		validateCaseA(out)
	case "sample_2.pdf":
		validateCaseB(out)
	case "sample_3.pdf":
		validateCaseC(out)
	}
}

/*
	Validator functions
Each function validates a specific asset for this plugin.

	// TODO: This validations needs a lot of improvements
*/

func validateCaseA(out *goca.Output) {
	if out.MainType != "PDF" {
		T.Errorf("expected PDF but found %s", out.MainType)
	}
	if out.CreateDate != "2009-03-22T18:48:09+01:00" {
		T.Errorf("expected 2009-03-22T18:48:09+01:00 but found %s", out.CreateDate)
	}
	if out.ModifyDate != "2009-03-22T18:48:32+01:00" {
		T.Errorf("expected 2009-03-22T18:48:32+01:00 but found %s", out.ModifyDate)
	}
	if out.MetadataDate != "2009-03-22T18:48:32+01:00" {
		T.Errorf("expected 2009-03-22T18:48:32+01:00 but found %s", out.MetadataDate)
	}
	if out.CreatorTool != "Acrobat PDFMaker 9.0 para PowerPoint" {
		T.Errorf("expected Acrobat PDFMaker 9.0 para PowerPoint but found %s", out.CreatorTool)
	}
	if out.DocumentID != "uuid:013743e9-9796-4d7b-a3b6-f26fee2d1e30" {
		T.Errorf("expected uuid:013743e9-9796-4d7b-a3b6-f26fee2d1e30 but found %s", out.DocumentID)
	}
	if out.InstanceID != "uuid:56388999-c380-4045-9604-0cf12bd13107" {
		T.Errorf("expected uuid:56388999-c380-4045-9604-0cf12bd13107 but found %s", out.InstanceID)
	}
	if out.ContentType != "application/pdf" {
		T.Errorf("expected application/pdf but found %s", out.ContentType)
	}
	if out.Title != "" {
		T.Errorf("expected \"\" but found %s", out.Title)
	}
	if out.Producer != "Adobe PDF Library 9.0" {
		T.Errorf("expected Adobe PDF Library 9.0 but found %s", out.Producer)
	}
	if out.Lang != "" {
		T.Errorf("expected \"\" but found %s", out.Lang)
	}
	if out.Genre != "" {
		T.Errorf("expected \"\" but found %s", out.Genre)
	}
	if out.Artist != "" {
		T.Errorf("expected \"\" but found %s", out.Artist)
	}
	if out.AlbumArtist != "" {
		T.Errorf("expected \"\" but found %s", out.AlbumArtist)
	}
	if out.Album != "" {
		T.Errorf("expected \"\" but found %s", out.Album)
	}
	if out.Year != "" {
		T.Errorf("expected \"\" but found %s", out.Year)
	}
	if out.Month != "" {
		T.Errorf("expected \"\" but found %s", out.Month)
	}
	if out.Day != "" {
		T.Errorf("expected \"\" but found %s", out.Day)
	}
	if out.Comment != "" {
		T.Errorf("expected \"\" but found %s", out.Comment)
	}
	if out.Composer != "" {
		T.Errorf("expected \"\" but found %s", out.Composer)
	}
	if out.Lyrics != "" {
		T.Errorf("expected \"\" but found %s", out.Lyrics)
	}
	if out.Keywords != "" {
		T.Errorf("expected \"\" but found %s", out.Keywords)
	}
	if out.Description != "" {
		T.Errorf("expected \"\" but found %s", out.Description)
	}
	if out.ModifiedBy != "" {
		T.Errorf("expected \"\" but found %s", out.ModifiedBy)
	}
	if out.Category != "" {
		T.Errorf("expected \"\" but found %s", out.Category)
	}
	if out.DiscU != 0 {
		T.Errorf("expected 0 but found %d", out.DiscU)
	}
	if out.DiscD != 0 {
		T.Errorf("expected 0 but found %d", out.DiscD)
	}
	if out.DiscC != 0 {
		T.Errorf("expected 0 but found %d", out.DiscC)
	}
	if out.TrackU != 0 {
		T.Errorf("expected 0 but found %d", out.TrackU)
	}
	if out.TrackD != 0 {
		T.Errorf("expected 0 but found %d", out.TrackD)
	}
	if out.TrackC != 0 {
		T.Errorf("expected 0 but found %d", out.TrackC)
	}
}

func validateCaseB(out *goca.Output) {
	if out.MainType != "PDF" {
		T.Errorf("expected PDF but found %s", out.MainType)
	}
	if out.CreateDate != "2015-08-21T09:42:21+01:00" {
		T.Errorf("expected 2015-08-21T09:42:21+01:00 but found %s", out.CreateDate)
	}
	if out.ModifyDate != "2015-08-21T09:45:31+01:00" {
		T.Errorf("expected 2015-08-21T09:45:31+01:00 but found %s", out.ModifyDate)
	}
	if out.MetadataDate != "2015-08-21T09:45:31+01:00" {
		T.Errorf("expected 2015-08-21T09:45:31+01:00 but found %s", out.MetadataDate)
	}
	if out.CreatorTool != "Acrobat PDFMaker 15 for Word" {
		T.Errorf("expected Acrobat PDFMaker 15 for Word but found %s", out.CreatorTool)
	}
	if out.DocumentID != "uuid:32f55d7c-7ef1-46ca-85c0-b5b91509ff82" {
		T.Errorf("expected uuid:32f55d7c-7ef1-46ca-85c0-b5b91509ff82 but found %s", out.DocumentID)
	}
	if out.InstanceID != "uuid:4da38cf4-2b42-417c-b34e-f529c34ac6cc" {
		T.Errorf("expected uuid:4da38cf4-2b42-417c-b34e-f529c34ac6cc but found %s", out.InstanceID)
	}
	if out.ContentType != "application/pdf" {
		T.Errorf("expected application/pdf but found %s", out.ContentType)
	}
	if out.Title != "PDF Metadata Sample" {
		T.Errorf("expected PDF Metadata Sample but found %s", out.Title)
	}
	if out.Producer != "Adobe PDF Library 15.0" {
		T.Errorf("expected Adobe PDF Library 15.0 but found %s", out.Producer)
	}
	if out.Lang != "" {
		T.Errorf("expected \"\" but found %s", out.Lang)
	}
	if out.Genre != "" {
		T.Errorf("expected \"\" but found %s", out.Genre)
	}
	if out.Artist != "" {
		T.Errorf("expected \"\" but found %s", out.Artist)
	}
	if out.AlbumArtist != "" {
		T.Errorf("expected \"\" but found %s", out.AlbumArtist)
	}
	if out.Album != "" {
		T.Errorf("expected \"\" but found %s", out.Album)
	}
	if out.Year != "" {
		T.Errorf("expected \"\" but found %s", out.Year)
	}
	if out.Month != "" {
		T.Errorf("expected \"\" but found %s", out.Month)
	}
	if out.Day != "" {
		T.Errorf("expected \"\" but found %s", out.Day)
	}
	if out.Comment != "" {
		T.Errorf("expected \"\" but found %s", out.Comment)
	}
	if out.Composer != "" {
		T.Errorf("expected \"\" but found %s", out.Composer)
	}
	if out.Lyrics != "" {
		T.Errorf("expected \"\" but found %s", out.Lyrics)
	}
	if out.Keywords != "" {
		T.Errorf("expected \"\" but found %s", out.Keywords)
	}
	if out.Description != "" {
		T.Errorf("expected \"\" but found %s", out.Description)
	}
	if out.ModifiedBy != "" {
		T.Errorf("expected \"\" but found %s", out.ModifiedBy)
	}
	if out.Category != "" {
		T.Errorf("expected \"\" but found %s", out.Category)
	}
	if out.DiscU != 0 {
		T.Errorf("expected 0 but found %d", out.DiscU)
	}
	if out.DiscD != 0 {
		T.Errorf("expected 0 but found %d", out.DiscD)
	}
	if out.DiscC != 0 {
		T.Errorf("expected 0 but found %d", out.DiscC)
	}
	if out.TrackU != 0 {
		T.Errorf("expected 0 but found %d", out.TrackU)
	}
	if out.TrackD != 0 {
		T.Errorf("expected 0 but found %d", out.TrackD)
	}
	if out.TrackC != 0 {
		T.Errorf("expected 0 but found %d", out.TrackC)
	}
}

func validateCaseC(out *goca.Output) {
	if out.MainType != "PDF" {
		T.Errorf("expected PDF but found %s", out.MainType)
	}
	if out.CreateDate != "2015-08-21T09:42:21+01:00" {
		T.Errorf("expected 2015-08-21T09:42:21+01:00 but found %s", out.CreateDate)
	}
	if out.ModifyDate != "2015-08-21T09:45:31+01:00" {
		T.Errorf("expected 2015-08-21T09:45:31+01:00 but found %s", out.ModifyDate)
	}
	if out.MetadataDate != "2015-08-21T09:45:31+01:00" {
		T.Errorf("expected 2015-08-21T09:45:31+01:00 but found %s", out.MetadataDate)
	}
	if out.CreatorTool != "Acrobat PDFMaker 15 for Word" {
		T.Errorf("expected Acrobat PDFMaker 15 for Word but found %s", out.CreatorTool)
	}
	if out.DocumentID != "uuid:32f55d7c-7ef1-46ca-85c0-b5b91509ff82" {
		T.Errorf("expected uuid:32f55d7c-7ef1-46ca-85c0-b5b91509ff82 but found %s", out.DocumentID)
	}
	if out.InstanceID != "uuid:4da38cf4-2b42-417c-b34e-f529c34ac6cc" {
		T.Errorf("expected uuid:4da38cf4-2b42-417c-b34e-f529c34ac6cc but found %s", out.InstanceID)
	}
	if out.ContentType != "application/pdf" {
		T.Errorf("expected application/pdf but found %s", out.ContentType)
	}
	if out.Title != "PDF Metadata Sample" {
		T.Errorf("expected PDF Metadata Sample but found %s", out.Title)
	}
	if out.Producer != "Adobe PDF Library 15.0" {
		T.Errorf("expected Adobe PDF Library 15.0 but found %s", out.Producer)
	}
	if out.Lang != "" {
		T.Errorf("expected \"\" but found %s", out.Lang)
	}
	if out.Genre != "" {
		T.Errorf("expected \"\" but found %s", out.Genre)
	}
	if out.Artist != "" {
		T.Errorf("expected \"\" but found %s", out.Artist)
	}
	if out.AlbumArtist != "" {
		T.Errorf("expected \"\" but found %s", out.AlbumArtist)
	}
	if out.Album != "" {
		T.Errorf("expected \"\" but found %s", out.Album)
	}
	if out.Year != "" {
		T.Errorf("expected \"\" but found %s", out.Year)
	}
	if out.Month != "" {
		T.Errorf("expected \"\" but found %s", out.Month)
	}
	if out.Day != "" {
		T.Errorf("expected \"\" but found %s", out.Day)
	}
	if out.Comment != "" {
		T.Errorf("expected \"\" but found %s", out.Comment)
	}
	if out.Composer != "" {
		T.Errorf("expected \"\" but found %s", out.Composer)
	}
	if out.Lyrics != "" {
		T.Errorf("expected \"\" but found %s", out.Lyrics)
	}
	if out.Keywords != "" {
		T.Errorf("expected \"\" but found %s", out.Keywords)
	}
	if out.Description != "" {
		T.Errorf("expected \"\" but found %s", out.Description)
	}
	if out.ModifiedBy != "" {
		T.Errorf("expected \"\" but found %s", out.ModifiedBy)
	}
	if out.Category != "" {
		T.Errorf("expected \"\" but found %s", out.Category)
	}
	if out.DiscU != 0 {
		T.Errorf("expected 0 but found %d", out.DiscU)
	}
	if out.DiscD != 0 {
		T.Errorf("expected 0 but found %d", out.DiscD)
	}
	if out.DiscC != 0 {
		T.Errorf("expected 0 but found %d", out.DiscC)
	}
	if out.TrackU != 0 {
		T.Errorf("expected 0 but found %d", out.TrackU)
	}
	if out.TrackD != 0 {
		T.Errorf("expected 0 but found %d", out.TrackD)
	}
	if out.TrackC != 0 {
		T.Errorf("expected 0 but found %d", out.TrackC)
	}
}
