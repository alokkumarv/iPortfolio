package models

type User struct {
	Name     string `json:"name"`
	LinkedIn string `json:"linkedin"`
	GitHub   string `json:"github"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	About    About  `json:"about"`
}

type About struct {
	Intro string `json:"intro"`
	Title string `json:"title"`
}

type Skills struct {
}
type Resume struct {
}
type Portfolio struct {
}
type Services struct {
}
