package constant

const (
	HALF_SPACE = " "

	VALID             = "有効"
	IN_VALID          = "無効"
	PUBLISHED         = "公開"
	REMOVED           = "削除"
	ARTICLES_PER_PAGE = 5

	SUCCESS_MESSAGE           = "OK"
	ERR_MESSAGE_PREFIX        = "エラー："
	ERR_INVALID_REQUEST_PARAM = ERR_MESSAGE_PREFIX + "リクエストパラメータが不正です"
	ERR_ARTICLE_NOT_FOUND     = ERR_MESSAGE_PREFIX + "記事が見つかりません"
	ERR_USER_NOT_FOUND        = ERR_MESSAGE_PREFIX + "ユーザが見つかりません"
	ERR_USER_EXISTED          = ERR_MESSAGE_PREFIX + "既に登録済みのユーザ名です。別のユーザ名を入力してください。"
	ERR_TOKEN_NOT_FOUND       = ERR_MESSAGE_PREFIX + "JWTトークンが見つかりません"
	ERR_INVALID_TOKEN         = ERR_MESSAGE_PREFIX + "不正なJWTトークンです"
	ERR_SIGNUP_FAILED         = ERR_MESSAGE_PREFIX + "ユーザ名またはパスワードが間違っています"

	ERR_SQL_MESSAGE       = LOG_ERROR_MARK + HALF_SPACE + "SQL ERROR:" + HALF_SPACE + "%s"
	ERR_SQL_MESSAGE_DEBUG = LOG_DEBUG_MARK + HALF_SPACE + "SQL ERROR:" + HALF_SPACE + "%+v"
	ERR_APP_ERROR         = LOG_ERROR_MARK + HALF_SPACE + "APP ERROR:" + HALF_SPACE + "%s"
	ERR_APP_ERROR_DEBUG   = LOG_DEBUG_MARK + HALF_SPACE + "APP ERROR:" + HALF_SPACE + "%+v"

	LOG_INFO int = iota
	LOG_ERROR

	LOG_FILE_PATH    = "logs/system.log"
	LOG_SEPARATOR    = HALF_SPACE
	LOG_METHOD_BEGIN = "[BEGIN]"
	LOG_METHOD_END   = "[END]"
	LOG_ERROR_MARK   = "[ERROR]"
	LOG_DEBUG_MARK   = "[DEBUG]"

	USER_ID_COLUMN = "user_id"
)
