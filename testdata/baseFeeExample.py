"""
A test shows basefee transaction example
"""

from ethereum_test_tools import (
	Account,
	Environment,
	StateTest,
	TestAddress,
	Transaction,
	Yul,
	test_from,
)

@test_from("Berlin")
def basefeeExample(fork):
    """
    A test shows basefee transaction example
    """

    env = Environment()
    
    pre = {
        "095e7baea6a6c7c4c2dfeb977efac326af552d87" : Account(
            code=lll("""{
   ; Can also add lll style comments here
   [[0]] (ADD 1 1) 
}
"""),
            nonce=0,
            balance=1000000000000000000,
        )
        "a94f5374fce5edbc8e2a8697c15331677e6ebf0b" : Account(
            code="0x",
            nonce=0,
            balance=1000000000000000000,
        )

    }

    tx = Transaction(
        code=lll(""":label declaredKeyWrite :raw 0x00"""),
        access_list=AccessList(
            address="0x095e7baea6a6c7c4c2dfeb977efac326af552d87",
            storage_keys=[
                "0x00", 
                "0x01", 
                
            ],
        ),
        nonce=0,
        max_fee_per_gas=0x12A05F200,
        max_priority_fee_per_gas=2,
        gas_limit=4000000,
        to="095e7baea6a6c7c4c2dfeb977efac326af552d87",
        value=100000,
        secret_key="45a915e4d060149eb4365960e6a7a45f334393093061116b197e3240065ff2d8",
    )


    post = {
        "095e7baea6a6c7c4c2dfeb977efac326af552d87" : Account(
            code="",
            storage={
                0x00: 2,
            },
        )

    }


    yield StateTest(env=env, pre=pre, post=post, txs=[tx])
    