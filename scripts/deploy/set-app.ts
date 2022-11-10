import { ethers } from 'hardhat';
import { Deployer } from './deployer';

const defaultSetAppParameters: SetAppParameters = {
  erc20Bridge: '0x0000000000000000000000000000000000000001',
  bridgeApp: '0x0000000000000000000000000000000000000001',
  chainId: 953842,

  displayLogs: false,
};

export async function setApp(options?: SetAppOptions): Promise<void> {
  const params = resolveParameters(options);
  const deployer = new Deployer(params.displayLogs);

  deployer.log('Setting contract address in bridge app\n');

  deployer.log(`ChainId: ${params.chainId} Address: ${params.erc20Bridge}\n`);

  const bridgeApp = await ethers.getContractAt('IBridgeApp', params.bridgeApp);
  await bridgeApp.setContractAddress(params.chainId, params.erc20Bridge);

  deployer.log('Successfully set contract address in bridge app\n');

  return;
}

function resolveParameters(options?: SetAppOptions): SetAppParameters {
  let parameters = defaultSetAppParameters;

  if (options === undefined) {
    return parameters;
  }

  if (options.erc20Bridge !== undefined) {
    parameters.erc20Bridge = options.erc20Bridge;
  }

  if (options.chainId !== undefined) {
    parameters.chainId = options.chainId;
  }

  if (options.bridgeApp !== undefined) {
    parameters.bridgeApp = options.bridgeApp;
  }

  if (options.displayLogs !== undefined) {
    parameters.displayLogs = options.displayLogs;
  }

  return parameters;
}

export interface SetAppParameters {
  erc20Bridge: string;
  chainId: number;
  bridgeApp: string;

  displayLogs: boolean;
}

export interface SetAppOptions {
  erc20Bridge?: string;
  chainId?: number;
  bridgeApp?: string;

  displayLogs?: boolean;
}
