package constant

const (
	VALID = "有効"
	IN_VALID = "無効"
	PUBLISHED = "公開"
	REMOVED = "削除"
	ARTICLES_PER_PAGE = 5
	ERR_INVALID_REQUEST_PARAM = "リクエストパラメータが不正です"

	LOG_INFO int = iota
	LOG_ERROR

	LOG_SEPARATOR = " "
	LOG_METHOD_BEGIN = "[BEGIN]"
	LOG_METHOD_END = "[END]"
)
