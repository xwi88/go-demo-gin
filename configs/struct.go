// Package configs global config struct
package configs

import (
	"time"

	"github.com/kdpujie/log4go"
	"github.com/xwi88/kit4go/json"
)

// Config global config struct
type Config struct {
	NodeInfo NodeInfo         `json:"node_info"`
	APP      APP              `json:"app"`
	Kafka    Kafka            `json:"kafka"`
	Log4go   log4go.LogConfig `json:"log4go"`
}

func (c Config) String() string {
	_b, _ := json.Marshal(c)
	return string(_b)
}

// NodeInfo node info
type NodeInfo struct {
	HostName   string    `json:"hostname"`
	PublicIP   string    `json:"pub_ip"`
	ServerIP   string    `json:"server_ip"`
	PID        int       `json:"pid"`
	NumCPU     int       `json:"num_cpu"`
	GOMAXPROCS int       `json:"go_max_procs"`
	StartTime  time.Time `json:"start_time"`
	WorkDir    string    `json:"work_dir"`
	ConfigFile string    `json:"config_file"`
}

// APP app config
type APP struct {
	AppEnv                         string        `json:"app_env" mapstructure:"app_env"`
	AppName                        string        `json:"app_env" mapstructure:"app_name"`
	Mode                           string        `json:"mode" mapstructure:"mode"`
	Addr                           string        `json:"addr" mapstructure:"addr"`
	CronTaskSyncCacheExpression    string        `json:"cron_task_sync_cache_expression" mapstructure:"cron_task_sync_cache_expression"`
	CronTaskResetStatusExpression  string        `json:"cron_task_reset_status_expression" mapstructure:"cron_task_reset_status_expression"`
	CronTaskQueryBalanceExpression string        `json:"cron_task_query_balance_expression" mapstructure:"cron_task_query_balance_expression"`
	ExportConfigPath               string        `json:"export_config_path" mapstructure:"export_config_path"`
	ExportConfigForce              bool          `json:"export_config_force" mapstructure:"export_config_force"`
	ExportConfig                   bool          `json:"export_config" mapstructure:"export_config"`
	ExportConfigUnique             bool          `json:"export_config_unique" mapstructure:"export_config_unique"`
	TickerInterval                 time.Duration `json:"ticker_interval" mapstructure:"ticker_interval"`

	ReadTimeout  time.Duration `json:"read_timeout" mapstructure:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout" mapstructure:"write_timeout"`
}

// Kafka kafka config
type Kafka struct {
	Producer KafkaProducer `json:"producer" mapstructure:"producer"`
}

// KafkaProducer kafka producer
type KafkaProducer struct {
	Async           bool          `json:"async" mapstructure:"async"`
	Debug           bool          `json:"debug" mapstructure:"debug"`
	Enable          bool          `json:"enable" mapstructure:"enable"`
	SpecifyVersion  bool          `json:"specify_version" mapstructure:"specify_version"`
	ReturnSuccesses bool          `json:"return_successes" mapstructure:"return_successes"`
	Errors          bool          `json:"errors" mapstructure:"errors"`
	BufferSize      int           `json:"buffer_size" mapstructure:"buffer_size"`
	RetryMax        int           `json:"retry_max" mapstructure:"retry_max"`
	Brokers         []string      `json:"brokers" mapstructure:"brokers"`
	Key             string        `json:"key" mapstructure:"key"`
	Topic           string        `json:"topic" mapstructure:"topic"`
	Version         string        `json:"version" mapstructure:"version"`
	Timeout         time.Duration `json:"timeout" mapstructure:"timeout"`
}
