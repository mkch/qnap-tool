package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(),
		`Usage of %s:

Utilities to manage QNAP NAS.

qnap-tool [flags] action arg

Actions:
  version: Show version number of this program.
  wake:  Wake up a NAS by Wake-on-LAN.
    qnap-tool wake MAC-ADDRESS
  shutdown:  Shuts down a NAS.
    qnap-tool [-user -password -port -https] shutdown HOST-OF-NAS.
  
Flags:
`, os.Args[0])
	flag.PrintDefaults()
}
