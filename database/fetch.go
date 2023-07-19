package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func FetchSymbolData(symbol string) (map[string]interface{}, error) {
	/**
	* Getting our API Key from .env
	**/
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Failed to load env: ", err)
		return nil, err
	}

	apiKey := os.Getenv("API_KEY")

	if apiKey == "" {
		fmt.Println("No API Key found in ENV")
		return nil, err
	}

	/**
	* Making our GET request
	**/
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&SYMBOL=%S&apikey=%s", symbol, apiKey)

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Request failed. Error: ", err)
	}
	defer response.Body.Close()

	/**
	* If we have a response, map the data and return it
	**/
	if response.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(response.Body)

		if err != nil {
			fmt.Println("Failed to read response body. Error: ", err)
			return nil, err
		}

		var data map[string]interface{}

		err = json.Unmarshal(body, &data)

		if err != nil {
			fmt.Println("Failed to parse JSON. Error: ", err)
			return nil, err
		}

		return data, nil

	} else {
		return nil, err
	}
}
