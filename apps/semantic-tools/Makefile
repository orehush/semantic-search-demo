# Makefile for semantic-tools

# Define the paths to the necessary tools and files
PYTHON=python3
PIP=pip
DOCKER=docker
DOCKER_COMPOSE=docker-compose
UVICORN=uvicorn
TEST_CMD=pytest
LINT_CMD=flake8
DOCKER_IMAGE=semantic-tools:latest
VENV_DIR=venv

# Define the paths for src and tests
SRC_DIR=src
TESTS_DIR=src/tests

install: $(VENV_DIR)/bin/activate

$(VENV_DIR)/bin/activate: requirements.txt
	python3 -m venv $(VENV_DIR)
	$(PIP) install --upgrade pip
	$(PIP) install -r requirements.txt
	@touch $(VENV_DIR)/bin/activate

# Run PEP8 and linting checks
check:
	$(PIP) install flake8
	$(LINT_CMD) --exclude=venv $(SRC_DIR)

# Run FastAPI server locally
run:
	PYTHONPATH=$(SRC_DIR) $(UVICORN) src.main:app --reload --host 0.0.0.0 --port 8000

# Run tests locally
test:
	PYTHONPATH=$(SRC_DIR) $(PIP) install -r requirements.txt
	PYTHONPATH=$(SRC_DIR) $(PIP) install pytest
	PYTHONPATH=$(SRC_DIR) $(TEST_CMD) $(TESTS_DIR)

# Build Docker image and run FastAPI server in Docker container
docker_run:
	$(DOCKER) build -t $(DOCKER_IMAGE) .
	$(DOCKER) run -d --name semantic-tools-container -p 8000:8000 $(DOCKER_IMAGE)

# Run tests in Docker container
docker_test:
	$(DOCKER) build -t semantic-tools-test .
	$(DOCKER) run --rm -e PYTHONPATH=/app/src semantic-tools-test /bin/sh -c "pip install pytest && pytest /app/src/tests"
