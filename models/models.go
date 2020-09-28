package models

var Address = []string{"https://www.reddit.com/r/worldnews/.rss"}

type Auth struct {
	Name string `xml:"name"`
}

type Item struct {
	Author  Auth   `xml:"author"`
	Content string `xml:"content"`
	Link    string `xml:"link,attr"`
	Updated string `xml:"updated"`
	Title   string `xml:"title"`
}

type Rss struct {
	Category string `xml:"category"`
	Updated  string `xml:"updated"`
	Title    string `xml:"title"`
	Subtitle string `xml:"subtitle"`
	Entry    []Item `xml:"entry"`
}
