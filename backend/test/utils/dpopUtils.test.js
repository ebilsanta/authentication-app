const { generateEphemeralKeys, generateDpop, generateJKTThumbprint, generateClientAssertion } = require("../../src/utils/dpopUtils");
const crypto = require('crypto');

describe('generateEphemeralKeys', () => {
  test('should return an object with public and private keys', async () => {
    const keys = await generateEphemeralKeys();

    expect(keys).toHaveProperty('publicKey');
    expect(keys).toHaveProperty('privateKey');
  });

  test('should return keys in the specified format', async () => {
    const keys = await generateEphemeralKeys();

    // Check if the public key is in PEM format
    expect(keys.publicKey).toContain('-----BEGIN PUBLIC KEY-----');
    expect(keys.publicKey).toContain('-----END PUBLIC KEY-----');

    // Check if the private key is in PEM format
    expect(keys.privateKey).toContain('-----BEGIN EC PRIVATE KEY-----');
    expect(keys.privateKey).toContain('-----END EC PRIVATE KEY-----');
  });
});

describe('generateJKTThumbprint', () => {
  test('should generate the correct thumbprint', async () => {
    // Example PEM-encoded public key
    const mockPublicKey = `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEaE5kZnlJ0Vc+kF8P0wAMuVNtzMQi
qBYx8KG3J0kKkOB8mOJiRAJjED5CIq3it7X5pbS2c1EbXjublw4nOPyyMQ==
-----END PUBLIC KEY-----
    `
    const expectedThumbprint = 'O-xVod82HsvtmDUyR3IMsjxV6htfXjoRazeZRmqgpns';
    
    const thumbprint = await generateJKTThumbprint(mockPublicKey);

    expect(thumbprint).toBe(expectedThumbprint);
  });
});

const ephemeralKeyPair =  {
  publicKey: '-----BEGIN PUBLIC KEY-----\n' +
    'MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEQpsaqsSBBcLIunsPRJ6wR22sDczh\n' +
    'VxgM4VhmQoiSJM6cW101oh7GfXldIEqQidUmnSMyb8AuJoGqiRiimOeTPg==\n' +
    '-----END PUBLIC KEY-----\n',
  privateKey: '-----BEGIN EC PRIVATE KEY-----\n' +
    'MHcCAQEEIJfC7C65wSA6grpU/3GSJiAiMZ9w8ZqxJUir5d6beDiioAoGCCqGSM49\n' +
    'AwEHoUQDQgAEQpsaqsSBBcLIunsPRJ6wR22sDczhVxgM4VhmQoiSJM6cW101oh7G\n' +
    'fXldIEqQidUmnSMyb8AuJoGqiRiimOeTPg==\n' +
    '-----END EC PRIVATE KEY-----\n'
}

describe('generateDpop', () => {
  test('should generate DPoP token with access token hash', async () => {
    const url = 'https://example.com';
    const ath = 'Eh9xM2CDQyrI3OMLJJJDhtmQeakdYYTDckNtklRQOVw';
    const method = 'POST';
    
    const dpop = await generateDpop(url, ath, method, ephemeralKeyPair);

    expect(typeof dpop).toBe('string');
  });

  test('should generate DPoP token without access token hash', async () => {
    const url = 'https://example.com';
    const ath = undefined; // No authentication header
    const method = 'GET';

    const dpop = await generateDpop(url, ath, method, ephemeralKeyPair);

    expect(typeof dpop).toBe('string');
  });
});

describe('generateClientAssertion', () => {  
  test('should generate client assertion with valid input', async () => {
    const url = 'https://example.com';
    const clientId = 'client123';
    const privateSigningKey = '-----BEGIN EC PRIVATE KEY-----\n' +
    'MHcCAQEEIJfC7C65wSA6grpU/3GSJiAiMZ9w8ZqxJUir5d6beDiioAoGCCqGSM49\n' +
    'AwEHoUQDQgAEQpsaqsSBBcLIunsPRJ6wR22sDczhVxgM4VhmQoiSJM6cW101oh7G\n' +
    'fXldIEqQidUmnSMyb8AuJoGqiRiimOeTPg==\n' +
    '-----END EC PRIVATE KEY-----\n';
    const jktThumbprint = 'yGtLGpDsckhNmFVNkMu1aAqL3DzBA2ArGwmwIPUSIAI';

    jest.mock('../../src/utils/dpopUtils', () => {
      const originalModule = jest.requireActual('../../src/utils/dpopUtils');
      return {
        ...originalModule,
        generateRandomString: jest.fn(() => 'kRxyB3pNzQ27aWm1ho4G6gIDKQFw8SU5Tc2JXeVL')
      };
    });

    const clientAssertion = await generateClientAssertion(url, clientId, privateSigningKey, jktThumbprint);
    console.log(clientAssertion);

    const expectedClientAssertion = "eyJ0eXAiOiJKV1QiLCJhbGciOiJFUzI1NiIsImtpZCI6ImlrVmw4OG90RE9rUW55UDdINE9kNE1sM3dDa0Q3Q3d1S2xvRlF1eTFjdGMifQ.eyJzdWIiOiJjbGllbnQxMjMiLCJqdGkiOiJacW96N0dvYU0xS0d5VGVQZFl0M08xOTF0ZHNmc2VLaHg5Q2FiWmRSIiwiYXVkIjoiaHR0cHM6Ly9leGFtcGxlLmNvbSIsImlzcyI6ImNsaWVudDEyMyIsImlhdCI6MTY5NzkxMzUzMiwiZXhwIjoxNjk3OTEzODMyLCJjbmYiOnsiamt0IjoieUd0TEdwRHNja2hObUZWTmtNdTFhQXFMM0R6QkEyQXJHd213SVBVU0lBSSJ9fQ.i8IxaajeqeuRrIdEPcM3sHcHbz1Nt50kwBNj4UzcHfOTchEN_fTRjDxnLPDWbugnTxsrn1brGpismyocRLRN_g";

    expect(clientAssertion).toBe(expectedClientAssertion);
  });
});
