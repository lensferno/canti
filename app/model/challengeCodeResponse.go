package model

type ChallengeCodeResponse struct {
	Challenge string `json:"challenge"`

	Error    string `json:"error"`
	ErrorMsg string `json:"error_msg"`
	Res      string `json:"res"`
	SrunVer  string `json:"srun_ver"`
	St       int    `json:"st"`
}

//{
//    "challenge": "3c6d08d667d0ee0ccad77c55b19d3e4ab2552f7163ec40a9389095a18f86c398",
//    "client_ip": "10.16.1.9",
//    "ecode": 0,
//    "error": "ok",
//    "error_msg": "",
//    "expire": "52",
//    "online_ip": "10.16.1.9",
//    "res": "ok",
//    "srun_ver": "SRunCGIAuthIntfSvr V1.18 B20211105",
//    "st": 1668219964
//}
