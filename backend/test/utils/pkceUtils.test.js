const {
  base64URLEncode,
  generateCodeVerifier,
  generateCodeChallenge
} = require('../../src/utils/pkceUtils'); // Replace with the correct file path

const crypto = require('./cryptoMock'); // Import the mock

jest.mock('crypto', () => require('./cryptoMock')); // Mock the crypto module

describe('base64URLEncode', () => {
  test('encodes a string to base64 URL format', () => {
    const url = 'https://elearn.smu.edu.sg/d2l/le/content/360462/viewContent/2272113/View';
    const buffer = Buffer.from(url);
    const result = base64URLEncode(buffer);
    expect(result).toBe('aHR0cHM6Ly9lbGVhcm4uc211LmVkdS5zZy9kMmwvbGUvY29udGVudC8zNjA0NjIvdmlld0NvbnRlbnQvMjI3MjExMy9WaWV3');
  });
});

describe('generateCodeVerifier', () => {
  test('generates a code verifier', () => {
    const result = generateCodeVerifier();
    expect(result).toBe('bW9ja2VkX3JhbmRvbV9ieXRlcw');
  });
});

describe('generateCodeChallenge', () => {
  test('generates a code challenge', () => {
    const codeVerifier = 'bW9ja2VkX3JhbmRvbV9ieXRlcw';
    const result = generateCodeChallenge(codeVerifier);
    expect(result).toBe('bW9ja2VkX2hhc2hfY29tcGxleA');
  });
});
