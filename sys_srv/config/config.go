//Author: Ryan SU
//Email: yuansu.china.work@gmail.com

package config

type ServerConfig struct {
	Name            string       `json:"name"`
	Host            string       `json:"host"`
	Port            uint64       `json:"port"`
	Tags            []string     `json:"tags"`
	Lang            string       `json:"lang"`
	MysqlConfig     MysqlConfig  `json:"mysql"`
	ConsulConfig    ConsulConfig `json:"consul"`
	PasswordEncrypt string       `json:"pass_encrypt"`
}

type MysqlConfig struct {
	Host     string `json:"host"`
	Port     int32  `json:"port"`
	Db       string `json:"db"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type ConsulConfig struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}
