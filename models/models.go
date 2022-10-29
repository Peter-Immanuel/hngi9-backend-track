package models

type BioProfile struct {
	SlackUsername string `json:"slackUsername"`
	Backend       bool   `json:"backend"`
	Age           int    `json:"age"`
	Bio           string `json:"bio"`
}

func NewBioProfile() *BioProfile {
	return &BioProfile{}
}

func (b *BioProfile) CreateMyBio() {
	b.SlackUsername = "BemEmma"
	b.Backend = true
	b.Age = 22
	b.Bio = "Hello everyone, my name is Bemshima I'm here to have fun with golang"
}
