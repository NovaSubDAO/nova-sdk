import eth_abi
from web3.exceptions import ContractLogicError

from nova_sdk.config import load_config_and_web3

config, w3 = load_config_and_web3()


def deposit(assets_, receiver_):
    function_signature = "deposit(uint256,address)"
    function_selector = w3.keccak(text=function_signature)[:4]

    receiver = w3.to_checksum_address(receiver_)

    encoded_arguments = eth_abi.encode(["uint256", "address"], [assets_, receiver])

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
