package requests

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func parseResponsesFile(file string) ([]JsonResponse, error) {
	data, readErr := ioutil.ReadFile(file)

	if readErr != nil {
		log.Printf("Reading file '%s' error: %s\n", file, readErr.Error())
		return nil, readErr
	}

	var responses []JsonResponse

	parseErr := json.Unmarshal(data, &responses)
	if parseErr != nil {
		log.Printf("Parsing json error: %s\n", parseErr.Error())
		return nil, parseErr
	}

	return responses, nil
}


func Load(file string) (JsonPaths, error) {
	data, readErr := ioutil.ReadFile(file)

	if readErr != nil {
		return nil, readErr
	}

	var responses JsonPaths

	parseErr := json.Unmarshal(data, &responses)
	if parseErr != nil {
		return nil, parseErr
	}

	for index, response := range responses {
		if response.ResponsesFile != nil {
			log.Printf("Appending data to responses from '%s'\n", *response.ResponsesFile)
			if parsedResponses, err := parseResponsesFile(*response.ResponsesFile); err == nil {
				for _, item := range parsedResponses {
					responses[index].Responses = append(responses[index].Responses, item)
				}
			} else {
				log.Panic(err)
			}
		}
	}

	return responses, nil
}
