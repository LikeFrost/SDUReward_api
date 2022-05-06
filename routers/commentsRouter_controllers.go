package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["api/controllers:RewardController"] = append(beego.GlobalControllerRouter["api/controllers:RewardController"],
        beego.ControllerComments{
            Method: "Add",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:RewardController"] = append(beego.GlobalControllerRouter["api/controllers:RewardController"],
        beego.ControllerComments{
            Method: "GetAllReward",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:RewardController"] = append(beego.GlobalControllerRouter["api/controllers:RewardController"],
        beego.ControllerComments{
            Method: "GetReward",
            Router: `/:rewardId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:RewardController"] = append(beego.GlobalControllerRouter["api/controllers:RewardController"],
        beego.ControllerComments{
            Method: "DeleteReward",
            Router: `/:rewardId`,
            AllowHTTPMethods: []string{"Delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:RewardController"] = append(beego.GlobalControllerRouter["api/controllers:RewardController"],
        beego.ControllerComments{
            Method: "GetRewardByTag",
            Router: `/byTag/:tag`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:RewardController"] = append(beego.GlobalControllerRouter["api/controllers:RewardController"],
        beego.ControllerComments{
            Method: "GetUserReward",
            Router: `/byUserId/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:SuggestionController"] = append(beego.GlobalControllerRouter["api/controllers:SuggestionController"],
        beego.ControllerComments{
            Method: "AddSuggestion",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:UserController"] = append(beego.GlobalControllerRouter["api/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateUser",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:UserController"] = append(beego.GlobalControllerRouter["api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:UserController"] = append(beego.GlobalControllerRouter["api/controllers:UserController"],
        beego.ControllerComments{
            Method: "DeleteUser",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"Delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:UserController"] = append(beego.GlobalControllerRouter["api/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAllUsers",
            Router: `/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:UserController"] = append(beego.GlobalControllerRouter["api/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:UserController"] = append(beego.GlobalControllerRouter["api/controllers:UserController"],
        beego.ControllerComments{
            Method: "LogUp",
            Router: `/logup`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
