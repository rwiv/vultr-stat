package boot

import (
	"fmt"
	"os"
	"vultr-stat/pkg/client/client"
	"vultr-stat/pkg/lib/string/format"
)

type AppRunner struct {
}

func NewAppRunner() AppRunner {
	return AppRunner{}
}

func (r *AppRunner) Run() {
	if len(os.Args) < 1 {
		fmt.Println("not found args")
		return
	}

	vultr := client.NewVultrClient()

	if os.Args[1] == "os" {
		res, err := vultr.Os()
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, osInfo := range res.Os {
			json := format.ToJsonPretty(osInfo)
			fmt.Println(json)
		}
	}
}
