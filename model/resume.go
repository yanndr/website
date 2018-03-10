package model

import "time"

//Profile represent your profesional profile.
type Profile struct {
	ID        int
	Firstname string
	Lastname  string
	Telephone string
	Email     string
	Skills    []*Skill
	Projects  []*Project
	Employers []*Employer
}

//ProfileRepository is the interface for the profile repository operations.
type ProfileRepository interface {
	Store(profile *Profile) error
	Get() (*Profile, error)
}

//Project represnet any type of project that you accomplished.
type Project struct {
	Name        string
	Description string
	DateStarted time.Time
	DateEnded   time.Time
	Employer    *Employer
	Skills      []*Skill
}

//Employer represent a employer that you work with.
type Employer struct {
	Name        string
	DateStarted time.Time
	DateEnded   time.Time
	Pojects     []*Project
}

//Skill represent a skill that you have or develop.
type Skill struct {
	Name        string
	Description string
}
