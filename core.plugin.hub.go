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

// core.plugin.hub.go loads plugins in pluginHub

import (
	"github.com/gocaio/goca/plugin"
)

// PluginHub is the analyzer master
type PluginHub struct {
	plugins map[string]*plugin.Plugin
}

// NewPluginHub creates an empty plugin hub
func NewPluginHub() *PluginHub {
	return &PluginHub{}
}

// Init initializes the plugin hub
func (p *PluginHub) Init() {
	plgs := plugin.GetPlugins()
	p.plugins = map[string]*plugin.Plugin{}

	for k, v := range plgs {
		logDebug("Loading plugin for " + k)
		p.plugins[k] = v
	}
}

// InitWith initializes the plugin hub only with selected plugins
func (p *PluginHub) InitWith(plugins []string) {
	if len(plugins) == 1 && plugins[0] == "all" {
		p.Init()
	} else {
		plgs := plugin.GetPlugins()
		p.plugins = map[string]*plugin.Plugin{}

		for k, v := range plgs {
			for _, sp := range plugins {
				if k == sp {
					logDebug("Loading plugin for " + k)
					p.plugins[k] = v
				}
			}
		}
	}
}

// Lookup helps to get a plugin from PluginHub according to its mime type
func (p *PluginHub) Lookup(mime string) (plug *plugin.Plugin) {
	plug = p.plugins[mime]
	return
}
