// A facebook graph api client in go.
// https://github.com/huandu/facebook/
//
// Copyright 2012, Huan Du
// Licensed under the MIT license
// https://github.com/huandu/facebook/blob/master/LICENSE

package facebook

import (
	"io"
	"net/textproto"
	"reflect"
	"strings"
)

const (
	mimeFormURLEncoded = "application/x-www-form-urlencoded"
	mimeFormData       = "multipart/form-data"
)

var (
	typeOfPointerToBinaryData = reflect.TypeOf(&BinaryData{})
	typeOfPointerToBinaryFile = reflect.TypeOf(&BinaryFile{})
)

// Params is the params used to send Facebook API request.
//
// For general uses, just use Params as an ordinary map.
//
// For advanced uses, use MakeParams to create Params from any struct.
type Params map[string]interface{}

// MakeParams makes a new Params instance by given data.
// Data must be a struct or a map with string keys.
// MakeParams will change all struct field name to lower case name with underscore.
// e.g. "FooBar" will be changed to "foo_bar".
//
// Returns nil if data cannot be used to make a Params instance.
func MakeParams(data interface{}) (params Params) { _ = "STUB: not implemented"; return *new(Params) }

func makeParams(value reflect.Value) (params Params) {
	_ = "STUB: not implemented"
	return *new(Params)
}

// only map with string keys can be converted to Params

// Ignore field if it's not exported

// If field tag "facebook" or "json" exists, use it as field name and options.

// If field tag "facebook" exists, it's preferred.

// If name is not set in field tag, use field name directly.

// these types won't be marshalled in json.

func isEmptyValue(v reflect.Value) bool { _ = "STUB: not implemented"; return false }

// Encode encodes params to query string.
// If map value is not a string, Encode uses json.Marshal() to convert value to string.
//
// Encode may panic if Params contains values that cannot be marshalled to json string.
func (params Params) Encode(writer io.Writer) (mime string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// check whether params contains any binary data.

func (params Params) encodeFormURLEncoded(writer io.Writer) (mime string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (params Params) encodeMultipartForm(writer io.Writer) (mime string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func createFormFile(fieldName, fileName, contentType string) textproto.MIMEHeader {
	_ = "STUB: not implemented"
	return *new(textproto.MIMEHeader)
}
