package adventure

type Story map[string]Chapter

type Chapter struct {
	Intro struct {
		Title     string   `json:"title"`
		Paragraph []string `json:"story"`
		Options   []Option `json:"options"`
	}
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
