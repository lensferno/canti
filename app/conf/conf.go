package conf

const (
	LoginWebMethod   = "web"
	LoginPPPOEMethod = "pppoe"
)

type Config struct {
	Username  string `json:"username" yaml:"username"`
	Password  string `json:"password" yaml:"password"`
	Method    string `json:"method" yaml:"method"`
	Reconnect bool   `json:"reconnect" yaml:"reconnect"`
	Silence   bool   `json:"silence" yaml:"silence"`
}
