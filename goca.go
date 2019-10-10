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

func init() {
	// Generate the base goca folder depending on the OS
	baseFolder = "~/.goca/"
	// baseFolder = "C:/goca"

	// The root command shows the help and the banner
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().StringVarP(&loglevel, "loglevel", "l", loglevel, "Log verbosity [debug|info|warning|error]")
	rootCmd.PersistentFlags().StringVarP(&userAgent, "userAgent", "u", userAgent, "The UserAgent to set on the request headers")
	rootCmd.PersistentFlags().StringVarP(&baseFolder, "baseFolder", "b", baseFolder, "Goca's base folder for conf and downloads")
	rootCmd.PersistentFlags().StringArrayVarP(&selectedPlugins, "plugins", "p", selectedPlugins, "Plugins to run through selected command")

	// Database command and its specific flags
	rootCmd.AddCommand(databaseCmd)
	databaseCmd.Flags().StringVarP(&databaseFile, "file", "f", databaseFile, "Database file")
	databaseCmd.MarkFlagRequired("file")

	// Scrapper command and its specific flags
	rootCmd.AddCommand(scrapperCmd)
	scrapperCmd.Flags().StringVarP(&domainToScrap, "domain", "d", domainToScrap, "Tells Goca the domain to scrap")
	scrapperCmd.Flags().IntVarP(&scrapperThreshold, "threshold", "t", scrapperThreshold, "This makes Goca wait [t] seconds between URLs")
	scrapperCmd.Flags().IntVarP(&scrapperDepth, "depth", "D", scrapperDepth, "The depth of the scrapping")
	scrapperCmd.Flags().IntP("threads", "T", 3, "The number of threads used for the file download")
	scrapperCmd.Flags().BoolP("save", "s", false, "Save the downloaded files to disk")
	scrapperCmd.MarkFlagRequired("domain")

	// Crawler command and its specific flags
	rootCmd.AddCommand(dorkerCmd)
	dorkerCmd.Flags().StringVarP(&termToDork, "term", "t", termToDork, "Term for the dork query")
	dorkerCmd.Flags().IntVarP(&maxPagesToDork, "pages", "p", maxPagesToDork, "Maximum search engine result pages to scrap")
	dorkerCmd.Flags().IntVarP(&scrapperThreshold, "threshold", "T", scrapperThreshold, "This makes Goca wait [t] seconds between URLs")
	dorkerCmd.Flags().BoolP("save", "s", false, "Save the downloaded files to disk")
	dorkerCmd.MarkFlagRequired("term")

	// Plugin command and its specific flags
	rootCmd.AddCommand(pluginCmd)
	pluginCmd.Flags().BoolP("list", "L", false, "List avaliable plugins")

}

func main() {
	execute()
}
