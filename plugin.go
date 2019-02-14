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
	"github.com/h2non/filetype"
	"github.com/h2non/filetype/matchers"
	log "github.com/sirupsen/logrus"
)

var (
	// plugins is the internal map of initialized plugins
	// plugins are stored based on type
	plugins = make(map[string]map[string]Plugin)

	mimeAssociation = newMimeAssoc()
)

// RegisterPlugin is a function that allows plugins to get registered in Goca
func RegisterPlugin(name string, plugin Plugin) {
	if name == "" {
		panic("plugin must have a name")
	}
	if _, ok := plugins[plugin.Type]; !ok {
		plugins[plugin.Type] = make(map[string]Plugin)
	}
	if _, dup := plugins[plugin.Type][name]; dup {
		panic("plugin named " + name + " already registered " + plugin.Type)
	}
	plugins[plugin.Type][name] = plugin

	if (plugin.Assoc == nil || len(plugin.Assoc) == 0) && plugin.Matcher == nil {
		panic("plugin named " + name + "has no mime type associated or matcher defined")
	}

	if plugin.Assoc == nil || len(plugin.Assoc) == 0 {
		log.Warnf("plugin %s has not association with any mime type", name)
	} else {
		mimeAssociation.addMIME(plugin.Type, plugin.Assoc)
	}
	if plugin.Matcher != nil {
		// Register custom matchers
		for _, mime := range plugin.Assoc {
			plugType := filetype.NewType(plugin.Type, mime)
			filetype.AddMatcher(plugType, plugin.Matcher)
		}
	}
}

func listPluginTypes() []string {
	var types []string
	for t := range plugins {
		types = append(types, t)
	}
	return types
}

// IsPluginTypeValid validates plugin against registered plugins
func IsPluginTypeValid(plugType string) bool {
	for _, t := range listPluginTypes() {
		if plugType == t {
			return true
		}
	}
	return plugType == "*"
}

// ListPlugins list loaded plugins
func ListPlugins() map[string][]string {
	in := func(a string, b []string) bool {
		for _, i := range b {
			if a == i {
				return true
			}
		}
		return false
	}
	var pluginsList map[string][]string
	pluginsList = make(map[string][]string)
	for typ, plugs := range plugins {
		if _, ok := pluginsList[typ]; !ok {
			pluginsList[typ] = []string{}
		}
		for name := range plugs {
			if !in(name, pluginsList[typ]) {
				pluginsList[typ] = append(pluginsList[typ], name)
			}
		}
	}
	return pluginsList
}

// Plugin defines the plugin entrypoint
type Plugin struct {
	Type    string
	Assoc   []string
	Action  SetupFunc
	Matcher matchers.Matcher
}

// SetupFunc defines how plugin entrypoint looks like
type SetupFunc func(Manager) error

type mimeAssoc map[string][]string

func newMimeAssoc() mimeAssoc {
	return make(map[string][]string)
}

func (m mimeAssoc) addMIME(typ string, mimes []string) {
	for _, mime := range mimes {
		if !m.existMime(typ, mime) {
			m[typ] = append(m[typ], mime)
		}
	}
}

func (m mimeAssoc) existMime(typ, mime string) bool {
	if typ == "*" {
		return true
	}
	if _, ok := m[typ]; ok {
		for _, mt := range m[typ] {
			if mt == mime {
				return true
			}
		}
	}
	return false
}

func (m mimeAssoc) getAssoc(typ string) []string {
	if _, ok := m[typ]; ok {
		return m[typ]
	}
	return nil
}

func PluginRecover(plugName, funcName string) {
	if r := recover(); r != nil {
		log.Errorf("[%s] Recovered in %s: %s", plugName, funcName, r)
	}
}
