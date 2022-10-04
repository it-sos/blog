package repositories

import (
	"testing"

	"github.com/it-sos/blog/datamodels"
	"github.com/it-sos/golibs/db/mysql"
)

func Test_categoryRelRepository_Delete(t *testing.T) {
	var (
		expected = true
		cid      = uint(1)
		aid      = uint(2)
	)
	mysql.NewMysql().Exec("truncate table category_rel")
	id := RCategoryRel.Insert(&datamodels.CategoryRel{Cid: cid, Aid: aid})
	actual := RCategoryRel.Delete(id)
	if actual != expected {
		t.Errorf("RCategoryRel.Delete(%d) = %t; expected %t", id, actual, expected)
	}
}

func Test_categoryRelRepository_DeleteByAid(t *testing.T) {
	var (
		expected = true
		cid      = uint(1)
		aid      = uint(2)
	)
	mysql.NewMysql().Exec("truncate table category_rel")
	RCategoryRel.Insert(&datamodels.CategoryRel{Cid: cid, Aid: aid})
	actual := RCategoryRel.DeleteByAid(aid)
	if actual != expected {
		t.Errorf("RCategoryRel.DeleteByAid(%d) = %t; expected %t", aid, actual, expected)
	}
}

func Test_categoryRelRepository_DeleteByAidAndCid(t *testing.T) {
	var (
		expected = true
		cid      = uint(1)
		aid      = uint(2)
	)
	mysql.NewMysql().Exec("truncate table category_rel")
	RCategoryRel.Insert(&datamodels.CategoryRel{Cid: cid, Aid: aid})
	actual := RCategoryRel.DeleteByAidAndCid(aid, cid)
	if actual != expected {
		t.Errorf("RCategoryRel.DeleteByAidAndCid(cid=%d, aid=%d) = %t; expected %t", aid, cid, actual, expected)
	}
}

func Test_categoryRelRepository_DeleteByCid(t *testing.T) {
	var (
		expected = true
		cid      = uint(1)
		aid      = uint(2)
	)
	mysql.NewMysql().Exec("truncate table category_rel")
	RCategoryRel.Insert(&datamodels.CategoryRel{Cid: cid, Aid: aid})
	actual := RCategoryRel.DeleteByAidAndCid(aid, cid)
	if actual != expected {
		t.Errorf("RCategoryRel.DeleteByAidAndCid(aid=%d, cid=%d) = %t; expected %t", aid, cid, actual, expected)
	}
}

func Test_categoryRelRepository_GetCountByCid(t *testing.T) {
	var (
		expected = uint(1)
		cid      = uint(1)
		aid      = uint(2)
	)
	mysql.NewMysql().Exec("truncate table category_rel")
	RCategoryRel.Insert(&datamodels.CategoryRel{Cid: cid, Aid: aid})
	actual := RCategoryRel.GetCountByCid(cid)
	if actual != expected {
		t.Errorf("RCategoryRel.GetCountByCid(%d) = %d; expected %d", cid, actual, expected)
	}
}

func Test_categoryRelRepository_Select(t *testing.T) {
	var (
		expected = uint(2)
		cid      = uint(1)
		aid      = uint(2)
	)
	mysql.NewMysql().Exec("truncate table category_rel")
	RCategoryRel.Insert(&datamodels.CategoryRel{Cid: cid, Aid: aid})
	actual, _ := RCategoryRel.Select(&datamodels.CategoryRel{Cid: cid, Aid: aid})

	if actual.Aid != expected {
		t.Errorf("RCategoryRel.Select(%v) = %v; expected aid=%v", &datamodels.CategoryRel{Cid: cid, Aid: aid}, actual, expected)
	}
}

func Test_categoryRelRepository_SelectManyByAid(t *testing.T) {
	var (
		expected = uint(2)
		cid      = uint(1)
		aid      = uint(2)
	)
	mysql.NewMysql().Exec("truncate table category_rel")
	RCategoryRel.Insert(&datamodels.CategoryRel{Cid: cid, Aid: aid})
	actual := RCategoryRel.SelectManyByAid(aid)
	if actual[0].Aid != expected {
		t.Errorf("RCategoryRel.SelectManyByAid(%d) = %v; expected aid=%d", cid, actual, expected)
	}
}

func Test_categoryRelRepository_SelectManyByCid(t *testing.T) {
	var (
		expected = uint(1)
		cid      = uint(1)
		aid      = uint(2)
	)
	mysql.NewMysql().Exec("truncate table category_rel")
	RCategoryRel.Insert(&datamodels.CategoryRel{Cid: cid, Aid: aid})
	actual := RCategoryRel.SelectManyByCid(cid)
	if actual[0].Cid != expected {
		t.Errorf("RCategoryRel.SelectManyByCid(%d) = %v; expected cid=%d", cid, actual, expected)
	}
}
