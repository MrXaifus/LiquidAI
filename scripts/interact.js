// interact.js
// LiquidAI Smart Contract Interaction Script

const Web3 = require('web3');
const fs = require('fs');
require('dotenv').config();

// Load environment variables
const RPC_URL = process.env.RPC_URL;
const PRIVATE_KEY = process.env.PRIVATE_KEY;
const CONTRACT_ADDRESS = process.env.CONTRACT_ADDRESS;
const ABI_PATH = './LiquidAgent.json'; // Contract ABI JSON (compiled)

// Initialize Web3 provider
const web3 = new Web3(new Web3.providers.HttpProvider(RPC_URL));

// Load ABI
const abi = JSON.parse(fs.readFileSync(ABI_PATH));

// Load contract
const contract = new web3.eth.Contract(abi, CONTRACT_ADDRESS);

// Load account
const account = web3.eth.accounts.privateKeyToAccount(PRIVATE_KEY);
web3.eth.accounts.wallet.add(account);
web3.eth.defaultAccount = account.address;

// Utility function: send transaction
async function sendTx(method, value = '0') {
  const tx = {
    from: account.address,
    to: CONTRACT_ADDRESS,
    data: method.encodeABI(),
    gas: await method.estimateGas({ from: account.address }),
    gasPrice: await web3.eth.getGasPrice(),
    value: web3.utils.toWei(value, 'ether')
  };

  const receipt = await web3.eth.sendTransaction(tx);
  console.log(`✅ Tx Successful! Hash: ${receipt.transactionHash}`);
}

// Record AI Prediction to Smart Contract
async function recordPrediction(predictionCode) {
  console.log(`🚀 Recording AI prediction: ${predictionCode}`);
  const method = contract.methods.recordPrediction(predictionCode);
  await sendTx(method);
}

// Execute Trade from Smart Contract
async function executeTrade(targetAddress, amountEth) {
  console.log(`🚀 Executing trade: Sending ${amountEth} ETH to ${targetAddress}`);
  const method = contract.methods.executeTrade(targetAddress);
  await sendTx(method, amountEth);
}

// Main execution flow
(async () => {
  try {
    console.log('🔵 Starting interaction with LiquidAgent smart contract...');
    
    // Example: record a prediction (e.g., code 42 = bullish trend)
    await recordPrediction(42);

    // Example: execute a trade sending 0.05 ETH to a target address
    const targetAddress = '0xDEADBEEFDEADBEEFDEADBEEFDEADBEEFDEADBEEF'; // Replace with actual target
    await executeTrade(targetAddress, '0.05');

    console.log('✅ All operations completed successfully!');
  } catch (error) {
    console.error('❌ Error during interaction:', error);
  }
})();
