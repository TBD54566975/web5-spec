import "dotenv/config";
import { Octokit } from "@octokit/core";

const octokit = new Octokit({
  auth: process.env["GH_TOKEN"], // Your GitHub token from environment variable
});

const owner = process.env["GH_OWNER"]; // Owner of the repository from environment variable

const projectId = "PVT_kwDOBaHX684AX9Bo"; // SDK Development

const repos = [
  { repo: "tbdex-rs", labels: ["cicd"] },
  { repo: "tbdex-js", labels: ["cicd"] },
  { repo: "tbdex-kt", labels: ["cicd"] },
  { repo: "web5-kt", labels: ["cicd"] },
  { repo: "web5-js", labels: ["cicd"] },
  { repo: "web5-rs", labels: ["cicd"] },
];

const issueTitle = "Setup artifacts publishing";

const issueBody = `
We need to have an easy automated workflow to publish artifacts to the corresponding sdk registry (npm, maven, crates etc.), and also publish to the internal Artifactory registry too
`;

async function createIssue(repoName, issueTitle, issueBody, labels) {
  try {
    console.info(`Creating issue in ${repoName}...`)
    const issueResponse = await octokit.request(
      "POST /repos/{owner}/{repo}/issues",
      {
        owner,
        repo: repoName,
        title: issueTitle,
        body: issueBody,
        labels,
      }
    );
    console.info(`Issue created in ${repoName}: ${issueResponse.data.html_url}`);

    await addIssueToProject(issueResponse.data.node_id);
  } catch (error) {
    console.error(
      `Error creating and associating issue in ${repoName}: ${error.message}`
    );
  }
}

async function addIssueToProject(issueId) {
  try {
    const mutation = `
        mutation ($projectId: ID!, $contentId: ID!) {
            addProjectV2ItemById(input: {projectId: $projectId, contentId: $contentId}) {
                item {
                    id
                }
            }
        }
    `;
    await octokit.graphql(mutation, {
      projectId,
      contentId: issueId,
    });
    console.info(`Issue added to project: ${issueId}`);
  } catch (error) {
    console.error(`Error adding issue to project: ${error.message}`);
  }
}

async function createIssues() {
  for (const repo of repos) {
    await createIssue(repo.repo, issueTitle, issueBody, repo.labels);
  }
}

/**
 * Used to list all projects ids,so we can update the script
 */
async function listOrgProjects() {
  const query = `
      query ($orgName: String!) {
          organization(login: $orgName) {
              projectsV2(first: 100) {
                  nodes {
                      id
                      title
                      url
                  }
              }
          }
      }
  `;

  try {
      const response = await octokit.graphql(query, { orgName: owner });
      const projects = response.organization.projectsV2.nodes;
      projects.forEach(project => {
          console.info(`${project.id} - Project Title: ${project.title}, URL: ${project.url}`);
      });
  } catch (error) {
      console.error(`Error fetching projects: ${error.message}`);
  }
}

createIssues().catch(console.error);
// listOrgProjects().catch(console.error);
