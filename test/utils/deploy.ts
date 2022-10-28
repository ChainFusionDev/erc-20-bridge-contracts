import { ethers } from 'hardhat';
import { BigNumber } from 'ethers';
import { deploySystemContracts, SystemDeploymentOptions, SystemDeploymentResult } from '../../scripts/deploy/system';
import { BridgeDeploymentOptions, BridgeDeploymentResult, deployBridgeContracts } from '../../scripts/deploy/chain';
import { MockMintableBurnableToken, MockToken, ERC20Bridgeable, MockRelayBridge } from '../../typechain';

export async function deploySystem(options?: SystemDeploymentOptions): Promise<SystemDeploymentResult> {
  return await deploySystemContracts(options);
}

export async function deployBridge(options?: BridgeDeploymentOptions): Promise<BridgeDeploymentResult> {
  return await deployBridgeContracts(options);
}

export async function deployBridgeWithMocks(
  options?: BridgeDeploymentOptions
): Promise<BridgeWithMocksDeploymentResult> {
  const chainId = BigNumber.from(853);

  const MockRelayBridge = await ethers.getContractFactory('MockRelayBridge');
  const mockRelayBridge = await MockRelayBridge.deploy();
  await mockRelayBridge.deployed();

  if (options === undefined) {
    options = {};
  }

  if (options.relayBridge === undefined) {
    options.relayBridge = mockRelayBridge.address;
  }

  const deployment = await deployBridgeContracts(options);

  const MockToken = await ethers.getContractFactory('MockToken');
  const mockToken = await MockToken.deploy('Mock Token', 'MOCK', BigNumber.from('10000000000000000000000000'));
  await mockToken.deployed();

  await deployment.tokenManager.setToken(mockToken.address, 1);

  const MockMintableBurnableToken = await ethers.getContractFactory('MockMintableBurnableToken');
  const mockMintableBurnableToken = await MockMintableBurnableToken.deploy('Mintable Mock Token', 'MINT');
  await mockMintableBurnableToken.deployed();

  await deployment.tokenManager.setToken(mockMintableBurnableToken.address, 2);

  const ERC20Bridgeable = await ethers.getContractFactory('ERC20Bridgeable');
  const erc20Bridgeable = await ERC20Bridgeable.deploy('CFN Token', 'CFN');
  await erc20Bridgeable.deployed();

  return {
    mockChainId: chainId,
    mockToken: mockToken,
    mockMintableBurnableToken: mockMintableBurnableToken,
    erc20Bridgeable: erc20Bridgeable,
    mockRelayBridge: mockRelayBridge,
    ...deployment,
  };
}

export interface BridgeWithMocksDeploymentResult extends BridgeDeploymentResult {
  mockChainId: BigNumber;
  mockToken: MockToken;
  mockMintableBurnableToken: MockMintableBurnableToken;
  erc20Bridgeable: ERC20Bridgeable;
  mockRelayBridge: MockRelayBridge;
}
