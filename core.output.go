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

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/gocaio/goca/rsrc"
	jsonc "github.com/nwidger/jsoncolor"
)

// core.output.go is in charge of generating the result output, this output can
// be the database, the stdout, a file, etc.

var (
	targetFolder string
)

// saveFile save the file in disk
func saveFile(mime string, data []byte) {
	if _, err := os.Stat(targetFolder); os.IsNotExist(err) {
		err = os.MkdirAll(targetFolder, os.ModePerm)
		logFatal(err)
	}
	fname := fmt.Sprintf("%s_%s.%s", mime, newULID(), mimeToExtension[mime])
	fname = strings.ReplaceAll(fname, "/", "_")
	fname = strings.ReplaceAll(fname, "\\", "_")

	err := ioutil.WriteFile(path.Join(targetFolder, path.Clean(fname)), data, 0644)
	logFatal(err)
}

func processOutput(out *rsrc.Output) {
	var data []byte
	var err error

	if runtime.GOOS == "windows" {
		data, err = json.MarshalIndent(out, "", "\t")
	} else {
		data, err = jsonc.MarshalIndent(out, "", "\t")
	}

	if err != nil {
		logError(err)
	} else {
		fmt.Println(string(data))
	}
}
