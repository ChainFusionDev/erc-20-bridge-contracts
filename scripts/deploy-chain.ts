import { deployBridgeContracts } from './deploy/chain';

async function main() {
  const verify = (process.env.VERIFY || '').trim().toLowerCase() === 'true';
  await deployBridgeContracts({ displayLogs: true, verify: verify });
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
