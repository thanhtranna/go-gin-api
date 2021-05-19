package code

// Return structure on error
type Failure struct {
	Code    int    `json:"code"`    // business code
	Message string `json:"message"` // description information
}

const (
	// Service level error code
	ServerError        = 10101
	TooManyRequests    = 10102
	ParamBindError     = 10103
	AuthorizationError = 10104
	CallHTTPError      = 10105
	ResubmitError      = 10106
	ResubmitMsg        = 10107
	HashIdsDecodeError = 10108
	SignatureError     = 10109

	// Business module level error code
	// User module
	IllegalUserName = 20101
	UserCreateError = 20102
	UserUpdateError = 20103
	UserSearchError = 20104

	// Authorized caller
	AuthorizedCreateError    = 20201
	AuthorizedListError      = 20202
	AuthorizedDeleteError    = 20203
	AuthorizedUpdateError    = 20204
	AuthorizedDetailError    = 20205
	AuthorizedCreateAPIError = 20206
	AuthorizedListAPIError   = 20207
	AuthorizedDeleteAPIError = 20208

	// Asdministrator
	AdminCreateError             = 20301
	AdminListError               = 20302
	AdminDeleteError             = 20303
	AdminUpdateError             = 20304
	AdminResetPasswordError      = 20305
	AdminLoginError              = 20306
	AdminLogOutError             = 20307
	AdminModifyPasswordError     = 20308
	AdminModifyPersonalInfoError = 20309

	// Configuration
	ConfigEmailError        = 20401
	ConfigSaveError         = 20402
	ConfigRedisConnectError = 20403
	ConfigMySQLConnectError = 20404
	ConfigMySQLInstallError = 20405
	ConfigGoVersionError    = 20406

	// Utility Toolbox
	SearchRedisError = 20501
	ClearRedisError  = 20502
	SearchRedisEmpty = 20503
	SearchMySQLError = 20504

	// Menu Bar
	MenuCreateError = 20601
	MenuUpdateError = 20602
	MenuListError   = 20603
	MenuDeleteError = 20604
	MenuDetailError = 20605
)

var codeText = map[int]string{
	ServerError:        "Internal Server Error",
	TooManyRequests:    "Too Many Requests",
	ParamBindError:     "The parameter information is incorrect",
	AuthorizationError: "The signature information is incorrect",
	CallHTTPError:      "Failed to call the third-party HTTP interface",
	ResubmitError:      "Resubmit Error",
	ResubmitMsg:        "Do not submit repeatedly",
	HashIdsDecodeError: "ID parameter is wrong",
	SignatureError:     "Signature Error",

	IllegalUserName: "Illegal User Name",
	UserCreateError: "Failed to create user",
	UserUpdateError: "Failed to update user",
	UserSearchError: "Failed to query user",

	AuthorizedCreateError:    "Failed to create authorized",
	AuthorizedListError:      "Failed to get the authorized list page",
	AuthorizedDeleteError:    "Failed to delete authorized",
	AuthorizedUpdateError:    "Failed to update authorized",
	AuthorizedDetailError:    "Failed to get calauthorizedler details",
	AuthorizedCreateAPIError: "Failed to create authorized API address",
	AuthorizedListAPIError:   "Failed to obtain the authorized's API address list",
	AuthorizedDeleteAPIError: "Failed to delete authorized API address",

	AdminCreateError:             "Failed to create administrator",
	AdminListError:               "Failed to get the administrator list page",
	AdminDeleteError:             "Failed to delete administrator",
	AdminUpdateError:             "Failed to update the administrator",
	AdminResetPasswordError:      "Failed to reset password",
	AdminLoginError:              "Login failed",
	AdminLogOutError:             "Failed to exit",
	AdminModifyPasswordError:     "Failed to modify password",
	AdminModifyPersonalInfoError: "Failed to modify personal information",

	ConfigEmailError:        "Failed to modify mailbox configuration",
	ConfigSaveError:         "Failed to write configuration file",
	ConfigRedisConnectError: "Redis connection failed",
	ConfigMySQLConnectError: "MySQL connection failed",
	ConfigMySQLInstallError: "MySQL failed to initialize data",
	ConfigGoVersionError:    "Go Version does not meet the requirements",

	SearchRedisError: "Failed to query Redis Key",
	ClearRedisError:  "Failed to clear Redis Key",
	SearchRedisEmpty: "The Redis Key queried does not exist",
	SearchMySQLError: "Failed to query MySQL",

	MenuCreateError: "Failed to create menu",
	MenuUpdateError: "Failed to update menu",
	MenuDeleteError: "Failed to delete menu",
	MenuListError:   "Failed to get the menu list page",
	MenuDetailError: "Failed to get menu details",
}

func Text(code int) string {
	return codeText[code]
}
