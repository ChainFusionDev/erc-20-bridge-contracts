import { deploySystemContracts } from './deploy/system';
import { readContractsConfig } from './deploy/config';

async function main() {
  const verify = (process.env.VERIFY || '').trim().toLowerCase() === 'true';

  const contractsConfig = await readContractsConfig();
  const bridgeAppFactory = contractsConfig.bridgeAppFactory ?? process.env.BRIDGE_APP_FACTORY;

  await deploySystemContracts({ displayLogs: true, verify, bridgeAppFactory: bridgeAppFactory });
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
