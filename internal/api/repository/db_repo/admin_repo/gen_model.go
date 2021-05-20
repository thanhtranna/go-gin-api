package admin_repo

import "time"

// administrator table
//go:generate gormgen -structs Admin -input .
type Admin struct {
	Id          int32     // primary key
	Username    string    // username
	Password    string    // password
	Nickname    string    // nickname
	Mobile      string    // phone number
	IsUsed      int32     // is used 1: yes -1: no
	IsDeleted   int32     // delete 1: yes -1: no
	CreatedAt   time.Time `gorm:"time"` // created time
	CreatedUser string    // founder
	UpdatedAt   time.Time `gorm:"time"` // update time
	UpdatedUser string    // updater
}
