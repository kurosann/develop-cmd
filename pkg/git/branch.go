package git

import (
	"context"
	"strings"

	"develop-cmd/pkg/C"
	"develop-cmd/pkg/logger"
)

// GetUserName 获取 git 用户名
func GetUserName(ctx context.Context) (string, error) {
	out, err := C.CmdOutByte(ctx, "git", "config", "user.name")
	if err != nil {
		logger.Info("获取 git 用户名失败: %v", err)
		return "", err
	}
	username := strings.TrimSpace(string(out))
	logger.Info("获取 git 用户名: %s", username)
	return username, nil
}

// checkRemoteBranch 检查远程分支是否存在
func checkRemoteBranch(ctx context.Context, branch string) (bool, error) {
	logger.Info("开始检查远程分支 %s 是否存在", branch)
	output, err := C.CmdOutByte(ctx, "git", "ls-remote", "--heads", "origin", branch)
	if err != nil {
		logger.Info("检查远程分支 %s 失败: %v", branch, err)
		return false, err
	}
	exists := strings.Contains(string(output), branch)
	logger.Info("远程分支 %s 检查结果: %v", branch, exists)
	return exists, nil
}

// hasUpstream 检查分支是否有上游分支
func hasUpstream(ctx context.Context, branch string) (bool, error) {
	_, err := C.CmdOutByte(ctx, "git", "rev-parse", "--abbrev-ref", branch+"@{upstream}")
	hasUpstream := err == nil
	return hasUpstream, nil
}

// HandleBranch 处理分支的创建、切换和推送
func HandleBranch(ctx context.Context, branch string) error {
	logger.Info("开始处理分支 %s", branch)

	// 查看当前分支
	currentBranch, err := C.CmdOutByte(ctx, "git", "branch", "--show-current")
	if err != nil {
		logger.Info("获取当前分支失败: %v", err)
		return err
	}
	currentBranchName := strings.TrimSpace(string(currentBranch))
	logger.Info("当前分支: %s", currentBranchName)

	// 如果已经在目标分支上，检查是否有上游分支
	if currentBranchName == branch {
		logger.Info("当前已在目标分支 %s 上", branch)
		hasUpstream, err := hasUpstream(ctx, branch)
		if err != nil {
			return err
		}
		if hasUpstream {
			logger.Info("分支 %s 已有上游分支，无需推送", branch)
			return nil
		}
		logger.Info("分支 %s 无上游分支，开始推送", branch)
		return Push(ctx, branch)
	}

	// 检查本地分支是否存在
	_, err = C.CmdOutByte(ctx, "git", "show-ref", "--verify", "--quiet", "refs/heads/"+branch)
	localBranchExists := err == nil
	logger.Info("本地分支 %s 检查结果: %v", branch, localBranchExists)

	if !localBranchExists {
		// 只有在本地分支不存在时才检查远程
		remoteBranchExists, err := checkRemoteBranch(ctx, branch)
		if err != nil {
			return err
		}

		// 本地分支不存在，先创建本地分支
		logger.Info("开始创建本地分支 %s", branch)
		_, err = C.CmdOutByte(ctx, "git", "checkout", "-b", branch)
		if err != nil {
			logger.Info("创建本地分支 %s 失败: %v", branch, err)
			return err
		}
		logger.Info("创建本地分支 %s 成功", branch)

		if remoteBranchExists {
			// 远程分支存在，先 fetch 所有分支
			logger.Info("开始 fetch 所有分支")
			_, err = C.CmdOutByte(ctx, "git", "fetch", "origin")
			if err != nil {
				logger.Info("fetch 分支失败: %v", err)
				return err
			}
			// 设置跟踪
			logger.Info("开始设置分支 %s 跟踪远程分支", branch)
			_, err = C.CmdOutByte(ctx, "git", "branch", "--set-upstream-to", "origin/"+branch)
			if err != nil {
				logger.Info("设置分支跟踪失败: %v", err)
				return err
			}
			logger.Info("设置分支 %s 跟踪远程分支成功", branch)
		} else {
			// 远程分支不存在，推送新分支到远程
			logger.Info("开始推送新分支 %s 到远程", branch)
			return Push(ctx, branch)
		}
	} else {
		// 本地分支存在，先检查是否有上游分支
		hasUpstream, err := hasUpstream(ctx, branch)
		if err != nil {
			return err
		}

		// 切换到分支
		logger.Info("开始切换到分支 %s", branch)
		_, err = C.CmdOutByte(ctx, "git", "checkout", branch)
		if err != nil {
			logger.Info("切换到分支 %s 失败: %v", branch, err)
			return err
		}
		logger.Info("切换到分支 %s 成功", branch)

		// 如果没有上游分支，则推送
		if !hasUpstream {
			logger.Info("分支 %s 无上游分支，开始推送", branch)
			return Push(ctx, branch)
		}
	}

	logger.Info("分支 %s 处理完成", branch)
	return nil
}
