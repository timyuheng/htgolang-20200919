package models

const (
	// 0 0 0 0 0
	// v c d m q
	PermAdmin uint = 31
	PermSuper uint = 27
	PermUser  uint = 17
)

type RolePerm struct {
	View   bool //
	Create bool // UserController.Create
	Delete bool // UserController.Delete
	Modify bool // UserController.Edit
	Query  bool // UserController.Query
}
