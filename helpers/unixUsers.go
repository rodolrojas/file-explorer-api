package helpers

import (
	"fmt"
	"io/fs"
	"os/user"
	"syscall"
)

func GetUnixUserGroup(file fs.FileInfo) (string, string, error) {

	SysInfo := file.Sys()
	stat, ok := SysInfo.(*syscall.Stat_t)
	if !ok {
		return "", "", fmt.Errorf("failed to assert file.Sys() to *syscall.Stat_t")
	}
	uid := fmt.Sprintf("%d", stat.Uid)
	gid := fmt.Sprintf("%d", stat.Gid)
	u, err := user.LookupId(uid)
	if err != nil {
		return "", "", err
	}
	g, err := user.LookupGroupId(gid)
	if err != nil {
		return "", "", err
	}
	return u.Username, g.Name, nil
}

func GetUnixGroup(gid int) string {
	// Implementation for getting Unix group by GID
	return "group" // Placeholder
}