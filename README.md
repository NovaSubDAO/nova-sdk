# nova-python-sdk

[![Release](https://img.shields.io/github/v/release/solidi-labs/nova-python-sdk)](https://img.shields.io/github/v/release/solidi-labs/nova-python-sdk)
[![Build status](https://img.shields.io/github/actions/workflow/status/solidi-labs/nova-python-sdk/main.yml?branch=main)](https://github.com/solidi-labs/nova-python-sdk/actions/workflows/main.yml?query=branch%3Amain)
[![codecov](https://codecov.io/gh/solidi-labs/nova-python-sdk/branch/main/graph/badge.svg)](https://codecov.io/gh/solidi-labs/nova-python-sdk)
[![Commit activity](https://img.shields.io/github/commit-activity/m/solidi-labs/nova-python-sdk)](https://img.shields.io/github/commit-activity/m/solidi-labs/nova-python-sdk)
[![License](https://img.shields.io/github/license/solidi-labs/nova-python-sdk)](https://img.shields.io/github/license/solidi-labs/nova-python-sdk)

This is a template repository for Python projects that use Poetry for their dependency management.

- **Github repository**: <https://github.com/Solidi-Labs/nova-python-sdk/>
- **Documentation** <https://solidi-labs.github.io/nova-python-sdk/>

## Getting started with your project

Install the environment and the pre-commit hooks with

```bash
make install
```

To finalize the set-up for publishing to PyPi or Artifactory, see [here](https://fpgmaas.github.io/cookiecutter-poetry/features/publishing/#set-up-for-pypi).
For activating the automatic documentation with MkDocs, see [here](https://fpgmaas.github.io/cookiecutter-poetry/features/mkdocs/#enabling-the-documentation-on-github).
To enable the code coverage reports, see [here](https://fpgmaas.github.io/cookiecutter-poetry/features/codecov/).

## Releasing a new version

- Create an API Token on [Pypi](https://pypi.org/).
- Add the API Token to your projects secrets with the name `PYPI_TOKEN` by visiting [this page](https://github.com/solidi-labs/nova-python-sdk/settings/secrets/actions/new).
- Create a [new release](https://github.com/solidi-labs/nova-python-sdk/releases/new) on Github.
- Create a new tag in the form `*.*.*`.

For more details, see [here](https://fpgmaas.github.io/cookiecutter-poetry/features/cicd/#how-to-trigger-a-release).

---

Repository initiated with [fpgmaas/cookiecutter-poetry](https://github.com/fpgmaas/cookiecutter-poetry).
