package models

type BioProfile struct {
	SlackUsername string
	Backend       bool
	Age           int
	Bio           string
}

func (b BioProfile) CreateMyBio() BioProfile {
	return BioProfile{
		"BemEmma",
		true,
		22,
		"Hello everyone, my name is Bemshima I'm here to have fun with golang",
	}
}
