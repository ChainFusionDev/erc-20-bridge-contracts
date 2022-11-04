import { deployBridgeContracts } from './deploy/chain';

async function main() {
  const verify = (process.env.VERIFY || '').trim().toLowerCase() === 'true';
  const foundationAddress = process.env.FOUNDATION_ADDRESS;
  const bridgeAppAddress = process.env.BRIDGE_APP_ADDRESS;
  const relayBridge = process.env.RELAY_BRIDGE_ADDRESS;

  await deployBridgeContracts({ displayLogs: true, verify: verify, relayBridge: relayBridge, foundationAddress: foundationAddress, bridgeAppAddress: bridgeAppAddress });
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
