package model

import "time"

//ProfileRepository is the interface for the profile repository operations.
type ProfileRepository interface {
	Store(profile *Profile) error
	Get() (*Profile, error)
}

//Profile represent your profesional profile.
type Profile struct {
	Firstname   string      `json:"firstname"`
	Lastname    string      `json:"lastname"`
	Title       string      `json:"title"`
	Telephone   string      `json:"telephone"`
	Email       string      `json:"email"`
	GithubURL   string      `json:"githubURL"`
	LinkedinURL string      `json:"linkedinURL"`
	TwitterURL  string      `json:"twitterURL"`
	FacebookURL string      `json:"facebookURL"`
	Bio         []string    `json:"bio"`
	Skills      []*Skill    `json:"skills"`
	Projects    []*Project  `json:"projects"`
	Employers   []*Employer `json:"employers"`
}

//Project represnet any type of project that you accomplished.
type Project struct {
	Name        string    `json:"name"`
	Description []string  `json:"description"`
	DateStarted time.Time `json:"dateStarted"`
	DateEnded   time.Time `json:"dateEnded"`
	Employer    *Employer `json:"employer"`
	Skills      []*Skill  `json:"skills"`
}

//Employer represent a employer that you work with.
type Employer struct {
	Name        string     `json:"name"`
	Description []string   `json:"description"`
	DateStarted time.Time  `json:"dateStarted"`
	DateEnded   time.Time  `json:"dateEnded"`
	Projects    []*Project `json:"projects"`
}

//Skill represent a skill that you have or develop.
type Skill struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
