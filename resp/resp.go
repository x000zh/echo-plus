package resp

import (
    "net/http"
    "github.com/labstack/echo/v4"
    //"encoding/json"
)

//BaseResp - JSON Response
type BaseResp struct {
    Code int64 `json:"code"`
    Data interface{} `json:"data"`
    Desc string `json:"desc"`
}


//JSONResp - 返回json结果
func JSONResp( ctx echo.Context, code int64, data interface{}) error {
    return ctx.JSON(http.StatusOK, & BaseResp{
        Code: code,
        Data: data,
        Desc: "success",
    })
}

//JSONFail - 返回对应的错误
func JSONFail( ctx echo.Context, code int64, desc string ) error {
    return ctx.JSON(http.StatusOK, & BaseResp{
        Code: code,
        Data: nil,
        Desc: desc,
    })
}

//JSONFatal - 返回严重服务器错误
func JSONFatal( ctx echo.Context, code int64, desc string ) error {
    return ctx.JSON(http.StatusInternalServerError, & BaseResp{
        Code: code,
        Data: nil,
        Desc: desc,
    })
}
