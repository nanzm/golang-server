package config

type LogConfig struct {
	File string
}

type SecretConfig struct {
	Secret string
}

type LogStore struct {
	Enable string
}

type SlsLog struct {
	Endpoint  string
	AccessKey string
	Secret    string
	Project   string
	LogStore  string
	Topic     string
	Source    string
}

type Elastic struct {
	Addresses []string
	Username  string
	Password  string
	Index     string
}

type GormConfig struct {
	Driver string
	DSN    string
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

type OssConfig struct {
	Endpoint  string
	Bucket    string
	AccessKey string
	Secret    string
}

type MailConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

type DingDingRobot struct {
	AccessToken string
	Secret      string
}

type NsqConfig struct {
	Address string
	Topic   string
	Channel string
}
