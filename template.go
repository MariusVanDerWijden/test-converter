package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
	"sort"
	"text/template"
)

type kv struct {
	Key   string
	Value string
}

func fillTest(t Test, testname string) ([]byte, error) {
	testTemplate, err := os.ReadFile("templates/test.txt")
	if err != nil {
		return nil, err
	}

	tmpl, err := template.New("TestTemplate").Parse(string(testTemplate))
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(nil)

	preState, err := fillAccounts(t.Pre)
	if err != nil {
		return nil, err
	}

	transactions, err := fillTransactions(t.Transactions)
	if err != nil {
		return nil, err
	}

	postState, err := fillAccounts(t.Post[0].Result)

	f := struct {
		Description  string
		EarliestFork string
		TestName     string
		PreState     string
		Transaction  string
		PostState    string
	}{
		Description:  t.Info.Comment,
		EarliestFork: "Berlin",
		TestName:     testname,
		PreState:     preState,
		Transaction:  transactions,
		PostState:    postState,
	}

	if err := tmpl.Execute(buf, f); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func fillAccounts(accounts map[string]Account) (string, error) {
	testTemplate, err := os.ReadFile("templates/account.txt")
	if err != nil {
		return "", err
	}

	tmpl, err := template.New("Account").Parse(string(testTemplate))
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer(nil)

	// Make accounts order deterministic
	names := make([]string, 0)
	for name := range accounts {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		account := accounts[name]
		storage, err := fillStorage(account.Storage)
		if err != nil {
			return "", err
		}
		values := make([]kv, 0)
		setArr(&values, "code", handleCode(account.Code))
		setArr(&values, "nonce", account.Nonce)
		setArr(&values, "balance", account.Balance)
		setArr(&values, "storage", storage)

		f := struct {
			Name   string
			Values []kv
		}{
			Name:   name,
			Values: values,
		}
		// Add padding to the buffer
		buf.Write([]byte("        "))
		if err := tmpl.Execute(buf, f); err != nil {
			return "", err
		}
	}
	return buf.String(), nil
}

func fillTransactions(tx Transaction) (string, error) {
	testTemplate, err := os.ReadFile("templates/transaction.txt")
	if err != nil {
		return "", err
	}

	tmpl, err := template.New("Account").Parse(string(testTemplate))
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer(nil)

	// if we have no txdata, we need to add some empty data
	if len(tx.Data) == 0 {
		tx.Data = append(tx.Data, DataWrapper{
			d: Data{
				Data:       "",
				AccessList: []AccessList{},
			},
		})
	}
	// One transaction field can have multiple transaction descriptions
	for i := 0; i < len(tx.Data); i++ {
		values := make([]kv, 0)
		setArr(&values, "code", handleCode(tx.Data[i].d.Data))
		if len(tx.Data[i].d.AccessList) > 0 {
			al, err := fillAccesslist(tx.Data[i].d.AccessList)
			if err != nil {
				return "", err
			}
			setArr(&values, "access_list", al)
		}
		setArr(&values, "nonce", tx.Nonce)
		setArr(&values, "max_fee_per_gas", tx.MaxFeePerGas)
		setArr(&values, "max_priority_fee_per_gas", tx.MaxPriorityFeePerGas)
		setValueIfExists(&values, "gas_limit", tx.GasLimit, i)
		setArr(&values, "gas_price", tx.GasPrice)
		setArr(&values, "to", stringify(tx.To))
		setValueIfExists(&values, "value", tx.Value, i)
		setArr(&values, "secret_key", stringify(tx.SecretKey))

		f := struct {
			Values []kv
		}{
			Values: values,
		}
		if err := tmpl.Execute(buf, f); err != nil {
			return "", err
		}
	}
	return buf.String(), nil
}

func fillAccesslist(als []AccessList) (string, error) {
	testTemplate, err := os.ReadFile("templates/accesslist.txt")
	if err != nil {
		return "", err
	}

	tmpl, err := template.New("AccessList").Parse(string(testTemplate))
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer(nil)

	for _, al := range als {
		f := struct {
			Address     string
			StorageKeys []string
		}{
			Address:     al.Address,
			StorageKeys: al.StorageKeys,
		}
		if err := tmpl.Execute(buf, f); err != nil {
			return "", err
		}
	}
	return buf.String(), nil
}

func fillStorage(storage map[string]string) (string, error) {
	if len(storage) == 0 {
		return "", nil
	}
	testTemplate, err := os.ReadFile("templates/storage.txt")
	if err != nil {
		return "", err
	}

	tmpl, err := template.New("Storage").Parse(string(testTemplate))
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer(nil)

	// Make storage order deterministic
	keys := make([]string, 0)
	for key := range storage {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	values := make([]kv, 0)
	for _, key := range keys {
		value := storage[key]
		setArr(&values, key, value)
	}
	f := struct {
		Values []kv
	}{
		Values: values,
	}
	if err := tmpl.Execute(buf, f); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func stringify(str string) string {
	return fmt.Sprintf("\"%s\"", str)
}

func handleCode(code string) string {
	// Code can be empty
	if len(code) < 2 {
		return stringify(code)
	}
	// Code can be hex string, drop 0x
	if _, err := hex.DecodeString(code[2:]); err == nil {
		return stringify(code)
	}
	// Code can be LLL
	return fmt.Sprintf("lll(\"\"\"%v\"\"\")", code)
}

func setArr(array *[]kv, key, value string) {
	if value != "" {
		*array = append(*array, kv{
			Key:   key,
			Value: value,
		})
	}
}

func setValueIfExists(array *[]kv, key string, value []string, index int) {
	// No value exists, return
	if len(value) == 0 {
		return
	}
	// index is less than the max length, set
	if index < len(value) {
		setArr(array, key, value[index])
	}
	// index is not available, set [0] value
	setArr(array, key, value[0])
}
