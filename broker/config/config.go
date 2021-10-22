package config

// Config represent service config.
type Config struct {
	GRPCAddress string `env:"GRPC_ADDR" valid:",required"`
}

// Validate config struct.
// func (c Config) Validate() error {
// 	_, err := valid.ValidateStruct(c)
// 	return err
// }
