package models

type Users struct {
	Name  string `json:"name" valid:"required"`
	Email string `json:"email"valid:"email,required"`
	Age   int    `json:"age"valid:"required"`
}
