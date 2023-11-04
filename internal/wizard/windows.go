package main

import (
	"fmt"
	"os"
)

const (
	defaultProgramFiles    = "C:\\Program Files"
	defaultProgramFilesx86 = "C:\\Program Files (x86)"
)

func DefaultInstallLocation(appName string, x86 bool) string {
	var prefix string
	if x86 {
		prefix = envVarOrDefault("PROGRAMFILES(x86)", defaultProgramFilesx86)
	} else {
		prefix = envVarOrDefault("PROGRAMFILES", defaultProgramFiles)
	}
	return fmt.Sprintf("%s\\%s", prefix, appName)
}

func envVarOrDefault(env, deflt string) string {
	if s := os.Getenv(env); s != "" {
		return s
	}
	return deflt
}
