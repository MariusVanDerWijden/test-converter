"""
A test for (add 1 1) opcode result
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
def add11_yml(fork):
    """
    A test for (add 1 1) opcode result
    """

    env = Environment()
    
    pre = {
        "095e7baea6a6c7c4c2dfeb977efac326af552d87" : Account(
            code="""lll({
   ; Can also add lll style comments here
   [[0]] (ADD 1 1) 
}
)""",
            nonce=0,
            balance=1000000000000000000,
        )
        "2adc25665018aa1fe0e6bc666dac8fc2697ff9ba" : Account(
            code="",
            nonce=1,
            balance=0,
        )
        "a94f5374fce5edbc8e2a8697c15331677e6ebf0b" : Account(
            code="0x",
            nonce=0,
            balance=1000000000000000000,
        )

    }

    tx = Transaction(
        code="",
        nonce=0,
        gas_limit=400000,
        gas_price=10,
        to="095e7baea6a6c7c4c2dfeb977efac326af552d87",
        value=100000,
        secret_key="45a915e4d060149eb4365960e6a7a45f334393093061116b197e3240065ff2d8",
    )


    post = {
        "095e7baea6a6c7c4c2dfeb977efac326af552d87" : Account(
            code="0x600160010160005500",
            storage={
                0x00: 2,
            },
        )
        "2adc25665018aa1fe0e6bc666dac8fc2697ff9ba" : Account(
            code="",
            nonce=1,
        )
        "a94f5374fce5edbc8e2a8697c15331677e6ebf0b" : Account(
            code="0x",
            nonce=1,
        )
        "e94f5374fce5edbc8e2a8697c15331677e6ebf0b" : Account(
            code="",
        )

    }


    yield StateTest(env=env, pre=pre, post=post, txs=[tx])
    