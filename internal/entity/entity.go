package entity

type Error string

// Declare error messege
const (
	ErrPermissionNotAllowed = Error("permission.not_allowed")

	//User Error
	ErrUserNotExist            = Error("domain.user.error.not_exist")
	ErrUserAlreadyExist        = Error("domain.user.error.email_or_username_alredy_exist")
	ErrUsersCredentialNotExist = Error("domain.user.error.credential_not_exist")
)

func (e Error) Error() string {
	return string(e)
}
