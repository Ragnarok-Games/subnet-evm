// Code generated
// This file is a generated precompile contract config with stubbed abstract functions.
// The file is generated by a template. Please inspect every code and comment in this file before use.

package txwhitelist

import (
	"math/big"

	"github.com/ava-labs/subnet-evm/precompile/contract"
	"github.com/ethereum/go-ethereum/common"
)

// CUSTOM CODE STARTS HERE
// Reference imports to suppress errors from unused imports. This code and any unnecessary imports can be removed.
var (
	_ = big.NewInt
	_ = common.Big0
	_ = contract.LogGas
)

/* NOTE: Events can only be emitted in state-changing functions. So you cannot use events in read-only (view) functions.
Events are generally emitted at the end of a state-changing function with AddLog method of the StateDB. The AddLog method takes 4 arguments:
	1. Address of the contract that emitted the event.
	2. Topic hashes of the event.
	3. Encoded non-indexed data of the event.
	4. Block number at which the event was emitted.
The first argument is the address of the contract that emitted the event.
Topics can be at most 4 elements, the first topic is the hash of the event signature and the rest are the indexed event arguments. There can be at most 3 indexed arguments.
Topics cannot be fully unpacked into their original values since they're 32-bytes hashes.
The non-indexed arguments are encoded using the ABI encoding scheme. The non-indexed arguments can be unpacked into their original values.
Before packing the event, you need to calculate the gas cost of the event. The gas cost of an event is the base gas cost + the gas cost of the topics + the gas cost of the non-indexed data.
See Get{EvetName}EventGasCost functions for more details.
You can use the following code to emit an event in your state-changing precompile functions (generated packer might be different)):
topics, data, err := PackMyEvent(
	topic1,
	topic2,
	data1,
	data2,
)
if err != nil {
	return nil, remainingGas, err
}
accessibleState.GetStateDB().AddLog(
	ContractAddress,
	topics,
	data,
	accessibleState.GetBlockContext().Number().Uint64(),
)
*/

// TxWhitelistTxWhitelistChanged represents a TxWhitelistChanged non-indexed event data raised by the TxWhitelist contract.
type TxWhitelistChangedEventData struct {
	Status bool
}

// GetTxWhitelistChangedEventGasCost returns the gas cost of the event.
// The gas cost of an event is the base gas cost + the gas cost of the topics + the gas cost of the non-indexed data.
// The base gas cost and the gas cost of per topics are fixed and can be found in the contract package.
// The gas cost of the non-indexed data depends on the data type and the data size.
func GetTxWhitelistChangedEventGasCost(data TxWhitelistChangedEventData) uint64 {
	gas := contract.LogGas // base gas cost

	// Add topics gas cost (2 topics)
	// Topics always include the signature hash of the event. The rest are the indexed event arguments.
	gas += contract.LogTopicGas * 2

	// CUSTOM CODE STARTS HERE
	// TODO: calculate gas cost for packing the data.status according to the type.
	// Keep in mind that the data here will be encoded using the ABI encoding scheme.
	// So the computation cost might change according to the data type + data size and should be charged accordingly.
	// i.e gas += LogDataGas * uint64(len(data.status))
	gas += contract.LogDataGas // * ...
	// CUSTOM CODE ENDS HERE

	// CUSTOM CODE STARTS HERE
	// TODO: do any additional gas cost calculation here (only if needed)
	return gas
}

// PackTxWhitelistChangedEvent packs the event into the appropriate arguments for TxWhitelistChanged.
// It returns topic hashes and the encoded non-indexed data.
func PackTxWhitelistChangedEvent(contractAddress common.Address, method []byte, data TxWhitelistChangedEventData) ([]common.Hash, []byte, error) {
	return TxWhitelistABI.PackEvent("TxWhitelistChanged", contractAddress, method, data.Status)
}

// UnpackTxWhitelistChangedEventData attempts to unpack non-indexed [dataBytes].
func UnpackTxWhitelistChangedEventData(dataBytes []byte) (TxWhitelistChangedEventData, error) {
	eventData := TxWhitelistChangedEventData{}
	err := TxWhitelistABI.UnpackIntoInterface(&eventData, "TxWhitelistChanged", dataBytes)
	return eventData, err
}