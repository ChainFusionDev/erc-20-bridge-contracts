import { Signer } from 'ethers';
import { promises as fs } from 'fs';
import { config, ethers } from 'hardhat';
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

const nativeChainName = 'ChainFusion';

export async function createNativeChainConfig(contractsConfig: ContractsConfig): Promise<NativeChainConfig> {
  return new Promise((resolve, reject) => {
    if (contractsConfig.networkName === undefined) {
      reject();
    } else {
      for (const chain of config.etherscan.customChains) {
        if (chain.network != contractsConfig.networkName) {
          continue;
        }

        const networkConfig = config.networks[contractsConfig.networkName] as HttpNetworkConfig;

        var nativeChain: NativeChainConfig = {
          chainId: chain.chainId,
          identifier: nativeChainName.toLowerCase(),
          name: nativeChainName,
          rpc: networkConfig.url,
          explorer: chain.urls.browserURL,
          nativeCurrency: getNativeCurrency(),
        };

        resolve(nativeChain);
      }
    }
  });
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
        reject();
      }

      var explorer = '';

      for (const chain of config.etherscan.customChains) {
        if (chain.network != contractChainConfig.networkName) {
          continue;
        }

        explorer = chain.urls.browserURL;
      }

      const chainId = networkConfig.chainId ?? 1;
      const name = contractChainConfig.networkName.charAt(0).toUpperCase() + contractChainConfig.networkName.slice(1);

      resolve({
        chainId: chainId,
        identifier: `${contractChainConfig.networkName}`.toLowerCase(),
        name: name,
        rpc: networkConfig.url,
        explorer: explorer,
        nativeCurrency: getNativeCurrency(),
        erc20BridgeAddress: contractChainConfig.erc20Bridge,
      });
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
      if (config.identifier == tokenConfig.identifier && contractChainConfig.networkName !== undefined) {
        config.chains[contractChainConfig.networkName.toLowerCase()] = address;
      }
    }

    return;
  } catch (error) {
    return;
  }
}

async function createTokenConfig(
  contractChainConfig: ContractsConfig,
  address: string,
  signer: Signer
): Promise<TokenConfig> {
  try {
    var chains: Chains = {};

    if (contractChainConfig.networkName === undefined) {
      throw new Error('Failed to create token config');
    }

    const tokenFactory = await ethers.getContractFactory('Token', signer);
    const tokenContract = tokenFactory.attach(address);

    const symbol = await tokenContract.symbol();
    const name = await tokenContract.name();
    const decimals = await tokenContract.decimals();

    chains[contractChainConfig.networkName.toLowerCase()] = address;

    return {
      identifier: symbol.toLowerCase(),
      name: name,
      symbol: symbol,
      decimals: decimals,
      chains: chains,
    };
  } catch (error) {
    throw new Error('Failed to create token config');
  }
}

export async function writeDAppConfig(dAppConfig: DAppConfig): Promise<void> {
  try {
    await fs.writeFile(`config.json`, JSON.stringify(dAppConfig, null, 2));
  } catch (error) {
    console.error(error);
  }
}

function getNativeCurrency(): NativeCurrency {
  return { name: nativeChainName, symbol: 'CFN', decimals: 18 };
}

function contains(tokenConfigs: TokenConfig[], tokenConfig: TokenConfig): boolean {
  for (const config of tokenConfigs) {
    if (config.identifier == tokenConfig.identifier) {
      return true;
    }
  }

  return false;
}
