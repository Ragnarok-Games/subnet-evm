// Code generated
// This file is a generated precompile contract test with the skeleton of test functions.
// The file is generated by a template. Please inspect every code and comment in this file before use.

package txwhitelist

import (
	"math/big"
	"testing"

	"github.com/ava-labs/subnet-evm/core/state"
	"github.com/ava-labs/subnet-evm/precompile/allowlist"
	"github.com/ava-labs/subnet-evm/precompile/testutils"
	"github.com/ava-labs/subnet-evm/vmerrs"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

var (
	_ = vmerrs.ErrOutOfGas
	_ = big.NewInt
	_ = common.Big0
	_ = require.New
)

// These tests are run against the precompile contract directly with
// the given input and expected output. They're just a guide to
// help you write your own tests. These tests are for general cases like
// allowlist, readOnly behaviour, and gas cost. You should write your own
// tests for specific cases.
var (
	tests = map[string]testutils.PrecompileTest{
		"calling getTxWhitelistStatus from NoRole should succeed": {
			Caller:     allowlist.TestNoRoleAddr,
			BeforeHook: allowlist.SetDefaultRoles(Module.Address),
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// populate test input here
				testInput := GetTxWhitelistStatusInput{}
				input, err := PackGetTxWhitelistStatus(testInput)
				require.NoError(t, err)
				return input
			},
			// This test is for a successful call. You can set the expected output here.
			// CUSTOM CODE STARTS HERE
			ExpectedRes: func() []byte {

				var output bool // CUSTOM CODE FOR AN OUTPUT
				output = false  // CUSTOM CODE FOR AN OUTPUT
				packedOutput, err := PackGetTxWhitelistStatusOutput(output)
				if err != nil {
					panic(err)
				}
				return packedOutput
			}(),
			SuppliedGas: GetTxWhitelistStatusGasCost,
			ReadOnly:    false,
			ExpectedErr: "",
		},
		"calling getTxWhitelistStatus from Enabled should succeed": {
			Caller:     allowlist.TestEnabledAddr,
			BeforeHook: allowlist.SetDefaultRoles(Module.Address),
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// populate test input here
				testInput := GetTxWhitelistStatusInput{}
				input, err := PackGetTxWhitelistStatus(testInput)
				require.NoError(t, err)
				return input
			},
			// This test is for a successful call. You can set the expected output here.
			// CUSTOM CODE STARTS HERE
			ExpectedRes: func() []byte {

				var output bool // CUSTOM CODE FOR AN OUTPUT
				output = false  // CUSTOM CODE FOR AN OUTPUT
				packedOutput, err := PackGetTxWhitelistStatusOutput(output)
				if err != nil {
					panic(err)
				}
				return packedOutput
			}(),
			SuppliedGas: GetTxWhitelistStatusGasCost,
			ReadOnly:    false,
			ExpectedErr: "",
		},
		"calling getTxWhitelistStatus from Manager should succeed": {
			Caller:     allowlist.TestManagerAddr,
			BeforeHook: allowlist.SetDefaultRoles(Module.Address),
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// populate test input here
				testInput := GetTxWhitelistStatusInput{}
				input, err := PackGetTxWhitelistStatus(testInput)
				require.NoError(t, err)
				return input
			},
			// This test is for a successful call. You can set the expected output here.
			// CUSTOM CODE STARTS HERE
			ExpectedRes: func() []byte {

				var output bool // CUSTOM CODE FOR AN OUTPUT
				output = false  // CUSTOM CODE FOR AN OUTPUT
				packedOutput, err := PackGetTxWhitelistStatusOutput(output)
				if err != nil {
					panic(err)
				}
				return packedOutput
			}(),
			SuppliedGas: GetTxWhitelistStatusGasCost,
			ReadOnly:    false,
			ExpectedErr: "",
		},
		"calling getTxWhitelistStatus from Admin should succeed": {
			Caller:     allowlist.TestAdminAddr,
			BeforeHook: allowlist.SetDefaultRoles(Module.Address),
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// populate test input here
				testInput := GetTxWhitelistStatusInput{}
				input, err := PackGetTxWhitelistStatus(testInput)
				require.NoError(t, err)
				return input
			},
			// This test is for a successful call. You can set the expected output here.
			// CUSTOM CODE STARTS HERE
			ExpectedRes: func() []byte {

				var output bool // CUSTOM CODE FOR AN OUTPUT
				output = false  // CUSTOM CODE FOR AN OUTPUT
				packedOutput, err := PackGetTxWhitelistStatusOutput(output)
				if err != nil {
					panic(err)
				}
				return packedOutput
			}(),
			SuppliedGas: GetTxWhitelistStatusGasCost,
			ReadOnly:    false,
			ExpectedErr: "",
		},
		"insufficient gas for getTxWhitelistStatus should fail": {
			Caller: common.Address{1},
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// populate test input here
				testInput := GetTxWhitelistStatusInput{}
				input, err := PackGetTxWhitelistStatus(testInput)
				require.NoError(t, err)
				return input
			},
			SuppliedGas: GetTxWhitelistStatusGasCost - 1,
			ReadOnly:    false,
			ExpectedErr: vmerrs.ErrOutOfGas.Error(),
		},
		"calling setTxWhitelistStatus from NoRole should fail": {
			Caller:     allowlist.TestNoRoleAddr,
			BeforeHook: allowlist.SetDefaultRoles(Module.Address),
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// populate test input here
				testInput := SetTxWhitelistStatusInput{}
				input, err := PackSetTxWhitelistStatus(testInput)
				require.NoError(t, err)
				return input
			},
			SuppliedGas: SetTxWhitelistStatusGasCost,
			ReadOnly:    false,
			ExpectedErr: ErrCannotSetTxWhitelistStatus.Error(),
		},
		"calling setTxWhitelistStatus from Enabled should succeed": {
			Caller:     allowlist.TestEnabledAddr,
			BeforeHook: allowlist.SetDefaultRoles(Module.Address),
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// populate test input here
				testInput := SetTxWhitelistStatusInput{}
				input, err := PackSetTxWhitelistStatus(testInput)
				require.NoError(t, err)
				return input
			},
			// This test is for a successful call. You can set the expected output here.
			// CUSTOM CODE STARTS HERE
			ExpectedRes: func() []byte {
				// this function does not return an output, leave this one as is
				packedOutput := []byte{}
				return packedOutput
			}(),
			SuppliedGas: SetTxWhitelistStatusGasCost,
			ReadOnly:    false,
			ExpectedErr: "",
		},
		"calling setTxWhitelistStatus from Manager should succeed": {
			Caller:     allowlist.TestManagerAddr,
			BeforeHook: allowlist.SetDefaultRoles(Module.Address),
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// populate test input here
				testInput := SetTxWhitelistStatusInput{}
				input, err := PackSetTxWhitelistStatus(testInput)
				require.NoError(t, err)
				return input
			},
			// This test is for a successful call. You can set the expected output here.
			// CUSTOM CODE STARTS HERE
			ExpectedRes: func() []byte {
				// this function does not return an output, leave this one as is
				packedOutput := []byte{}
				return packedOutput
			}(),
			SuppliedGas: SetTxWhitelistStatusGasCost,
			ReadOnly:    false,
			ExpectedErr: "",
		},
		"calling setTxWhitelistStatus from Admin should succeed": {
			Caller:     allowlist.TestAdminAddr,
			BeforeHook: allowlist.SetDefaultRoles(Module.Address),
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// populate test input here
				testInput := SetTxWhitelistStatusInput{}
				input, err := PackSetTxWhitelistStatus(testInput)
				require.NoError(t, err)
				return input
			},
			// This test is for a successful call. You can set the expected output here.
			// CUSTOM CODE STARTS HERE
			ExpectedRes: func() []byte {
				// this function does not return an output, leave this one as is
				packedOutput := []byte{}
				return packedOutput
			}(),
			SuppliedGas: SetTxWhitelistStatusGasCost,
			ReadOnly:    false,
			ExpectedErr: "",
		},
		"readOnly setTxWhitelistStatus should fail": {
			Caller: common.Address{1},
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// populate test input here
				testInput := SetTxWhitelistStatusInput{}
				input, err := PackSetTxWhitelistStatus(testInput)
				require.NoError(t, err)
				return input
			},
			SuppliedGas: SetTxWhitelistStatusGasCost,
			ReadOnly:    true,
			ExpectedErr: vmerrs.ErrWriteProtection.Error(),
		},
		"insufficient gas for setTxWhitelistStatus should fail": {
			Caller: common.Address{1},
			InputFn: func(t testing.TB) []byte {
				// CUSTOM CODE STARTS HERE
				// populate test input here
				testInput := SetTxWhitelistStatusInput{}
				input, err := PackSetTxWhitelistStatus(testInput)
				require.NoError(t, err)
				return input
			},
			SuppliedGas: SetTxWhitelistStatusGasCost - 1,
			ReadOnly:    false,
			ExpectedErr: vmerrs.ErrOutOfGas.Error(),
		},
	}
)

