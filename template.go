package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"os"
	"text/template"
)

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

	for name, account := range accounts {

		storage, err := fillStorage(account)
		if err != nil {
			return "", err
		}
		values := make(map[string]string)
		values["code"] = handleCode(account.Code)
		values["nonce"] = account.Nonce
		values["storage"] = storage

		f := struct {
			Name   string
			Values map[string]string
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

	// One transaction field can have multiple transaction descriptions
	for i := 0; i < len(tx.Data); i++ {
		values := make(map[string]string)
		values["data"] = tx.Data[i]
		values["nonce"] = tx.Nonce
		values["gasLimit"] = tx.GasLimit[i]
		values["gasPrice"] = tx.GasPrice
		values["to"] = stringify(tx.To)
		values["value"] = tx.Value[i]
		values["secretKey"] = stringify(tx.SecretKey)

		f := struct {
			Values map[string]string
		}{
			Values: values,
		}
		if err := tmpl.Execute(buf, f); err != nil {
			return "", err
		}
	}
	return buf.String(), nil
}

func fillStorage(account Account) (string, error) {
	return "", nil
}

func stringify(str string) string {
	return fmt.Sprintf("\"%s\"", str)
}

func handleCode(code string) string {
	// Code can be hex string
	if _, err := hex.DecodeString(code); err == nil {
		return stringify(code)
	}
	// Code can be LLL
	return fmt.Sprintf("\"\"\"lll(%v)\"\"\"", code)
}
