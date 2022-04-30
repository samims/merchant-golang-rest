package config

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"github.com/samims/merchant-api/logger"
	"github.com/spf13/viper"
)

type PostgresConfig interface {
	Host() string
	Port() string
	Database() string
	User() string
	Password() string
	ConnectionURL() string
	GetDB() orm.Ormer
}

type postgresConfig struct {
	env *viper.Viper
}

func (cfg *postgresConfig) Host() string {
	cfg.env.AutomaticEnv()
	return cfg.env.GetString("postgres_host")
}

func (cfg *postgresConfig) Port() string {
	cfg.env.AutomaticEnv()
	return cfg.env.GetString("postgres_port")
}

func (cfg *postgresConfig) Database() string {
	cfg.env.AutomaticEnv()
	return cfg.env.GetString("postgres_db")
}

//
func (cfg *postgresConfig) User() string {
	cfg.env.AutomaticEnv()
	return cfg.env.GetString("postgres_user")
}

//
func (cfg *postgresConfig) Password() string {
	cfg.env.AutomaticEnv()
	return cfg.env.GetString("postgres_password")
}

// ConnectionURL returns connection url for postgresql database
func (cfg *postgresConfig) ConnectionURL() string {
	connectionUrl := fmt.Sprintf(`postgres://%v:%v@%v:%v/%v?sslmode=disable`, cfg.User(), cfg.Password(), cfg.Host(), cfg.Port(), cfg.Database())
	logger.Log.Info(connectionUrl)
	return connectionUrl
}

//"postgres://merchant_db_user:merchant_db_password@localhost:5435/?sslmode=disable

func (cfg *postgresConfig) GetDB() orm.Ormer {
	logger.Log.Info("Connecting to Database.....")
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		logger.Log.Fatal(err)
	}
	err = orm.RegisterDataBase("default", "postgres", cfg.ConnectionURL())
	if err != nil {
		logger.Log.Fatal(err)
	}
	db := orm.NewOrm()
	db.Using("default")
	// auto migration turned on
	orm.RunSyncdb("default", true, true)
	logger.Log.Info("Database connected successfully!!!!")
	return db

}

func NewPostgresConfig(env *viper.Viper) PostgresConfig {
	logger.Log.Info("DB config reading...")
	return &postgresConfig{
		env: env,
	}
}
