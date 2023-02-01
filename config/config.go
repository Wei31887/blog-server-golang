package config

type Config struct {
	Server Server	`yaml:"server"`
	Db Db			`yaml:"db"`
	Mylog Mylog		`yaml:"mylog"`
	Cache Cache		`yaml:"cache"`
	JWT	  JWT		`yaml:"jwt"`
}

type Server struct {
	Address string
	Model string
}

type Db struct {
	Dialects string
  	Host string
  	Port int
  	Db string
  	Username string
  	Password string
  	Charset string

	MaxIdle int		`yaml:"max-idle-conns"`
	MaxOpen int		`yaml:"max-open-conns"`
}

type Mylog struct{
  	Path string
  	Name string
	Model string
	Format string
	Level string
}

type Cache struct {
  	Expire int
  	Clearup int
}

type JWT struct {
	SigningKey 	string	`yaml:"signing-key"`
	ExpireTime 	string	`yaml:"expires-time"`
	Issuer 		string 	`yaml:"issuer"`
}
