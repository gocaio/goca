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
	"fmt"
	"os"
	"path"

	"github.com/spf13/viper"
)

const (
	gocaFolder = ".goca"
)

var (
	// Version is the semver setted with ldflags
	Version string
	// Codename is the friendly version name setted with ldflags
	Codename string
	// BuildHash represents the git commit from which is Goca is compiled
	BuildHash string
	//BuildTime is the Goca compile timestamp
	BuildTime string

	// Root flags
	loglevel        = "info"
	userAgent       = "Goca Metadata Scanner " + Version
	baseFolder      string
	selectedPlugins = []string{"all"}
	threads         = 1

	// Database flags
	databaseFile = "goca.db"

	// Scrapper flags
	domainToScrap     = ""
	scrapperThreshold = 0
	scrapperDepth     = 1

	// Dorker flags
	termToDork     = ""
	maxPagesToDork = 1
)

func main() {
	initializeBaseFolders()

	// Generate the base goca folder depending on the OS
	uh, err := os.UserHomeDir()
	logFatal(err)
	baseFolder = path.Join(uh, gocaFolder)

	viper.SetConfigName("gocacfg")  // name of config file (without extension)
	viper.AddConfigPath(baseFolder) // call multiple times to add many search paths
	viper.AddConfigPath(".")        // optionally look for config in the working directory
	viper.SetConfigType("toml")

	if err := viper.ReadInConfig(); err != nil { // Find and read the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println("Goca configuration file not found")
			logError("Goca configuration file not found")
		} else {
			// Config file was found but another error was produced
			fmt.Println("An error occurred while reading Goca configuration file")
			logError("An error occurred while reading Goca configuration file")
		}
	}

	logInfo(fmt.Sprintf("Using %s as configuration file", viper.ConfigFileUsed()))

	// The root command shows the help and the banner
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().StringVarP(&loglevel, "loglevel", "L", loglevel, "Log verbosity [debug|info|warning|error]")
	rootCmd.PersistentFlags().StringVarP(&userAgent, "userAgent", "U", userAgent, "The UserAgent to set on the request headers")
	rootCmd.PersistentFlags().StringVarP(&baseFolder, "baseFolder", "B", baseFolder, "Goca's base folder for conf and downloads")
	rootCmd.PersistentFlags().IntVarP(&threads, "threads", "T", threads, "The number of threads used for the file download")
	rootCmd.PersistentFlags().StringSliceVarP(&selectedPlugins, "plugins", "P", selectedPlugins, "Plugins to run through selected command")

	// Viper configuration flags for root command
	viper.BindPFlag("global.basefolder", rootCmd.PersistentFlags().Lookup("baseFolder"))
	viper.BindPFlag("global.useragent", rootCmd.PersistentFlags().Lookup("userAgent"))
	viper.BindPFlag("global.plugins", rootCmd.PersistentFlags().Lookup("plugins"))
	viper.BindPFlag("global.threads", rootCmd.PersistentFlags().Lookup("threads"))
	viper.BindPFlag("global.loglevel", rootCmd.PersistentFlags().Lookup("loglevel"))

	// Database command and its specific flags
	rootCmd.AddCommand(databaseCmd)
	databaseCmd.Flags().StringVarP(&databaseFile, "file", "f", databaseFile, "Database file")
	// databaseCmd.MarkFlagRequired("database.file") // This will be checked in the command itself

	// Viper configuration flags for database command
	viper.BindPFlag("databaseFile", databaseCmd.Flags().Lookup("file"))

	// Scrapper command and its specific flags
	rootCmd.AddCommand(scrapperCmd)
	scrapperCmd.Flags().StringVarP(&domainToScrap, "domain", "d", domainToScrap, "Tells Goca the domain to scrap")
	scrapperCmd.Flags().IntVarP(&scrapperThreshold, "threshold", "t", scrapperThreshold, "This makes Goca wait [t] seconds between URLs")
	scrapperCmd.Flags().IntVarP(&scrapperDepth, "depth", "D", scrapperDepth, "The depth of the scrapping")
	scrapperCmd.Flags().BoolP("save", "s", false, "Save the downloaded files to disk")
	//scrapperCmd.MarkFlagRequired("domain") // This will be checked in the command itself

	// Viper configuration flags for scrapper command
	viper.BindPFlag("scrapper.domain", scrapperCmd.Flags().Lookup("domain"))
	viper.BindPFlag("scrapper.threshold", scrapperCmd.Flags().Lookup("threshold"))
	viper.BindPFlag("scrapper.depth", scrapperCmd.Flags().Lookup("depth"))
	viper.BindPFlag("scrapper.save", scrapperCmd.Flags().Lookup("save"))

	// Crawler command and its specific flags
	rootCmd.AddCommand(dorkerCmd)
	dorkerCmd.Flags().StringVarP(&termToDork, "term", "q", termToDork, "Term for the dork query")
	dorkerCmd.Flags().IntVarP(&maxPagesToDork, "pages", "p", maxPagesToDork, "Maximum search engine result pages to dork")
	dorkerCmd.Flags().IntVarP(&scrapperThreshold, "threshold", "t", scrapperThreshold, "This makes Goca wait [t] seconds between URLs")
	dorkerCmd.Flags().StringSliceP("engines", "e", []string{"all"}, "Engines to drok through")
	dorkerCmd.Flags().BoolP("listEngines", "l", false, "List the avaliable engines")
	dorkerCmd.Flags().BoolP("save", "s", false, "Save the downloaded files to disk")
	// dorkerCmd.MarkFlagRequired("term") // This will be checked in the command itself

	// Viper configuration flags for dorker command
	viper.BindPFlag("dorker.term", dorkerCmd.Flags().Lookup("term"))
	viper.BindPFlag("dorker.maxpages", dorkerCmd.Flags().Lookup("pages"))
	viper.BindPFlag("dorker.threshold", dorkerCmd.Flags().Lookup("threshold"))
	viper.BindPFlag("dorker.engines", dorkerCmd.Flags().Lookup("engines"))
	viper.BindPFlag("dorker.save", dorkerCmd.Flags().Lookup("save"))

	// Plugin command and its specific flags
	rootCmd.AddCommand(pluginCmd)
	pluginCmd.Flags().BoolP("list", "l", false, "List avaliable plugins")

	// Viper configuration flags for plugin command
	viper.BindPFlag("plugin.list", pluginCmd.Flags().Lookup("list"))

	// Run Goca
	execute()
}

func initializeBaseFolders() {
	// Ensure Goca directory is created, otherwise create it
	if _, err := os.Stat(baseFolder); os.IsNotExist(err) {
		os.MkdirAll(baseFolder, os.ModePerm)
	}
	targetFolder = path.Join(baseFolder, newULID())
}
