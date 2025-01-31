package teconfig

func DefaultConfig() TeConfig {
	return TeConfig{
		Listen:   defaultListenConfig(),
		DBConfig: defaultPgConfig(),
	}
}
