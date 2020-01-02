package samesite_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/VojtechVitek/samesite"
)

func ExampleCookieSetSameSiteNone() {
	// import "github.com/VojtechVitek/samesite"

	var ( // Imagine we're in a http.Handler.
		w http.ResponseWriter
		r *http.Request
	)

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

func TestCookieSetSameSiteNone(t *testing.T) {
	tt := []struct {
		userAgent                string
		setSameSiteNoneAttribute bool
	}{
		{"Mozilla/5.0 (Linux; Android 4.4.2; SM-T330 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.93 Safari/537.36", true},       // Chrome 43.
		{"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML like Gecko) Chrome/45.0.2454.85 Safari/537.36 115Browser/6.0.3", true},            // Chrome 45.
		{"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0. 2661.102 Safari/537.36", true},                         // Chrome 50.
		{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36", false},              // Chrome 51. Incompatible.
		{"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.59 Safari/537.36 115Browser/8.6.2", false},          // Chrome 54. Incompatible.
		{"Mozilla/5.0 (Linux; Android 7.0) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Focus/1.0 Chrome/59.0.3029.83 Mobile Safari/537.36", false}, // Chrome 59. Incompatible.
		{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0. 3165.0 Safari/537.36", false},                     // Chrome 62. Incompatible.
		{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36", false},              // Chrome 66. Incompatible.
		{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36", true},                     // Chrome 69.
		{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36", true},                // Chrome 79.
	}

	for _, tc := range tt {
		if (samesite.None(tc.userAgent) == http.SameSiteNoneMode) != tc.setSameSiteNoneAttribute {
			t.Errorf("unexpected SameSite attribute on %q", tc.userAgent)
		}
	}
}
