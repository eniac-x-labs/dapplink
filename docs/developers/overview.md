<!--
order: 1
-->

# Developers

An overview of the available documentation for developers. {synopsis}

## List of Resources

Please find the following resources for in-depth information:

- **[Quick Connect](./connect.md)**: List of publicly available endpoints.
- **[dapplink Clients](./clients.md)**: Description of available clients.
- **[Block Explorers](./explorers.md)**: List of block explorers available for dapplink.
- **[Testnet Faucet](./faucet.md)**: Explaination of faucet use to obtain testnet tokens.
- **Localnet**: Instructions on how to configure a local instance of the dapplink blockchain.
    - **[Single Node](./localnet/single_node.md)**: Run a single local node.
    - **[Multi Node](./localnet/multi_node.md)**: Run a local testnet with multiple nodes.
    - **[Testnet](./localnet/testnet_cmd.md)**: Use the testnet command of the dapplink daemon.
- **Libraries**:
    - **[dapplinkJS](./libraries/dapplinkjs.md)**: Javascript library for dapplink.

### Remote Procedure Calls (RPCs)

As dapplink lives at the crossroads of Cosmos and Ethereum, there are RPC connections available for all corresponding interfaces:

- **[JSON-RPC Server](./json-rpc/server.md)**: General information about the JSON-RPC server provided by dapplink.
- **[Running The Server](./json-rpc/running_server.md)**: Instructions on how to set up the server when running a node.
- **[Namespaces](./json-rpc/namespaces.md)**: Description of the available JSON-RPC namespaces.
- **[JSON-RPC Methods](./json-rpc/endpoints.md)**: List of supported JSON-RPC endpoints and methods.
- **[Events](./json-rpc/events.md)**: Information about the available events and instructions to subscribe to them.
- **[Cosmos gRPC & REST](https://api.dapplink.dev/)**: Documentation of the available gRPC implementation on dapplink.
- **[Tendermint RPC](https://docs.tendermint.com/v0.34/rpc/)**: Documentation for the RPC protocols supported by Tendermint.

### Tutorials For Ethereum Devs

To welcome Ethereum developers, there is a selection of tutorials available showcasing some widely used developer tools:

- **[Remix](./tools/remix.md)**: Instructions to deploy a smart contract on dapplink with [Remix IDE](http://remix.ethereum.org/).
- **[Hardhat](./tools/hardhat.md)**: Instructions to deploy a smart contract on dapplink with the [Hardhat](https://hardhat.org/) suite of development tools.
- **[Truffle](./tools/truffle.md)**: Instructions to deploy a smart contract on dapplink using the [Truffle](https://www.trufflesuite.com/truffle) framework.
