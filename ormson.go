package main

import (
	"fmt"
	"ormson/generater"
	oio "ormson/io"
)

func main() {
	steps()
}

func steps() {
	//生成文件夹 Controllers Services Steps Storages Atoms
	initDir()
	//处理生成
	generater.HandleGenerateAtoms("Admin")
	generater.HandleGenerateStorages("Admin")
	generater.HandleGenerateSteps("Admin")
	generater.HandleGenerateServices("Admin")
	generater.HandleGenerateControllers("Admin")
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
