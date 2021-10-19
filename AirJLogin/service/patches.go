package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

type item struct {
	Acc  string `json:"acc"`
	Pass string `json:"pass"`
}

type Data struct {
	item
	Alternatives []item `json:"alternatives"`
}

var path string
var homeDir string

func init() {
	cur, err := user.Current() // Get the User document path
	if err != nil {
		log.Fatal("User Path Get Failed!")
	}
	homeDir = cur.HomeDir
	path = cur.HomeDir + "/AirJLogin/data.json"
}

// ReadContent from patch file
func ReadContent() (Data, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return Data{}, err
	}
	var data Data
	err = json.Unmarshal([]byte(content), &data)
	if err != nil {
		log.Printf("Unmarshal data error: %v\n", err.Error())
		return Data{}, err
	}
	return data, nil
}

//ReadLoginData from specific patch file
func ReadLoginData() (acc string, pass string) {
	data, err := ReadContent()

	// first login
	if err != nil {
		_ = os.MkdirAll(homeDir+"\\AirJLogin", 0644)
		log.Println(`Patch File "data.json" Not Found!`)
		data := &Data{}
		// input
		log.Println("You first Login, please input your account and password.")
		fmt.Print("Account:")
		_, _ = fmt.Scanf("%s\n", &data.Acc)
		fmt.Print("Password:")
		_, _ = fmt.Scanf("%s\n", &data.Pass)
		// Write to file
		it := item{
			Acc:  data.Acc,
			Pass: data.Pass,
		}
		data.Alternatives = append(data.Alternatives, it)
		info, _ := json.Marshal(data)
		_ = ioutil.WriteFile(path, info, 0644)
		log.Println("Please fill the patch file.")
		return data.Acc, data.Pass
	}

	if len(data.Acc) > 0 && len(data.Pass) > 0 {
		return data.Acc, data.Pass
	} else {
		log.Fatal("Account and Password not set!")
		return "", ""
	}
}

func AddLoginData(acc string, pass *string) bool {
	data, err := ReadContent()
	if err != nil {
		log.Printf("Read login data error: %v\n", err.Error())
		return false
	}

	// Update data
	has := false
	for i, item := range data.Alternatives {
		if item.Acc == acc {
			if len(*pass) >= 6 {
				data.Alternatives[i].Pass = *pass
			} else {
				*pass = item.Pass // 返回密码
			}
			has = true
			break
		}
	}

	if !has {
		if len(*pass) >= 6 {
			data.Alternatives = append(data.Alternatives, item{
				Acc:  acc,
				Pass: *pass,
			})
		} else {
			log.Println("New account password is invalid or not specified")
			return false
		}

	}
	// Add new account
	data.Acc, data.Pass = acc, *pass

	info, _ := json.Marshal(data)

	err = ioutil.WriteFile(path, info, 0644)
	if err != nil {
		return false
	}
	return true
}

func DelAccount(acc string) bool {
	data, err := ReadContent()
	if err != nil {
		return false
	}

	if acc == data.Acc {
		log.Println("Your are using this account.")
		return false
	}

	has := false

	for i, item := range data.Alternatives {
		if item.Acc == acc {
			if i == len(data.Alternatives)-1 {
				data.Alternatives = data.Alternatives[0 : len(data.Alternatives)-1]
			} else {
				data.Alternatives = append(data.Alternatives[0:i], data.Alternatives[i+1:]...)
			}
			has = true
			break
		}
	}

	if !has {
		log.Println("This account not exists.")
		return false
	}

	info, _ := json.Marshal(data)

	err = ioutil.WriteFile(path, info, 0644)
	if err != nil {
		return false
	}
	return true
}
