# NeXJSON
Easy JSON parsing, stringifying, and accesing

# Installing
```
go get github.com/NeutronX-dev/NeXJSON
```

# Importing
```go
import "github.com/NeutronX-dev/NeXJSON"
```

# Usage
* `NeXJSON.Parse(ToParse string) interface{}`
* * For Example `"{\"Value\": 0}"`
* `NeXJSON.Stringify(JSON interface{}) string`
* * The parameter of this function is the return of `NeXJSON.Parse`
* `NeXJSON.Access(JSON interface{}, props string)`
* * props e.x.: `"Value.MyOtherValue"`
