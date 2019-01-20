package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type TuPath struct {
	Id int64
	TuName string
	TuPath string
	Md5 string
	Status int
}

func init()  {
	// register model
	orm.RegisterModel(new(TuPath))
}

// TableName 获取对应数据库表名.
func (tu *TuPath) TableName() string {
	return "tu_path"
}

// TableEngine 获取数据使用的引擎.
func (tu *TuPath) TableEngine() string {
	return "INNODB"
}

func NewTuPath() *TuPath {
	return &TuPath{}
}

//获取箱码信息
func (tu *TuPath) Find(filter string, param string) ([]orm.Params, error) {
	var tus []orm.Params
	tableName := tu.TableName()

	o := orm.NewOrm()
	num, err := o.Raw("SELECT * FROM " + tableName + " where "+filter+" = " + param).Values(&tus)
	
    if err != nil {
        beego.Debug("数据查询错误", err)
    } else {
        beego.Debug(filter+"：", param, "数量：", num)
    }

	return tus, err
}

func (tu *TuPath) Inserts() (id int64)  {
	o := orm.NewOrm()

	id, err := o.Insert(tu)
	if err == nil {
		return id
	}
	return 0
}

func (tu *TuPath) Update(id int64) (error)  {
	o := orm.NewOrm()

	tu.Id = id
	//以处理
	tu.Status = 3
	_, err := o.Update(tu)
	if err != nil {
		return err
	}
	return nil
}

