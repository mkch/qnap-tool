package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(),
		`Usage of %s:
		
qnap-tool [flags] action arg

Actions:
  weak:  Wake up a NAS through Wake-on-LAN.
    qnap-tool wake MAC-ADDRESS
  shutdown:  Shuts down a NAS.
    qnap-tool [-user -password -port -https] shutdown HOST-OF-NAS.
  
Flags:
`, os.Args[0])
	flag.PrintDefaults()
}
