package funnel

import (
	"io"
	"net/http"
	"regexp"
)

var Addresses = []string{
	"https://www.google.com",
	"https://www.apple.com",
	"https://michaelwhatcott.com",
	"https://www.catfacts.co",
	"https://stackoverflow.com",
	"https://duckduckgo.com",
	"https://regex101.com",
	"https://trello.com",
	"https://bitbucket.org",
	"https://github.com",
}

func ScrapeTitle(address string) string {
	response, err := http.Get(address)
	if err != nil {
		return err.Error()
	}
	if response.StatusCode != http.StatusOK {
		return response.Status
	}
	content, err := io.ReadAll(response.Body)
	if err != nil {
		return err.Error()
	}
	re, err := regexp.Compile(`\<title\>(.*)\<\/title\>`) // H̸̡̪̯ͨ͊̽̅̾̎Ȩ̬̩̾͛ͪ̈́̀́͘ ̶̧̨̱̹̭̯ͧ̾ͬC̷̙̲̝͖ͭ̏ͥͮ͟Oͮ͏̮̪̝͍M̲̖͊̒ͪͩͬ̚̚͜Ȇ̴̟̟͙̞ͩ͌͝S̨ͯ̿̔̀ ̥̫͎̭ͅ... https://stackoverflow.com/a/1732454
	if err != nil {
		return err.Error()
	}
	matches := re.FindStringSubmatch(string(content))
	if len(matches) > 0 {
		return matches[1]
	}
	return ""
}
