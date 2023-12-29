package tfac

import (
	"fmt"
	"vultr-stat/pkg/lib/date"

	"vultr-stat/pkg/client"
)

type InstanceTableFactory struct {
}

func NewInstanceTableFactory() InstanceTableFactory {
	return InstanceTableFactory{}
}

func (f *InstanceTableFactory) Columns() []string {
	return []string{
		"STATUS", "POWER", "IP", "REGION", "CREATED",
		"CORE", "RAM", "DISK", "BANDWIDTH", "OS",
	}
}

func (f *InstanceTableFactory) Rows(instances []*client.Instance) [][]string {
	var rows [][]string
	for _, instance := range instances {
		create, err := date.ByRFC3339String(instance.DateCreated)
		if err != nil {
			panic(err)
		}
		row := []string{
			instance.Status,
			instance.PowerStatus,
			instance.MainIp,
			instance.Region,
			date.ToPrettyString(create),
			fmt.Sprintf("%v", instance.VcpuCount),
			fmt.Sprintf("%vMB", instance.Ram),
			fmt.Sprintf("%vGB", instance.Disk),
			fmt.Sprintf("%vTB", instance.AllowedBandwidth),
			instance.Os,
		}
		rows = append(rows, row)
	}
	return rows
}
