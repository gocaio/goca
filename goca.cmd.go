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

// goca.cmd.go is the cobra command builder and initialization.

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

func execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "goca",
	Short: "Goca is a file metadata extractor with a powerful scraping and dorking engine",
	Long:  gocaBanner,
	// Run:   func(cmd *cobra.Command, args []string) {},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Goca",
	// Long:  `All software has versions. This is Goca's`,
	Run: func(cmd *cobra.Command, args []string) {
		if Version != "" || Codename != "" {
			fmt.Printf("Goca %s -- %s\n", Version, Codename)
			if BuildHash != "" && BuildTime != "" {
				fmt.Printf("Build timestamp: %s\nCommit: %s\n", BuildTime, BuildHash)
			}
			fmt.Printf("Go version: %s\nGo compiler: %s\nPlatform: %s/%s", runtime.Version(), runtime.Compiler, runtime.GOOS, runtime.GOARCH)
		} else {
			fmt.Println("Goca v0.0.0 -- Dev build")
		}
	},
}

var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "Database administration command",
	// Long:  `All software has versions. This is Goca's`,
	Run: databaseCmdFunc,
}

var scrapperCmd = &cobra.Command{
	Use:   "scrapper",
	Short: "This mode is used to scrap a website for documents",
	// Long:  `All software has versions. This is Goca's`,
	PreRun: initializeGoca,
	Run:    scrapperCmdFunc,
}

var dorkerCmd = &cobra.Command{
	Use:   "dorker",
	Short: "This mode uses multiple search engines for the metadata search",
	// Long:  `All software has versions. This is Goca's`,
	PreRun: initializeGoca,
	Run:    dorkerCmdFunc,
}

var pluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "Plugin management command",
	// Long:  `All software has versions. This is Goca's`,
	Run: pluginCmdFunc,
}
