package dummy2

import (
	"fmt"

	"github.com/gocaio/goca/plugin"
	"github.com/gocaio/goca/rsrc"
)

func init() {
	plg := plugin.NewPlugin("application/zip")
	plg.Name = "Dummy plugin 2"
	plg.Description = "This plugin is as dumb as a stone, enjoy :)"
	plg.Check = Check
	plg.Run = Run

	fmt.Println("Hello World from Dummy Plugin 2")
}

// Check is the mimetype filter function, it returns true if it suports the
// given file.
func Check() bool { return true }

// Run starts the analyzer and returns an goca output
func Run(in []byte) (output *rsrc.Output) {
	fmt.Println("Hello world, I'm 'Dummy plugin 2'")
	return output
}
