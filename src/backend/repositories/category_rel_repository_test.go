package repositories

import (
	"gitee.com/itsos/golibs/v2/db/mysql"
	"gitee.com/itsos/studynotes/datamodels"
	"testing"
)

func Test_categoryRelRepository_Delete(t *testing.T) {
	var (
		expected = true
		in1      = uint(1)
		in2      = uint(1)
	)
	mysql.NewMysql().Exec("truncate table category_rel")
	id := RCategoryRel.Insert(&datamodels.CategoryRel{Cid: in1, Aid: in2})
	actual := RCategoryRel.Delete(id)
	if actual != expected {
		t.Errorf("RCategoryRel.Delete(%d) = %t; expected %t", id, actual, expected)
	}
}

func Test_categoryRelRepository_DeleteByAid(t *testing.T) {
	var (
		expected = true
		in1      = uint(1)
		in2      = uint(1)
	)
	mysql.NewMysql().Exec("truncate table category_rel")
	RCategoryRel.Insert(&datamodels.CategoryRel{Cid: in1, Aid: in2})
	actual := RCategoryRel.DeleteByAid(in2)
	if actual != expected {
		t.Errorf("RCategoryRel.DeleteByAid(%d) = %t; expected %t", in2, actual, expected)
	}
}

func Test_categoryRelRepository_DeleteByAidAndCid(t *testing.T) {
	var (
		expected = true
		in1      = uint(1)
		in2      = uint(1)
	)
	mysql.NewMysql().Exec("truncate table category_rel")
	RCategoryRel.Insert(&datamodels.CategoryRel{Cid: in1, Aid: in2})
	actual := RCategoryRel.DeleteByAidAndCid(in1, in2)
	if actual != expected {
		t.Errorf("RCategoryRel.DeleteByAidAndCid(%d, %d) = %t; expected %t", in1, in2, actual, expected)
	}
}

func Test_categoryRelRepository_DeleteByCid(t *testing.T) {
	var (
		expected = true
		in1      = uint(1)
		in2      = uint(1)
	)
	mysql.NewMysql().Exec("truncate table category_rel")
	RCategoryRel.Insert(&datamodels.CategoryRel{Cid: in1, Aid: in2})
	actual := RCategoryRel.DeleteByAidAndCid(in1, in2)
	if actual != expected {
		t.Errorf("RCategoryRel.DeleteByAidAndCid(%d, %d) = %t; expected %t", in1, in2, actual, expected)
	}
}

func Test_categoryRelRepository_GetCountByCid(t *testing.T) {
	var (
		expected = uint(1)
		in1      = uint(1)
		in2      = uint(1)
	)
	mysql.NewMysql().Exec("truncate table category_rel")
	RCategoryRel.Insert(&datamodels.CategoryRel{Cid: in1, Aid: in2})
	actual := RCategoryRel.GetCountByCid(in1)
	if actual != expected {
		t.Errorf("RCategoryRel.GetCountByCid(%d) = %d; expected %d", in1, actual, expected)
	}
}

func Test_categoryRelRepository_Select(t *testing.T) {
	var (
		expected = uint(1)
		in1      = uint(1)
		in2      = uint(1)
	)
	mysql.NewMysql().Exec("truncate table category_rel")
	RCategoryRel.Insert(&datamodels.CategoryRel{Cid: in1, Aid: in2})
	actual, _ := RCategoryRel.Select(&datamodels.CategoryRel{Cid: in1, Aid: in2})

	if actual.Aid != expected {
		t.Errorf("RCategoryRel.Select(%v) = %v; expected aid=%v", &datamodels.CategoryRel{Cid: in1, Aid: in2}, actual, expected)
	}
}

func Test_categoryRelRepository_SelectManyByAid(t *testing.T) {
	var (
		expected = true
		in1      = uint(1)
		in2      = uint(1)
	)
	RCategoryRel.Insert(&datamodels.CategoryRel{Cid: in1, Aid: in2})
	actual := RCategoryRel.DeleteByAidAndCid(in1, in2)
	if actual != expected {
		t.Errorf("RCategoryRel.DeleteByAidAndCid(%d, %d) = %t; expected %t", in1, in2, actual, expected)
	}
}

func Test_categoryRelRepository_SelectManyByCid(t *testing.T) {
	var (
		expected = true
		in1      = uint(1)
		in2      = uint(1)
	)
	RCategoryRel.Insert(&datamodels.CategoryRel{Cid: in1, Aid: in2})
	actual := RCategoryRel.DeleteByAidAndCid(in1, in2)
	if actual != expected {
		t.Errorf("RCategoryRel.DeleteByAidAndCid(%d, %d) = %t; expected %t", in1, in2, actual, expected)
	}
}

func Test_categoryRelRepository_Update(t *testing.T) {
	var (
		expected = true
		in1      = uint(1)
		in2      = uint(1)
	)
	RCategoryRel.Insert(&datamodels.CategoryRel{Cid: in1, Aid: in2})
	actual := RCategoryRel.DeleteByAidAndCid(in1, in2)
	if actual != expected {
		t.Errorf("RCategoryRel.DeleteByAidAndCid(%d, %d) = %t; expected %t", in1, in2, actual, expected)
	}
}
