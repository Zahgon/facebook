// A facebook graph api client in go.
// https://github.com/huandu/facebook/
//
// Copyright 2012, Huan Du
// Licensed under the MIT license
// https://github.com/huandu/facebook/blob/master/LICENSE

package facebook

func camelCaseToUnderScore(str string) string { _ = "STUB: not implemented"; return "" }

// find next non-upper-case character and insert `_` properly.
// it's designed to convert `HTTPServer` to `http_server`.
// if there are more than 2 adjacent upper case characters in a word,
// treat them as an abbreviation plus a normal word.
