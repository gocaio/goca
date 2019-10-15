package plugin

import (
	"github.com/gocaio/goca/dork"
	"github.com/gocaio/goca/rsrc"
)

// plugin.plugin.go defines the public plugin interface

// plugins is a temporary instantiated plugin holder
var plugins map[string]*Plugin

// Plugin defines the base plugin type
type Plugin struct {
	mimetype    string                    // The mimetype to subscribe to in the queue
	Name        string                    // A small plugin name for informational purposes
	Description string                    // A small plugin description for informational purposes
	Dorks       []*dork.Dork              // A list of dorks to catch files
	Check       func([]byte) bool         // The magic bytes checker for the plugin
	Run         func([]byte) *rsrc.Output // The analysis function
}

func init() {
	plugins = map[string]*Plugin{}
}

// NewPlugin returns a new instance of a plugin object
func NewPlugin(mimeType string) *Plugin {
	plg := &Plugin{
		mimetype: mimeType,
	}

	plugins[mimeType] = plg

	return plg
}

// GetPlugins returns the plugin holder
func GetPlugins() map[string]*Plugin { return plugins }
