package logic

import (
	"fmt"
	"strings"
)

func CreateInserts(table string, columns []string, values [][]string) []string {
	var sqls []string
	for _, val := range values {
		sqls = append(sqls, fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s);", table, columnsToString(columns), valuesToString(val)))
	}
	return sqls
}

// カラムをpostgresで格納可能なstringに変換します
func columnsToString(columns []string) string {
	return strings.Join(columns, ",")
}

// 値をpostgresで格納可能なstringに変換します
func valuesToString(vals []string) string {
	for i := range vals {
		if strings.HasPrefix(vals[i], "'") && strings.HasSuffix(vals[i], "'") {
			vals[i] = fmt.Sprintf("'\\\\'%s\\\\''", vals[i])
		}else if vals[i] == "" {
			vals[i] = "NULL"
		} else {
			vals[i] = fmt.Sprintf("%s", vals[i])
		}
	}
	return strings.Join(vals, ",")
}
