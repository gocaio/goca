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

// goca.cmd.dorker.go is the cobra command selector entrypoint for the dorker.

import (
	"fmt"
	"time"

	"github.com/gocaio/goca/dork"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func dorkerCmdFunc(cmd *cobra.Command, args []string) {
	setLogLevel(cmd)

	// One shot arguments
	ld, err := cmd.Flags().GetBool("listEngines")
	logFatal(err)
	if ld {
		listEngines()
		return
	}

	// Global arguments
	userAgent := viper.GetString("global.useragent")
	if userAgent == "" {
		userAgent, err = cmd.Flags().GetString("userAgent")
		logFatal(err)
	}

	threads := viper.GetInt("global.threads")
	if threads < 1 {
		logError("The number of threads to run can't be less than 1")
		return
	}

	// Initializing pluginHub with selected plugins
	plugins := viper.GetStringSlice("global.plugins")
	if len(plugins) == 0 || (len(plugins) == 1 && plugins[0] == "false") {
		plugins, err = cmd.Flags().GetStringSlice("plugins")
		logFatal(err)
	}

	// Command arguments
	term := viper.GetString("dorker.term")
	if term == "" {
		term, err = cmd.Flags().GetString("term")
		logFatal(err)
		if term == "" {
			logError("Goca needs a term to dork")
			return
		}
	}

	maxPages := viper.GetInt("dorker.maxpages")
	if maxPages < 1 {
		logError("pages can't be less than 0")
		return
	}

	threshold := viper.GetInt("dorker.threshold")
	if threshold < 0 {
		logError("threshold can't be less than 0")
		return
	}

	engines := viper.GetStringSlice("dorker.engines")
	if len(engines) == 0 || (len(engines) == 1 && engines[0] == "false") {
		engines, err = cmd.Flags().GetStringSlice("engines")
		logFatal(err)
	}

	saveFiles := viper.GetBool("dorker.save")

	if userAgent == "random" {
		userAgent = pickupRandomUA()
	}

	// Initialize pluginhub
	pluginHub.InitWith(plugins)

	// Setup the task manager instance
	mq := NewGotam()

	// Setup the plugin hub instance
	plgHub := new(PluginHub)
	plgHub.Init()

	// Setup Goca instance
	g := &Goca{
		UserAgent: userAgent,
		Threshold: threshold,
		Term:      term,
		Pages:     maxPages,
		Engines:   getEngines(engines),
		Stats: Stats{
			Start:         time.Now(),
			MimeTypeCount: make(map[string]int),
		},
	}

	dorker := NewDorker(g)
	dorker.AddDorksFromPluginHub(plgHub)
	err = dorker.Run()
	logFatal(err)

	// We don't want to process same links twice
	links := uniqueLinks(dorker.Links())

	if len(links) == 0 {
		logWarning("Term dorked without results, try with another term or dorks")
		return
	}

	logInfo(fmt.Sprintf("Found a total of %d links", len(links)))
	logDebug(fmt.Sprintf("Links: %v", links))
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
		g.Stats.MimeTypeCount[k] = len(v)
		for i := range v {
			mq.Push(k, v[i])
			if saveFiles {
				// TODO: Review how files are saved in disk, folder structure and so
				go saveFile(k, v[i])
			}
		}
	}

	logDebug(fmt.Sprintf("MQ registered a total of %d topics", mq.Len()))
	logDebug(fmt.Sprintf("List of topics %v", mq.Mimes()))

	for _, m := range mq.Mimes() {
		plug := pluginHub.Lookup(m)
		if plug != nil {
			totalTopics := mq.QLen(m)
			for i := 0; i < totalTopics; i++ {
				out := plug.Run(mq.Pop(m))
				// TODO: This can be done in parallel
				processOutput(out)
			}
		} else {
			logWarning("There is no plugin for mimeType " + m)
		}
	}

	g.Stats.Stop = time.Now()
	delta := g.Stats.Stop.Sub(g.Stats.Start)
	logDebug(fmt.Sprintf("Scan took %s", delta.String()))
}

func listEngines() {
	fmt.Println("Engines list")
	for _, e := range dork.EngineList {
		fmt.Printf("  - %s\n", e)
	}
}

func getEngines(el []string) (del []dork.Engine) {
	if len(el) == 1 && el[0] == "all" {
		logInfo("Selected engines: all")
		for e := range dork.EngineList {
			del = append(del, e)
		}
		return
	}
	sen := []string{}
	for e, n := range dork.EngineList {
		for _, en := range el {
			if n == en {
				del = append(del, e)
				sen = append(sen, dork.EngineList[e])
			}
		}
	}
	logInfo(fmt.Sprintf("Selected engines: %v", sen))
	return
}

func strInStrMap(a string, b map[string]string) (ok bool) {
	_, ok = b[a]
	return ok
}
