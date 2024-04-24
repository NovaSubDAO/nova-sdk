import os

from nova_sdk.exceptions import EnvVarCantBeEmpty


def get_env(key: str) -> str | None:
    value = os.environ.get(key)
    if value == "":
        raise EnvVarCantBeEmpty(key)
    return value
