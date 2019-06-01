package paging

import (
	"demoGo/libraries"
	"github.com/go-xorm/xorm"
	"math"
	"strings"
)

type PaginationInterface interface {
	SetQueryBuilder(paging *Pagination, counter *xorm.Session, query *xorm.Session, criteria []string) (*Pagination, *xorm.Session, *xorm.Session)
	GetSearchingBy(paging *Pagination) []map[string]string
	GetSorting(paging *Pagination) []map[string]string
	GetRowCount(pagination *Pagination) int64
}

type Pagination struct {
	Size      int         `json:"size"`
	Page      int         `json:"page"`
	PageCount int64       `json:"page_count"`
	RowCount  int64       `json:"row_count"`
	SortBy    []string    `json:"sort_by"`
	SearchBy  []string    `json:"search_by"`
	Rows      interface{} `json:"rows"`
}

var (
	pagingInfo Pagination
)

type Paging struct {
}

func GetPage() PaginationInterface {
	pagination := &Paging{}
	return pagination
}

func (p Paging) GetSearchingBy(paging *Pagination) []map[string]string {
	var searchBy []map[string]string
	for _, row := range paging.SearchBy {
		mapSearch := make(map[string]string)
		search := strings.Split(row, ":")
		if len(search) <= 1 {
			searchBy = append(searchBy, nil)
		} else {
			key := search[0]
			value := search[1]
			mapSearch[key] = value
			searchBy = append(searchBy, mapSearch)
		}
	}
	return searchBy
}

func (p Paging) GetSorting(paging *Pagination) []map[string]string {
	var sortBy []map[string]string
	for _, row := range paging.SortBy {
		mapSearch := make(map[string]string)
		search := strings.Split(row, ":")
		if len(search) <= 1 {
			sortBy = append(sortBy, nil)
		} else {
			key := search[0]
			value := search[1]
			mapSearch[key] = value
			sortBy = append(sortBy, mapSearch)
		}
	}
	return sortBy
}

/*
injecting the additional query for queryStatement for pagination feature
*/
func (p Paging) SetQueryBuilder(paging *Pagination, counter *xorm.Session,
	queryStatement *xorm.Session, criteria []string) (*Pagination, *xorm.Session, *xorm.Session) {
	var size, offset int = 10, 1
	for _, searchValue := range GetPage().GetSearchingBy(paging) {
		for key, value := range searchValue {
			if libraries.Contains(criteria, key) {
				queryStatement.Where(key+" like ?", "%"+value+"%")
				counter.Where(key+" like ?", "%"+value+"%")
			}
		}
	}
	for _, sortBy := range GetPage().GetSorting(paging) {
		for key, value := range sortBy {
			if libraries.Contains(criteria, key) {
				if value == "desc" {
					queryStatement.Desc(key)
				} else {
					queryStatement.Asc(key)
				}
			}
		}
	}
	if paging.Size < 0 {
		paging.Size = size
	}
	/*
		offseting query with paging and limiting
	*/
	offset = paging.Size * (paging.Page - 1)
	queryStatement.Limit(paging.Size, offset)
	return paging, counter, queryStatement
}

func (p Paging) GetRowCount(pagination *Pagination) int64 {
	return int64(math.Ceil(float64(pagination.RowCount) / float64(pagination.Size)))
}
