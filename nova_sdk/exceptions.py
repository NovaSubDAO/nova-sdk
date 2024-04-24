class EnvVarCantBeEmpty(Exception):
    def __init__(self, key):
        super().__init__(f"You should pass an non-empty envionment variable for {key}.")


class Web3ShouldBeConnected(Exception):
    def __init__(self):
        super().__init__("Failed to connect to the network")
