import 'dotenv/config';
import { Octokit } from '@octokit/core';

const octokit = new Octokit({
  auth: process.env['GH_TOKEN'] // Your GitHub token from environment variable
});

const owner = process.env['GH_OWNER']; // Owner of the repository from environment variable
const repo = process.env['GH_REPO']; // Name of the repository from environment variable

const newMilestones = [
  { "title": "Eagle" },
  { "title": "ABC" }
  // Add more milestones here if needed
];

async function fetchExistingMilestones() {
  const response = await octokit.request('GET /repos/{owner}/{repo}/milestones', {
    owner,
    repo,
    headers: {
      'X-GitHub-Api-Version': '2022-11-28'
    }
  });

  return response.data.map(milestone => milestone.title);
}

async function createMilestones() {
  const existingMilestones = await fetchExistingMilestones();

  for (const milestone of newMilestones) {
    // Check if the milestone already exists
    if (existingMilestones.includes(milestone.title)) {
      console.log(`Milestone "${milestone.title}" already exists.`);
      continue; // Skip creating this milestone
    }

    try {
      const response = await octokit.request('POST /repos/{owner}/{repo}/milestones', {
        owner,
        repo,
        title: milestone.title,
        headers: {
          'X-GitHub-Api-Version': '2022-11-28'
        }
      });

      if (response.status === 201) {
        console.log(`Created milestone: ${milestone.title}`);
      } else {
        console.log(`Failed to create milestone: ${milestone.title}. Status: ${response.status}`);
      }
    } catch (error) {
      console.error(`Failed to create milestone: ${milestone.title}. Error: ${error.message}`);
    }
  }
}

createMilestones().catch(console.error);
