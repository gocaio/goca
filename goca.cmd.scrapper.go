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

// goca.cmd.scrapper.go is the cobra command selector entrypoint for the scrapper.

import (
	"fmt"
	"time"

	// "github.com/gocaio/goca/plugin"

	"github.com/spf13/cobra"
)

func scrapperCmdFunc(cmd *cobra.Command, args []string) {
	setLogLevel(cmd)

	ua, err := cmd.Flags().GetString("userAgent")
	logFatal(err)

	dm, err := cmd.Flags().GetString("domain")
	logFatal(err)

	th, err := cmd.Flags().GetInt("threshold")
	logFatal(err)
	if th < 0 {
		logError("threshold can't be less than 0")
		return
	}

	dp, err := cmd.Flags().GetInt("depth")
	logFatal(err)
	if dp < 0 {
		logError("depth can't be less than 0")
		return
	}

	threads, err := cmd.Flags().GetInt("threads")
	logFatal(err)
	if threads < 1 {
		logError("The number of threads to run can't be less than 1")
		return
	}

	// Initializing pluginHub with selected plugins
	plugins, err := cmd.Flags().GetStringArray("plugins")
	pluginHub.InitWith(plugins)

	// Setup Goca instance
	g := &Goca{
		UserAgent: ua,
		Domain:    dm,
		Threshold: th,
		Depth:     dp,
		// BaseFolder: "",
		// DB:         "",
		// Term:       "",
		// Pages:      "",
		Stats: Stats{
			Start:         time.Now(),
			MimeTypeCount: make(map[string]int),
		},
	}

	// Setup the task manager instance
	mq := NewGotam()

	// Setup the plugin hub instance
	plgHub := new(PluginHub)
	plgHub.Init()

	// Setup the scrapper instance
	s := NewScrapper(g)

	err = s.Run()
	logFatal(err)

	links := uniqueLinks(s.Links())

	if len(links) == 1 {
		logWarning("No crawleable links, try with other domain")
		return
	}

	logInfo(fmt.Sprintf("Found a total of %d links", len(links)))
	logDebug(fmt.Sprintf("Starting the file download with %d threads", threads))
	logWarning("Falling back to single thread, multi-threading NYI")

	d := NewDownloader(g, threads)
	d.Links = links
	err = d.Run()
	logFatal(err)

	files := d.Files()

	logDebug("Sending the files to processing pipeline")

	for k, v := range files {
		logDebug(fmt.Sprintf("Found %d files of type %s", len(v), k))
		g.Stats.MimeTypeCount[sanitizeMime(k)] = len(v)
		for i := range v {
			mq.Push(k, v[i])
		}
	}

	logDebug(fmt.Sprintf("MQ registered a total of %d topics", mq.Len()))
	logDebug(fmt.Sprintf("List of topics %v", mq.Mimes()))

	for _, m := range mq.Mimes() {
		plug := pluginHub.Lookup(m)
		if plug != nil {
			for i := 0; i < mq.QLen(m); i++ {
				plug.Run(mq.Pop(m))
			}
		} else {
			logWarning("There is no plugin for mimeType " + m)
		}
	}

	g.Stats.Stop = time.Now()
	delta := g.Stats.Stop.Sub(g.Stats.Start)
	logDebug(fmt.Sprintf("Scan took %s", delta.String()))
}
