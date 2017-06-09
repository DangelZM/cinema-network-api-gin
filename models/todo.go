package models

type TodoModel struct {
	Id int `json:"id"`
	Title string `json:"title"`
}

type TodosModel []TodoModel
