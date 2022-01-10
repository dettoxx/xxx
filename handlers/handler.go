package handlers

import (
	"Users/rajtajale/Proj/configs"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetProvider(appConfig *configs.Config) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		//TODO handle the req and make necessary service calls to ProviderRTR
		response := ProviderRtrResponse{}


		//set the response type
		rw.Header().Set("Content-Type", "application/json")

		//encode the response
		if err := json.NewEncoder(rw).Encode(response); err != nil {
			fmt.Errorf("error encoding the response", err.Error())
		}
	}
}