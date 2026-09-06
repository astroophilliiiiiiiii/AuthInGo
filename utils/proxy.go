package utils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func ProxyToService(targetBaseUrl string, pathPrefix string) http.HandlerFunc {
	target, err := url.Parse(targetBaseUrl) // converted the string into proper url format

	if err != nil {
		fmt.Println("Error parsing target URL:", err)
		return nil
	}

	proxy := httputil.NewSingleHostReverseProxy(target) // readymade reverse proxy engine -- jo requests ko targeturl pe forward

	originalDirector := proxy.Director // saved the original setting that was coming from the client

	// modifying the settings
	proxy.Director = func(r *http.Request) {
		originalDirector(r) // Jo kaam reverse proxy normally karta tha, woh pehle kar do

		r.URL.Path = strings.TrimPrefix(r.URL.Path, pathPrefix) // changing the url too -- removing patPrefic
		r.Host = target.Host                                    // changing the host -- backend service ka host dere

		// Pehle dabbe se int64 nikal (kyunki tune pichhe wahi daala tha)
		if userId, ok := r.Context().Value("userID").(int64); ok {

			// Ab us int64 ko string me convert kar le, kyunki Headers sirf string lete hain
			userIdString := fmt.Sprintf("%d", userId)

			r.Header.Set("X-User_ID", userIdString) // set in the header
		}

	}
	//Request ko actual backend server tak forward karo aur backend ka response client ko wapas do
	return proxy.ServeHTTP // returning the reverse proxy server
}
