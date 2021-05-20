package menu_repo

import "time"

// left menu bar table
//go:generate gormgen -structs Menu -input .
type Menu struct {
	Id          int32     // primary key
	Pid         int32     // parent id
	Name        string    // menu name
	Link        string    // link address
	Icon        string    // icon
	Level       int32     // menu type 1: first level menu 2: second level menu
	IsUsed      int32     // enable 1: yes -1: no
	IsDeleted   int32     // delete 1: yes -1: no
	CreatedAt   time.Time `gorm:"time"` // created time
	CreatedUser string    // founder
	UpdatedAt   time.Time `gorm:"time"` // update time
	UpdatedUser string    // updater
}
