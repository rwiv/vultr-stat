package tfac

import (
	"fmt"

	"vultr-stat/pkg/client"
)

type OsTableFactory struct {
}

func NewOsTableFactory() OsTableFactory {
	return OsTableFactory{}
}

func (f *OsTableFactory) Columns() []string {
	return []string{"ID", "NAME", "ARCH", "FAMILY"}
}

func (f *OsTableFactory) Rows(osList []*client.OS) [][]string {
	var rows [][]string
	for _, osInfo := range osList {
		row := []string{
			fmt.Sprintf("%v", osInfo.Id),
			osInfo.Name,
			osInfo.Arch,
			osInfo.Family,
		}
		rows = append(rows, row)
	}
	return rows
}
