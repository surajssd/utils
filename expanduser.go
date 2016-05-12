package utils

import (
    "strings"
    "os/user"
    "path/filepath"
)

func ExpandUser(path string) string {

    if strings.HasPrefix(path, "~") == false {
        return path
    }

    usr, _ := user.Current()
    path = filepath.Join(usr.HomeDir, path[1:])
    return path
}
