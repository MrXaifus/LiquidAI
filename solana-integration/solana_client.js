// solana_client.js
// LiquidAI - Solana Blockchain Client
// Provides Solana RPC connection and basic account queries.

const { Connection, PublicKey, clusterApiUrl } = require('@solana/web3.js');

class SolanaClient {
  /**
   * Initialize Solana client
   * @param {string} endpoint - Custom RPC endpoint or cluster name (e.g., 'mainnet-beta')
   */
  constructor(endpoint = 'mainnet-beta') {
    this.connection = new Connection(
      endpoint.startsWith('http') ? endpoint : clusterApiUrl(endpoint),
      'confirmed'
    );
  }

  /**
   * Get SOL balance of a wallet
   * @param {string} walletAddress - Base58 encoded public key
   * @returns {Promise<number>} balance in SOL
   */
  async getBalance(walletAddress) {
    try {
      const pubKey = new PublicKey(walletAddress);
      const lamports = await this.connection.getBalance(pubKey);
      return lamports / 1e9; // Convert lamports to SOL
    } catch (err) {
      throw new Error(`Failed to get balance: ${err.message}`);
    }
  }

  /**
   * Fetch recent transaction signatures for a wallet
   * @param {string} walletAddress
   * @param {number} limit - How many transactions to fetch
   * @returns {Promise<Array>}
   */
  async getRecentTransactions(walletAddress, limit = 5) {
    try {
      const pubKey = new PublicKey(walletAddress);
      const signatures = await this.connection.getSignaturesForAddress(pubKey, { limit });
      return signatures;
    } catch (err) {
      throw new Error(`Failed to fetch transactions: ${err.message}`);
    }
  }
}

module.exports = SolanaClient;
