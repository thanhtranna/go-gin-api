package authorized_repo

import "time"

// Authorized caller table
//go:generate gormgen -structs Authorized -input .
type Authorized struct {
	Id                int32     // primary key
	BusinessKey       string    // caller key
	BusinessSecret    string    // caller secret
	BusinessDeveloper string    // caller developer
	Remark            string    // remarks
	IsUsed            int32     // whether to enable 1: yes -1: no
	IsDeleted         int32     // whether to delete 1: yes -1: no
	CreatedAt         time.Time `gorm:"time"` // created time
	CreatedUser       string    // created by
	UpdatedAt         time.Time `gorm:"time"` // updated time
	UpdatedUser       string    // updated by
}
