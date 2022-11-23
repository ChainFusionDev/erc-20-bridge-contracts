import { BigNumber } from 'ethers';
import { network } from 'hardhat';
import { deployBridgeContracts } from './deploy/chain';
import { readChainContractsConfig, readContractsConfig, updateContractsConfig, writeChainContractsConfig } from './deploy/config';

async function main() {
  const verify = (process.env.VERIFY || '').trim().toLowerCase() === 'true';
  const chainId = network.config.chainId ?? 1;
  const privateKey = process.env.PRIVATE_KEY;

  const homeContractsConfig = await readContractsConfig();
  const homeNetwork = homeContractsConfig.networkName;
  const bridgeApp = homeContractsConfig.bridgeApp ?? process.env.BRIDGE_APP_ADDRESS;

  const contractsConfig = await readChainContractsConfig(chainId)
  const foundationAddress = contractsConfig.foundationAddress ?? process.env.FOUNDATION_ADDRESS;
  const relayBridge = contractsConfig.relayBridge ?? process.env.RELAY_BRIDGE_ADDRESS;
  const signerStorage = contractsConfig.signerStorage ?? process.env.SIGNER_STORAGE;
  const homeChainId = BigNumber.from(process.env.HOME_CHAIN_ID ?? '1');

  const res = await deployBridgeContracts({
    displayLogs: true,
    setApp: true,
    privateKey,
    homeNetwork,
    verify,
    signerStorage,
    relayBridge,
    foundationAddress,
    bridgeApp,
    homeChainId,
  });

  updateContractsConfig(contractsConfig, res);
  await writeChainContractsConfig(chainId, contractsConfig);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
