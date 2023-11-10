import 'dotenv/config';
import { Octokit } from '@octokit/core';

const octokit = new Octokit({
  auth: process.env['GH_TOKEN'] // Your GitHub token from environment variable
});

const owner = process.env['GH_OWNER']; // Owner of the repository from environment variable
const repo = process.env['GH_REPO']; // Name of the repository from environment variable

const labelsToCreate = [
  { "name": "dsa", "color": "7FDBFF", "description": "Cryptographic Digital Signature Algorithms" },
  { "name": "key-mgmt", "color": "0074D9", "description": "Key Management" },
  { "name": "did:web", "color": "2ECC40", "description": "did:web" },
  { "name": "did:jwk", "color": "FFDC00", "description": "did:jwk" },
  { "name": "did:dht", "color": "FF851B", "description": "did:dht" },
  { "name": "did:key", "color": "F012BE", "description": "did:key" },
  { "name": "did:ion", "color": "B10DC9", "description": "did:ion" },
  { "name": "did-doc-validation", "color": "3D9970", "description": "DID Document & Resolution Validation" },
  { "name": "w3c-vc-dm-1.1", "color": "39CCCC", "description": "W3C Verifiable Credential Data Model 1.1" },
  { "name": "w3c-vc-dm-2.0", "color": "01FF70", "description": "W3C Verifiable Credential Data Model 2.0" },
  { "name": "sd-jwt", "color": "85144B", "description": "SD-JWT / SD-JWT-VC" },
  { "name": "pd-v2", "color": "F9A602", "description": "Presentation Definition V2" },
  { "name": "tbdex-message", "color": "70DB93", "description": "tbDEX Message" },
  { "name": "tbdex-resource", "color": "5B2C6F", "description": "tbDEX Resource" },
  { "name": "tbdex-offering", "color": "E59866", "description": "tbDEX Offering Resource" },
  { "name": "tbdex-rfq", "color": "1F618D", "description": "tbDEX RFQ Message" },
  { "name": "tbdex-quote", "color": "186A3B", "description": "tbDEX Quote Message" },
  { "name": "tbdex-order", "color": "28B463", "description": "tbDEX Order Message" },
  { "name": "tbdex-orderstatus", "color": "D68910", "description": "tbDEX Order-Status Message" },
  { "name": "tbdex-close", "color": "34495E", "description": "tbDEX Close Message" },
  { "name": "tbdex-server", "color": "3498DB", "description": "HTTP server for tbDEX PFIs " },
  { "name": "tbdex-client", "color": "E74C3C", "description": "HTTP client for tbDEX wallets" }
];

async function fetchExistingLabels() {
  const response = await octokit.request('GET /repos/{owner}/{repo}/labels', {
    owner,
    repo,
    headers: {
      'X-GitHub-Api-Version': '2022-11-28'
    }
  });

  return response.data.map(label => label.name);
}

async function createLabels() {
  const existingLabels = await fetchExistingLabels();

  for (const label of labelsToCreate) {
    // Check if the label already exists
    if (existingLabels.includes(label.name)) {
      console.log(`Label "${label.name}" already exists.`);
      continue; // Skip creating this label
    }

    try {
      const response = await octokit.request('POST /repos/{owner}/{repo}/labels', {
        owner,
        repo,
        name: label.name,
        color: label.color,
        description: label.description,
        headers: {
          'X-GitHub-Api-Version': '2022-11-28'
        }
      });

      if (response.status === 201) {
        console.log(`Created label: ${label.name}`);
      } else {
        console.log(`Failed to create label: ${label.name}. Status: ${response.status}`);
      }
    } catch (error) {
      console.error(`Failed to create label: ${label.name}. Error: ${error.message}`);
    }
  }
}

createLabels().catch(console.error);
