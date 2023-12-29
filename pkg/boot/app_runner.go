package boot

import (
	"fmt"
	"os"

	"vultr-stat/pkg/client"
	"vultr-stat/pkg/tfac"
	"vultr-stat/pkg/tfac/tw"
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
		f := tfac.NewOsTableFactory()
		t := tw.GetTable(f.Columns(), f.Rows(res.Os))
		t.Render()
	}
}
