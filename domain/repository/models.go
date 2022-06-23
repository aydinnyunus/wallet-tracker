package repository

import (
	"github.com/go-playground/validator"
)

// global constants for file
const (
	semVerRegExp = `^(?P<major>0|[1-9]\d*)\.(?P<minor>0|[1-9]\d*)\.(?P<patch>0|[1-9]\d*)(?:-(?P<prerelease>(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$`
)

// global variables (not cool) for this package
var (
	validate = validator.New()
)

// Adding custom validator operators for our usecase

// Database stores credentials and configurations about strixeye agent database.
type Database struct {
	DBAddr               string `mapstructure:"DB_ADDR" json:"db_addr"  yaml:"db_addr" flag:"db-addr"`
	DBUser               string `mapstructure:"DB_USER" json:"db_user" validate:"omitempty" yaml:"db_user" flag:"db-user"`
	DBPass               string `mapstructure:"DB_PASS" json:"db_pass" validate:"omitempty" yaml:"db_pass" flag:"db-pass"`
	DBName               string `mapstructure:"DB_NAME" json:"db_name" validate:"omitempty" yaml:"db_name" flag:"db-name"`
	DBPort               string `mapstructure:"DB_PORT" json:"db_port" validate:"port" yaml:"db_port" flag:"db-port"`
	OverrideRemoteConfig bool   `mapstructure:"DB_OVERRIDE" json:"override_remote_config" yaml:"override_remote_config" flag:"override-remote-config"`
	testContainerName    string `flag:"test_container_name"`
}

// Validate checks for the fields of given instance.
// check for struct type definition for more documentation about fields and their validation functions.
func (d Database) Validate() error {
	return validate.Struct(d)
}
