package generater

import (
	"fmt"
	"os"
	"strings"
)

//HandleGenerateAtoms ...
func HandleGenerateAtoms(name, ns string) {
	var fileName = "./dist/" + ns + ".Atoms/" + name + ".cs"
	var code = strings.Replace(templeteAtoms(), "{ns}", ns, -1)
	code = strings.Replace(code, "{name}", name, -1)

	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString(code)
}

//HandleGenerateStorages ...
func HandleGenerateStorages(name, ns string) {
	var fileName = "./dist/" + ns + ".Storages/" + name + "Storage.cs"
	var code = strings.Replace(templeteStorages(), "{ns}", ns, -1)
	code = strings.Replace(code, "{name}", name, -1)

	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString(code)
}

//HandleGenerateSteps ..
func HandleGenerateSteps(name, ns string) {
	var fileName = "./dist/" + ns + ".Steps/" + name + "Step.cs"
	var code = strings.Replace(templeteSteps(), "{ns}", ns, -1)
	code = strings.Replace(code, "{name}", name, -1)

	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString(code)
}

//HandleGenerateServices ...
func HandleGenerateServices(name, ns string) {
	var fileName = "./dist/" + ns + ".Services/" + name + "Service.cs"
	var code = strings.Replace(templeteServices(), "{ns}", ns, -1)
	code = strings.Replace(code, "{name}", name, -1)

	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString(code)
}

//HandleGenerateControllers ...
func HandleGenerateControllers(name, ns string) {
	var fileName = "./dist/" + ns + ".Web/Controllers/" + name + "Controller.cs"
	var code = strings.Replace(templeteControllers(), "{ns}", ns, -1)
	code = strings.Replace(code, "{name}", name, -1)

	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString(code)
}

//HandleGenerateCore ...
func HandleGenerateCore(ns string) {
	var fileName = "./dist/" + ns + ".Storages/StorageCore/DbCtx.cs"
	var code = strings.Replace(templeteCore(), "{ns}", ns, -1)

	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString(code)
}

func templeteStorages() string {
	return `using System;
using System.Collections.Generic;
using System.Data.SqlClient;
using Orm.Son.Core;
using {ns}.Atoms;
using {ns}.Storages.StorageCore;

namespace {ns}.Storages
{
	public class {name}Storage
	{
		public int Add({name} entity)
		{
			using (var db = new DbCtx())
			{
                return db.Insert(entity);
			}
		}

        public int Del(int id)
        {
            using (var db = new DbCtx())
            {
                return db.Delete<{name}>(id);
            }
        }

        public int Edit({name} entity)
        {
            using (var db = new DbCtx())
            {
                return db.Update(entity);
            }
        }

        public {name} GetById(int id)
        {
            using (var db = new DbCtx())
            {
                return db.Find<{name}>(id);
            }
        }
    }
}`

}

func templeteAtoms() string {
	return `using System;
using Orm.Son.Mapper;

namespace {ns}.Atoms
{
	public class {name}
	{
	}
}`
}

func templeteSteps() string {
	return `using System;
using System.Collections.Generic;
using System.Data.SqlClient;
using Orm.Son.Core;
using {ns}.Storages;

namespace {ns}.Steps
{
	public class {name}Step
	{
	}
}`
}

func templeteServices() string {
	return `using System;
using System.Collections.Generic;
using System.Data.SqlClient;
using {ns}.Steps;

namespace {ns}.Services
{
	public class {name}Service
	{
	}
}`
}

func templeteControllers() string {
	return `using System;
using System.Collections.Generic;
using System.Data.SqlClient;
using System.Web.Mvc;
using {ns}.Services;

namespace {ns}.Controllers 
{
	public class {name}Controller : Controller
	{
	}
}`
}

func templeteCore() string {
	return `using System;
using Orm.Son.Core;

namespace {ns}.Storages.StorageCore
{
	public class DbCtx : SonConnection
	{
		public DbCtx() : base("connString")
		{
		}
	}
}`
}
