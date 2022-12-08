package database

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Config struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     int    `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Dbname   string `mapstructure:"DB_NAME"`
}

func (config Config) ConnectGenerator(typeDataBase string) string {
	if strings.ToLower(typeDataBase) == "postgres" {
		return fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%d sslmode=disable`,
			config.Host, config.User, config.Password, config.Dbname, config.Port)
	} else {
		return "No provide Database"
	}
}

func ConnectToDB(config Config) (db *gorm.DB) {
	db, err := gorm.Open(postgres.Open(config.ConnectGenerator("postgres")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "testdata.",
		},
	})
	config.migrateDB()
	if err != nil {
		return nil
	}
	return db
}

func CreateConfig(typeDataBase string) Config {
	config := Config{}
	if typeDataBase == "PG" {
		config = LoadConfig("./")
	}
	return config
}

// load config from .env to app
func LoadConfig(path string) Config {
	config := Config{}
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return Config{}
	}
	err = viper.Unmarshal(&config)
	return config
}

// init data base in docker
func (config *Config) migrateDB() {
	migrations := &migrate.FileMigrationSource{
		Dir: "./migrations",
	}
	db, err := sql.Open("postgres", config.ConnectGenerator("postgres"))
	if err != nil {
		panic(err)
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
