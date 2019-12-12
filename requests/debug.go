package requests

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func (pathsMap * DynamicHandlerMap) RunDebugHandler(writer http.ResponseWriter, request * http.Request) {
    writer.Header().Set("Content-Type", "application/json")

    data, error := json.Marshal(pathsMap)

    if error != nil {
        writer.Write([]byte(fmt.Sprintf("Error: %s", error.Error())))
    } else {
        writer.Write(data)
    }
}
