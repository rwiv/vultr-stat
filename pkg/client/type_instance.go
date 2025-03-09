package client

import (
	"time"

	"github.com/rwiv/vultr-stat/pkg/lib/date"
)

type InstanceResponse struct {
	Instances []*Instance `json:"instances"`
	Meta      *Meta       `json:"meta"`
}

type Instance struct {
	Id               string   `json:"id"`
	Os               string   `json:"os"`
	Ram              int32    `json:"ram"`
	Disk             int32    `json:"disk"`
	MainIp           string   `json:"main_ip"`
	VcpuCount        int32    `json:"vcpu_count"`
	Region           string   `json:"region"`
	Plan             string   `json:"plan"`
	DateCreated      string   `json:"date_created"`
	Status           string   `json:"status"`
	AllowedBandwidth int64    `json:"allowed_bandwidth"`
	NetmaskV4        string   `json:"netmask_v4"`
	GatewayV4        string   `json:"gateway_v4"`
	PowerStatus      string   `json:"power_status"`
	ServerStatus     string   `json:"server_status"`
	V6Network        string   `json:"v6_network"`
	V6MainIp         string   `json:"v6_main_ip"`
	V6NetworkSize    int64    `json:"v6_network_size"`
	Label            string   `json:"label"`
	InternalIp       string   `json:"internal_ip"`
	Kvm              string   `json:"kvm"`
	Hostname         string   `json:"hostname"`
	Tag              string   `json:"tag"`
	Tags             []string `json:"tags"`
	OsId             int64    `json:"os_id"`
	AppId            int64    `json:"app_id"`
	ImageId          string   `json:"image_id"`
	FirewallGroupId  string   `json:"firewall_group_id"`
	Features         []any    `json:"features"`
	UserScheme       string   `json:"user_scheme"`
}

func (ins *Instance) ToInfo(usedBandwidth *int64) (*InstanceInfo, error) {
	createdAt, err := date.ByRFC3339String(ins.DateCreated)
	if err != nil {
		return nil, err
	}

	return &InstanceInfo{
		Id:               ins.Id,
		Os:               ins.Os,
		Ram:              ins.Ram,
		Disk:             ins.Disk,
		MainIp:           ins.MainIp,
		VcpuCount:        ins.VcpuCount,
		Region:           ins.Region,
		Plan:             ins.Plan,
		DateCreated:      createdAt,
		Status:           ins.Status,
		AllowedBandwidth: ins.AllowedBandwidth,
		NetmaskV4:        ins.NetmaskV4,
		GatewayV4:        ins.GatewayV4,
		PowerStatus:      ins.PowerStatus,
		ServerStatus:     ins.ServerStatus,
		OsId:             ins.OsId,
		AppId:            ins.AppId,
		ImageId:          ins.ImageId,
		FirewallGroupId:  ins.FirewallGroupId,
		UsedBandwidth:    usedBandwidth,
	}, nil
}

type InstanceInfo struct {
	Id               string    `json:"id"`
	Os               string    `json:"os"`
	Ram              int32     `json:"ram"`
	Disk             int32     `json:"disk"`
	MainIp           string    `json:"main_ip"`
	VcpuCount        int32     `json:"vcpu_count"`
	Region           string    `json:"region"`
	Plan             string    `json:"plan"`
	DateCreated      time.Time `json:"date_created"`
	Status           string    `json:"status"`
	AllowedBandwidth int64     `json:"allowed_bandwidth"`
	NetmaskV4        string    `json:"netmask_v4"`
	GatewayV4        string    `json:"gateway_v4"`
	PowerStatus      string    `json:"power_status"`
	ServerStatus     string    `json:"server_status"`
	OsId             int64     `json:"os_id"`
	AppId            int64     `json:"app_id"`
	ImageId          string    `json:"image_id"`
	FirewallGroupId  string    `json:"firewall_group_id"`

	UsedBandwidth *int64 `json:"used_bandwidth"`
}
