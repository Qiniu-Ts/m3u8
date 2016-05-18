package main 

import (
    "fmt"
    "os"
    "bufio"
    "regexp"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "strings"
    "bytes"

    "qiniupkg.com/api.v7/kodo"
	"qiniupkg.com/api.v7/conf"
    "qiniupkg.com/api.v7/kodocli"
)

type Config struct {
	Domain string `json:"domain"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
	Bucket string `json:"bucket"`
	BucketBak string `json:"bucket_bak"`
	KeysFileLoc string `json:"keys_file_loc"`
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("./rdm3u8 <config_file_path>")
		return
	}
	c, err := loadCfg(os.Args[1])
	fmt.Println(c)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	f, err := os.Open(c.KeysFileLoc)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer f.Close()

    r := bufio.NewReader(f)
    scanner := bufio.NewScanner(r)

	for scanner.Scan() {

		key := scanner.Text()
		u := c.url(key)

		data, err := m3u8RmDomain(u)
		if err != nil {
			fmt.Println(err)
	        os.Exit(2)
		}

		if c.BucketBak != "" {
			c.Copy(key)
		}

		err = c.Up(key, data)
		if err != nil {
			fmt.Println(err)
			os.Exit(3)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func loadCfg(path string) (cfg Config, err error) {

	file, err := os.Open(path)
	if err != nil {
		return
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	return
}

func (c *Config) url(key string) string {
	u := c.Domain + "/" + key
	if !strings.HasPrefix(u, "http") {
		u = "http://" + u
	}
	return u
}

func (c *Config) Up(key string, data []byte) error {

    //初始化AK，SK
    conf.ACCESS_KEY = c.AccessKey
    conf.SECRET_KEY = c.SecretKey

    //创建一个Client
    cli := kodo.New(0, nil)

    //设置上传的策略
    policy := &kodo.PutPolicy{
        Scope:   c.Bucket + ":" + key,
        Expires: 3600,
    }

    //生成一个上传token
    token := cli.MakeUptoken(policy);

    //构建一个uploader
    zone := 0
    uploader := kodocli.NewUploader(zone, nil)

    datar := bytes.NewReader(data)

    size := int64(len(data))
    //调用PutFile方式上传，这里的key需要和上传指定的key一致
    return uploader.Put(nil, nil, token, key, datar, size, nil)
}

func (c *Config) Copy(key string) error {

	conf.ACCESS_KEY = c.AccessKey
	conf.SECRET_KEY = c.SecretKey

	//new一个Bucket管理对象
	cli := kodo.New(0, nil)
	p := cli.Bucket(c.Bucket)

	//调用Copy方法移动文件
	return p.Conn.Call(nil, nil, "POST", p.Conn.RSHost + kodo.URICopy(c.Bucket, key, c.BucketBak, key))
}


func m3u8RmDomain(u string)(body []byte, err error) {

	resp, err := http.Get(u)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	exp, err := regexp.CompilePOSIX("^http://[^/]*")
	if err != nil {
		return
	}

	body = exp.ReplaceAll(body, []byte(""))
	return
}
