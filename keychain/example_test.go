package keychain_test

import (
	"bitbucket.org/tlee/netgo/keychain"
	"net/http"
)

func ExampleCredentials() {
	// Get credentials
	HOSTNAME := "bitbucket.org"
	acct, pw, err := keychain.Credentials(HOSTNAME)
	handle(err, "Failed to get credentials for "+HOSTNAME)

	// Prepare HTTP request to router
	req, err := http.NewRequest("GET", PAGE, nil)
	handle(err, "Failed to create new request to "+PAGE)
	req.SetBasicAuth(acct, pw)

	resp, err := http.DefaultClient.Do(req)
}
