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
		"IDX",
		"AVAILABLE",
		"IP",
		"ID",
		"CREATED",
		"POWER",
		"STATUS",
	}
}

func (f *InstanceTableFactory) SimpleRows(instances []*client.InstanceInfo) [][]string {
	var rows [][]string
	for i, instance := range instances {
		row := []string{
			fmt.Sprintf("%v", i+1),
			fmt.Sprintf("%vGB", instance.AllowedBandwidth-*instance.UsedBandwidth),
			instance.MainIp,
			strings.Split(instance.Id, "-")[0] + "...",
			date.ToPrettyString(instance.DateCreated),
			instance.PowerStatus,
			instance.Status,
		}
		rows = append(rows, row)
	}
	return rows
}

func (f *InstanceTableFactory) DetailColumns() []string {
	return []string{
		"IDX",
		"AVAILABLE",
		"ALLOWED",
		"USAGE",
		"IP",
		"ID",
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
	for i, instance := range instances {
		row := []string{
			fmt.Sprintf("%v", i+1),
			fmt.Sprintf("%vGB", instance.AllowedBandwidth-*instance.UsedBandwidth),
			fmt.Sprintf("%vGB", instance.AllowedBandwidth),
			fmt.Sprintf("%vGB", *instance.UsedBandwidth),
			instance.MainIp,
			instance.Id,
			date.ToPrettyString(instance.DateCreated),
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
