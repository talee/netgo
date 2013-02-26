package main;

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

/** Log error with msg and die. */
func handle(err error, msg string) {
	if err != nil {
		if len(msg) != 0 {
			fmt.Fprint(os.Stderr, "ERROR: ")
			fmt.Fprintln(os.Stderr, msg)
		}
		log.Fatal(err)
	}
}
const (
	ROOT = "http://10.0.0.1/"
	DEVICES = "DEV_device.htm"
	LOG = "fwLog.cgi"
)

/**
* Authenticate to local NETGEAR router and get attached devices.
*/
func main() {
	PAGE := ""
	if len(os.Args) > 1 {
		switch os.Args[1] {
			case "log":
				PAGE = LOG
			case "devices":
				PAGE = DEVICES
		}
	}
	if len(PAGE) == 0 {
		fmt.Fprintln(os.Stderr, "netgo [log, devices]")
		os.Exit(1)
	}

	fmt.Println("Authenticate to local NETGEAR router and get attached devices.\n")

	fmt.Println("Preparing auth request...")
	p, err := ioutil.ReadFile("../.p")  // Readonly
	handle(err, "Failed to read file for auth")
	// Get page based on command-line arguments

	req, err := http.NewRequest("GET", ROOT + PAGE, nil)
	handle(err, "Failed to create new request to " + ROOT + PAGE)
	req.Header.Add("Authorization", "Basic " + string(p))
	p = nil

	resp, err := http.DefaultClient.Do(req)
	handle(err, "Failed to get " + ROOT + PAGE)

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintln(os.Stderr, "ERROR: Bad HTTP status:", resp.StatusCode)
		log.Fatal(err)
	}

	_, err = io.Copy(os.Stdout, resp.Body)
	handle(err, "Failed to copy response to os.Stdout")
}
