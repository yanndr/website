package viewmodel

import (
	"log"
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
	Pages       []Titler
}

type Titler interface {
	Title() string
}

type Page struct {
	title string
	Text  []string
}

func (p *Page) Title() string {
	return p.title
}

type KnowledgePage struct {
	Page
	Experience []string
}

type WorkExperience struct {
	KnowledgePage
	From time.Time
	To   time.Time
}

var VM = new()

func new() *Home {
	const shortForm = "2006-01-02"
	from, err := time.Parse(shortForm, "2011-06-01")
	if err != nil {
		log.Panic(err)
	}
	to, err := time.Parse(shortForm, "2017-12-31")
	if err != nil {
		log.Panic(err)
	}

	return &Home{
		Lastname:    "Druffin",
		Firstname:   "Yann",
		Title:       "Software Developer",
		GithubURL:   "https://github.com/yanndr",
		LinkedinURL: "https://www.linkedin.com/in/yann-druffin-431b3823",
		TwitterURL:  "https://twitter.com/ydruffin",
		Pages: []Titler{
			&Page{
				title: "Software Developer",
				Text: []string{"I am a software application developer with sixteen years of experience in a variety of work environments, including the business, financial and research sectors.",
					"Some of the approaches I have found useful in my work are Agile development methods, test driven development and MVC design patterns.",
					"I am a very team-oriented person in my work. I really enjoy the flexibility to get to know a project from different perspectives, as well as the productive exchange of multiple points of view.",
				},
			},
			&KnowledgePage{
				Page: Page{
					title: "Technologies",
				},
				Experience: []string{"C#", "Go", "Javascript"},
			},
			&WorkExperience{
				KnowledgePage: KnowledgePage{
					Page: Page{
						title: "Western Union",
					},
				},
				From: from,
				To:   to,
			},
		},
	}
}
