# NeXJSON
Easy JSON parsing, stringifying, and accesing, and replacing.

# Installing
```
go get github.com/NeutronX-dev/NeXJSON
```

# Importing
```go
import "github.com/NeutronX-dev/NeXJSON"
```

# Usage
## Parse
### Code
This will convert a normal string into a usable JSON.
```go
MyJSON, err := NeXJSON.Parse(`{"value": {"hello": "world"}}`)
HandleError(err) // Any code that handles Errors...
```
## Stringify
Will turn a usable JSON into a string. (to save for example)
### Code
```go
MyJSON, err := NeXJSON.Parse(`{"value": {"hello": "world"}}`)
HandleError(err) // Any code that handles Errors...
JSONstring, err := NeXJSON.Stringify(MyJSON)
HandleError(err) // Any code that handles Errors...
fmt.Println(JSONstring)
```
### Output
```
{"value":{"hello":"world"}}
```
## Access
### Code
```go
MyJSON, err := NeXJSON.Parse(`{"value": {"hello": "world"}}`)
HandleError(err) // Any code that handles Errors...
value := NeXJSON.Access(MyJSON, "value.hello")
fmt.Println(value)

/*
	Note: If the property does
		  not exists it might
		  panic.
*/
```
### Output
```
world
```
## Replace
```go
MyJSON, err := NeXJSON.Parse(`{"value": {"hello": "world"}}`)
HandleError(err) // Any code that handles Errors..
fmt.Printf("Before Replacing: %v\n", NeXJSON.Access(MyJSON, "value.hello"))
NeXJSON.Replace(MyJSON, "value.hello", "New Value")
fmt.Printf("After Replacing: %v\n", NeXJSON.Access(MyJSON, "value.hello"))

/*
	Note: If the property does
		  not exists it might
		  panic.
*/
```
### Output
```
Before Replacing: world
After Replacing: New Value
```