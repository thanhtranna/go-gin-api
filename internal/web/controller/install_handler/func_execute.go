package install_handler

import (
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/api/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/web/controller/install_handler/mysql_table"
	"github.com/xinliangnote/go-gin-api/pkg/env"
	"github.com/xinliangnote/go-gin-api/pkg/errno"

	"github.com/go-redis/redis/v7"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type initExecuteRequest struct {
	RedisAddr string `form:"redis_addr"` // Connection address, for example: 127.0.0.1:6379
	RedisPass string `form:"redis_pass"` // connection password
	RedisDb   string `form:"redis_db"`   // connect to db

	MySQLAddr string `form:"mysql_addr"`
	MySQLUser string `form:"mysql_user"`
	MySQLPass string `form:"mysql_pass"`
	MySQLName string `form:"mysql_name"`
}

func (h *handler) Execute() core.HandlerFunc {
	return func(c core.Context) {
		req := new(initExecuteRequest)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		versionStr := runtime.Version()
		version := cast.ToFloat32(versionStr[2:6])
		if version < 1.14 {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ConfigGoVersionError,
				code.Text(code.ConfigGoVersionError)),
			)
			return
		}

		outPutString := ""

		cfg := configs.Get()
		redisClient := redis.NewClient(&redis.Options{
			Addr:         req.RedisAddr,
			Password:     req.RedisPass,
			DB:           cast.ToInt(req.RedisDb),
			MaxRetries:   cfg.Redis.MaxRetries,
			PoolSize:     cfg.Redis.PoolSize,
			MinIdleConns: cfg.Redis.MinIdleConns,
		})

		if err := redisClient.Ping().Err(); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ConfigRedisConnectError,
				code.Text(code.ConfigRedisConnectError)).WithErr(err),
			)
			return
		}

		defer redisClient.Close()

		outPutString += "The Redis configuration has been checked to be available.\n"

		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
			req.MySQLUser,
			req.MySQLPass,
			req.MySQLAddr,
			req.MySQLName,
			true,
			"Local")

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			//Logger: logger.Default.LogMode(logger.Info), // Log configuration
		})

		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ConfigMySQLConnectError,
				code.Text(code.ConfigMySQLConnectError)).WithErr(err),
			)
			return
		}

		db.Set("gorm:table_options", "CHARSET=utf8mb4")

		dbClient, _ := db.DB()
		defer dbClient.Close()

		outPutString += "The MySQL configuration has been checked to be available.\n"

		viper.SetConfigName(env.Active().Value() + "_configs")
		viper.SetConfigType("toml")
		viper.AddConfigPath("./configs")
		viper.Set("redis.addr", req.RedisAddr)
		viper.Set("redis.pass", req.RedisPass)
		viper.Set("redis.db", req.RedisDb)

		viper.Set("mysql.read.addr", req.MySQLAddr)
		viper.Set("mysql.read.user", req.MySQLUser)
		viper.Set("mysql.read.pass", req.MySQLPass)
		viper.Set("mysql.read.name", req.MySQLName)

		viper.Set("mysql.write.addr", req.MySQLAddr)
		viper.Set("mysql.write.user", req.MySQLUser)
		viper.Set("mysql.write.pass", req.MySQLPass)
		viper.Set("mysql.write.name", req.MySQLName)

		if viper.WriteConfig() != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ConfigSaveError,
				code.Text(code.ConfigSaveError)).WithErr(err),
			)
			return
		}

		outPutString += "Configuration items Redis and MySQL are successfully configured.\n"

		if err = db.Exec(mysql_table.CreateAuthorizedTableSql()).Error; err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ConfigMySQLInstallError,
				"MySQL "+err.Error()).WithErr(err),
			)
			return
		}
		outPutString += "Initialize MySQL data table: authorized successfully.\n"

		if err = db.Exec(mysql_table.CreateAuthorizedTableDataSql()).Error; err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ConfigMySQLInstallError,
				"MySQL "+err.Error()).WithErr(err),
			)
			return
		}
		outPutString += "Initialize MySQL data table: Authorized default data is successful.\n"

		if err = db.Exec(mysql_table.CreateAuthorizedAPITableSql()).Error; err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ConfigMySQLInstallError,
				"MySQL "+err.Error()).WithErr(err),
			)
			return
		}
		outPutString += "Initialize MySQL data table: authorized_api succeeded.\n"

		if err = db.Exec(mysql_table.CreateAuthorizedAPITableDataSql()).Error; err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ConfigMySQLInstallError,
				"MySQL "+err.Error()).WithErr(err),
			)
			return
		}
		outPutString += "Initialize MySQL data table: authorized_api default data is successful.\n"

		if err = db.Exec(mysql_table.CreateAdminTableSql()).Error; err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ConfigMySQLInstallError,
				"MySQL "+err.Error()).WithErr(err),
			)
			return
		}
		outPutString += "Initialize MySQL data table: admin succeeded.\n"

		if err = db.Exec(mysql_table.CreateAdminTableDataSql()).Error; err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ConfigMySQLInstallError,
				"MySQL "+err.Error()).WithErr(err),
			)
			return
		}
		outPutString += "Initialize MySQL data table: admin default data is successful.\n"

		if err = db.Exec(mysql_table.CreateMenuTableSql()).Error; err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ConfigMySQLInstallError,
				"MySQL "+err.Error()).WithErr(err),
			)
			return
		}
		outPutString += "Initialize MySQL data table: menu succeeded.\n"

		// Generate install completion flag
		f, err := os.Create(configs.ProjectInstallFile())
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ConfigMySQLInstallError,
				"create lock file err:  "+err.Error()).WithErr(err),
			)
			return
		}
		defer f.Close()

		c.Payload(outPutString)
	}
}
