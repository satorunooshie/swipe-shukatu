package errors

var (
	UserAlreadyExists = newBadRequest("user_registration_error", "すでに登録済みです")
	UserNotFound      = newBadRequest("user_not_found", "ユーザーが見つかりません")
	UserNotActive     = newBadRequest("user_not_active", "ユーザーが見つかりません")
)
