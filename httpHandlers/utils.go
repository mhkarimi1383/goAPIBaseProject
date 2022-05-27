package httpHandlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/mhkarimi1383/goAPIBaseProject/types"
)

func responseWriter[R string | types.UntypedMap | types.Response](w http.ResponseWriter, response *R, status int) error {
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
