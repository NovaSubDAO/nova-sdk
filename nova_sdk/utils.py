import os

from nova_sdk.exceptions import EnvVarCantBeEmpty


def get_env(key):
    value = os.environ.get(key)
    if value == "":
        raise EnvVarCantBeEmpty
    return value
