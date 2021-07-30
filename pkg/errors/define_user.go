package errors

var (
	// UserAlreadyExists is the error returned when a user already exists.
	UserAlreadyExists = newBadRequest("user_registration_error", "すでに登録済みです")
	// UserNotFound is the error returned when a user is not found.
	UserNotFound = newBadRequest("user_not_found", "ユーザーが見つかりません")
	// UserNotActive is the error returned when a user is not active.
	UserNotActive = newBadRequest("user_not_active", "ユーザーが見つかりません")
)
