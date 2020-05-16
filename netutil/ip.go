package netutil

import (
    "github.com/labstack/echo/v4"
    "strings"
)

//GetRemoteIP - 获得用户的正确ip
func GetRemoteIP(ctx echo.Context) string {
    req := ctx.Request()
    headers := req.Header
    ip := headers.Get("X-Real-IP")
    if len(ip)<1 {
        ip = getIpFromRemoteAddr(req.RemoteAddr, ":")
    }
    return ip
}

func getIpFromRemoteAddr(ip string, sep string) string {
    parts := strings.Split(ip, sep)
    return parts[0]
}
