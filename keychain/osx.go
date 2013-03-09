/*
Copyright 2013 Thomas Lee

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
package keychain

import (
	"fmt"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

// Gets the username and password for the given hostname e.g. "example.com"
func Credentials(server string) (username string, password string, err error) {
	switch runtime.GOOS {
	case "darwin":
		return osx(server)
	}
	return "", "", fmt.Errorf("Keychain for %v isn't implemented yet.", runtime.GOOS)
}

func osx(server string) (username string, password string, err error) {
	// Get password from keychain
	pwCmd := exec.Command("security", "find-internet-password", "-ws", server)
	pwOut, err := pwCmd.CombinedOutput()
	if err != nil {
		return "", "", fmt.Errorf("Failed to get the password from the keychain for %v.\n%v", server, err)
	}
	pw := strings.TrimRight(string(pwOut), "\n")

	// Get username
	acctCmd := exec.Command("security", "find-internet-password", "-s", server)
	acctOut, err := acctCmd.CombinedOutput()
	if err != nil {
		return "", "", fmt.Errorf("Failed to get the username from the keychain for %v.\n%v", server, err)
	}

	// Get account username from output
	acctRegex := regexp.MustCompile("acct\"<blob>=\"[^\"]+\"")
	acctOut = acctRegex.Find(acctOut)
	if acctOut == nil {
		return "", "", fmt.Errorf("Failed to find account username for %v.", server)
	}
	acct := strings.Split(string(acctOut), "\"")[2]
	return acct, pw, err
}
