package requests

import (
    "errors"
    "io/ioutil"
    "net/http"
    "net/url"
    "os"
)

type JsonResponseHeaders struct {
    Name    string  `json:"name"`
    Value   string   `json:"value"`
}

type JsonResponse struct {
    Headers         []JsonResponseHeaders   `json:"headers"`
    Code            *int                    `json:"code,omitempty"`
    Body            *string	                `json:"body,omitempty"`
    BodyFromFile    *string                 `json:"body_from_file,omitempty"`
    BodyFromUrl     *string                 `json:"body_from_url,omitempty"`
}

type JsonPath struct {
    Path            string          `json:"path"`
    AlwaysShow      *int            `json:"always_show,omitempty"`
    Responses       []JsonResponse  `json:"responses"`
    ResponsesFile   *string	        `json:"responses_file,omitempty"`
}

type JsonPaths []JsonPath

type Request struct {
    Method   string        `json:"method"`
    Url      *url.URL      `json:"url"`
    Host     string        `json:"host"`
    Body     string        `json:"body"`
    Header   http.Header   `json:"header"`
}

type DynamicHandler struct {
    Path     JsonPath   `json:"path"`
    Hit      int        `json:"hit"`
    Requests []Request  `json:"requests"`
}

type DynamicHandlerMap map[string]*DynamicHandler

func (response JsonResponse) GetBody() ([]byte, error) {
    if response.BodyFromUrl != nil {
        httpResp, error := http.Get(*response.BodyFromUrl)
        if error != nil {
            return nil, error
        }
        
        return ioutil.ReadAll(httpResp.Body)
    }
    
    if response.BodyFromFile != nil {
        if _, err := os.Stat(*response.BodyFromFile); os.IsNotExist(err) {
            return nil, errors.New("unable to read file " + *response.BodyFromFile)
        }

        return ioutil.ReadFile(*response.BodyFromFile)
    }

    if response.Body == nil {
        return nil, errors.New("body field not found")
    }

    return []byte(*response.Body), nil
}

func MakeRequest(r * http.Request) Request {
    body, err := ioutil.ReadAll(r.Body)

    if err != nil {
        body = nil
    }

    return Request{
        Method: r.Method,
        Url: r.URL,
        Host: r.Host,
        Body: string(body),
        Header: r.Header,
    }
}