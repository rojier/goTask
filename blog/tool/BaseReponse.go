package tool

type Reponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewBaseReponse(code int, msg string, data interface{}) Reponse {
	return Reponse{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func ReponseSuccess(data interface{}) Reponse {
	return Reponse{
		Code: SUCCESS,
		Msg:  SUCCESS_MSG,
		Data: data,
	}
}

func ReponseError() Reponse {
	return Reponse{
		Code: ERROR,
		Msg:  ERROR_MSG,
	}
}
func ReponseErrorMsg(errorMsg string, errorCode int) Reponse {
	return Reponse{
		Code: errorCode,
		Msg:  errorMsg,
	}
}
