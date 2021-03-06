// Reference URL: https://github.com/andrewkroh/sys/blob/master/windows/svc/example/manage.go

// +build windows

package win

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tengge1/studio/cmd"
	"golang.org/x/sys/windows/svc/mgr"
)

// startCmd start Studio service on Windows.
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start service on Windows",
	Long:  `Start service on Windows`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := startService(ServiceName); err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	cmd.AddCommand(startCmd)
}

func startService(name string) error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()
	s, err := m.OpenService(name)
	if err != nil {
		return fmt.Errorf("could not access service: %v", err)
	}
	defer s.Close()
	err = s.Start("is", "manual-started")
	if err != nil {
		return fmt.Errorf("could not start service: %v", err)
	}
	return nil
}
