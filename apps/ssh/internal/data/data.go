package data

import (
	_ "embed"
	"encoding/json"
)

//go:embed cv.json
var cvJSON []byte

type CV struct {
	Basics                      Basics         `json:"basics"`
	Education                   []Education    `json:"education"`
	Skills                      []string       `json:"skills"`
	Talks                       []Talk         `json:"talks"`
	Work                        []Work         `json:"work"`
	ExperiencesInOrganization   []Organization `json:"experiences_in_organization"`
	Projects                    []Project      `json:"projects"`
	Achievements                []Achievement  `json:"achievements"`
	Certifications              []Certification `json:"certifications"`
	Socials                     []Social       `json:"socials"`
}

type Basics struct {
	Name     string   `json:"name"`
	Label    string   `json:"label"`
	Email    string   `json:"email"`
	Phone    string   `json:"phone"`
	Website  string   `json:"website"`
	Summary  string   `json:"summary"`
	Location Location `json:"location"`
}

type Location struct {
	City        string `json:"city"`
	CountryCode string `json:"countryCode"`
	Region      string `json:"region"`
}

type Education struct {
	Institution string `json:"institution"`
	Area        string `json:"area"`
	StudyType   string `json:"studyType"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
}

type Talk struct {
	Title   string `json:"title"`
	Event   string `json:"event"`
	Date    string `json:"date"`
	Summary string `json:"summary"`
}

type Work struct {
	Company    string   `json:"company"`
	Position   string   `json:"position"`
	StartDate  string   `json:"startDate"`
	EndDate    string   `json:"endDate"`
	Highlights []string `json:"highlights"`
}

type Organization struct {
	Organization string `json:"organization"`
	Position     string `json:"position"`
	StartDate    string `json:"startDate"`
	EndDate      string `json:"endDate"`
	Summary      string `json:"summary"`
}

type Project struct {
	Name       string   `json:"name"`
	URL        string   `json:"url"`
	Date       string   `json:"date"`
	EndDate    string   `json:"endDate,omitempty"`
	Summary    string   `json:"summary"`
	Highlights []string `json:"highlights"`
	Stack      string   `json:"stack"`
}

type Achievement struct {
	Title      string   `json:"title"`
	Date       string   `json:"date"`
	Summary    string   `json:"summary"`
	Highlights []string `json:"highlights"`
	Stack      string   `json:"stack,omitempty"`
}

type Certification struct {
	Name  string `json:"name"`
	Date  string `json:"date"`
	Score string `json:"score"`
}

type Social struct {
	Network  string `json:"network"`
	Username string `json:"username"`
	URL      string `json:"url"`
}

func LoadCV() (*CV, error) {
	var cv CV
	if err := json.Unmarshal(cvJSON, &cv); err != nil {
		return nil, err
	}
	return &cv, nil
}
