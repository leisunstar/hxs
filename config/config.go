package config

var (
	// 应用配置
	App = &app{}
)

// 配置
type app struct {
	// app
	Debug          bool   `conf:"debug"`
	DateLayout     string `conf:"date.layout"`
	DatetimeLayout string `conf:"datetime.layout"`
	LogsConf       string `conf:"logs.conf"`

	// db
	DbMysqlDsn                string `conf:"db.mysql.dsn"`
	DbMysqlMaxConn            int    `conf:"db.mysql.max.conn"`
	DbRedisAddr               string `conf:"db.redis.addr"`
	DbRedisMaxConn            int    `conf:"db.redis.max.conn"`
	DbRedisPrefix             string `conf:"db.redis.prefix"`
	DbRedis1DayExpiresSecond  int    `conf:"db.redis.1day.second"`
	DbRedis30DayExpiresSecond int    `conf:"db.redis.30day.second"`

	// rpc
	RPCAddr string `conf:"rpc.addr"`

	// http pprof
	HttpPprofAddr string `conf:"http.pprof.addr"`

	// snowflake
	SnowflakeDataCenterId int64 `conf:"snowflake_data_center_id"`
	SnowflakeWorkerId     int64 `conf:"snowflake_worker_id"`

	// host
	ListenHost string `conf:"listen.host"`
}
