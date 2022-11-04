import { deploySystemContracts } from './deploy/setapp';

async function main() {
    const verify = (process.env.VERIFY || '').trim().toLowerCase() === 'true';
    const erc20Bridge = process.env.ERC20BRIDGE_ADDRESS;
    const bridgeApp = process.env.BRIDGE_APP_ADDRESS;
    const chainId = process.env.CHAIN_ID;

    await deploySystemContracts({ displayLogs: true, verify, erc20Bridge: erc20Bridge, bridgeApp: bridgeApp, chainId: chainId });
}

main().catch((error) => {
    console.error(error);
    process.exitCode = 1;
});
