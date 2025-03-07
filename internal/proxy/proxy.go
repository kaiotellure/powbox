package proxy

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

var providers = []string{
	"https://api.proxyscrape.com/v2/?request=displayproxies&protocol=http&timeout=10000&country=all",
}

func gather_proxies(provider string) (list []string, err error) {
	r, err := http.Get(provider)
	if err != nil {
		return
	}
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	list = strings.Split(string(b), "\n")
	return
}

// Fetches proxies from all providers and merges into a single list, if provider fails, the error is appended and moves to next provider
func Fetch() (list []string, errs []error) {
	for _, provider := range providers {
		proxies, err := gather_proxies(provider)

		if err != nil {
			errs = append(errs, err)
			continue
		}

		list = append(list, proxies...)
	}

	return
}

// Returns an http client wrapped by the provided proxy
func Client(proxy string) (client *http.Client, err error) {
	proxyUrl, err := url.Parse("http://" + proxy)
	if err != nil {
		return
	}

	client = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}
	return
}
