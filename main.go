package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type EndpointType struct {
	Path    string
	Method  string
	Content map[string]any
}

type ConfigType struct {
	Port      int
	Endpoints []EndpointType
}

const configFileName = "config.json"

func main() {
	configData, err := os.ReadFile(configFileName)
	if err != nil {
		panic(err)
	}
	var config ConfigType
	err = json.Unmarshal(configData, &config)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, r.URL)
		w.Header().Set("Access-Control-Allow-Origin", "https://dcs-learning.cnu.ac.kr")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With")
		for i := range config.Endpoints {
			e := config.Endpoints[i]
			if e.Path != r.URL.Path || r.Method != e.Method {
				continue
			}
			fmt.Println(e.Method, e.Path)
			b, err := json.Marshal(e.Content)
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")

			_, err = w.Write(b)
			if err != nil {
				panic(err)
			}
		}
	})

	err = http.ListenAndServe(fmt.Sprintf("localhost:%d", config.Port), nil)
	if err != nil {
		panic(err)
	}
}
