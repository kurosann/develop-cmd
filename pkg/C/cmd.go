package C

import (
	"bytes"
	"context"
	"develop-cmd/pkg/logger"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

var _ io.Writer = (*multiPip)(nil)

type multiPip struct {
	writer     io.Writer
	copyWriter io.Writer
}

func (p multiPip) Write(b []byte) (n int, err error) {
	n, err = p.writer.Write(b)
	if err != nil {
		return
	}
	return p.copyWriter.Write(b)
}

func CmdOutByte(ctx context.Context, name string, args ...string) ([]byte, error) {
	fileDir, _ := ctx.Value("dir").(string)
	buf := bytes.NewBuffer(nil)
	cmd := exec.CommandContext(ctx, name, args...)
	logger.Info("执行命令: %s %s", name, strings.Join(args, " "))
	cmd.Stdout = buf
	cmd.Stderr = buf
	cmd.Env = os.Environ()
	cmd.Dir = fileDir
	if err := cmd.Run(); err != nil {
		return buf.Bytes(), fmt.Errorf("%s: %s", err, buf.String())
	}
	return buf.Bytes(), nil
}

func CmdStream(ctx context.Context, name string, args ...string) error {
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
