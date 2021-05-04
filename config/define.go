package config

type MailConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type GormConfig struct {
	Driver string `yaml:"driver"`
	DSN    string `yaml:"dsn"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type NsqConfig struct {
	Address string `yaml:"address"`
	Topic   string `yaml:"topic"`
	Channel string `yaml:"channel"`
}

type OssConfig struct {
	Endpoint  string `yaml:"endpoint"`
	Bucket    string `yaml:"bucket"`
	AccessKey string `yaml:"accessKey"`
	Secret    string `yaml:"secret"`
}

type SlsLog struct {
	Endpoint  string `yaml:"endpoint"`
	AccessKey string `yaml:"accessKey"`
	Secret    string `yaml:"secret"`
	Project   string `yaml:"project"`
	LogStore  string `yaml:"logStore"`
	Topic     string `yaml:"topic"`
	Source    string `yaml:"source"`
}

type Elastic struct {
	Addresses []string `yaml:"addresses"`
	Username  string   `yaml:"username"`
	Password  string   `yaml:"password"`
	Index     string   `yaml:"index"`
}

type Robot struct {
	AccessToken string `yaml:"accessToken"`
	Secret      string `yaml:"secret"`
}
