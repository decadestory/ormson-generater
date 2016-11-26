package main

import (
	"flag"
	"fmt"
	"ormson/generater"
	oio "ormson/io"
	"os"
)

var g string

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Input Like: ormson -g Test")
		return
	}

	flag.StringVar(&g, "g", "empty", "输入表名")
	flag.Parse()
	steps(g)
	fmt.Println("---------------------")
	fmt.Println("Generate Complete !!")
	fmt.Println("---------------------")
}

func steps(name string) {
	//生成文件夹 Controllers Services Steps Storages Atoms
	initDir()
	//处理生成
	generater.HandleGenerateAtoms(name)
	generater.HandleGenerateStorages(name)
	generater.HandleGenerateSteps(name)
	generater.HandleGenerateServices(name)
	generater.HandleGenerateControllers(name)
	generater.HandleGenerateCore()
}

func initDir() {
	var dist = "./dist/"
	ok := oio.CreateDir(dist)
	if ok {
		fmt.Println(dist)
	}

	dirPaths := map[string]string{
		"Controller":  dist + "Controllers/",
		"Service":     dist + "Services/",
		"Step":        dist + "Steps/",
		"Storage":     dist + "Storages/",
		"Atom":        dist + "Atoms/",
		"StorageCore": dist + "Storages/StorageCore/",
	}

	for _, v := range dirPaths {
		ok := oio.CreateDir(v)
		if ok {
			fmt.Println(v)
		}
	}
}
