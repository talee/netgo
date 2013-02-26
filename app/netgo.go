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
)

/**
* Authenticate to local NETGEAR router and get attached devices.
*/
func main() {
	fmt.Println("Authenticate to local NETGEAR router and get attached devices.\n")

	fmt.Println("Preparing auth request...")
	p, err := ioutil.ReadFile("../.p")  // Readonly
	handle(err, "Failed to read file for auth")

	req, err := http.NewRequest("GET", ROOT + DEVICES, nil)
	handle(err, "Failed to create new request to " + ROOT + DEVICES)
	req.Header.Add("Authorization", "Basic " + string(p))
	p = nil

	resp, err := http.DefaultClient.Do(req)
	handle(err, "Failed to get " + ROOT + DEVICES)

	if resp.StatusCode != http.StatusOK {
		fmt.Fprintln(os.Stderr, "ERROR: Bad HTTP status:", resp.StatusCode)
		log.Fatal(err)
	}

	_, err = io.Copy(os.Stdout, resp.Body)
	handle(err, "Failed to copy response to os.Stdout")
}
