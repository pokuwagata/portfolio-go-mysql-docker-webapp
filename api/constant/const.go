package constant

const (
	HALF_SPACE = " "

	VALID = "有効"
	IN_VALID = "無効"
	PUBLISHED = "公開"
	REMOVED = "削除"
	ARTICLES_PER_PAGE = 5

	SUCCESS_MESSAGE = "OK"
	ERR_INVALID_REQUEST_PARAM = "リクエストパラメータが不正です"
	ERR_SQL_MESSAGE = LOG_ERROR_MARK + HALF_SPACE + "SQL ERROR:" + HALF_SPACE + "%+v"
	ERR_SQL_MESSAGE_DEBUG = LOG_DEBUG_MARK + HALF_SPACE + "SQL ERROR:" + HALF_SPACE + "%+v"

	LOG_INFO int = iota
	LOG_ERROR

	LOG_SEPARATOR = HALF_SPACE
	LOG_METHOD_BEGIN = "[BEGIN]"
	LOG_METHOD_END = "[END]"
	LOG_ERROR_MARK = "[ERROR]"
	LOG_DEBUG_MARK = "[DEBUG]"
)