// TestTxWhitelistRun tests the Run function of the precompile contract.
func TestTxWhitelistRun(t *testing.T) {
	// Run tests with allowlist tests.
	// This adds allowlist tests to your custom tests
	// and runs them all together.
	// Even if you don't add any custom tests, keep this. This will still
	// run the default allowlist tests.
	allowlist.RunPrecompileWithAllowListTests(t, Module, state.NewTestStateDB, tests)
}

// TestPackUnpackTxWhitelistChangedEventData tests the Pack/UnpackTxWhitelistChangedEventData.
func TestPackUnpackTxWhitelistChangedEventData(t *testing.T) {
	// CUSTOM CODE STARTS HERE
	// set test inputs with proper values here
	var contractAddress common.Address = common.Address{}
	var methodInput []byte = []byte{}

	dataInput := TxWhitelistChangedEventData{
		Status: false,
	}

	_, data, err := PackTxWhitelistChangedEvent(
		contractAddress,
		methodInput,
		dataInput,
	)
	require.NoError(t, err)

	unpacked, err := UnpackTxWhitelistChangedEventData(data)
	require.NoError(t, err)
	require.Equal(t, dataInput, unpacked)
}

func BenchmarkTxWhitelist(b *testing.B) {
	// Benchmark tests with allowlist tests.
	// This adds allowlist tests to your custom tests
	// and benchmarks them all together.
	// Even if you don't add any custom tests, keep this. This will still
	// run the default allowlist tests.
	allowlist.BenchPrecompileWithAllowList(b, Module, state.NewTestStateDB, tests)
}
