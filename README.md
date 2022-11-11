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

In This example we are deploying to `goerli` testnet. To deploy to different chain, `--network` parameter should be changed to `ropsten`, etc.

Deployment involves running three scripts in sequence. The first script is deployed once and used to create the BridgeApp allows to register an app and set its addresses for each chain.

```
$ npx hardhat --network goerli deploy scripts/deploy-system.ts
```
The second script is deployed each time for a new chain, deploys the main contracts that are responsible for the bridge logic and sets the BridgeApp address in the ERC20Bridge contract.

```
$ npx hardhat --network goerli deploy scripts/deploy-chain.ts
```

The third script is also deployed every time for a new chain, sets the Chain ID and address of the deployed ERC20Bridge in the BridgeApp contract, and completes the registration of the application in the new chain.

```
$ npx hardhat --network goerli deploy scripts/set-app.ts 
```
