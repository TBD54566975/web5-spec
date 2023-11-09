# GitHub Repository Scripts

This directory contains scripts that help with individual SDK repo management.

## Scripts Description

| Script Name                 | Description                                                                   |
| --------------------------- | ----------------------------------------------------------------------------- |
| `npm run create-labels`     | Creates GitHub labels from a predefined list if they don't already exist.     |
| `npm run create-milestones` | Creates GitHub milestones from a predefined list if they don't already exist. |

## Installation and Usage

To run the scripts, first install the necessary dependencies:

```bash
npm install
```

## Environment Variables

Before running the scripts, ensure that a `.env` file exists and has the following variables set. You can copy `.env.example`:

| Environment Variable | Description                                  |
| -------------------- | -------------------------------------------- |
| `GH_TOKEN`           | Your GitHub personal access token.           |
| `GH_OWNER`           | The owner username of the GitHub repository. |
| `GH_REPO`            | The name of the GitHub repository.           |
