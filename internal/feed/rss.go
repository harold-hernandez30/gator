package feed

type RSSFeed struct {
	Channel struct {
		Title string `xml:"title"`
		Link string `xml:"link"`
		Description string `xml:"description"`
		Items []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title string `xml:"title"`
	Link string `xml:"link"`
	Description string `xml:"description"`
	PublicationDate string `xml:"pubDate"`
}