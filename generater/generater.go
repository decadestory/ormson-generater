package generater

import (
	"fmt"
	"ormson/dbstruct"
	"os"
	"strings"
)

//HandleGenerateAtoms ...
func HandleGenerateAtoms(name, ns string) {
	var fileName = "./dist/" + ns + ".Atoms/" + name + ".cs"
	var code = strings.Replace(templeteAtoms(), "{ns}", ns, -1)
	code = strings.Replace(code, "{name}", name, -1)
	var text = dbstruct.GetAtomLines(name)
	code = strings.Replace(code, "{text}", text, -1)

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

//HandleGenerateMolecules ...
func HandleGenerateMolecules(name, ns string) {
	var fileName = "./dist/" + ns + ".Molecules/" + name + "Molecule.cs"
	var code = strings.Replace(templeteMolecules(), "{ns}", ns, -1)
	code = strings.Replace(code, "{name}", name+"Molecule", -1)
	var text = dbstruct.GetAtomLines(name)
	code = strings.Replace(code, "{text}", text, -1)

	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString(code)
}

//HandleGenerateLibs ...
func HandleGenerateLibs(ns string) {
	var fileName = "./dist/" + ns + ".Libs/DataResult.cs"
	var code = strings.Replace(templeteLibs(), "{ns}", ns, -1)

	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString(code)
}

//HandleGenerateLibForPage ...
func HandleGenerateLibForPage(ns string) {
	var fileName = "./dist/" + ns + ".Libs/PageResult.cs"
	var code = strings.Replace(templeteLibForPage(), "{ns}", ns, -1)

	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString(code)
}

//HandleGenerateUtils ...
func HandleGenerateUtils(ns string) {
	var fileName = "./dist/" + ns + ".Web/Utils/ApiExceptionAttribute.cs"
	var code = strings.Replace(templeteAPIExceptionAttribute(), "{ns}", ns, -1)

	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString(code)
	handleGenerateMvcUtils(ns)
}

func handleGenerateMvcUtils(ns string) {
	var fileName = "./dist/" + ns + ".Web/Utils/MvcExceptionAttribute.cs"
	var code = strings.Replace(templeteMvcExceptionAttribute(), "{ns}", ns, -1)

	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	f.WriteString(code)
}

func templeteMvcExceptionAttribute() string {
	return `using System.Text;
using System.Web.Mvc;
using {ns}.Libs;

namespace {ns}.Web.Utils
{
    public class MvcExceptionAttribute : ActionFilterAttribute, IExceptionFilter
    {
        public void OnException(ExceptionContext context)
        {
            if (context.Exception != null)
            {
                var error = new DataResult<string>(false, context.Exception.Message);
                context.Result = new JsonResult
                {
                    Data = error,
                    ContentType = "application/json; charset=utf-8",
                    ContentEncoding = Encoding.UTF8
                };
            }

            context.ExceptionHandled = true;
        }
    }
}`
}

func templeteAPIExceptionAttribute() string {
	return `using Newtonsoft.Json;
using {ns}.Libs;
using System;
using System.IO;
using System.Net;
using System.Net.Http;
using System.Text;
using System.Web.Http.Filters;

namespace {ns}.Web.Utils
{
    public class ApiExceptionAttribute : ExceptionFilterAttribute
    {
        public override void OnException(HttpActionExecutedContext actionExecutedContext)
        {
            if (actionExecutedContext.Exception is NotImplementedException)
            {
                actionExecutedContext.Response = new HttpResponseMessage(HttpStatusCode.NotImplemented);
            }
            else if (actionExecutedContext.Exception is TimeoutException)
            {
                actionExecutedContext.Response = new HttpResponseMessage(HttpStatusCode.RequestTimeout);
            }
            else
            {
                var error = new DataResult<string>(false, actionExecutedContext.Exception.Message);
                actionExecutedContext.Response = ToResponseMessage(error);

            }

            base.OnException(actionExecutedContext);
        }

       
        protected HttpResponseMessage ToResponseMessage(DataResult<string> error)
        {
            var respJson = JsonConvert.SerializeObject(error);
            var bs = Encoding.UTF8.GetBytes(respJson);
            var rm = new HttpResponseMessage();
            {
                MemoryStream ms = new MemoryStream(bs);
                rm.Content = new StreamContent(ms);
                rm.Content.Headers.Add("Content-Type", "application/json; charset=utf-8"); 
            }
            return rm;
        }
    }
}`
}

func templeteLibs() string {
	return `namespace {ns}.Libs
{
    public class DataResult<T>
    {
        public bool IsOk { get; set; }
        public T Data { get; set; }
        public string Msg { get; set; }
        public object ExtData { get; set; }

        public DataResult()
        {
            IsOk = true;
        }
        public DataResult(T data)
        {
            IsOk = true;
            Data = data;
        }
        public DataResult(T data, object extData)
        {
            IsOk = true;
            Data = data;
            ExtData = extData;
        }
        public DataResult(bool isOk, string msg)
        {
            IsOk = isOk;
            Msg = msg;
        }
        public DataResult(bool isOk, string msg, object extData)
        {
            IsOk = isOk;
            Msg = msg;
            ExtData = extData;
        }

    }
}
`
}

func templeteLibForPage() string {
	return `using System.Collections.Generic;
    
namespace {ns}.Libs
{
    public class PageResult<T>
    {
        public int Total { get; set; }
        public List<T> DataList { get; set; }
    }
}
`
}

func templeteMolecules() string {
	return `using System;

namespace {ns}.Molecules
{
	public class {name}
	{
{text}
	}
}`
}

func templeteStorages() string {
	return `using System;
using System.Collections.Generic;
using System.Data.SqlClient;
using Orm.Son.Core;
using {ns}.Atoms;
using {ns}.Storages.StorageCore;
using {ns}.Molecules;

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

        public Tuple<List<{name}Molecule>, int> FindPage(int page, int pageSize,string keyWord)
        {
            using (var db = new DbCtx())
            {
                if (keyWord == null) keyWord = "";
                return db.FindPage<{name}, {name}Molecule>(t => t.Id > 0, t => t.Id, page, pageSize, true);
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
{text}
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
using {ns}.Atoms;
using {ns}.Steps;
using {ns}.Storages;
using {ns}.Libs;
using {ns}.Molecules;
using Orm.Son.Mapper;

namespace {ns}.Services
{
    public class {name}Service
    {
        {name}Storage store = new {name}Storage();

        public DataResult<string> Add({name}Molecule model)
        {
            var entity = EntityMapper.Mapper<{name}Molecule, {name}>(model);
            entity.AddTime = DateTime.Now;
            entity.EditTime = DateTime.Now;
            return store.Add(entity) > 0
                ? new DataResult<string>(true,"添加成功")
                : new DataResult<string>(false, "添加失败");
        }

        public DataResult<string> Del(int id)
        {
            return store.Del(id) > 0
                ? new DataResult<string>(true, "删除成功")
                : new DataResult<string>(false, "删除失败");
        }

        public DataResult<string> Edit({name}Molecule model)
        {
            var entity = EntityMapper.Mapper<{name}Molecule, {name}>(model);
            entity.EditTime = DateTime.Now;
            return store.Edit(entity) > 0
                ? new DataResult<string>(true, "更新成功")
                : new DataResult<string>(false, "更新失败");
        }

        public DataResult<{name}Molecule> GetById(int id)
        {
            var reslut = store.GetById(id);
            var entity = EntityMapper.Mapper<{name},{name}Molecule>(reslut);
            return new DataResult<{name}Molecule>(entity);
        }

        public DataResult<PageResult<{name}Molecule>> FindPage(int page,int pageSize,string keyWord)
        {
            var reslut = store.FindPage(page, pageSize, keyWord);
            return new DataResult<PageResult<{name}Molecule>>
            {
                Data = new PageResult<{name}Molecule> { Total = reslut.Item2, DataList = reslut.Item1 }
            };
        }
    }
}`
}

func templeteControllers() string {
	return `using System;
using System.Web.Http;
using {ns}.Services;
using {ns}.Molecules;
using {ns}.Libs;
using System.Collections.Generic;
using {ns}.Web.Utils;

namespace {ns}.Web.Controllers 
{
	public class {name}Controller : ApiController
    {
        {name}Service service = new {name}Service();

		[HttpPost,ApiException]
        public DataResult<string> Add({name}Molecule model)
        {
            return service.Add(model);
        }

		[HttpPost,ApiException]
        public DataResult<string> Del(int id)
        {
            return service.Del(id);
        }
		
		[HttpPost,ApiException]
        public DataResult<string> Edit({name}Molecule model)
        {
            return service.Edit(model);
        }

		[HttpGet,ApiException]
        public DataResult<{name}Molecule> GetById(int id)
        {
            return service.GetById(id);
        }

		[HttpGet,ApiException]
        public DataResult<PageResult<{name}Molecule>> FindPage(int id = 1,int pageSize=10, string keyWord="")
        {
            return service.FindPage(id, pageSize, keyWord);
        }
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
