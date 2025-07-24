package git

import (
	"context"

	"develop-cmd/pkg/C"
)

// Push 执行 git push 操作
func Push(ctx context.Context, branch string) error {
	if branch == "" {
		_, err := C.CmdOutByte(ctx, "git", "push")
		return err
	}
	_, err := C.CmdOutByte(ctx, "git", "push", "--set-upstream", "origin", branch)
	return err
}

// Pull 执行 git pull 操作
func Pull(ctx context.Context) error {
	_, err := C.CmdOutByte(ctx, "git", "pull")
	return err
}

// Merge 执行 git merge 操作
func Merge(ctx context.Context, branch string) error {
	_, err := C.CmdOutByte(ctx, "git", "merge", branch)
	return err
}

// Clone 执行 git clone 操作
func Clone(ctx context.Context, repo string) error {
	_, err := C.CmdOutByte(ctx, "git", "clone", repo)
	return err
}
