package controllers

type CreateUserInput struct {
	Name string `json:"name" binding:"required"`
	Book string `json:"book"`
}
