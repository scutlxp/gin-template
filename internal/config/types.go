package config

// Config 配置类
type Config struct {
	APIServerConf APIServerConf `yaml:"api-server"`
}

type APIServerConf struct {
	Port                    int    `yaml:"port"`
	Ip                      string `yaml:"ip"`
	GracefulShutdownTimeSec int    `yaml:"gracefulShutdownTimeSec"`
}
