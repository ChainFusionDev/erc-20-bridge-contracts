import { deploySystemContracts } from './deploy/system';

async function main() {
  const verify = (process.env.VERIFY || '').trim().toLowerCase() === 'true';
  const stakingKeys = !process.env.STAKING_KEYS ? [] : (process.env.STAKING_KEYS).trim().split(',');
  const bridgeAppFactory = process.env.BRIDGE_APP_FACTORY;

  await deploySystemContracts({ displayLogs: true, verify, stakingKeys, bridgeAppFactory: bridgeAppFactory });
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
