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
        "2adc25665018aa1fe0e6bc666dac8fc2697ff9ba" : Account(
            code: ""
            nonce: 1
            storage: 
        )
        "095e7baea6a6c7c4c2dfeb977efac326af552d87" : Account(
            code: """lll({
   ; Can also add lll style comments here
   [[0]] (ADD 1 1) 
}
)"""
            nonce: 0
            storage: 
        )
        "a94f5374fce5edbc8e2a8697c15331677e6ebf0b" : Account(
            code: """lll(0x)"""
            nonce: 0
            storage: 
        )

    }

    tx = Transaction(
        data: 
        gasLimit: 400000
        gasPrice: 10
        nonce: 0
        secretKey: "45a915e4d060149eb4365960e6a7a45f334393093061116b197e3240065ff2d8"
        to: "095e7baea6a6c7c4c2dfeb977efac326af552d87"
        value: 100000
    )


    post = {
        "095e7baea6a6c7c4c2dfeb977efac326af552d87" : Account(
            code: """lll(0x600160010160005500)"""
            nonce: 
            storage: 
        )
        "2adc25665018aa1fe0e6bc666dac8fc2697ff9ba" : Account(
            code: ""
            nonce: 1
            storage: 
        )
        "a94f5374fce5edbc8e2a8697c15331677e6ebf0b" : Account(
            code: """lll(0x)"""
            nonce: 1
            storage: 
        )
        "e94f5374fce5edbc8e2a8697c15331677e6ebf0b" : Account(
            code: ""
            nonce: 
            storage: 
        )

    }


    yield StateTest(env=env, pre=pre, post=post, txs=[tx])
    