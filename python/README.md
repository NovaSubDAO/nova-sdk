# nova-sdk

[![Release](https://img.shields.io/github/v/release/novasubdao/nova-sdk)](https://img.shields.io/github/v/release/novasubdao/nova-sdk)
[![Build status](https://img.shields.io/github/actions/workflow/status/novasubdao/nova-sdk/main.yml?branch=main)](https://github.com/novasubdao/nova-sdk/actions/workflows/main.yml?query=branch%3Amain)
[![codecov](https://codecov.io/gh/novasubdao/nova-sdk/branch/main/graph/badge.svg)](https://codecov.io/gh/novasubdao/nova-sdk)
[![Commit activity](https://img.shields.io/github/commit-activity/m/novasubdao/nova-sdk)](https://img.shields.io/github/commit-activity/m/novasubdao/nova-sdk)
[![License](https://img.shields.io/github/license/novasubdao/nova-sdk)](https://img.shields.io/github/license/novasubdao/nova-sdk)

This is a template repository for Python projects that use Poetry for their dependency management.

- **Github repository**: <https://github.com/novasubdao/nova-sdk/>
- **Documentation** <https://novasubdao.github.io/nova-sdk/>

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
- Add the API Token to your projects secrets with the name `PYPI_TOKEN` by visiting [this page](https://github.com/novasubdao/nova-sdk/settings/secrets/actions/new).
- Create a [new release](https://github.com/novasubdao/nova-sdk/releases/new) on Github.
- Create a new tag in the form `*.*.*`.

For more details, see [here](https://fpgmaas.github.io/cookiecutter-poetry/features/cicd/#how-to-trigger-a-release).

---

Repository initiated with [fpgmaas/cookiecutter-poetry](https://github.com/fpgmaas/cookiecutter-poetry).
