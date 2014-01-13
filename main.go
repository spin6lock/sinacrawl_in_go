package main
import (
	"fmt"
	"bufio"
	"os"
	"log"
	"net/http"
	"io/ioutil"
)

var account string
var password string

func ReadConfig() {
	file, err := os.Open("config.cfg")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	account = scanner.Text();
	scanner.Scan()
	password = scanner.Text();
}

func main() {
	ReadConfig()
	fmt.Println(account)
	resp, _ := http.Get("http://login.weibo.cn/login/?ns=1&revalid=2")
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", string(body))
}
