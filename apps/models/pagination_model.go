package models

import (
	"strings"
)

type PaginationInterface interface {
	GetSearchingBy(*Pagination) []map[string]string
	GetSorting(*Pagination) []map[string]string
}

type Pagination struct {
	Size      int
	Page      int
	PageCount int
	RowCount  int
	SortBy    []string
	SearchBy  []string
}

func (Pagination) GetSearchingBy(paging *Pagination) []map[string]string {
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

func (Pagination) GetSorting(paging *Pagination) []map[string]string {
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
