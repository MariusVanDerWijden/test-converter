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

type AccessList struct {
	Address     string   `yaml:"address"`
	StorageKeys []string `yaml:"storageKeys"`
}

type Data struct {
	Data       string       `yaml:"data"`
	AccessList []AccessList `yaml:"accessList"`
}

type Transaction struct {
	Data      []*Data  `yaml:"data"` // Data can be both '' which is a string and Data
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
