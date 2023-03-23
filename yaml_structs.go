package main

type Environment struct {
	Coinbase     string `yaml:"currentCoinbase"`
	Difficulty   string `yaml:"currentDifficulty"`
	GasLimit     string `yaml:"currentGasLimit"`
	Number       string `yaml:"currentNumber"`
	Timestamp    string `yaml:"currentTimestamp"`
	PreviousHash string `yaml:"previousHash"`
}

type Index struct {
	Data    int
	Gas     int
	Value   int
	Network []string
	Result  map[string]Account
}

type Info struct {
	Comment string `yaml:"comment"`
}

type Test struct {
	Info         Info               `yaml:"_info"`
	Env          Environment        `yaml:"env"`
	Pre          map[string]Account `yaml:"pre"`
	Transactions Transaction        `yaml:"transaction"`
	Post         []Index            `yaml:"expect"`
}

type Account struct {
	Code           string
	Storage        map[string]string
	Nonce          string
	ShouldNotExist int
}

type Transaction struct {
	Data      []string `yaml:"data"`
	GasLimit  []string `yaml:"gasLimit"`
	GasPrice  string   `yaml:"gasPrice"`
	Nonce     string   `yaml:"nonce"`
	To        string   `yaml:"to"`
	Value     []string `yaml:"value"`
	SecretKey string   `yaml:"secretKey"`
}

type Overall struct {
	Tests map[string]Test
}
