class BankTokenStore {
  instance = null;

  constructor() {
    if (!BankTokenStore.instance) {
      BankTokenStore.instance = this;
      this.tokens = {};
    }
    return BankTokenStore.instance;
  }

  // Set tokens
  setTokens(sessionId, accessToken, idToken) {
    this.tokens[sessionId] = {
      accessToken,
      idToken,
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

module.exports = BankTokenStore;
