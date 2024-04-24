from dataclasses import dataclass

from web3 import Web3

from nova_sdk.exceptions import Web3ShouldBeConnected
from nova_sdk.utils import get_env

SDAI_ADDRESSES_MAP = {"ETHEREUM": "0x83F20F44975D03b1b09e64809B757c47f942BEeA"}


@dataclass
class Config:
    private_key: str
    rpc_endpoint: str
    chain_id: str
    sdai_address: str


def load_config():
    config = Config(
        private_key=get_env("PRIVATE_KEY"),
        rpc_endpoint=get_env("RPC_ENDPOINT"),
        chain_id=get_env("CHAIN_ID"),
        sdai_address=SDAI_ADDRESSES_MAP[get_env("NETWORK")],
    )
    return config


def load_config_and_web3():
    config = load_config()

    w3 = Web3(Web3.HTTPProvider(config.rpc_endpoint))
    if not w3.is_connected():
        raise ValueError(Web3ShouldBeConnected)

    config.sdai_address = w3.to_checksum_address(config.sdai_address)

    account = w3.eth.account.from_key(config.private_key)
    w3.eth.default_account = account.address

    return config, w3
