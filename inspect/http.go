// Prints out the properties of interfaces in the net/http package.
package inspect

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
)

// Prints field names and values to the given output.
func Response(resp *http.Response, out io.Writer) {
	val := reflect.ValueOf(*resp)
	typeOf := val.Type()
	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i).Interface()
		if header, isHeader := fieldVal.(http.Header); isHeader {
			fmt.Fprintf(out, "%v:\n", typeOf.Field(i).Name)
			Header(header, out)
		} else {
			fmt.Fprintf(out, "%v: %v\n", typeOf.Field(i).Name, fieldVal)
		}
	}
}

// Prints http.Header with indents for cookies.
func Header(header http.Header, out io.Writer) {
	for key := range header {
		isMultiVal := len(header[key]) > 1
		if isMultiVal {
			fmt.Fprintf(out, "  %v:\n", key)
		}
		for _, val := range header[key] {
			if isMultiVal {
				fmt.Fprintf(out, "    %v\n", val)
			} else {
				fmt.Fprintf(out, "  %v: %v\n", key, val)
			}
		}
	}
}
