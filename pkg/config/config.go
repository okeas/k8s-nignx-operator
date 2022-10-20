package config

import (
	"fmt"
	"io/ioutil"
	v1 "k8s.io/api/networking/v1"
	"log"
	"sigs.k8s.io/yaml"
)

// server
type Server struct {
	Port int	// 代理启动的端口
}

// yaml文件对象
type SysConfigStruct struct {
	Server Server
	//Ingress v1.IngressSpec
	Ingress []v1.Ingress
}

var SysConfig = new(SysConfigStruct)

func InitConfig() {
	config, err := ioutil.ReadFile("./app.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(config, SysConfig)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(SysConfig.Ingress)
	ParseRule()
}