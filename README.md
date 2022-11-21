# ChainFusion Contracts

This repository contains Solidity contracts with bridge logic

## Development

### Install Packages

```
$ npm install
```

### Compile Contracts

```
$ npm run compile
```

### Run Tests

```
$ npm run test
```

## Deployment

First of all, copy `.env.example` into `.env` and set up all required variables inside

### Deploy Contracts

In this example we are deploying to `ternopil` testnet. To deploy to different chain, `--network` parameter should be changed.

Deployment involves running two scripts in sequence. The first script is deployed once, on ChainFusion network and used to create the BridgeApp that allows to register an app and set its addresses for each chain.

```
$ npx hardhat --network ternopil run scripts/deploy-chainfusion.ts
```
The second script is deployed each time for a new chain, deploys the main contracts that are responsible for the bridge logic and sets the BridgeApp address in the ERC20Bridge contract.

```
$ npx hardhat --network bg1 run scripts/deploy-chain.ts
```

The third script is also deployed every time for a new chain, sets the Chain ID and address of the deployed ERC20Bridge in the BridgeApp contract, and completes the registration of the application in the new chain.

```
$ CHAIN_ID=123 npx hardhat --network bg1 run scripts/set-app.ts
```

### Verify Contracts

To verify contracts, you need to specify network, contract address and constructor parameters (if present).

```
$ npx hardhat --network goerli verify <CONTRACT_ADDRESS> <CONSTRUCTOR_PARAMETERS>
```

Or you can set `VERIFY` variable to `true` while deploying contracts to automatically verify them afterwards.

```
$ VERIFY=true npx hardhat --network ternopil run scripts/deploy-chainfusion.ts
```
