// qnap-tool is a command line program to manage QNAP NAS.
// See qnap-tool --help for details.
package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mkch/qnap-tool/qnap"
	"golang.org/x/term"
)

func main() {
	var port uint
	var user, password string
	var https bool
	flag.Usage = usage
	flag.UintVar(&port, "port", 5000, "The port of the admin web page URL")
	flag.StringVar(&user, "user", "", "The user name to login")
	flag.StringVar(&password, "password", "", "The password of user to login")
	flag.BoolVar(&https, "https", false, "Whether the admin page is https enabled")
	flag.Parse()

	if flag.NArg() > 2 {
		errorExit("too many arguments", 1)
	} else if flag.NArg() < 2 {
		errorExit("too few arguments", 1)
	}

	var action = flag.Args()[0]
	var arg = flag.Args()[1]

	switch action {
	case "wake":
		if err := qnap.Wake(arg); err != nil {
			errorExit(err.Error(), 1)
		}
	case "shutdown":
		var userSet, passwordSet bool
		flag.Visit(func(f *flag.Flag) {
			switch f.Name {
			case "password":
				passwordSet = true
			case "user":
				userSet = true
			}
		})
		if !userSet {
			fmt.Fprintln(os.Stdout, "Please enter username:")
			// fmt.Scanln returns "unexpected newline" error if only enter is pressed.
			// user is untouched in this case.
			fmt.Scanln(&user)
		}

		if !passwordSet {
			fmt.Fprintln(os.Stdout, "Please enter password:(not visible)")
			if input, err := term.ReadPassword(int(os.Stdin.Fd())); err != nil {
				log.Fatal(err)
			} else {
				password = string(input)
			}
		}
		var scheme = "http"
		if https {
			scheme = "https"
		}
		var baseUrl = fmt.Sprintf("%v://%v:%v", scheme, arg, port)
		if err := qnap.Shutdown(baseUrl, user, password); err != nil {
			code := 1
			if errors.Is(err, qnap.ErrAuthFailed) {
				code = 2
			}
			errorExit(err.Error(), code)
		}
	default:
		errorExit("invalid acton: "+action, 1)
	}
}

func errorExit(msg string, exitCode int) {
	os.Stderr.WriteString(msg + "\n")
	os.Exit(exitCode)
}
