package utils

import (
	"QuakeAPI/model"
	"bufio"
	"flag"
	"os"
)

func ParseInput() model.Input {
	var (
		userInfo bool
		key      string
		search   string
		help     bool
		output   string
		total    int
		result   = model.Input{}
		keyFile  = "key.ini"
	)
	flag.StringVar(&key, "key", "", "SetUp Your API Key.")
	flag.IntVar(&total, "total", 100, "Number Of Queries You Want.")
	flag.StringVar(&search, "search", "", "Input Search String.")
	flag.StringVar(&output, "output", "quake.txt", "Output File.")
	flag.BoolVar(&userInfo, "userinfo", false, "Show Your User Information.")
	flag.BoolVar(&help, "help", false, "Show Help Information.")
	flag.Parse()

	if len(key) > 0 {
		println("[*] Your API key is:", key)
		if os.WriteFile(keyFile, []byte(key), 0644) != nil {
			println("[-] Setup API key error")
		} else {
			println("[+] Setup API key success")
		}
		println()
		return result
	} else if help == true {
		flag.PrintDefaults()
		return result
	} else {
		keyByte, err := os.ReadFile(keyFile)
		if err != nil {
			return result
		}

		if len(search) == 0 {
			// 创建一个 Scanner 来从标准输入读取
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			if scanner.Err() != nil {
				return result
			} else {
				search = scanner.Text()
			}
		}
		result.UserInfo = userInfo
		result.Key = string(keyByte)
		result.Search = search
		result.Output = output
		result.Total = total
		return result
	}
}
