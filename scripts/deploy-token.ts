import { network } from 'hardhat';
import { deployToken, setTokenToTokenManager, setTokenFee, addTokenToERC20BridgeMediator, setLimitPerToken, TokenDeploymentOptions, addLiquidity } from './deploy/token';
import {
  readChainContractsConfig,
  readContractsConfig,
  addTokenToContractsConfig,
  writeChainContractsConfig,
} from './deploy/config';

async function main() {
    const tokenName = (process.env.TOKEN_NAME || '');
    const tokenSymbol = (process.env.TOKEN_SYMBOL || '');
    const privateKey = (process.env.PRIVATE_KEY || '');

    const chainId = network.config.chainId ?? 1;

    const contractsChainConfig = await readChainContractsConfig(chainId);
    const contractsConfig = await readContractsConfig();

    var options: TokenDeploymentOptions = {
      privateKey: privateKey, 
      symbol: tokenSymbol, 
      name: tokenName, 
      homeNetwork: contractsConfig.networkName,  
      tokenManagerAddress: contractsChainConfig.tokenManager, 
      feeManagerAddress: contractsChainConfig.feeManager, 
      bridgeValidatorFeePoolAddress: contractsChainConfig.bridgeValidatorFeePool, 
      erc20BridgeMediatorAddress: contractsConfig.erc20BridgeMediator,
      liquidityPoolsAddress: contractsChainConfig.liquidityPools,
    }

    const token = await deployToken(options);

    addTokenToContractsConfig(contractsChainConfig, tokenSymbol, token.address);
    await writeChainContractsConfig(chainId, contractsChainConfig);

    await setTokenToTokenManager(token.address, options);
    await setTokenFee(token.address, options);
    await setLimitPerToken(token.address, options);
    await addTokenToERC20BridgeMediator(token.address, options);
    await addLiquidity(token, options);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
