package viewmodel

import (
	"time"

	"github.com/yanndr/website/pkg/model"
)

type Home struct {
	Lastname    string
	Firstname   string
	Title       string
	GithubURL   string
	LinkedinURL string
	TwitterURL  string
	FacebookURL string
	Pages       []Sectioner
}

type Sectioner interface {
	Title() string
}

type Section struct {
	title string
	Text  []string
}

func (s *Section) Title() string {
	return s.title
}

type Knowledge struct {
	Section
	Experience []string
}

type WorkExperience struct {
	Section

	From           time.Time
	To             time.Time
	CompanyName    string
	CompanyLogoURL string
}

func NewHome(p *model.Profile) *Home {
	h := Home{
		Lastname:    p.Lastname,
		Firstname:   p.Firstname,
		Title:       p.Title,
		GithubURL:   p.GithubURL,
		LinkedinURL: p.LinkedinURL,
		TwitterURL:  p.TwitterURL,
	}

	h.Pages = append(h.Pages, &Section{
		title: p.Title,
		Text:  p.Bio,
	})

	h.Pages = append(h.Pages)

	for _, e := range p.Employers {
		h.Pages = append(h.Pages, &WorkExperience{
			Section: Section{
				title: e.Name,
				Text:  e.Description,
			},
			From: e.DateStarted,
			To:   e.DateEnded,
		})
	}

	return &h
}
