import { ethers, network } from 'hardhat';
import { deployBridgeContracts } from './deploy/chain';
import { readChainContractsConfig } from './deploy/config';

async function main() {
  const verify = (process.env.VERIFY || '').trim().toLowerCase() === 'true';
  const chainId = network.config.chainId ?? 1;

  const contractsConfig = await readChainContractsConfig(chainId)
  const foundationAddress = contractsConfig.foundationAddress ?? process.env.FOUNDATION_ADDRESS;
  const bridgeAppAddress = contractsConfig.bridgeAppAddress ?? process.env.BRIDGE_APP_ADDRESS;
  const relayBridge = contractsConfig.relayBridge ?? process.env.RELAY_BRIDGE_ADDRESS;

  await deployBridgeContracts({ displayLogs: true, verify: verify, relayBridge: relayBridge, foundationAddress: foundationAddress, bridgeAppAddress: bridgeAppAddress });
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
