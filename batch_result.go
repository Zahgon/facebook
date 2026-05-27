// A facebook graph api client in go.
// https://github.com/huandu/facebook/
//
// Copyright 2012, Huan Du
// Licensed under the MIT license
// https://github.com/huandu/facebook/blob/master/LICENSE

package facebook

type batchResultHeader struct {
	Name  string `facebook:",required"`
	Value string `facebook:",required"`
}

type batchResultData struct {
	Code    int `facebook:",required"`
	Headers []batchResultHeader
	Body    string `facebook:",required"`
}

func newBatchResult(res Result) (*BatchResult, error) { _ = "STUB: not implemented"; return nil, nil }

// add headers to result.
