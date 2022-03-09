package params

type GetUserRsp struct {
	Result   string `json:"result"`
	Msg      string `json:"msg"`
	Account  string `json:"account"`
	Password string `json:"password"`
}
