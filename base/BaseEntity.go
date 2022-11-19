package base

type ResponseEntity struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ResponseListEntity struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Total int         `json:"total"`
}

type BadResponseEntity struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
