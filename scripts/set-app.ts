import { setApp } from './deploy/set-app';
import { readChainContractsConfig, readContractsConfig } from './deploy/config';

async function main() {
    const chainId = parseInt(process.env.CHAIN_ID ?? '1');

    const contractsChainConfig = await readChainContractsConfig(chainId);
    const erc20Bridge = contractsChainConfig.erc20Bridge ?? process.env.ERC20BRIDGE_ADDRESS;

    const contractsConfig = await readContractsConfig();
    const bridgeApp = contractsConfig.bridgeApp ?? process.env.BRIDGE_APP_ADDRESS;

    await setApp({ displayLogs: true, erc20Bridge, bridgeApp, chainId });
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
