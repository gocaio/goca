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
	"bytes"

	"github.com/dhowden/tag"
	"github.com/gocaio/Goca"
	log "github.com/sirupsen/logrus"
)

const plugName = "mp3"

func init() {
	goca.RegisterPlugin(plugName, goca.Plugin{
		Type: "mp3",
		Assoc: []string{"audio/mpeg", // mp3
			"audio/x-flac",     // FLAC
			"audio/mpa-robust", // MP4
			"audio/MPA",        // MP4
			"application/ogg",  // OGG
			"audio/ogg",        // OGG
			"video/ogg"},       // OGG
		Action:  setup,
		Matcher: nil,
	})
}

var mp3 *mp3MetaExtractor

func setup(m goca.Manager) error {
	mp3 = new(mp3MetaExtractor)
	mp3.Manager = m
	mp3.Subscribe(goca.Topics["NewTarget"], mp3.readMP3)
	return nil
}

type mp3MetaExtractor struct {
	goca.Manager
}

func (mp3 *mp3MetaExtractor) readMP3(target string, data []byte) {
	log.Debugf("[MP3] Received Data Length: %d - TARGET: %s\n", len(data), target)

	defer goca.PluginRecover("MP3", "readMP3")

	buf := bytes.NewReader(data)
	m, err := tag.ReadFrom(buf)
	if err == nil {
		out := goca.NewOutput()
		out.MainType = "mp3"
		out.Target = target
		out.Title = m.Title()
		out.Album = m.Album()
		out.Artist = m.Artist()
		out.AlbumArtist = m.AlbumArtist()
		out.Comment = m.Comment()
		out.Composer = m.Composer()
		out.DiscD, out.DiscU = m.Disc()
		out.Genre = m.Genre()
		out.Lyrics = m.Lyrics()
		out.TrackD, out.TrackU = m.Track()

		// TODO: Process picture information

		mp3.Publish(goca.Topics["NewOutput"], plugName, target, out)
	} else {
		log.Debugf("[MP3] - Err: %s\n", err.Error())
	}
}
