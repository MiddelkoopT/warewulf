package list

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/warewulf/warewulf/internal/app/wwctl/table"
	apiprofile "github.com/warewulf/warewulf/internal/pkg/api/profile"
	"github.com/warewulf/warewulf/internal/pkg/api/routes/wwapiv1"
	"github.com/warewulf/warewulf/internal/pkg/wwlog"
)

func CobraRunE(vars *variables) func(cmd *cobra.Command, args []string) (err error) {
	return func(cmd *cobra.Command, args []string) (err error) {
		if len(args) > 0 && strings.Contains(args[0], ",") {
			args = strings.FieldsFunc(args[0], func(r rune) bool { return r == ',' })
		}
		req := wwapiv1.GetProfileList{
			ShowAll:  vars.showAll,
			ShowYaml: vars.showYaml,
			ShowJson: vars.showJson,
			Profiles: args,
		}
		profileInfo, err := apiprofile.ProfileList(&req)
		if err != nil {
			return
		}
		if len(profileInfo.Output) > 0 {
			if vars.showYaml || vars.showJson {
				wwlog.Info(profileInfo.Output[0])
			} else {
				t := table.New(cmd.OutOrStdout())
				t.AddHeader(table.Prep(strings.Split(profileInfo.Output[0], ":=:"))...)
				for _, val := range profileInfo.Output[1:] {
					t.AddLine(table.Prep(strings.Split(val, ":=:"))...)
				}
				t.Print()
			}
		}
		return
	}
}
