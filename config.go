package gocfg

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
)

type Cfg struct {
	Name          string
	Path          string
	DefalutConfig interface{}
}

var Config *Cfg

func init() {
	Config = new(Cfg)
	Config.Name = "config.json"
}

func (c *Cfg) SetPath(p string) {
	c.Path = p
}

func (c *Cfg) SetName(n string) {
	c.Name = n
}

func (c *Cfg) SetDefaultConfig(v interface{}) {
	c.DefalutConfig = v
}

func (c *Cfg) LoadDefault(v interface{}) {
	if c.DefalutConfig != nil {
		c.Save(c.DefalutConfig)
		c.Load(v)
	} else {
		log.Println("没有设置默认配置！")
	}
}

func (c *Cfg) Load(v interface{}) {
	d, err := ioutil.ReadFile(filepath.Join(c.Path, c.Name))
	if err != nil {
		log.Println("没有发现json文件！如果存在默认配置，将使用默认配置。")
		c.LoadDefault(v)
		return
	}
	err = json.Unmarshal(d, v)
	if err != nil {
		log.Println("解析json文件失败！如果存在默认配置，将使用默认配置。")
		c.LoadDefault(v)
	}
}

func (c *Cfg) Save(v interface{}) {
	c.Backup(v, c.Name, c.Path)
}

func (c *Cfg) Backup(v interface{}, Name, Path string) {
	d, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Println("编码json文件出错！")
	}
	err = ioutil.WriteFile(filepath.Join(Path, Name), d, 0666)
	if err != nil {
		log.Println("保存json文件出错！")
	}
}
