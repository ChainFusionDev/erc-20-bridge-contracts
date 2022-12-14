import { Signer } from 'ethers';
import { promises as fs } from 'fs';
import { config, ethers, network } from 'hardhat';
import { HttpNetworkConfig } from 'hardhat/types';
import { ContractsConfig, readChainContractsConfigByPath } from './config';

export interface DAppConfig {
  nativeChain: NativeChainConfig | undefined;
  nativeContracts: NativeContractsConfig | undefined;
  chains: Array<ChainConfig> | undefined;
  tokens: Array<TokenConfig> | undefined;
}

export interface NativeChainConfig {
  chainId: number;
  identifier: string;
  name: string;
  rpc: string;
  explorer: string;
  nativeCurrency: NativeCurrency;
}

export interface NativeContractsConfig {
  staking: string;
  eventRegistry: string;
}

export interface ChainConfig {
  chainId: number;
  identifier: string;
  name: string;
  rpc: string;
  explorer: string;
  nativeCurrency: NativeCurrency;
  erc20BridgeAddress: string;
}

export interface TokenConfig {
  identifier: string;
  name: string;
  symbol: string;
  decimals: number;
  chains: Chains;
}

export interface NativeCurrency {
  name: string;
  symbol: string;
  decimals: number;
}

export interface Chains {
  [key: string]: string | undefined;
}

export async function createDAppConfig(
  nativeChain: NativeChainConfig,
  nativeContracts: NativeContractsConfig,
  chains: Array<ChainConfig>,
  tokens: Array<TokenConfig>
): Promise<DAppConfig> {
  return { nativeChain, nativeContracts, chains, tokens };
}

export async function createNativeChainConfig(): Promise<NativeChainConfig> {
  const networkConfig = config.networks[network.name] as HttpNetworkConfig;

  const chainId = networkConfig.chainId ?? 1;
  const url = networkConfig.url;

  var nativeChain: NativeChainConfig = {
    chainId: chainId,
    identifier: network.name.toLowerCase(),
    name: network.name,
    rpc: url,
    explorer: 'test',
    nativeCurrency: { name: 'test', symbol: 'test', decimals: 1 },
  };

  return nativeChain;
}

export async function createNativeContractsConfig(contractsConfig: ContractsConfig): Promise<NativeContractsConfig> {
  return new Promise((resolve, reject) => {
    if (contractsConfig.staking === undefined || contractsConfig.eventRegistry === undefined) {
      reject();
    } else {
      resolve({ staking: contractsConfig.staking, eventRegistry: contractsConfig.eventRegistry });
    }
  });
}

export async function createChainConfig(contractChainConfig: ContractsConfig): Promise<ChainConfig> {
  return new Promise((resolve, reject) => {
    if (contractChainConfig.networkName === undefined || contractChainConfig.erc20Bridge === undefined) {
      reject();
    } else {
      const networkConfig = config.networks[contractChainConfig.networkName] as HttpNetworkConfig;

      if (networkConfig === undefined) {
        return;
      }

      const chainId = networkConfig.chainId ?? 1;

      var chain: ChainConfig = {
        chainId: chainId,
        identifier: `chainfusion-${contractChainConfig.networkName}`.toLowerCase(),
        name: contractChainConfig.networkName,
        rpc: networkConfig.url,
        explorer: 'test',
        nativeCurrency: { name: 'test', symbol: 'test', decimals: 18 },
        erc20BridgeAddress: contractChainConfig.erc20Bridge,
      };

      resolve(chain);
    }
  });
}

export async function createTokenConfigs(files: string[]): Promise<TokenConfig[]> {
  var tokenConfigs: Array<TokenConfig> = [];

  for (const path of files) {
    const contractChainConfig = await readChainContractsConfigByPath(path);

    if (contractChainConfig.networkName === undefined) {
      continue;
    }

    const networkConfig = config.networks[contractChainConfig.networkName] as HttpNetworkConfig;

    const provider = new ethers.providers.JsonRpcProvider(networkConfig.url, networkConfig.chainId);
    const signer = provider.getSigner('0x0000000000000000000000000000000000000001');

    for (const [key, value] of Object.entries(contractChainConfig)) {
      var reg = new RegExp('token');

      if (!reg.test(key) || value === undefined) {
        continue;
      }

      await updateTokenConfigs(tokenConfigs, contractChainConfig, value, signer);
    }
  }

  return tokenConfigs;
}

async function updateTokenConfigs(
  tokenConfigs: TokenConfig[],
  contractChainConfig: ContractsConfig,
  address: string,
  signer: Signer
): Promise<void> {
  try {
    const tokenConfig = await createTokenConfig(contractChainConfig, address, signer);

    if (!contains(tokenConfigs, tokenConfig)) {
      tokenConfigs.push(tokenConfig);

      return;
    }

    for (const config of tokenConfigs) {
      if (config.identifier == tokenConfig.identifier) {
        config.chains[`chainfusion-${contractChainConfig.networkName}`.toLowerCase()] = address;
      }
    }

    return;
  } catch (error) {
    return;
  }
}

function createTokenConfig(
  contractChainConfig: ContractsConfig,
  address: string,
  signer: Signer
): Promise<TokenConfig> {
  return new Promise(async (resolve, reject) => {
    try {
      const tokenFactory = await ethers.getContractFactory('Token', signer);
      const tokenContract = tokenFactory.attach(address);

      const symbol = await tokenContract.symbol();
      const name = await tokenContract.name();
      const decimals = await tokenContract.decimals();

      var chains: Chains = {};
      chains[`chainfusion-${contractChainConfig.networkName}`.toLowerCase()] = address;

      var tokenConfig: TokenConfig = {
        identifier: symbol.toLowerCase(),
        name: name,
        symbol: symbol,
        decimals: decimals,
        chains: chains,
      };

      resolve(tokenConfig);
    } catch (error) {
      reject();
    }
  });
}

export async function writeDAppConfig(dAppConfig: DAppConfig): Promise<void> {
  try {
    await fs.writeFile(`config.json`, JSON.stringify(dAppConfig, null, 2));
  } catch (error) {
    console.error(error);
  }
}

function contains(tokenConfigs: TokenConfig[], tokenConfig: TokenConfig): boolean {
  for (const config of tokenConfigs) {
    if (config.identifier == tokenConfig.identifier) {
      return true;
    }
  }

  return false;
}
