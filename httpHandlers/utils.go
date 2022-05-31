// this file contains the utility functions for the http handlers
// here we going to implement some features from frameworks (e.g. response writer, validation, etc.)
package httpHandlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/mhkarimi1383/goAPIBaseProject/types"
)

// responseWriter is a function to send response to the client with the given status code
// and decide whether to send the response as json or string then set the content type
func responseWriter[R string | types.UntypedMap | types.Response](w http.ResponseWriter, response *R, status int) error {
	w.WriteHeader(status)
	if reflect.ValueOf(*response).Kind() == reflect.Struct || reflect.ValueOf(*response).Kind() == reflect.Map {
		err := json.NewEncoder(w).Encode(response)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			return errors.New("error in encoding response")
		}
	} else {
		_, err := fmt.Fprintf(w, "%s", *response)
		if err != nil {
			return errors.New("error in writing response")
		}
	}
	return nil
}
