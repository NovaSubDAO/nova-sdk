from nova_sdk.config import SDAI_ADDRESSES_MAP


def test_nb_supported_networks():
    assert len(SDAI_ADDRESSES_MAP) == 1
