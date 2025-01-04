package response

type Response struct {
	Response   interface{} `json:"response,omitempty"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
}

var Errors = map[int]string{
	500: "Что-то пошло не так!",
	400: "Некорректные данные!",
	401: "Пользователь не авторизован!",
	403: "Доступ запрещён!",
}
