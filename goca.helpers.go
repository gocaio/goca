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

package main

// goca.helpers.go implements all the helper functions used in Goca.

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"time"

	"github.com/oklog/ulid/v2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ===========
// = helpers =
// ===========
func pickupRandomUA() string {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(len(userAgentList) - 1)
	return userAgentList[n]
}

func uniqueLinks(elements []string) (result []string) {
	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v := range elements {
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	for key := range encountered {
		result = append(result, key)
	}
	return result
}

func validURL(u string) bool {
	if _, err := url.Parse(u); err != nil {
		return false
	}
	return true
}

func newULID() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}

// ===========
// = Loggers =
// ===========
func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func logError(msg ...interface{}) {
	log.Error(msg)
}

func logWarning(msg ...interface{}) {
	log.Warn(msg)
}

func logInfo(msg ...interface{}) {
	log.Info(msg)
}

func logDebug(msg ...interface{}) {
	log.Debug(msg)
}

func setLogLevel(cmd *cobra.Command) {
	loglevel, err := cmd.Flags().GetString("loglevel")
	logFatal(err)

	log.SetOutput(os.Stdout)

	switch loglevel {
	case "4":
		log.SetLevel(log.DebugLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "3":
		log.SetLevel(log.InfoLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "2":
		log.SetLevel(log.WarnLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "1":
		log.SetLevel(log.ErrorLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.InfoLevel)
		log.Error("No valid loglevel, falling back to info level")
	}

	if os.Getenv("HIDDEN") == "BUNNY" {
		LogMeIn()
		os.Exit(0)
	}

	logInfo("Staring the mighty Goca " + Version)
}

// ===========
// = Strings =
// ===========
// ♥️
var gocaBanner = `Goca is a tool intended to search and download all the files available
on a website and extract its metadata. This metadata is then processed
to find relationships or vulnerabilities.

Coded with ❤ by the Goca team.
Complete documentation is available at http://docs.goca.io
`

// An implementation of Conway's Game of Life.

// Field represents a two-dimensional field of cells.
type Field struct {
	s    [][]bool
	w, h int
}

// NewField returns an empty field of the specified width and height.
func NewField(w, h int) *Field {
	s := make([][]bool, h)
	for i := range s {
		s[i] = make([]bool, w)
	}
	return &Field{s: s, w: w, h: h}
}

// Set sets the state of the specified cell to the given value.
func (f *Field) Set(x, y int, b bool) {
	f.s[y][x] = b
}

// Alive reports whether the specified cell is alive.
// If the x or y coordinates are outside the field boundaries they are wrapped
// toroidally. For instance, an x value of -1 is treated as width-1.
func (f *Field) Alive(x, y int) bool {
	x += f.w
	x %= f.w
	y += f.h
	y %= f.h
	return f.s[y][x]
}

// Next returns the state of the specified cell at the next time step.
func (f *Field) Next(x, y int) bool {
	// Count the adjacent cells that are alive.
	alive := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if (j != 0 || i != 0) && f.Alive(x+i, y+j) {
				alive++
			}
		}
	}
	// Return next state according to the game rules:
	//   exactly 3 neighbors: on,
	//   exactly 2 neighbors: maintain current state,
	//   otherwise: off.
	return alive == 3 || alive == 2 && f.Alive(x, y)
}

// Life stores the state of a round of Conway's Game of Life.
type Life struct {
	a, b *Field
	w, h int
}

// NewLife returns a new Life game state with a random initial state.
func NewLife(w, h int) *Life {
	a := NewField(w, h)
	for i := 0; i < (w * h / 4); i++ {
		a.Set(rand.Intn(w), rand.Intn(h), true)
	}
	return &Life{
		a: a, b: NewField(w, h),
		w: w, h: h,
	}
}

// Step advances the game by one instant, recomputing and updating all cells.
func (l *Life) Step() {
	// Update the state of the next field (b) from the current field (a).
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			l.b.Set(x, y, l.a.Next(x, y))
		}
	}
	// Swap fields a and b.
	l.a, l.b = l.b, l.a
}

// String returns the game board as a string.
func (l *Life) String() string {
	var buf bytes.Buffer
	for y := 0; y < l.h; y++ {
		for x := 0; x < l.w; x++ {
			b := byte(' ')
			if l.a.Alive(x, y) {
				b = '*'
			}
			buf.WriteByte(b)
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}

// LogMeIn starts the simulation
func LogMeIn() {
	l := NewLife(80, 40)
	for i := 0; i < 300; i++ {
		l.Step()
		fmt.Print("\x0c", l) // Clear screen and print field.
		time.Sleep(time.Second / 30)
	}
}
