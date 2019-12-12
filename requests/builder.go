package requests

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

var mutex = &sync.Mutex{}

func (paths JsonPaths) Build() DynamicHandlerMap {
	var mapped = make(map[string]*DynamicHandler)

	for _, path := range paths {
		mapped[path.Path] = &DynamicHandler{
			Path:    path,
			Hit: 	 0, // this needs to change +1 for each hit
		}
	}

	return mapped
}

func (pathsMap * DynamicHandlerMap) pathExists(path string) bool {
	if _, fine := (* pathsMap)[path]; fine {
		return true
	}
	return false
}

func (pathsMap * DynamicHandlerMap) getResponse(path string) (JsonResponse, error) {
	response := (*pathsMap)[path]

	var realResponse JsonResponse;

	if response.Path.AlwaysShow != nil {
		realResponse = response.Path.Responses[*response.Path.AlwaysShow]
	} else {
		if len(response.Path.Responses) <= response.Hit {
			return realResponse, errors.New(fmt.Sprintf("Only %d supported responses found", response.Hit))
		}

		realResponse = response.Path.Responses[response.Hit]
	}

	return realResponse, nil;
}

func (pathsMap * DynamicHandlerMap) GetHandler(debug string) (func (w http.ResponseWriter, r * http.Request)) {
	log.Printf("Registering handlers with debug endpoint '%s'\n", debug)

	return func(w http.ResponseWriter, r *http.Request) {
		pathMap := *pathsMap
		path := r.URL.Path

		log.Printf("Request to: %s\n", path)
		if strings.Trim(path, "/") == debug {
			pathMap.RunDebugHandler(w, r)
			return
		}


		if !(pathMap).pathExists(r.URL.Path) {
			w.Write([]byte("Unable to find "+path))
			return
		}

		realResponse, err := pathMap.getResponse(path)

		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		body, bodyErr := realResponse.GetBody()

		if bodyErr != nil {
			w.Write([]byte(bodyErr.Error()))
			return
		}

		headers := w.Header()

		for _, header := range realResponse.Headers {
			headers.Set(header.Name, header.Value)
		}

		if realResponse.Code != nil {
			w.WriteHeader(*realResponse.Code)
		}

		w.Write(body)

		if _, valid := pathMap[path]; valid {
			mutex.Lock()
			(*pathsMap)[path].Hit = pathMap[path].Hit + 1
			(*pathsMap)[path].Requests = append((*pathsMap)[path].Requests, MakeRequest(r))
			mutex.Unlock()
		}

	}
}
