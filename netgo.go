package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"bitbucket.org/tlee/netgo/inspect"
	"strings"
)

const (
	ROOT    = "http://10.0.0.1/"
	DEVICES = "DEV_device.htm"
	LOG     = "fwLog.cgi"
)

// Authenticate to local NETGEAR router and get attached devices.
func main() {
	// Get arguments to do
	PAGE := getTargetURL()
	fmt.Println()

	// Get authentication tokens
	printTitle("Preparing auth request...")
	p, err := ioutil.ReadFile(".p")
	handle(err, "Failed to read file for auth")

	// Prepare HTTP request to router
	req, err := http.NewRequest("GET", PAGE, nil)
	handle(err, "Failed to create new request to "+PAGE)
	req.Header.Add("Authorization", "Basic "+string(p))
	p = nil

	// Authenticate and get router data
	printTitle("Requesting router data...")
	resp, err := http.DefaultClient.Do(req)
	handle(err, "Failed to get "+PAGE)

	// Respond to HTTP status response
	if resp.StatusCode == http.StatusUnauthorized {
		fmt.Println(resp.StatusCode, http.StatusText(resp.StatusCode), "challenge received.")
		fmt.Println("\nAttempting again...\n")
		// Try authenticating again. The router is required to challenge the client
		// if credentials have expired on the server side
		resp, err = http.DefaultClient.Do(req)
		handle(err, "Failed to get "+PAGE)

		// Exit if HTTP status code is bad
		handleBadHttpStatus(resp)
	} else {
		handleBadHttpStatus(resp)
	}

	// Successfully authenicated and received a good response
	_, err = io.Copy(os.Stdout, resp.Body)
	handle(err, "Failed to copy response to os.Stdout")
}

// Print failure and exit.
func handleBadHttpStatus(resp *http.Response) {
	if resp.StatusCode != http.StatusOK {
		printBadHttpStatus(resp)
		fmt.Fprint(os.Stderr, "\nFailed to get a good response.\n")
		os.Exit(1)
	}
}

// Log error with msg and die.
func handle(err error, msg string) {
	if err != nil {
		if len(msg) != 0 {
			fmt.Fprint(os.Stderr, "ERROR: ")
			fmt.Fprintln(os.Stderr, msg)
		}
		log.Fatal(err)
	}
}

// Print bad HTTP status error msg. Includes response contents.
func printBadHttpStatus(resp *http.Response) {
	fmt.Fprintln(os.Stderr, "ERROR: Bad HTTP status response:")
	inspect.Response(resp, os.Stderr)
}

// Formats given string with an underline using '-'.
func printTitle(title string) {
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", len(title)), "\n")
}

// Get URL from command-line args.
func getTargetURL() string {
	url := ROOT
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "log":
			url += LOG
		case "devices":
			url += DEVICES
		default:
			url = ""
		}
	}
	if len(url) == 0 {
		fmt.Fprintln(os.Stderr, "Authenticates to local NETGEAR router and gets various data.\n")
		fmt.Fprintln(os.Stderr, "usage: netgo [log|devices]")
		os.Exit(1)
	}
	return url
}
