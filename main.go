package main

import (
  "fmt"

	"github.com/fatih/color"
  "gopkg.in/alecthomas/kingpin.v2"
)

const (
	APPNAME			= "mkd v%s"
	VERSION     = "0.1"
)

const (
	CMDVERSION		= "version"
	CMDINIT       = "init"
	CMDCOMPILE    = "compile"
	CMDSTAGE      = "stage"
	CMDPUBLISH    = "publish"
)

var (
	
	v = kingpin.Command("version", "Print app version")
	i = kingpin.Command("init", "Initialize repository")
	c = kingpin.Command("compile", "Compile templates")
	s = kingpin.Command("stage", "Stage website locally for testing")
	p = kingpin.Command("publish", "Upload static files to production")

)

func version() string {
  return fmt.Sprintf(APPNAME, VERSION)
} // version

func main() {

	switch kingpin.Parse() {
  case CMDVERSION:
    color.Green(version())
	case CMDINIT:
	  color.Green("+ Cloning files from github.com")
	case CMDCOMPILE:
	  color.Green("+ Compiling files")
	case CMDSTAGE:
	  color.Green("+ Staging service locally")
	case CMDPUBLISH:
	  color.Green("+ Publishing site")
	}

} // main
