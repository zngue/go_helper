package pkg

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

/*
*@Author Administrator
*@Date 31/3/2021 10:14
*@desc
 */

var (
	MysqlConn *gorm.DB
	RedisConn *redis.Client
)

func NewRedis() (*redis.Client, error) {
	DBNUM := viper.GetInt("REDIS.DBNUM")
	DialTimeout := viper.GetInt64("REDIS.DialTimeout")
	ReadTimeout := viper.GetInt64("REDIS.ReadTimeout")
	WriteTimeout := viper.GetInt64("REDIS.WriteTimeout")
	PoolTimeout := viper.GetInt64("REDIS.PoolTimeout")
	PoolSize := viper.GetInt("REDIS.PoolSize")
	MinIdleConns := viper.GetInt("REDIS.MinIdleConns")
	RedisCon := redis.NewClient(&redis.Options{
		Addr:         viper.GetString("REDIS.HOST") + ":" + viper.GetString("REDIS.PORT"),
		Password:     viper.GetString("REDIS.PASSWORD"),
		DialTimeout:  time.Duration(DialTimeout) * time.Second,
		ReadTimeout:  time.Duration(ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(WriteTimeout) * time.Second,
		PoolSize:     PoolSize,
		PoolTimeout:  time.Duration(PoolTimeout) * time.Second,
		MinIdleConns: MinIdleConns,
		DB:           DBNUM,
	})
	_, err := RedisCon.Ping().Result()
	if err != nil {
		return nil, err
	}
	RedisConn = RedisCon
	return RedisCon, nil
}

/*
*@Author Administrator
*@Date 31/3/2021 10:16
*@desc
 */
func NewMysql(opt ...gorm.Option) (*gorm.DB, error) {
	mysqluser := viper.GetString("db.USERNAME")
	mysqlpass := viper.GetString("db.PASSWORD")
	mysqlurls := viper.GetString("db.HOST") + ":" + viper.GetString("db.PORT")
	mysqldb := viper.GetString("db.DATABASE")
	mysqlCharset := viper.GetString("db.CHARSET")
	mysqlParseTime := viper.GetString("db.ParseTime")
	loc := viper.GetString("db.LOC")
	dbs := mysqluser + ":" + mysqlpass + "@tcp(" + mysqlurls + ")/" + mysqldb + "?charset=" + mysqlCharset + "&parseTime=" + mysqlParseTime + "&loc=" + loc
	db, errDb := gorm.Open(mysql.Open(dbs), opt...)
	if errDb != nil {
		return nil, errDb
	}
	//设置日志
	mysqlDebug := viper.GetBool("db.Debug")
	if mysqlDebug {
		db = db.Debug()
	}
	MysqlConn = db
	return db, nil
}
