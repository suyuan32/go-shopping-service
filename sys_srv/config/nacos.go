//Author: Ryan SU
//Email: yuansu.china.work@gmail.com

package config

type NacosConfig struct {
	Name         string       `mapstructure:"name"`
	Host         string       `mapstructure:"host"`
	Port         uint64       `mapstructure:"port"`
	Namespace    string       `mapstructure:"namespace"`
	User         string       `mapstructure:"user"`
	Password     string       `mapstructure:"password"`
	DataId       string       `mapstructure:"data_id"`
	Group        string       `mapstructure:"group"`
	ClientConfig ClientConfig `mapstructure:"client"`
}

type ClientConfig struct {
	LogDir   string `mapstructure:"log_dir"`
	CacheDir string `mapstructure:"cache_dir"`
}
