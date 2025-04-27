// token_tracker.js
// LiquidAI - Solana Token Tracker
// Provides simple SPL Token balance and metadata lookup.

const { Connection, PublicKey } = require('@solana/web3.js');
const { TOKEN_PROGRAM_ID, getAccount, getMint } = require('@solana/spl-token');

class TokenTracker {
  constructor(endpoint = 'mainnet-beta') {
    this.connection = new Connection(
      endpoint.startsWith('http') ? endpoint : clusterApiUrl(endpoint),
      'confirmed'
    );
  }

  /**
   * Fetch SPL token balance of a specific token account
   * @param {string} tokenAccountAddress
   * @returns {Promise<number>}
   */
  async getTokenBalance(tokenAccountAddress) {
    try {
      const pubKey = new PublicKey(tokenAccountAddress);
      const accountInfo = await getAccount(this.connection, pubKey);
      const balance = Number(accountInfo.amount);
      return balance;
    } catch (err) {
      throw new Error(`Failed to get token balance: ${err.message}`);
    }
  }

  /**
   * Fetch SPL token metadata (like decimals, supply)
   * @param {string} mintAddress
   * @returns {Promise<object>}
   */
  async getTokenMetadata(mintAddress) {
    try {
      const mintPubKey = new PublicKey(mintAddress);
      const mintInfo = await getMint(this.connection, mintPubKey);

      return {
        supply: mintInfo.supply.toString(),
        decimals: mintInfo.decimals,
        mintAuthority: mintInfo.mintAuthority.toBase58(),
      };
    } catch (err) {
      throw new Error(`Failed to get token metadata: ${err.message}`);
    }
  }
}

module.exports = TokenTracker;

