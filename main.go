package main

import (
	_ "demo/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

//type CommentItem struct{
//	UserName string
//	Score float64
//	GameLength string
//	Comment string
//}

func init() {
	//beego配置文件设置
	beego.BConfig.RunMode = "dev"
	//数据库配置
	//orm.RegisterDataBase("default", "mysql", "root:my110120130@tcp(127.0.0.1:3306)/test")
}

func main() {
	//beego orm操作示例

	//o := orm.NewOrm()
	//o.Using("default")
	//
	////test1 :=new(models.Testormtable)
	////test1.Test1 = "test1 col1"
	////test1.Test2 = "test2 col2"
	////
	////fmt.Print(o.Insert(test1))
	////o.Delete(&models.Testormtable{Id : 1})
	//
	////beego orm操作，以下操作为基础CRUD，大部分以主键为条件操作
	//test2 := new(models.Testormtable)
	//test2.Test1 = "orm operate test1"
	//test2.Test2 = "orm operate test2"
	////数据库添加条目，返回当前条目的主键，出现错误则传入err
	//tid, err := o.Insert(test2)
	////若没有出现错误则获取新增数据的主键id
	//if err == nil {
	//	fmt.Println(tid)
	//	test2.Id = int(tid)
	//}
	////通过主键读取数据库中对应数据信息并打印，直接使用了已声明的err
	//err = o.Read(test2)
	//if err == nil {
	//	fmt.Println(test2.Test1, test2.Test2)
	//}
	//
	//test2.Test1 = "updated col"
	//fmt.Println(o.Update(test2))
	//err = o.Read(test2)
	//if err == nil {
	//	printObject(*test2)
	//}
	//
	//fmt.Println(o.Delete(test2))

	beego.Run()
}

//func printObject(o interface{}) {
//	//go语言反射，获取结构体中的所有值
//	val := reflect.ValueOf(o)
//	for i := 0; i < val.NumField(); i++ {
//		//输出结构体中的数据类型，类型名称，值
//		fmt.Printf("%v %v:%v\n", val.Type().Field(i).Type, val.Type().Field(i).Name, val.Field(i))
//	}
//	//fmt.Println(val)
//}
