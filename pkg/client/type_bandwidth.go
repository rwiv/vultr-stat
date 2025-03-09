package client

import "fmt"

type InstanceBandwidthResponse struct {
	Bandwidth map[string]BandwidthInfo `json:"bandwidth"`
}

type BandwidthInfo struct {
	IncomingBytes float64 `json:"incoming_bytes"`
	OutgoingBytes float64 `json:"outgoing_bytes"`
}

type BandwidthInfoStr struct {
	IncomingBytes string `json:"incoming_bytes"`
	OutgoingBytes string `json:"outgoing_bytes"`
}

func (b BandwidthInfo) ToGb() BandwidthInfo {
	b.IncomingBytes /= 1024 * 1024 * 1024
	b.OutgoingBytes /= 1024 * 1024 * 1024
	return b
}

func (b BandwidthInfo) ToPretty() BandwidthInfoStr {
	return BandwidthInfoStr{
		IncomingBytes: fmt.Sprintf("%.2fGB", b.IncomingBytes),
		OutgoingBytes: fmt.Sprintf("%.2fGB", b.OutgoingBytes),
	}
}

func (res *InstanceBandwidthResponse) Sum() BandwidthInfo {
	var sum BandwidthInfo
	for _, bandwidth := range res.Bandwidth {
		sum.IncomingBytes += bandwidth.IncomingBytes
		sum.OutgoingBytes += bandwidth.OutgoingBytes
	}
	return sum
}
