
package config

var config *Config

type Config struct {
    Secret string
}
func LoadConfig() *Config{
	//
	return &Config{}
}

func init(){
	config = LoadConfig()
}