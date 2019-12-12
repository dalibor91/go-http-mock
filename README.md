### Mocking of http webserver 

It can be used for mocking simple http servers it enables you to return multiple responses for one endpoing.

Defining headers is also possible 

```bash
Available commands 
   -file        File with responses to mock
   -host        Ip or hostname to start webserver 
   -port        Port to be used for webserver 
   -debug       Endpoint that will show you debug data
```

Example of mocked requests can be found in `examples/test.json`

To run it just run 
```bash
go run main.go -file examples/test.json
```

Or 

```bash
go run main.go -file example/html/config.json
```


If you don't have go you can use *docker* just run 

```bash
bin/docker-run go run main.go -file examples/test.json
``` 

and check `localhost:8787` or if you want to bind it to specific port run 

```bash 
PORT=8989 bin/docker-run go run main.go -file examples/test.json
```

and check `localhost:8989`


To mock requests you need to create `json` file with urls that you want to mock
Example:
```json
[
  {
    "path": "/", 
    "always_show": 0,
    "responses": [
      {
        "headers": [
          { "name": "Content-Type",  "value": "application/json" }
        ],
        "body": "{\"this\": \"is root page\"}"
      }
    ]
  }
]
```

Json file is an array with urls (paths) which you mocked, and responses for those paths. 

If your page has multiple responses (it changes for each request) add items `responses` array and for each request 
it will show next response up until there is no more responses in that case it will show 
`Only X supported responses found` as error message

if you page has only one response use `"always_show": 0` where `0` indicates index of which response you want to use.

Structure of json file can be described as following :

  - `path` *required* string - path that we mock, for example `/test`
  - `always_show` optional integer - index of response that will be used always as response
  - `responses` *required* array - responses that are mocked
     - `headers` *required* array - with headers 
        - `name` *required* string - name of header 
        - `value` *required* string - value of header
     - `body` optional string - body that will be shown 
     - `body_from_file` optional string - show body from specific file, in this case `body` is ignored
     - `body_from_url` optional string - show body from specific url
     - `code` optional integer - status code that will be sent
  - `responses_file` will append requests to above requests array from file here. (See `examples/separate_responses`)
     
Please check examples to see how this works.

### Debug 

To see all requests with request data you can check `__debug` endpoint or if you want to change this endpoint because it conflicts with some of your endpoint start program with `-debug [your-endpoint]`

### Tests 

Run:
```bash
go test ./test
```


### TODO:
- add app config also, where we can change http/https

- add virtual hosts support 


