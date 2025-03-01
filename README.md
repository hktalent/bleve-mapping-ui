# bleve mapping UI

A reusable, web-based editor and viewer UI for bleve IndexMapping
JSON based on angular JS, angular-bootstrap and the angular-ui-control.

解除与gorilla/mux的耦合
支持任意的http server中间件，例如：gin

## Demo

Build the sample webapp...

    go build ./cmd/sample

Run the sample webapp...

    ./sample

Browse to this URL...

    http://localhost:9090/sample.html

## Screenshot

![screenshot](https://raw.githubusercontent.com/blevesearch/bleve-mapping-ui/master/docs/screenshot.png)

## License

Apache License Version 2.0

## For bleve mapping UI developers

### Code generation

There's static bindata resources, which can be regenerated using...
```
go install github.com/elazarl/go-bindata-assetfs/...
go generate

or set 
var StaticBleveMappingPath = flag.String("staticBleveMapping", "./static/", "optional path to static-bleve-mapping directory for web resources")
```
### Unit tests

There are some "poor man's" unit tests, which you can run by visiting...

    http://localhost:9090/mapping_test.html

