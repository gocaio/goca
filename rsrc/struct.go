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

package rsrc

// core.output.struct.go defines the Output object for the plugins.

// import (
// 	"encoding/json"
// 	"fmt"
// 	"runtime"

// 	jsonc "github.com/nwidger/jsoncolor"
// )

// Output is a flat structure which represents all possible values that
// Goca can extract. Each plugin will generate one Output per object processed
type Output struct {
	// Scan info metadata
	MainType string `json:"main_type,omitempty"` // Kind of file represents the Output
	Target   string `json:"target,omitempty"`    // ?
	Length   int    `json:"length,omitempty"`    // Analyzed file length
	Sha256   string `json:"sha256,omitempty"`    // File sha256 hash

	// General file metadata
	CreateDate   string `json:"create_date,omitempty"`
	ModifyDate   string `json:"modify_date,omitempty"`
	MetadataDate string `json:"metadata_date,omitempty"`
	CreatorTool  string `json:"creator_tool,omitempty"`
	DocumentID   string `json:"document_id,omitempty"`
	InstanceID   string `json:"instance_id,omitempty"`
	ContentType  string `json:"content_type,omitempty"`
	Title        string `json:"title,omitempty"`
	Lang         string `json:"lang,omitempty"`
	Producer     string `json:"producer,omitempty"`
	Comment      string `json:"comment,omitempty"`
	Keywords     string `json:"keywords,omitempty"`
	Description  string `json:"description,omitempty"`
	ModifiedBy   string `json:"modified_by,omitempty"`
	Category     string `json:"category,omitempty"`

	// Audio Specific metadata
	Genre       string `json:"genre,omitempty"`
	Artist      string `json:"artist,omitempty"`
	AlbumArtist string `json:"album_artist,omitempty"`
	Album       string `json:"album,omitempty"`
	Year        string `json:"year,omitempty"`
	Month       string `json:"month,omitempty"`
	Day         string `json:"day,omitempty"`
	Composer    string `json:"composer,omitempty"`
	Lyrics      string `json:"lyrics,omitempty"`
	DiscU       int    `json:"disc_u,omitempty"`
	DiscD       int    `json:"disc_d,omitempty"`
	DiscC       int    `json:"disc_c,omitempty"`
	TrackU      int    `json:"track_u,omitempty"`
	TrackD      int    `json:"track_d,omitempty"`
	TrackC      int    `json:"track_c,omitempty"`

	// Geo metadata
	GeoX string `json:"geo_x,omitempty"`
	GeoY string `json:"geo_y,omitempty"`
	GeoZ string `json:"geo_z,omitempty"`
}

// NewOutput returns a new empty Output structure
func NewOutput() *Output {
	return &Output{}
}

// // FIXME: Old code (this should be handled in core.output.go)
// func (o *Output) process(out *Output) {
// 	var data []byte
// 	var err error

// 	if runtime.GOOS == "windows" {
// 		data, err = json.MarshalIndent(out, "", "\t")
// 	} else {
// 		data, err = jsonc.MarshalIndent(out, "", "\t")
// 	}

// 	if err != nil {
// 		logError(err)
// 	} else {
// 		fmt.Println(string(data))
// 	}
// }
