package main
import (
	"fmt"
	"bufio"
	"os"
	"log"
	"net/http"
	"io/ioutil"
	"regexp"
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
	//fmt.Printf("%s\n", string(body))
	input_box_name_pattern := regexp.MustCompile(`password_[0-9]+`)
	input_box_name := string(input_box_name_pattern.Find([]byte(body)))
	vk_box_value_pattern := regexp.MustCompile(`vk" value="([0-9]+_[a-z0-9]+_[0-9]+)"`)
	vk_box_value := vk_box_value_pattern.FindAllStringSubmatch(string(body), -1)[0][1]
	fmt.Println(input_box_name)
	fmt.Println(vk_box_value)
}
