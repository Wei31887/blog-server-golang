package config

type Config struct {
	Server Server	`yaml:"server"`
	Db Db			`yaml:"db"`
	Mylog Mylog		`yaml:"mylog"`
	Cache Cache		`yaml:"cache"`
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
  	Expier int
  	Clearup int
}

