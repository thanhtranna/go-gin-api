package authorized_api_repo

import "time"

// Authorized caller table
//go:generate gormgen -structs AuthorizedApi -input .
type AuthorizedApi struct {
	Id          int32     // primary key
	BusinessKey string    // caller key
	Method      string    // request method
	Api         string    // request address
	IsDeleted   int32     // whether to delete 1: yes -1 no
	CreatedAt   time.Time `gorm:"time"` // created time
	CreatedUser string    // created by
	UpdatedAt   time.Time `gorm:"time"` // updated time
	UpdatedUser string    // updated by
}
