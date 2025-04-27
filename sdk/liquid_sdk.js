// liquid_sdk.js
// LiquidAI - Developer SDK
// Provides a simple interface to interact with LiquidAI backend services.

const axios = require('axios');

class LiquidSDK {
  /**
   * Initializes the SDK
   * @param {string} apiUrl - LiquidAI API endpoint
   */
  constructor(apiUrl) {
    if (!apiUrl) {
      throw new Error('API URL is required to initialize LiquidSDK');
    }
    this.apiUrl = apiUrl.endsWith('/') ? apiUrl.slice(0, -1) : apiUrl;
  }

  /**
   * Query wallet balance
   * @param {string} walletAddress - Wallet address to query
   * @returns {Promise<Object>}
   */
  async queryBalance(walletAddress) {
    try {
      const res = await axios.get(`${this.apiUrl}/balance`, {
        params: { address: walletAddress },
      });
      return res.data;
    } catch (err) {
      throw new Error(`Failed to query balance: ${err.message}`);
    }
  }

  /**
   * Request a market prediction
   * @param {string} walletAddress - Wallet address for prediction
   * @returns {Promise<Object>}
   */
  async predictMarket(walletAddress) {
    try {
      const res = await axios.get(`${this.apiUrl}/predict`, {
        params: { address: walletAddress },
      });
      return res.data;
    } catch (err) {
      throw new Error(`Failed to predict market: ${err.message}`);
    }
  }

  /**
   * Register a new agent
   * @param {string} agentName - Name of the agent
   * @returns {Promise<Object>}
   */
  async registerAgent(agentName) {
    try {
      const res = await axios.post(`${this.apiUrl}/agent/register`, {
        name: agentName,
      });
      return res.data;
    } catch (err) {
      throw new Error(`Failed to register agent: ${err.message}`);
    }
  }
}

module.exports = LiquidSDK;
