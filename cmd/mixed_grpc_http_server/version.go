package main

import (
	"fmt"
	"log"
)

var BuildGitCommitHash string
var BuildTime string
var BuildGitUser string
var BuildGitEmail string

func init() {
	if BuildGitCommitHash == "" {
		log.Println("BuildInfo-BuildGitCommitHash Empty")
	}
	if BuildTime == "" {
		log.Println("BuildInfo-BuildTime Empty")
	}
	if BuildGitUser == "" {
		log.Println("BuildInfo-BuildGitUser Empty")
	}
	if BuildGitEmail == "" {
		log.Println("BuildInfo-BuildGitEmail Empty")
	}

	buildInfo := fmt.Sprintf("BuildInfo: BuildGitCommitHash=%s BuildTime=%s BuildGitUser=%s BuildGitEmail=%s", BuildGitCommitHash, BuildTime, BuildGitUser, BuildGitEmail)
	log.Println(buildInfo)
}
