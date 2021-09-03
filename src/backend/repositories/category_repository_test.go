package repositories

import (
	"gitee.com/itsos/golibs/v2/db/mysql"
	"gitee.com/itsos/studynotes/datamodels"
	"testing"
)

func Test_categoryRepository_Delete(t *testing.T) {
	var (
		expected = true
		name     = "name"
		ctype    = CategoryTypeTag
		info     = "info"
	)
	mysql.NewMysql().Exec("truncate table category")
	id := RCategory.Insert(&datamodels.Category{Name: name, Type: ctype, Info: info})
	actual := RCategory.Delete(id)
	if actual != expected {
		t.Errorf("RCategory.Delete(%d) = %t; expected %t", id, actual, expected)
	}
}

func Test_categoryRepository_ExistName(t *testing.T) {
	var (
		expected = true
		name     = "name"
		ctype    = CategoryTypeTag
		info     = "info"
	)
	mysql.NewMysql().Exec("truncate table category")
	id := RCategory.Insert(&datamodels.Category{Name: name, Type: ctype, Info: info})
	actual := RCategory.ExistName(name, ctype)
	if actual != expected {
		t.Errorf("RCategory.ExistName(%d) = %t; expected %t", id, actual, expected)
	}
}

func Test_categoryRepository_Select(t *testing.T) {
	var (
		expected = true
		name     = "name"
		ctype    = CategoryTypeTag
		info     = "info"
	)
	mysql.NewMysql().Exec("truncate table category")
	RCategory.Insert(&datamodels.Category{Name: name, Type: ctype, Info: info})
	res, actual := RCategory.Select(&datamodels.Category{Name: name})
	if actual != expected {
		t.Errorf("RCategory.Select(%v) = %v,%t; expected %t", &datamodels.Category{Name: name}, res, actual, expected)
	}
}

func Test_categoryRepository_SelectMany(t *testing.T) {
	var (
		expected = 1
		name     = "name"
		ctype    = CategoryTypeTag
		info     = "info"
	)
	mysql.NewMysql().Exec("truncate table category")
	RCategory.Insert(&datamodels.Category{Name: name, Type: ctype, Info: info})
	actual := RCategory.SelectMany()
	if len(actual) != expected {
		t.Errorf("RCategory.SelectMany() = %v; expected %d", actual, expected)
	}
}

func Test_categoryRepository_SelectManyByType(t *testing.T) {
	var (
		expected = 1
		name     = "name"
		ctype    = CategoryTypeTag
		info     = "info"
	)
	mysql.NewMysql().Exec("truncate table category")
	RCategory.Insert(&datamodels.Category{Name: name, Type: ctype, Info: info})
	actual := RCategory.SelectManyByType(ctype)
	if len(actual) != expected {
		t.Errorf("RCategory.SelectManyByType(%d) = %v; expected %d", ctype, actual, expected)
	}
}

func Test_categoryRepository_Update(t *testing.T) {
	var (
		expected = "name1"
		name     = "name"
		ctype    = CategoryTypeTag
		info     = "info"
		update   = datamodels.Category{Name: expected}
	)
	mysql.NewMysql().Exec("truncate table category")
	id := RCategory.Insert(&datamodels.Category{Name: name, Type: ctype, Info: info})
	RCategory.Update(id, &update)
	res, _ := RCategory.Select(&datamodels.Category{Id: id})
	if res.Name != expected {
		t.Errorf("RCategory.Update(%d, %v) = %s; expected %s", id, update, res.Name, expected)
	}
}
