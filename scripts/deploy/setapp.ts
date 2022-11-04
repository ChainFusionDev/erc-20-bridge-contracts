import { ethers } from 'hardhat';
import { Deployer } from './deployer';

import {} from '../../typechain';

const defaultSystemDeploymentParameters: SystemDeploymentParameters = {
  erc20Bridge: '0x5b7eFCb3ebbde03625A92C67a87c5a58F046e64f',
  bridgeApp: '0x03080d4C8340c9b23D2fe87169B1Df8EccFE4230',
  chainId: '953842',

  displayLogs: false,
  verify: false,
};

export async function deploySystemContracts(options?: SystemDeploymentOptions): Promise<SystemDeploymentResult> {
  const params = resolveParameters(options);
  const deployer = new Deployer(params.displayLogs);

  deployer.log('Deploying contracts\n');

  const res: SystemDeployment = {};

  const bridgeApp = await ethers.getContractAt('IBridgeApp', params.bridgeApp);
  await bridgeApp.setContractAddress(params.chainId, params.erc20Bridge);

  deployer.log('Successfully deployed contracts\n');

  deployer.log('Initializing contracts\n');

  deployer.log('Successfully initialized contracts\n');

  if (params.verify) {
    await deployer.verifyObjectValues(res);
  }

  return {
    ...res,
    ...params,
  };
}

function resolveParameters(options?: SystemDeploymentOptions): SystemDeploymentParameters {
  let parameters = defaultSystemDeploymentParameters;

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

  return parameters;
}

export interface SystemDeploymentResult extends SystemDeployment, SystemDeploymentParameters {}

export interface SystemDeployment {}

export interface SystemDeploymentParameters {
  erc20Bridge: string;
  chainId: string;
  bridgeApp: string;

  displayLogs: boolean;
  verify: boolean;
}

export interface SystemDeploymentOptions {
  erc20Bridge?: string;
  chainId?: string;
  bridgeApp?: string;

  displayLogs?: boolean;
  verify?: boolean;
}
