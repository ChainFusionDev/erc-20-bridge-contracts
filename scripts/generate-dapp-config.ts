import { createDAppConfig, createNativeChainConfig, createNativeContractsConfig, createChainConfig, writeDAppConfig, ChainConfig, createTokenConfigs } from './deploy/dapp-config';
import { readContractsConfig, readChainContractsConfigByPath } from './deploy/config';
import glob from 'glob';

async function main() {
  const files = await getContractsFiles();

  const contractsConfig = await readContractsConfig();

  const nativeChain = await createNativeChainConfig();
  const nativeContracts = await createNativeContractsConfig(contractsConfig);

  var chains: Array<ChainConfig> = [];

  for (const path of files) {
    const contractsConfig = await readChainContractsConfigByPath(path);

    const chain = await createChainConfig(contractsConfig);
    chains.push(chain);

  }

  const tokens = await createTokenConfigs(files);
  const dAppConfig = await createDAppConfig(nativeChain, nativeContracts, chains, tokens)

  await writeDAppConfig(dAppConfig);
}

async function getContractsFiles(): Promise<string[]> {
  return new Promise((resolve, reject) => {
    glob('contracts-*.json', async function (err, files) {
      if (err) {
        reject(err);
      } else {
        resolve(files);
      }
    });
  })
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
