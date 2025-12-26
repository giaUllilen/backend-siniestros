package configs

import (
	"is-public-api/helpers/configloader"
)

type MongoConfig struct {
	Mongodb struct {
		Uri             string `yaml:"uri"`
		ApplicationName string `yaml:"application_name"`
		DatabaseName    string `yaml:"database_name"`
		Ssl             bool   `yaml:"ssl"`
		AuthMechanism   string `yaml:"auth_mechanism"`
		Credentials     struct {
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		} `yaml:"credentials"`
		Cluster struct {
			MaxWaitQueueSize uint64 `yaml:"max_wait_queue_size"`
		} `yaml:cluster`
		ConnectionPool struct {
			MinSize uint64 `yaml:"minSize"`
			MaxSize uint64 `yaml:"maxSize"`
		} `yaml:"connection_pool"`
	} `yaml:"mongodb"`
}

func (conf *MongoConfig) Merge(envCfg interface{}) configloader.ConfigurationProperties {

	envConfig := envCfg.(*MongoConfig)
	conf.Mongodb.Uri = configloader.GetVal(envConfig.Mongodb.Uri, conf.Mongodb.Uri).(string)
	conf.Mongodb.ApplicationName = configloader.GetVal(envConfig.Mongodb.ApplicationName, conf.Mongodb.ApplicationName).(string)
	conf.Mongodb.DatabaseName = configloader.GetVal(envConfig.Mongodb.DatabaseName, conf.Mongodb.DatabaseName).(string)
	conf.Mongodb.AuthMechanism = configloader.GetVal(envConfig.Mongodb.AuthMechanism, conf.Mongodb.AuthMechanism).(string)
	conf.Mongodb.Ssl = configloader.GetVal(envConfig.Mongodb.Ssl, conf.Mongodb.Ssl).(bool)
	conf.Mongodb.Credentials.Username = configloader.GetVal(envConfig.Mongodb.Credentials.Username, conf.Mongodb.Credentials.Username).(string)
	conf.Mongodb.Credentials.Password = configloader.GetVal(envConfig.Mongodb.Credentials.Password, conf.Mongodb.Credentials.Password).(string)
	conf.Mongodb.Cluster.MaxWaitQueueSize = configloader.GetVal(envConfig.Mongodb.Cluster.MaxWaitQueueSize, conf.Mongodb.Cluster.MaxWaitQueueSize).(uint64)
	conf.Mongodb.ConnectionPool.MinSize = configloader.GetVal(envConfig.Mongodb.ConnectionPool.MinSize, conf.Mongodb.ConnectionPool.MinSize).(uint64)
	conf.Mongodb.ConnectionPool.MaxSize = configloader.GetVal(envConfig.Mongodb.ConnectionPool.MaxSize, conf.Mongodb.ConnectionPool.MaxSize).(uint64)

	return conf
}
