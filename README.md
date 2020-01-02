# Set SameSite=None cookie attribute safely in Golang

[![GoDoc](https://godoc.org/github.com/VojtechVitek/samesite?status.svg)](https://godoc.org/github.com/VojtechVitek/samesite)

Set `SameSite=None` cookie attribute safely, so it's handled well by all incompatible
browsers, as listed at https://www.chromium.org/updates/same-site/incompatible-clients.

## Example:

```go
func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "name",
		Domain:   "example.com",
		Path:     "/",
		Secure:   true,                         // HTTPS only.
		SameSite: samesite.None(r.UserAgent()), // Set SameSite=None unless browser is incompatible.
		HttpOnly: true,
		MaxAge:   3600 * 24 * 365,
		Expires:  time.Now().AddDate(1, 0, 0),
		Value:    "value",
	}

	http.SetCookie(w, &cookie)
}
```
