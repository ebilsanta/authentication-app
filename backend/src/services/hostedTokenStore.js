class HostedTokenStore {
  instance = null;

  constructor() {
    if (!HostedTokenStore.instance) {
      HostedTokenStore.instance = this;
      this.tokens = {};
    }
    return HostedTokenStore.instance;
    
  }

  // Set tokens
  setTokens(sessionId, accessToken, refreshToken, dpopToken) {
    this.tokens[sessionId] = {
      accessToken,
      refreshToken,
      dpopToken,
    };
  }

  // Get tokens
  getTokens(sessionId) {
    return this.tokens[sessionId];
  }

  // Remove tokens
  removeTokens(sessionId) {
    delete this.tokens[sessionId];
  }

  // Check if a session ID exists
  hasSession(sessionId) {
    return this.tokens.hasOwnProperty(sessionId);
  }
}

module.exports = HostedTokenStore;
