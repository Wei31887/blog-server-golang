package config

type Config struct {
	Server Server	`yaml:"server"`
	Db Db			`yaml:"db"`
	Mylog Mylog		`yaml:"mylog"`
	Cache Cache		`yaml:"cache"`
	JWT	  JWT		`yaml:"jwt"`
}

type Server struct {
	Address string	`yaml:"address"`
	Model string	`yaml:"model"`
}

type Db struct {
	Dialects string	`yaml:"dialects"`
  	Host string		`yaml:"host"`
  	Port int		`yaml:"port"`
  	Db string		`yaml:"db"`
  	Username string	`yaml:"username"`
  	Password string	`yaml:"password"`
  	Charset string	`yaml:"charset"`

	MaxIdle int		`yaml:"max-idle-conns"`
	MaxOpen int		`yaml:"max-open-conns"`
}

type Mylog struct{
  	Path string	`yaml:"path"`
  	Name string	`yaml:"name"`
	Model string	`yaml:"model"`
	Format string	`yaml:"format"`
	Level string	`yaml:"level"`
}

type Cache struct {
  	Expire int	`yaml:"expire"`
  	Clearup int	`yaml:"clearup"`
}

type JWT struct {
	SigningKey 	string	`yaml:"signing-key"`
	ExpireTime 	string	`yaml:"expires-time"`
	Issuer 		string 	`yaml:"issuer"`
}
