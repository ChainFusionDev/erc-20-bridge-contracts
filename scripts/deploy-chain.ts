import { BigNumber } from 'ethers';
import { network } from 'hardhat';
import { deployBridgeContracts } from './deploy/chain';
import { readChainContractsConfig, updateContractsConfig, writeChainContractsConfig } from './deploy/config';

async function main() {
  const verify = (process.env.VERIFY || '').trim().toLowerCase() === 'true';
  const chainId = network.config.chainId ?? 1;

  const contractsConfig = await readChainContractsConfig(chainId)
  const foundationAddress = contractsConfig.foundationAddress ?? process.env.FOUNDATION_ADDRESS;
  const bridgeAppAddress = contractsConfig.bridgeAppAddress ?? process.env.BRIDGE_APP_ADDRESS;
  const relayBridge = contractsConfig.relayBridge ?? process.env.RELAY_BRIDGE_ADDRESS;
  const signerStorage = contractsConfig.signerStorage ?? process.env.SIGNER_STORAGE;
  const homeChainId = BigNumber.from(contractsConfig.homeChainId ?? process.env.HOME_CHAIN_ID);

  const res = await deployBridgeContracts({ displayLogs: true, verify, signerStorage, relayBridge, foundationAddress, bridgeAppAddress, homeChainId });

  updateContractsConfig(contractsConfig, res);
  await writeChainContractsConfig(chainId, contractsConfig);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
