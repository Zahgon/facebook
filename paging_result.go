// A facebook graph api client in go.
// https://github.com/huandu/facebook/
//
// Copyright 2012, Huan Du
// Licensed under the MIT license
// https://github.com/huandu/facebook/blob/master/LICENSE

package facebook

// PagingResult represents facebook API call result with paging information.
type PagingResult struct {
	session  *Session
	paging   pagingData
	previous string
	next     string
}

type pagingData struct {
	Data      []Result `facebook:",required"`
	Paging    *pagingNavigator
	UsageInfo *UsageInfo
}

type pagingNavigator struct {
	Previous string
	Next     string
}

func newPagingResult(session *Session, res Result) (*PagingResult, error) {
	_ = "STUB: not implemented"
	// quick check whether Result is a paging response.
	return nil, nil
}

// Data gets current data.
func (pr *PagingResult) Data() []Result { _ = "STUB: not implemented"; return nil }

// UsageInfo returns API usage information, including
// business use case, app, page, ad account rate limiting.
func (pr *PagingResult) UsageInfo() *UsageInfo { _ = "STUB: not implemented"; return nil }

// Decode decodes the current full result to a struct. See Result#Decode.
func (pr *PagingResult) Decode(v interface{}) (err error) { _ = "STUB: not implemented"; return nil }

// Previous reads previous page.
func (pr *PagingResult) Previous() (noMore bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Next reads next page.
func (pr *PagingResult) Next() (noMore bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// HasPrevious checks whether there is previous page.
func (pr *PagingResult) HasPrevious() bool { _ = "STUB: not implemented"; return false }

// HasNext checks whether there is next page.
func (pr *PagingResult) HasNext() bool { _ = "STUB: not implemented"; return false }

func (pr *PagingResult) navigate(url *string) (noMore bool, err error) {
	_ = "STUB: not implemented"
	return false,

		// add session information in paging url.
		nil
}

// Per #182, access_token is always useless.
// As we may need to keep other params, do a manual delete here.
