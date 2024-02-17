package main

import "errors"

func getSearchUrl(query string, engine string) (string, error) {
	switch engine {
	case "google":
		return "https://google.com/search?q=" + query, nil
	case "duckduckgo":
		return "https://duckduckgo.com/?q=" + query, nil
	case "marginalia":
		return "https://search.marginalia.nu/search?query=" + query, nil
	case "hn":
		return "https://google.com/search?q=site:+news.ycombinator.com+" + query, nil
	default:
		return "", errors.New("unknown engine")
	}
}
