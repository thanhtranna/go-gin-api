package admin_repo

import "time"

// Administrator table
//go:generate gormgen -structs Admin -input .
type Admin struct {
	Id          int32     // primary key
	Username    string    // username
	Password    string    // password
	Nickname    string    // nickname
	Mobile      string    // phone number
	IsUsed      int32     // whether to enable 1: Yes -1: No
	IsDeleted   int32     // whether to delete 1: yes -1: no
	CreatedAt   time.Time `gorm:"time"` // created time
	CreatedUser string    // created by
	UpdatedAt   time.Time `gorm:"time"` // updated time
	UpdatedUser string    // updated by
}
