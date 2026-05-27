// A facebook graph api client in go.
// https://github.com/huandu/facebook/
//
// Copyright 2012, Huan Du
// Licensed under the MIT license
// https://github.com/huandu/facebook/blob/master/LICENSE

package facebook

// App holds facebook application information.
type App struct {
	// Facebook app id
	AppId string

	// Facebook app secret
	AppSecret string

	// Facebook app redirect URI in the app's configuration.
	RedirectUri string

	// Enable appsecret proof in every API call to facebook.
	// Facebook document: https://developers.facebook.com/docs/graph-api/securing-requests
	EnableAppsecretProof bool

	// The session to send request when parsing tokens or code.
	// If it's not set, default session will be used.
	session *Session
}

// New creates a new App and sets app id and secret.
func New(appID, appSecret string) *App { _ = "STUB: not implemented"; return nil }

// AppAccessToken gets application access token, useful for gathering public information about users and applications.
func (app *App) AppAccessToken() string { _ = "STUB: not implemented"; return "" }

// SetSession is used to overwrite the default session used by the app
func (app *App) SetSession(s *Session) {
	_ = "STUB: not implemented"

	// ParseSignedRequest parses signed request.
	return
}

func (app *App) ParseSignedRequest(signedRequest string) (res Result, err error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

// note: here uses the payload base64 string, not decoded bytes

// ParseCode redeems code for a valid access token.
// It's a shorthand call to ParseCodeInfo(code, "").
//
// In facebook PHP SDK, there is a CSRF state to avoid attack.
// That state is not checked in this library.
// Caller is responsible to store and check state if possible.
func (app *App) ParseCode(code string) (token string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ParseCodeInfo redeems code for access token and returns extra information.
// The machineId is optional.
//
// See https://developers.facebook.com/docs/facebook-login/access-tokens#extending
func (app *App) ParseCodeInfo(code, machineID string) (token string, expires int, newMachineID string, err error) {
	_ = "STUB: not implemented"
	return "", 0, "", nil
}

// ExchangeToken exchanges a short-lived access token to a long-lived access token.
// Return new access token and its expires time.
func (app *App) ExchangeToken(accessToken string) (token string, expires int, err error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

// GetCode gets code from a long lived access token.
// Return the code retrieved from facebook.
func (app *App) GetCode(accessToken string) (code string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Session creates a session based on current App setting.
func (app *App) Session(accessToken string) *Session { _ = "STUB: not implemented"; return nil }

// SessionFromSignedRequest creates a session from a signed request.
// If signed request contains a code, it will automatically use this code
// to exchange a valid access token.
func (app *App) SessionFromSignedRequest(signedRequest string) (session *Session, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// it's ok without user id.

// cannot get "oauth_token"? try to get "code".

// no code? no way to continue.
