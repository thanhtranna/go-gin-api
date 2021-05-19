package gencode_handler

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"

	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type handlerExecuteRequest struct {
	Name string `form:"name"`
}

func (h *handler) HandlerExecute() core.HandlerFunc {
	return func(c core.Context) {
		req := new(handlerExecuteRequest)
		if err := c.ShouldBindPostForm(req); err != nil {
			c.Payload("The parameter is passed incorrectly")
			return
		}

		shellPath := fmt.Sprintf("./scripts/handlergen.sh %s", req.Name)

		// runtime.GOOS = linux or darwin
		command := exec.Command("/bin/bash", "-c", shellPath)

		if runtime.GOOS == "windows" {
			command = exec.Command("cmd", "/C", shellPath)
		}

		var stderr bytes.Buffer
		command.Stderr = &stderr

		output, err := command.Output()
		if err != nil {
			c.Payload(stderr.String())
			return
		}

		c.Payload(string(output))
	}
}
