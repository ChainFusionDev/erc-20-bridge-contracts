import { deploySystemContracts } from './deploy/chainfusion';
import { readContractsConfig, updateContractsConfig, writeContractsConfig } from './deploy/config';

async function main() {
  const verify = (process.env.VERIFY || '').trim().toLowerCase() === 'true';

  const contractsConfig = await readContractsConfig();
  const bridgeAppFactoryAddress = contractsConfig.bridgeAppFactory ?? process.env.BRIDGE_APP_FACTORY_ADDRESS;

  const res = await deploySystemContracts({ displayLogs: true, parallelDeployment: false, verify, bridgeAppFactoryAddress });
  updateContractsConfig(contractsConfig, res);
  await writeContractsConfig(contractsConfig);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
