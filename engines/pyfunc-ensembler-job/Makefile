SHELL := /bin/bash

PYTHON_VERSION ?= 3.7
APP_NAME := pyfunc-ensembler-job-py$(PYTHON_VERSION)
CONDA_ENV_NAME ?= $(APP_NAME)
ACTIVATE_ENV = source $$(conda info --base)/etc/profile.d/conda.sh ; conda activate $(CONDA_ENV_NAME)

.PHONY: setup
setup: $(CONDA_ENV_NAME)
$(CONDA_ENV_NAME):
	@conda env update -f env-$(PYTHON_VERSION).yaml -n $(CONDA_ENV_NAME) --prune
	$(ACTIVATE_ENV) && pip install -r requirements.dev.txt

.PHONY: type-check
type-check:
	@$(ACTIVATE_ENV) && mypy \
		--install-types \
		--non-interactive \
		--ignore-missing-imports \
		--allow-untyped-globals \
		ensembler

.PHONY: lint
lint:
	@$(eval black_version=$(shell pip show black | grep Version | grep -oE '[0-9].*'))
	@[ "${black_version}" == "22.6.0" ] || sh -c "pip install black==22.6.0"
	@black .

.PHONY: test
test: type-check
	@$(ACTIVATE_ENV) && \
		python -m pytest \
		--cov=ensembler \
		--ignore=env \
		-W ignore

.PHONY: build-image
build-image: version
	@mkdir -p temp-deps
	@cp -r ../../sdk temp-deps/
	@$(eval IMAGE_TAG = $(if $(DOCKER_REGISTRY),$(DOCKER_REGISTRY)/,)${APP_NAME}:${VERSION})
	@echo "Building docker image: ${IMAGE_TAG}"
	@docker build . --tag ${IMAGE_TAG} --build-arg PYTHON_VERSION=${PYTHON_VERSION}
	@rm -rf temp-deps

.PHONY: version
version:
	$(eval VERSION=$(if $(OVERWRITE_VERSION),$(OVERWRITE_VERSION),v$(shell ../../scripts/vertagen/vertagen.sh -p ${APP_NAME}/)))
	@echo "turing-pyfunc-ensembler-job version:" $(VERSION)
