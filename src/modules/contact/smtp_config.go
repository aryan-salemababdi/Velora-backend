package contact

type SMTPConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	From     string `yaml:"from"`
	Password string `yaml:"password"`
	To       string `yaml:"to"`
}
