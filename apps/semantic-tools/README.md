# Semantic Tools

## Overview

**Semantic Tools** is a FastAPI-based web service designed to generate synonyms for a given word or phrase using the NLTK (Natural Language Toolkit) library. The project provides a simple REST API endpoint where you can submit a phrase, and it will return a list of synonyms sorted by their relevance.

## Features

- **Synonym Generation**: Takes a phrase as input and returns a list of synonyms along with a relevance score.
- **FastAPI Framework**: Built using FastAPI, ensuring high performance and ease of use.
- **Dockerized**: The project is fully containerized, making it easy to deploy and run in any environment.

## Installation

### Prerequisites

- **Python 3.8+**
- **Docker** (optional, for running in a container)
- **Make** (to run the provided commands)

### Local Setup

1. **Clone the repository**:
   ```bash
   git clone https://github.com/orehush/semantic-search-demo
   cd apps/semantic-tools

2. **Install dependencies**
    ```bash
    make check

3. **Run the server**
    ```bash
    make run

### Docker Setup

1. **Build and run the Docker container:**
    ```bash
    make docker_run

### Usage

1. Generate synonyms:
- Endpoint: /semantic-tools/synonyms/
- Method: POST
- Content-Type: application/json
- Request Body:
    ```json
        {
            "phrase": "car"
        }
    ```
- Response:
    ```json
        {
            "synonyms": [
                {"synonym": "automobile", "score": 0.9},
                {"synonym": "vehicle", "score": 0.8}
            ]
        }
    ```

### Running Tests

To run the unit tests for the project, use the following command:

    ```bash
    make test
    ```

You can also run tests inside a Docker container:

    ```bash
    make docker_test
    ```
