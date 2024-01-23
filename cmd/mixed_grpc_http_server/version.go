package main

import "log/slog"

var BuildGitCommitHash string
var BuildTime string
var BuildGitUser string
var BuildGitEmail string

func init() {
	if BuildGitCommitHash == "" {
		slog.Warn("BuildInfo-BuildGitCommitHash Empty")
	}
	if BuildTime == "" {
		slog.Warn("BuildInfo-BuildTime Empty")
	}
	if BuildGitUser == "" {
		slog.Warn("BuildInfo-BuildGitUser Empty")
	}
	if BuildGitEmail == "" {
		slog.Warn("BuildInfo-BuildGitEmail Empty")
	}

	slog.Info("BuildInfo", "BuildGitCommitHash", BuildGitCommitHash, "BuildTime", BuildTime, "BuildGitUser", BuildGitUser, "BuildGitEmail", BuildGitEmail)
}
