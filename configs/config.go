package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)


// Guardar alguns dados que são de configuração
type confs struct {
	DBDriver      string           `mapstructure:"DB_DRIVER"`      // Qual o driver do banco de dados
	DBHost        string           `mapstructure:"DB_HOST"`        // Qual o host do banco de dados
	DBPort        string           `mapstructure:"DB_PORT"`        // Qual a porta do banco de dados
	DBUser        string           `mapstructure:"DB_USER"`        // Qual o usuário do banco de dados
	DBPass        string           `mapstructure:"DB_PASS"`        // Qual a senha do banco de dados
	DBName        string           `mapstructure:"DB_NAME"`        // Qual o nome do banco de dados
	DBCharset     string           `mapstructure:"DB_CHARSET"`     // Qual o charset do banco de dados
	AppPort       string           `mapstructure:"APP_PORT"`       // Qual a porta da aplicação
	JWTSecret     string           `mapstructure:"JWT_SECRET"`     // Qual a chave secreta do JWT
	JWTExpiration int              `mapstructure:"JWT_EXPIRATION"` // Qual o tempo de expiração do JWT
	TokenAuth     *jwtauth.JWTAuth // Qual o token de autenticação
}

// Se eu criar uma função init, ela vai ser executada antes de qualquer outra função

// É possível criar um arquivo de configuração e carregar esses valores de lá
func LoadConfig(path string) (*confs, error) {
	// Precisa carregas essas configurações para o resto da aplicação
	var cfg *confs
	// Carregar as configurações do arquivo
	// Se der algum erro, retornar o erro
	// Se não der erro, retornar as configurações
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, nil
}
