package models

type BioProfile struct {
	SlackUsername string `json:"slackUsername"`
	Backend       bool   `json:"backend"`
	Age           int    `json:"age"`
	Bio           string `json:"bio"`
}

type OperationType string

const (
	Addition OperationType = OperationType(iota)
	Subtraction
	Multiplication
)

func (operation OperationType) String() string {
	switch operation {
	case Addition:
		return "addition"
	case Subtraction:
		return "subtraction"
	case Multiplication:
		return "multiplication"
	}
	return "unknown"
}

type OperationRequest struct {
	OperationType OperationType `json:"operation_type"`
	X             int64         `json:"x"`
	Y             int64         `json:"y"`
}

type OperationResponse struct {
	SlackUsername string `json:"slackUsername"`
	Result        int64  `json:"result"`
	OperationType string `json:"operation_type"`
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
