// Create a file named cryptoMock.js
const crypto = jest.createMockFromModule('node:crypto');

crypto.randomBytes = jest.fn().mockImplementation((size, callback) => {
  const buffer = Buffer.from('mocked_random_bytes');
  if (callback) {
    callback(null, buffer);
  }
  return buffer;
});

crypto.createHash = jest.fn().mockReturnValue({
  update: jest.fn().mockReturnThis(),
  digest: jest.fn().mockReturnValue(Buffer.from('mocked_hash')),
});

module.exports = crypto;
