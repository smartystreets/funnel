package internet

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
	response, _ := http.Get(address)
	content, _ := io.ReadAll(response.Body)
	re, _ := regexp.Compile(`\<title\>(.*)\<\/title\>`) // H̸̡̪̯ͨ͊̽̅̾̎Ȩ̬̩̾͛ͪ̈́̀́͘ ̶̧̨̱̹̭̯ͧ̾ͬC̷̙̲̝͖ͭ̏ͥͮ͟Oͮ͏̮̪̝͍M̲̖͊̒ͪͩͬ̚̚͜Ȇ̴̟̟͙̞ͩ͌͝S̨ͯ̿̔̀ ̥̫͎̭ͅ... https://stackoverflow.com/a/1732454
	return re.FindStringSubmatch(string(content))[1]
}
