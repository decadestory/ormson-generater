package main

import (
	"flag"
	"fmt"
	_ "go-mssqldb"
	"ormson/generater"
	oio "ormson/io"
	"os"
	"strings"

	_ "github.com/Unknwon/goconfig"
)

var ns string
var g string

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Input Like: ormson -ns OrmSon.Project -g tableName")
		fmt.Println("Input Like: ormson -ns OrmSon.Project -g tableName1,tableName2,tableName3")
		return
	}

	flag.StringVar(&ns, "ns", "OrmSon.Project", "input namespace")
	flag.StringVar(&g, "g", "TableName", "input table name")
	flag.Parse()

	if len(ns) == 0 {
		ns = "OrmSon.Project"
	}

	params := strings.Split(g, ",")
	for _, v := range params {
		steps(v)
	}

	generater.HandleGenerateLibs(ns)
	generater.HandleGenerateUtils(ns)

	// ns = "TestSpace"
	// steps("Role")

	fmt.Println("---------------------")
	fmt.Println("Generate Complete !!")
	fmt.Println("---------------------")
}

func steps(name string) {
	//生成文件夹 Controllers Services Steps Storages Atoms
	initDir()
	//处理生成
	generater.HandleGenerateAtoms(name, ns)
	generater.HandleGenerateStorages(name, ns)
	generater.HandleGenerateSteps(name, ns)
	generater.HandleGenerateServices(name, ns)
	generater.HandleGenerateControllers(name, ns)
	generater.HandleGenerateCore(ns)
	generater.HandleGenerateMolecules(name, ns)

	fmt.Println(name, ",Generate Success")
}

func initDir() {
	var dist = "./dist/"
	ok := oio.CreateDir(dist)
	if !ok {
		fmt.Println("CreateDir Faild:", dist)
	}

	var distWeb = dist + ns + ".Web/"
	okWeb := oio.CreateDir(distWeb)
	if !okWeb {
		fmt.Println("CreateDir Faild:", distWeb)
	}

	dirPaths := map[string]string{
		"Controller":  distWeb + "Controllers/",
		"Util":        distWeb + "Utils/",
		"Service":     dist + ns + ".Services/",
		"Step":        dist + ns + ".Steps/",
		"Storage":     dist + ns + ".Storages/",
		"Atom":        dist + ns + ".Atoms/",
		"StorageCore": dist + ns + ".Storages/StorageCore/",
		"Libs":        dist + ns + ".Libs/",
		"Molecule":    dist + ns + ".Molecules/",
	}

	for _, v := range dirPaths {
		ok := oio.CreateDir(v)
		if !ok {
			fmt.Println("CreateDir Faild:", v)
		}
	}
}
