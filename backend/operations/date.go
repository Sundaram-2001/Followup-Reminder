package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func Date(res http.ResponseWriter, req *http.Request) {
	currentTime := time.Now()
	dateOnly := currentTime.Format("2006-01-02")
	response := map[string]interface{}{
		"date": dateOnly,
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Some error..")
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(jsonResponse)
}
