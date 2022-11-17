package network

type ChallengeCodeResponse struct {
	Challenge string `json:"challenge"`
	ClientIp  string `json:"client_ip"`
	Ecode     int    `json:"ecode"`
	Error     string `json:"error"`
	ErrorMsg  string `json:"error_msg"`
	Expire    string `json:"expire"`
	OnlineIp  string `json:"online_ip"`
	Res       string `json:"res"`
	SrunVer   string `json:"srun_ver"`
	St        int    `json:"st"`
}
