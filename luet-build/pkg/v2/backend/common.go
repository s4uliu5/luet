/*
Copyright © 2022-2024 Macaroni OS Linux
See AUTHORS and LICENSE for the license details and contributors.
*/
package backend

import (
	"fmt"
	"os/exec"

	"github.com/geaaru/luet/pkg/config"
	. "github.com/geaaru/luet/pkg/logger"
)

func runCommand(cmd *exec.Cmd) error {
	output := ""
	buffered := !config.LuetCfg.GetGeneral().ShowBuildOutput
	writer := NewBackendWriter(buffered)

	cmd.Stdout = writer
	cmd.Stderr = writer

	if buffered {
		Spinner(22)
		defer SpinnerStop()
	}

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed starting command: %s", err.Error())
	}

	err = cmd.Wait()
	if err != nil {
		output = writer.GetCombinedOutput()
		return fmt.Errorf("failed running command: %s\n%s", output, err.Error())
	}

	return nil
}
