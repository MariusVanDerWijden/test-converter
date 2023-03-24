package main

import (
	"strings"
	"testing"
)

func TestPreprocess(t *testing.T) {
	test := `# Multiple transaction description that will be applied on a pre state      
	transaction:
	  data:
	  - ''
	  # Be careful when adding gasLimits that are < then intrinsic gas cost. StateTests are not allowed to have invalid transactions
	  gasLimit:
	  - '400000'
	  gasPrice: '10'
	  # Nonce must be same as account nonce a94f5374fce5edbc8e2a8697c15331677e6ebf0b
	  nonce: '0'
	  to: 095e7baea6a6c7c4c2dfeb977efac326af552d87
	  value:
	  - '100000'
	  # secretKey is a privKey of a94f5374fce5edbc8e2a8697c15331677e6ebf0b
	  secretKey: "45a915e4d060149eb4365960e6a7a45f334393093061116b197e3240065ff2d8"'`

	want := `# Multiple transaction description that will be applied on a pre state      
	transaction:
	  
	  # Be careful when adding gasLimits that are < then intrinsic gas cost. StateTests are not allowed to have invalid transactions
	  gasLimit:
	  - '400000'
	  gasPrice: '10'
	  # Nonce must be same as account nonce a94f5374fce5edbc8e2a8697c15331677e6ebf0b
	  nonce: '0'
	  to: 095e7baea6a6c7c4c2dfeb977efac326af552d87
	  value:
	  - '100000'
	  # secretKey is a privKey of a94f5374fce5edbc8e2a8697c15331677e6ebf0b
	  secretKey: "45a915e4d060149eb4365960e6a7a45f334393093061116b197e3240065ff2d8"'`

	got := Preprocess(test)
	if !strings.EqualFold(got, want) {
		t.Fatalf("preprocessing failed, got %v want %v", got, want)
	}
}
