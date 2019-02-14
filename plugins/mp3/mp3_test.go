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

package mp3

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

// TestReadMP3 tests the read on MP3 files
func TestReadMP3(t *testing.T) {
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
	case "MP3_1.mp3":
		validateCaseA(out)
	case "MP3_2.mp3":
		validateCaseB(out)
	}
}

func validateCaseA(out *goca.Output) {
	if out.MainType != "mp3" {
		T.Errorf("expected p3 but found %s", out.MainType)
	}
	if out.CreateDate != "" {
		T.Errorf("expected \"\" but found %s", out.CreateDate)
	}
	if out.ModifyDate != "" {
		T.Errorf("expected \"\" but found %s", out.ModifyDate)
	}
	if out.MetadataDate != "" {
		T.Errorf("expected \"\" but found %s", out.MetadataDate)
	}
	if out.CreatorTool != "" {
		T.Errorf("expected \"\" but found %s", out.CreatorTool)
	}
	if out.DocumentID != "" {
		T.Errorf("expected \"\" but found %s", out.DocumentID)
	}
	if out.InstanceID != "" {
		T.Errorf("expected \"\" but found %s", out.InstanceID)
	}
	if out.ContentType != "" {
		T.Errorf("expected \"\" but found %s", out.ContentType)
	}
	if out.Title != "A Kind Of Magic" {
		T.Errorf("expected A Kind Of Magic but found %s", out.Title)
	}
	if out.Producer != "" {
		T.Errorf("expected \"\" but found %s", out.Producer)
	}
	if out.Lang != "" {
		T.Errorf("expected \"\" but found %s", out.Lang)
	}
	if out.Genre != "Hard rock" {
		T.Errorf("expected Hard rock but found %s", out.Genre)
	}
	if out.Artist != "Roger Taylor" {
		T.Errorf("expected Roger Taylor but found %s", out.Artist)
	}
	if out.AlbumArtist != "Queen" {
		T.Errorf("expected Queen but found %s", out.AlbumArtist)
	}
	if out.Album != "A Kind Of Magic" {
		T.Errorf("expected A Kind Of Magic but found %s", out.Album)
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
	if out.Lyrics != "It's a kind of magic\rIt's a kind of magic\rA kind of magic\rOne dream one soul one prize one goal\rOne golden glance of what should be\rIt's a kind of magic\rOne shaft of light that shows the way\rNo mortal man can win this day\rIt's a kind of magic\rThe bell that rings inside your mind\rIs challenging the doors of time\rIt's a kind of magic\rThe waiting seems eternity\rThe day will dawn of sanity\rIs this a kind of magic\rIt's a kind of magic\rThere can be only one\rThis rage that lasts a thousand years\rWill soon be done\rThis flame that burns inside of me\rI'm here in secret harmonies\rIt's a kind of magic\rThe bell that rings inside your mind\rIs challenging the doors of time\rIt's a kind of magic\rIt's a kind of magic\rThe rage that lasts a thousand years\rWill soon be will soon be\rWill soon be done\rThis is a kind of magic\rThere can be only one\rThis rage that lasts a thousand years\rWill soon be done-done\rMagic, it's a kind of magic\rIt's a kind of magic\rMagic magic magic magic\rHa ha ha it's magic\rIt's a kind of magic\x00" {
		T.Errorf("expected some lyric but found %s", out.Lyrics)
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
	if out.DiscU != 1 {
		T.Errorf("expected 1 but found %d", out.DiscU)
	}
	if out.DiscD != 1 {
		T.Errorf("expected 1 but found %d", out.DiscD)
	}
	if out.DiscC != 0 {
		T.Errorf("expected 0 but found %d", out.DiscC)
	}
	if out.TrackU != 9 {
		T.Errorf("expected 9 but found %d", out.TrackU)
	}
	if out.TrackD != 2 {
		T.Errorf("expected 2 but found %d", out.TrackD)
	}
	if out.TrackC != 0 {
		T.Errorf("expected 0 but found %d", out.TrackC)
	}
}

func validateCaseB(out *goca.Output) {
	if out.MainType != "mp3" {
		T.Errorf("expected mp3 but found %s", out.MainType)
	}
	if out.CreateDate != "" {
		T.Errorf("expected \"\" but found %s", out.CreateDate)
	}
	if out.ModifyDate != "" {
		T.Errorf("expected \"\" but found %s", out.ModifyDate)
	}
	if out.MetadataDate != "" {
		T.Errorf("expected \"\" but found %s", out.MetadataDate)
	}
	if out.CreatorTool != "" {
		T.Errorf("expected \"\" but found %s", out.CreatorTool)
	}
	if out.DocumentID != "" {
		T.Errorf("expected \"\" but found %s", out.DocumentID)
	}
	if out.InstanceID != "" {
		T.Errorf("expected \"\" but found %s", out.InstanceID)
	}
	if out.ContentType != "" {
		T.Errorf("expected \"\" but found %s", out.ContentType)
	}
	if out.Title != "Test Title" {
		T.Errorf("expected Test Title but found %s", out.Title)
	}
	if out.Producer != "" {
		T.Errorf("expected \"\" but found %s", out.Producer)
	}
	if out.Lang != "" {
		T.Errorf("expected \"\" but found %s", out.Lang)
	}
	if out.Genre != "Jazz" {
		T.Errorf("expected Jazz but found %s", out.Genre)
	}
	if out.Artist != "Test Artist" {
		T.Errorf("expected Test Artist but found %s", out.Artist)
	}
	if out.AlbumArtist != "" {
		T.Errorf("expected \"\" but found %s", out.AlbumArtist)
	}
	if out.Album != "Test Album" {
		T.Errorf("expected Test Album but found %s", out.Album)
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
	if out.Comment != "Test Comment" {
		T.Errorf("expected Test Comment but found %s", out.Comment)
	}
	if out.Composer != "" {
		T.Errorf("expected \"\" but found %s", out.Composer)
	}
	if out.Lyrics != "" {
		T.Errorf("expected some lyric but found %s", out.Lyrics)
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
	if out.TrackD != 3 {
		T.Errorf("expected 3 but found %d", out.TrackD)
	}
	if out.TrackC != 0 {
		T.Errorf("expected 0 but found %d", out.TrackC)
	}
}
