package models

type BioProfile struct {
	SlackUsername string
	Backend       bool
	Age           int
	Bio           string
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
