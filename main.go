package main

import (
	config "Users/rajtajale/Proj/configs"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	handler "Users/rajtajale/Proj/handlers"
	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//config path
const configPath = "configs/config.json"

func main() {

	//get the application configuration
	appConfig, err := getConfigFile(configPath)
	if err != nil {
		fmt.Errorf("error initializing the config file", err.Error())
	}

	router := mux.NewRouter().StrictSlash(true)

	router.Handle("/v2/claimRTR", handler.GetProvider(appConfig)).Methods(http.MethodGet)

	httpErr := http.ListenAndServe(fmt.Sprintf(":%d", appConfig.Port), gziphandler.GzipHandler(cors.Default().Handler(router)))
	if httpErr != nil {
		fmt.Errorf(httpErr.Error())
	}

}

func getConfigFile(configpath string) (*config.Config, error) {

	fmt.Printf("loading the config file: " + configpath)

	file, err := os.Open(configpath)
	if err != nil {
		fmt.Errorf("error opening the config file", err.Error())
		return nil, err
	}

	defer file.Close()

	configData, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Errorf("error reading the config file", err.Error())
		return nil, err
	}

	configBytes := bytes.NewReader(configData)

	con := &config.Config{}
	decoder := json.NewDecoder(configBytes)
	decoderErr := decoder.Decode(&con)
	if decoderErr != nil {
		fmt.Errorf("error decoding the config file", err.Error())
		return nil, decoderErr
	}

	return con, nil

}
