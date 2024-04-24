import eth_abi
from web3.exceptions import ContractLogicError

from nova_sdk.config import load_config_and_web3

config, w3 = load_config_and_web3()


def withdraw(assets_: int, receiver_: str, owner_: str) -> None:
    """
    Redeem sDAI shares versus the equivalent in DAI.

    This function constructs and sends a transaction for withdrawing assets from a specified owner address
    to a receiver address by encoding the necessary data and executing the transaction on the network.

    Args:
        assets_ (int): The amount of assets to be withdrawn.
        receiver_ (str): The Ethereum address of the receiver in string format.
        owner_ (str): The Ethereum address of the owner from whom the assets are withdrawn.

    Returns:
        None: This function does not return any value but prints the transaction receipt.

    Raises:
        ContractLogicError: If the gas estimation fails due to a contract logic error.

    The function computes the function selector based on the defined function signature, encodes the transaction
    arguments (assets amount, receiver address, and owner address), estimates the required gas, and sends a
    signed transaction. If there is an error during gas estimation, it uses a default gas limit.

    Example:
        withdraw(500, "0xABCDEF1234567890abcdef1234567890abcdefAB", "0xABCDEF1234567890abcdef1234567890abcdefAB")
    """
    function_signature = "withdraw(uint256,address,address)"
    function_selector = w3.keccak(text=function_signature)[:4]

    receiver = w3.to_checksum_address(receiver_)
    owner = w3.to_checksum_address(owner_)

    encoded_arguments = eth_abi.encode(["uint256", "address", "address"], [assets_, receiver, owner])

    data = function_selector + encoded_arguments

    try:
        gas_estimate = w3.eth.estimate_gas(
            {"from": w3.eth.default_account, "to": config.sdai_address, "value": 0, "data": data}
        )
    except ContractLogicError as e:
        print(f"Gas estimation failed due to contract logic error: {e}")
        gas_estimate = 20_00_000  # fallback gas limit

    transaction = {
        "to": config.sdai_address,
        "value": 0,
        "gasPrice": w3.eth.gas_price,
        "gas": gas_estimate,
        "nonce": w3.eth.get_transaction_count(w3.eth.default_account),
        "data": data,
        "chainId": config.chain_id,
    }

    signed_txn = w3.eth.account.sign_transaction(transaction, config.private_key)
    tx_hash = w3.eth.send_raw_transaction(signed_txn.rawTransaction)
    tx_receipt = w3.eth.wait_for_transaction_receipt(tx_hash)

    print(tx_receipt)
