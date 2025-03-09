package boot

import (
	"fmt"
	"os"

	"github.com/rwiv/vultr-stat/pkg/client"
	"github.com/rwiv/vultr-stat/pkg/common"
	"github.com/rwiv/vultr-stat/pkg/lib/string/format"
	"github.com/rwiv/vultr-stat/pkg/tfac"
	"github.com/rwiv/vultr-stat/pkg/tfac/tw"
)

type AppRunner struct {
}

func NewAppRunner() AppRunner {
	return AppRunner{}
}

func (r *AppRunner) Run() {
	env, err := common.ReadEnv()
	if err != nil {
		fmt.Println(err)
	}
	vultr := client.NewVultrClient(env.VultrApiKey)

	if len(os.Args) < 2 || os.Args[1] == "instances" {
		res, err := vultr.Instances()
		if err != nil {
			fmt.Println(err)
			return
		}
		f := tfac.NewInstanceTableFactory()
		t := tw.GetTable(f.Columns(), f.Rows(res.Instances))
		t.Render()
		return
	}

	if os.Args[1] == "instances-json" {
		res, err := vultr.Instances()
		if err != nil {
			fmt.Println(err)
			return
		}
		json := format.ToJsonPretty(res.Instances)
		fmt.Println(json)
		return
	}

	if os.Args[1] == "os" {
		res, err := vultr.Os()
		if err != nil {
			fmt.Println(err)
			return
		}
		f := tfac.NewOsTableFactory()
		t := tw.GetTable(f.Columns(), f.Rows(res.Os))
		t.Render()
		return
	}

	if os.Args[1] == "ac" {
		res, err := vultr.Account()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res.Account.ToInfo().ToPretty())
		return
	}
}
