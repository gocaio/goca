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

package goca

import (
	"encoding/json"
	"fmt"
	"runtime"

	jsoncolor "github.com/nwidger/jsoncolor"
)

// NewOutput returns a new empty Output structure
func NewOutput() *Output {
	return &Output{}
}

// Output is a flat structure which represents all possible values that
// Goca can extract. Each plugin will generate one Output per object processed
type Output struct {
	// MainType will set which kind of file represents the Output
	MainType     string  `json:"main_type,omitempty"`
	Target       string  `json:"target,omitempty"`
	CreateDate   string  `json:"create_date,omitempty"`
	ModifyDate   string  `json:"modify_date,omitempty"`
	MetadataDate string  `json:"metadata_date,omitempty"`
	CreatorTool  string  `json:"creator_tool,omitempty"`
	DocumentID   string  `json:"document_id,omitempty"`
	InstanceID   string  `json:"instance_id,omitempty"`
	ContentType  string  `json:"content_type,omitempty"`
	Title        string  `json:"title,omitempty"`
	Lang         string  `json:"lang,omitempty"`
	Producer     string  `json:"producer,omitempty"`
	Genre        string  `json:"genre,omitempty"`
	Artist       string  `json:"artist,omitempty"`
	AlbumArtist  string  `json:"album_artist,omitempty"`
	Album        string  `json:"album,omitempty"`
	Year         string  `json:"year,omitempty"`
	Month        string  `json:"month,omitempty"`
	Day          string  `json:"day,omitempty"`
	Comment      string  `json:"comment,omitempty"`
	Composer     string  `json:"composer,omitempty"`
	Lyrics       string  `json:"lyrics,omitempty"`
	Keywords     string  `json:"keywords,omitempty"`
	Description  string  `json:"description,omitempty"`
	ModifiedBy   string  `json:"modified_by,omitempty"`
	Category     string  `json:"category,omitempty"`
	DiscU        int     `json:"disc_u,omitempty"`
	DiscD        int     `json:"disc_d,omitempty"`
	DiscC        int     `json:"disc_c,omitempty"`
	TrackU       int     `json:"track_u,omitempty"`
	TrackD       int     `json:"track_d,omitempty"`
	TrackC       int     `json:"track_c,omitempty"`
	Duration     string  `json:"duration,omitempty"`
	Version      uint8   `json:"version,omitempty"`
	FrameCount   uint16  `json:"frame_count,omitempty"`
	FrameRate    float32 `json:"frame_rate,omitempty"`
	ImageWidth   float32 `json:"image_width,omitempty"`
	ImageHeight  float32 `json:"image_height,omitempty"`
}

func processOutput(module, url string, out *Output) {
	// This can be either save to database or display it on screen or send it to astilectron

	if runtime.GOOS == "windows" {
		data, err := json.MarshalIndent(out, "", "\t")
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		} else {
			fmt.Println(string(data))
		}
	} else {
		data, err := jsoncolor.MarshalIndent(out, "", "\t")
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		} else {
			fmt.Println(string(data))
		}
	}
}
