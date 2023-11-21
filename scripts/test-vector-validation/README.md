# Test Vector Validation

## Description

Validates test vectors in [`web5-test-vectors`](../../web5-test-vectors/) directory. Uses [vectors.schema.json](../../web5-test-vectors/vectors.schema.json) to validate.

> [!NOTE]
> Runs automatically anytime a change is made in [`web5-test-vectors`](../../web5-test-vectors/) or to anything in this directory

## Setup

### `node` and `npm`

This project is using a minimum of `node v20.3.0` and `npm v9.6.7`. You can verify your `node` and `npm` installation via the terminal:

```bash
$ node --version
v20.3.0
$ npm --version
9.6.7
```

If you don't have `node` installed. Feel free to choose whichever approach you feel the most comfortable with. If you don't have a preferred installation method, i'd recommend using `nvm` (aka node version manager). `nvm` allows you to install and use different versions of node. It can be installed by running `brew install nvm` (assuming that you have homebrew)

Once you have installed `nvm`, install the desired node version with `nvm install vX.Y.Z`.

```bash
npm install
```

## Run

```bash
node main.js
```
