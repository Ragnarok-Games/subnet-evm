// (c) 2022-2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: MIT

pragma solidity >=0.8.0;
import "./IAllowList.sol";

interface ITxWhitelistManager is IAllowList {
    event TxWhitelistChanged(address indexed contractAddress, bytes indexed method, bool status);

  // sayHello returns the stored greeting string
    function getTxWhitelistStatus(address contractAddress, bytes memory method) external view returns (bool status);

  // setGreeting  stores the greeting string
    function setTxWhitelistStatus(address contractAddress, bytes memory method, bool status) external;
}
