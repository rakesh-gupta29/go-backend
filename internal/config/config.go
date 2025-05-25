package config

func LoadAllConfigs() error {
	// TODO: logic to load the config from .env file

	loadAppConfig()
	LoadFiberConfig()

	return nil
}
