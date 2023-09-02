package portapps

import (
	"fmt"
	"os"

	"github.com/onaok/kernal/pkg/utl"
	"gopkg.in/yaml.v3"
)

// Config holds portapp configuration details
type Config struct {
	Common Common      `yaml:"common" mapstructure:"common"`
}

// Common holds common configuration data
type Common struct {
	Args       []string          `yaml:"args" mapstructure:"args"`
}

// loadConfig load common and app configuration
func (app *App) loadConfig(appcfg interface{}) (err error) {
	cfgPath := utl.PathJoin(app.RootPath, fmt.Sprintf("%s.yml", app.ID))
	app.config = &Config{
		Common: Common{
			Args:       []string{},
		},
	}

	// Write sample config
	raw, err := yaml.Marshal(app.config)

	// Read config
	raw, err = os.ReadFile(cfgPath)

	return yaml.Unmarshal(raw, &app.config)
}
