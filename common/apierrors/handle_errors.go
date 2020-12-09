package apierrors

import (
	"fmt"
	"net/http"
)

func HandleError(w http.ResponseWriter, custom string, navite error) {
	w.Header().Set("Content-Type", "application/json")
	clienError := JSONError{Message: custom}
	clienError.ToJSON(w, clienError)
	fmt.Printf("\n%s %s\n", custom, navite.Error())
}
