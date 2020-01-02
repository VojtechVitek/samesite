package samesite

import (
	"net/http"

	"github.com/avct/uasurfer"
)

// Returns SameSite=None cookie attribute based on the list of incompatible browsers,
// as described at https://www.chromium.org/updates/same-site/incompatible-clients.
//
// Imcompatible browser: returns 0 to drop the attribute
// Compatible browser: returns http.SameSiteNoneMode
func None(userAgent string) http.SameSite {
	ua := uasurfer.Parse(userAgent)

	// Versions of Chrome from Chrome 51 to Chrome 66 (inclusive on both ends). These Chrome versions will reject a cookie with `SameSite=None`. This also affects older versions of Chromium-derived browsers, as well as Android WebView. This behavior was correct according to the version of the cookie specification at that time, but with the addition of the new "None" value to the specification, this behavior has been updated in Chrome 67 and newer. (Prior to Chrome 51, the SameSite attribute was ignored entirely and all cookies were treated as if they were `SameSite=None`.)
	if ua.Browser.Name == uasurfer.BrowserChrome && ua.Browser.Version.Major >= 51 && ua.Browser.Version.Major <= 66 {
		return 0
	}

	// Versions of Safari and embedded browsers on MacOS 10.14 and all browsers on iOS 12. These versions will erroneously treat cookies marked with `SameSite=None` as if they were marked `SameSite=Strict`. This bug has been fixed on newer versions of iOS and MacOS.
	if ua.OS.Name == uasurfer.OSiOS && ua.OS.Version.Major == 12 {
		return 0
	}
	if ua.OS.Name == uasurfer.OSMacOSX && ua.OS.Version.Major == 10 && ua.OS.Version.Minor == 14 && ua.Browser.Name == uasurfer.BrowserSafari {
		return 0
	}

	// Versions of UC Browser on Android prior to version 12.13.2. Older versions will reject a cookie with `SameSite=None`. This behavior was correct according to the version of the cookie specification at that time, but with the addition of the new "None" value to the specification, this behavior has been updated in newer versions of UC Browser.
	if ua.OS.Name == uasurfer.OSAndroid && ua.Browser.Name == uasurfer.BrowserUCBrowser && ua.Browser.Version.Less(uasurfer.Version{Major: 12, Minor: 13, Patch: 2}) {
		return 0
	}

	return http.SameSiteNoneMode
}
