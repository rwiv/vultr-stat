package client

import (
	"vultr-stat/pkg/lib/string/format"
)

type FileInfo struct {
	PfPath          string  `json:"pfPath"`
	IsDirectory     bool    `json:"isDirectory"`
	EncFsPath       *string `json:"encFsPath,omitempty"`
	EncParentFsPath *string `json:"encParentFsPath,omitempty"`
	Size            *int64  `json:"size,omitempty"`
	Ctime           *string `json:"ctime,omitempty"`
	Mtime           *string `json:"mtime,omitempty"`
	Atime           *string `json:"atime,omitempty"`
	Details         *string `json:"details,omitempty"`
}

func (f FileInfo) String() string {
	return format.ToJson(f)
}
