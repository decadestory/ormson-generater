package generater

import (
	"fmt"
	"os"
)

var n1 = "\r\n"
var n2 = "\r\n\r\n"
var nt = "\t"
var ns = "OrmTest"

func HandleGenerateAtoms(name string) {
	var fileName = "./dist/Atoms/" + name + ".cs"
	var code = "using System;" + n1
	code += "using Orm.Son.Mapper;" + n2
	code += "namespace " + ns + ".Atoms" + n1 + "{" + n1
	code += nt + "public class " + name + n1 + nt + "{" + n1
	code += nt + "}" + n1
	code += "}"
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString(code)
}

func HandleGenerateStorages(name string) {
	var fileName = "./dist/Storages/" + name + "Storage.cs"
	var code = "using System;" + n1
	code += "using System.Collections.Generic;" + n1
	code += "using System.Data.SqlClient;" + n1
	code += "using Orm.Son.Core;" + n2
	code += "namespace " + ns + ".Storages" + n1 + "{" + n1
	code += nt + "public class " + name + "Storage" + n1 + nt + "{" + n1
	code += nt + "}" + n1
	code += "}"
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString(code)
}

func HandleGenerateSteps(name string) {
	var fileName = "./dist/Steps/" + name + "Step.cs"
	var code = "using System;" + n1
	code += "using System.Collections.Generic;" + n1
	code += "using System.Data.SqlClient;" + n1
	code += "using Orm.Son.Core;" + n1
	code += "using " + ns + ".Storages;" + n2
	code += "namespace " + ns + ".Steps" + n1 + "{" + n1
	code += nt + "public class " + name + "Step" + n1 + nt + "{" + n1
	code += nt + "}" + n1
	code += "}"
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString(code)
}

func HandleGenerateServices(name string) {
	var fileName = "./dist/Services/" + name + "Service.cs"
	var code = "using System;" + n1
	code += "using System.Collections.Generic;" + n1
	code += "using System.Data.SqlClient;" + n1
	code += "using " + ns + ".Steps;" + n2
	code += "namespace " + ns + ".Services" + n1 + "{" + n1
	code += nt + "public class " + name + "Service" + n1 + nt + "{" + n1
	code += nt + "}" + n1
	code += "}"
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString(code)
}

func HandleGenerateControllers(name string) {
	var fileName = "./dist/Controllers/" + name + "Controller.cs"
	var code = "using System;" + n1
	code += "using System.Collections.Generic;" + n1
	code += "using System.Data.SqlClient;" + n1
	code += "using " + ns + ".Services;" + n2
	code += "namespace " + ns + ".Controllers" + n1 + "{" + n1
	code += nt + "public class " + name + "Controller" + n1 + nt + "{" + n1
	code += nt + "}" + n1
	code += "}"
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString(code)
}

func HandleGenerateCore() {
	var fileName = "./dist/Storages/StorageCore/DbCtx.cs"
	var code = "using System;" + n1
	code += "using Orm.Son.Core;" + n2
	code += "namespace " + ns + ".Storages.StorageCore" + n1 + "{" + n1
	code += nt + "public class DbCtx : SonConnection" + n1 + nt + "{" + n1
	code += nt + nt + "public DbCtx() : base(\"connString\")" + n1 + nt + nt + "{" + n1
	code += nt + nt + "}" + n1
	code += nt + "}" + n1
	code += "}"
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString(code)
}
