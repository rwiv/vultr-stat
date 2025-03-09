package tfac

import (
	"fmt"
	"strings"

	"github.com/rwiv/vultr-stat/pkg/client"
	"github.com/rwiv/vultr-stat/pkg/lib/date"
)

type InstanceTableFactory struct {
}

func NewInstanceTableFactory() InstanceTableFactory {
	return InstanceTableFactory{}
}

func (f *InstanceTableFactory) SimpleColumns() []string {
	return []string{
		"ID",
		"IP",
		"ALLOWED",
		"USAGE",
		"CREATED",
		"POWER",
		"STATUS",
	}
}

func (f *InstanceTableFactory) SimpleRows(instances []*client.InstanceInfo) [][]string {
	var rows [][]string
	for _, instance := range instances {
		createdAt, err := date.ByRFC3339String(instance.DateCreated)
		if err != nil {
			panic(err)
		}
		row := []string{
			strings.Split(instance.Id, "-")[0] + "...",
			instance.MainIp,
			fmt.Sprintf("%vGB", instance.AllowedBandwidth),
			fmt.Sprintf("%vGB", *instance.UsedBandwidth),
			date.ToPrettyString(createdAt),
			instance.PowerStatus,
			instance.Status,
		}
		rows = append(rows, row)
	}
	return rows
}

func (f *InstanceTableFactory) DetailColumns() []string {
	return []string{
		"ID",
		"IP",
		"ALLOWED",
		"USAGE",
		"CREATED",
		"POWER",
		"STATUS",

		"REGION",
		"PLAN",
		"CORE",
		"RAM",
		"DISK",
		"OS",
	}
}

func (f *InstanceTableFactory) DetailRows(instances []*client.InstanceInfo) [][]string {
	var rows [][]string
	for _, instance := range instances {
		createdAt, err := date.ByRFC3339String(instance.DateCreated)
		if err != nil {
			panic(err)
		}
		row := []string{
			instance.Id,
			instance.MainIp,
			fmt.Sprintf("%vGB", instance.AllowedBandwidth),
			fmt.Sprintf("%vGB", *instance.UsedBandwidth),
			date.ToPrettyString(createdAt),
			instance.PowerStatus,
			instance.Status,

			instance.Region,
			instance.Plan,
			fmt.Sprintf("%v", instance.VcpuCount),
			fmt.Sprintf("%vMB", instance.Ram),
			fmt.Sprintf("%vGB", instance.Disk),
			instance.Os,
		}
		rows = append(rows, row)
	}
	return rows
}
