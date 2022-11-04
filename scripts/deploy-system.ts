import { deploySystemContracts } from './deploy/system';

async function main() {
  const verify = (process.env.VERIFY || '').trim().toLowerCase() === 'true';
  const bridgeAppFactory = process.env.BRIDGE_APP_FACTORY;

  await deploySystemContracts({ displayLogs: true, verify, bridgeAppFactory: bridgeAppFactory });
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
