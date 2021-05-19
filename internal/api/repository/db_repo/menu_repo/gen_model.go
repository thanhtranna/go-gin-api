package menu_repo

import "time"

// Left menu bar table
//go:generate gormgen -structs Menu -input .
type Menu struct {
	Id          int32     // primary key
	Pid         int32     // parent class id
	Name        string    // menu name
	Link        string    // link address
	Icon        string    // icon
	Level       int32     // menu type 1: first level menu 2: second level menu
	IsUsed      int32     // whether to enable: 1 yes -1 no
	IsDeleted   int32     // whether to delete: 1 yes -1 no
	CreatedAt   time.Time `gorm:"time"` // created time
	CreatedUser string    // created by
	UpdatedAt   time.Time `gorm:"time"` // updated time
	UpdatedUser string    // updated by
}
