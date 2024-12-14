// Code generated
// This file is a generated precompile contract config with stubbed abstract functions.
// The file is generated by a template. Please inspect every code and comment in this file before use.

package txwhitelist

import (
	"github.com/ava-labs/subnet-evm/precompile/allowlist"
	"github.com/ava-labs/subnet-evm/precompile/precompileconfig"

	"github.com/ethereum/go-ethereum/common"
)

var _ precompileconfig.Config = &Config{}

// Config implements the precompileconfig.Config interface and
// adds specific configuration for TxWhitelist.
type Config struct {
	allowlist.AllowListConfig
	precompileconfig.Upgrade
	// CUSTOM CODE STARTS HERE
	// Add your own custom fields for Config here
}

// NewConfig returns a config for a network upgrade at [blockTimestamp] that enables
// TxWhitelist with the given [admins], [enableds] and [managers] members of the allowlist .
func NewConfig(blockTimestamp *uint64, admins []common.Address, enableds []common.Address, managers []common.Address) *Config {
	return &Config{
		AllowListConfig: allowlist.AllowListConfig{
			AdminAddresses:   admins,
			EnabledAddresses: enableds,
			ManagerAddresses: managers,
		},
		Upgrade: precompileconfig.Upgrade{BlockTimestamp: blockTimestamp},
	}
}

// NewDisableConfig returns config for a network upgrade at [blockTimestamp]
// that disables TxWhitelist.
func NewDisableConfig(blockTimestamp *uint64) *Config {
	return &Config{
		Upgrade: precompileconfig.Upgrade{
			BlockTimestamp: blockTimestamp,
			Disable:        true,
		},
	}
}

// Key returns the key for the TxWhitelist precompileconfig.
// This should be the same key as used in the precompile module.
func (*Config) Key() string { return ConfigKey }

// Verify tries to verify Config and returns an error accordingly.
func (c *Config) Verify(chainConfig precompileconfig.ChainConfig) error {
	// Verify AllowList first
	if err := c.AllowListConfig.Verify(chainConfig, c.Upgrade); err != nil {
		return err
	}
	// CUSTOM CODE STARTS HERE
	// Add your own custom verify code for Config here
	// and return an error accordingly
	return nil
}

// Equal returns true if [s] is a [*Config] and it has been configured identical to [c].
func (c *Config) Equal(s precompileconfig.Config) bool {
	// typecast before comparison
	other, ok := (s).(*Config)
	if !ok {
		return false
	}
	// CUSTOM CODE STARTS HERE
	// modify this boolean accordingly with your custom Config, to check if [other] and the current [c] are equal
	// if Config contains only Upgrade and AllowListConfig  you can skip modifying it.
	equals := c.Upgrade.Equal(&other.Upgrade) && c.AllowListConfig.Equal(&other.AllowListConfig)
	return equals
}