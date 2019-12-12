package test

import (
    "../requests"
    "regexp"
    "testing"
)

func checkPath(paths []requests.JsonPath, path string, length int, lengthResponses int) bool {
    if len(paths) != length {
        return false
    }

    if lengthResponses > 0 && len(paths[0].Responses) != lengthResponses {
        return false
    }

    if paths[0].Path != path {
        return false
    }

    return true
}

func TestParsing1(t *testing.T) {
    paths, err := requests.Load("files/test1.json")
    if err != nil {
        t.Errorf("Unable to parse %s", err.Error())
    }

    if !checkPath(paths, "/test1", 1, 1) {
        t.Error("checkPaths failed for TestParsing1")
    }

    headers := paths[0].Responses[0].Headers[0]

    if headers.Name != "test" || headers.Value != "test1" {
        t.Error("headers not valid for json")
    }
}

func TestParsing2(t *testing.T) {
    paths, err := requests.Load("files/test2.json")
    if err != nil {
        t.Errorf("Unable to parse %s", err.Error())
    }

    if !checkPath(paths, "/test2", 1, 0) {
        t.Error("checkPaths failed for TestParsing1")
    }
}

func TestParsing3(t *testing.T) {
    paths, err := requests.Load("files/test3.json")
    if err != nil {
        t.Errorf("Unable to parse %s", err.Error())
    }

    if !checkPath(paths, "/test3", 1, 1) {
        t.Error("checkPaths failed for TestParsing1")
    }

    body, errRead :=  paths[0].Responses[0].GetBody()
    if errRead != nil {
        t.Errorf("Unable to read body %s", errRead.Error())
    }

    if string(body) != "test3" {
        t.Error("Body not correct for TestParsing3")
    }
}

func TestParsing4(t *testing.T) {
    paths, err := requests.Load("files/test4.json")
    if err != nil {
        t.Errorf("Unable to parse %s", err.Error())
    }

    if !checkPath(paths, "/test4", 1, 1) {
        t.Error("checkPaths failed for TestParsing1")
    }

    body, errRead :=  paths[0].Responses[0].GetBody()
    if errRead != nil {
        t.Errorf("Unable to read body %s", errRead.Error())
    }
    
    if found, regxErr := regexp.Match(`<html`, body); regxErr != nil || !found {
        t.Error("HTML body not correct for TestParsing4")
    }
}