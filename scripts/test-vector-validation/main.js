import fs from 'node:fs'
import path from 'path'
import Ajv from 'ajv'

import { fileURLToPath } from 'node:url'

const __filename = fileURLToPath(import.meta.url)
const __dirname = path.dirname(__filename)

const vectorsDir = `${__dirname}/../../test-vectors`

let vectorsSchema = fs.readFileSync(`${vectorsDir}/vectors.schema.json`, 'utf8')
vectorsSchema = JSON.parse(vectorsSchema)

const ajv = new Ajv()
const validate = ajv.compile(vectorsSchema)

function validateTestVectors() {
  const entries = fs.readdirSync(vectorsDir, { withFileTypes: true })

  function validateDescriptions(testData) {
    // accumulate duplicates inside a map
    const descriptions = new Set()
    for (const vector of testData.vectors) {
      if (descriptions.has(vector.description)) {
        console.log("Duplicate description found: \"" + vector.description + "\". Descriptions are meant to be unique.")
        process.exit(1)
      }
      descriptions.add(vector.description)
    }
  }

  for (const entry of entries) {
    if (!entry.isDirectory()) {
      continue
    }

    const featureDir = path.join(vectorsDir, entry.name)
    const files = fs.readdirSync(featureDir)

    for (const file of files) {
      if (path.extname(file) === '.json') {
        const filePath = path.join(featureDir, file)
        const fileContent = fs.readFileSync(filePath, 'utf8')
        const testData = JSON.parse(fileContent)

        if (!validate(testData)) {
          console.log(`Validation failed for ${filePath}:`, validate.errors)
          process.exit(1)
        }
        validateDescriptions(testData)
        console.log(`Validation passed for ${filePath}`)
      }
    }
  }
}

validateTestVectors()
