package viewmodel

import (
	"time"
)

type Home struct {
	Lastname    string
	Firstname   string
	Title       string
	GithubURL   string
	LinkedinURL string
	TwitterURL  string
	FacebookURL string
	Pages       []interface{}
}

type Page struct {
	Title string
	Text  string
}

type KnownledgePage struct {
	Page
	Experience []string
}

type WorkExperience struct {
	KnownledgePage
	From time.Time
	To   time.Time
}

var VM = new()

func new() *Home {
	return &Home{
		Lastname:    "Druffin",
		Firstname:   "Yann",
		Title:       "Software Developer",
		GithubURL:   "https://github.com/yanndr",
		LinkedinURL: "https://www.linkedin.com/in/yann-druffin-431b3823",
		TwitterURL:  "https://twitter.com/ydruffin",
		Pages: []interface{}{
			&Page{
				Title: "Software Developer",
				Text:  "blabla",
			},
			&KnownledgePage{
				Page: Page{
					Title: "Technologies",
				},
				Experience: []string{"C#", "Go", "Javascript"},
			},
		},
	}
}
