package configs

import "is-public-api/helpers/configloader"

type ConfigServer struct {
	Server struct {
		Port        int32  `yaml:"port"`
		ContextPath string `yaml:"context_path"`
		SubDomain   string `yaml:"subdomain"`
	} `yaml:"server"`
}

func (conf *ConfigServer) Merge(envCfg interface{}) configloader.ConfigurationProperties {
	envConfig := envCfg.(*ConfigServer)
	conf.Server.Port = configloader.GetVal(envConfig.Server.Port, conf.Server.Port).(int32)
	conf.Server.ContextPath = configloader.GetVal(envConfig.Server.ContextPath, conf.Server.ContextPath).(string)
	conf.Server.SubDomain = configloader.GetVal(envConfig.Server.SubDomain, conf.Server.SubDomain).(string)
	return conf
}
