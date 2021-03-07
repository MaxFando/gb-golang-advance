package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"testing"
)

func ExampleConfiguration_NewConnection() {
	os.Clearenv()

	os.Setenv("DEBUG", "false")
	os.Setenv("PORT", "80")
	os.Setenv("DB_URL", "postgres://user:password@localhost:5432/petstore")

	conf := &Configuration{}
	conf.NewConnection()
}

func ExampleValidateConfiguration() {
	conf := Configuration{
		Debug:     true,
		Port:      80,
		DBURL:     "mysql",
		connected: false,
	}

	_ = ValidateConfiguration(conf)
}

func TestValidUrl(t *testing.T) {
	os.Clearenv()
	os.Setenv("DEBUG", "true")
	os.Setenv("PORT", "80")
	os.Setenv("DB_URL", "mysql")

	t.Run("valid db url", func(t *testing.T) {
		conf := &Configuration{true, 80, "mysql", false}
		assert.Error(t, ValidateConfiguration(*conf))
	})
}

func TestValidateConfiguration(t *testing.T) {
	t.Run("not equal", func(t *testing.T) {
		want := Configuration{
			Debug:     true,
			Port:      80,
			DBURL:     "mysql",
			connected: false,
		}

		conf := Configuration{}
		conf.NewConnection()

		assert.NotEqual(t, reflect.ValueOf(want), reflect.ValueOf(conf))
	})
}
