package config

//数据数据库配置
type DBconf struct {
	Driver   string `yaml:"driver"`
	Hostname string `yaml:"hostname"`
	Hostport string `yaml:"hostport"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Prefix   string `yaml:"prefix"`
}

//应用配置
type App struct {
	Port                string `yaml:"port"`
	Version             string `yaml:"version"`
	Env                 string `yaml:"env"`
	Apisecret           string `yaml:"apisecret"`
	Allowurl            string `yaml:"allowurl"`
	TokenOutTime        string `yaml:"tokenouttime"`
	CPUnum              string `yaml:"cpunum"`
	Domain              string `yaml:"domain"`
	Vueobjroot          string `yaml:"vueobjroot"`
	CompanyPrivateHouse string `yaml:"companyPrivateHouse"`
	Rootview            string `yaml:"rootview"`
	RunlogType          string `yaml:"runlogtype"`
	NoVerifyTokenRoot   string `yaml:"noVerifyTokenRoot"`
	NoVerifyAPIRoot     string `yaml:"noVerifyAPIRoot"`
	NoVerifyToken       string `yaml:"noVerifyToken"`
	NoVerifyAPI         string `yaml:"noVerifyAPI"`
}

//JWT验证
type Jwt struct {
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"`
	JwtTtl int64  `mapstructure:"jwt_ttl" json:"jwt_ttl" yaml:"jwt_ttl"` // token 有效期（秒）
}

//日志文件
type Log struct {
	Level      string `mapstructure:"level" json:"level" yaml:"level"`
	RootDir    string `mapstructure:"root_dir" json:"root_dir" yaml:"root_dir"`
	Filename   string `mapstructure:"filename" json:"filename" yaml:"filename"`
	Format     string `mapstructure:"format" json:"format" yaml:"format"`
	ShowLine   bool   `mapstructure:"show_line" json:"show_line" yaml:"show_line"`
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"`
	MaxSize    int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"` // MB
	MaxAge     int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`    // day
	Compress   bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
}
