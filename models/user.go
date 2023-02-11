package models

type User struct {
	Id      int
	Name    string
	Surname string
	Date    int
}

type GetListRequest struct {
	Offset int
	Limit  int
}
