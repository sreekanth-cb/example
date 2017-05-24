package sample

type Social struct {
	Handle string `json:"handle"`
}
type CBPerson struct {
	LinkedIn *Social `json:"linkedin"`
}
type MediumPayload struct {
	Person *CBPerson `json:"person"`
}
