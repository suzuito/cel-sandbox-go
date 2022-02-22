package main

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	"google.golang.org/protobuf/proto"
)

type User struct {
	Name  string
	Team  string
	Roles []Role
}

func (u *User) TypeName() string {
	return "user"
}

type Role struct {
	Name        string
	Permissions []Permission
}

var (
	RoleAdmin = Role{
		Name: "チーム管理者",
		Permissions: []Permission{
			PermissionTeamDelete,
			PermissionTeamProfileEdit,
			PermissionTeamProfileRead,
			PermissionMemberCreate,
			PermissionMemberRead,
			PermissionMemberUpdate,
			PermissionMemberDelete,
			EventCreate,
			EventRead,
			EventUpdate,
			EventDelete,
			GameOrderCreate,
			GameOrderRead,
			GameOrderUpdate,
			GameOrderDelete,
		},
	}
	RoleManager = Role{
		Name: "マネージャー",
		Permissions: []Permission{
			PermissionTeamProfileEdit,
			PermissionTeamProfileRead,
			PermissionMemberCreate,
			PermissionMemberRead,
			PermissionMemberUpdate,
			PermissionMemberDelete,
			EventCreate,
			EventRead,
			EventUpdate,
			EventDelete,
			GameOrderRead,
		},
	}
	RoleDirector = Role{
		Name: "監督",
		Permissions: []Permission{
			PermissionTeamProfileRead,
			PermissionMemberRead,
			EventRead,
			GameOrderCreate,
			GameOrderRead,
			GameOrderUpdate,
			GameOrderDelete,
		},
	}
	RoleMember = Role{
		Name: "メンバー",
		Permissions: []Permission{
			PermissionTeamProfileRead,
			PermissionMemberRead,
			EventRead,
			GameOrderRead,
		},
	}
	RoleTemporaryMember = Role{
		Name: "体験入団",
		Permissions: []Permission{
			PermissionTeamProfileRead,
			GameOrderRead,
		},
	}
)

type Permission string

var (
	// チーム削除権限
	PermissionTeamDelete Permission = "team.delete"
	// チームのプロフィール編集権限
	PermissionTeamProfileEdit Permission = "taem.profile.edit"
	PermissionTeamProfileRead Permission = "taem.profile.edit"
	// チームメンバーの追加・編集・削除権限
	PermissionMemberCreate Permission = "member.create"
	PermissionMemberRead   Permission = "member.read"
	PermissionMemberUpdate Permission = "member.update"
	PermissionMemberDelete Permission = "member.delete"
	// チーム行事（試合、練習、飲み会、等々）の追加・編集・削除権限
	EventCreate Permission = "event.create"
	EventRead   Permission = "event.read"
	EventUpdate Permission = "event.update"
	EventDelete Permission = "event.delete"
	// 試合メンバーの追加・編集・削除権限
	GameOrderCreate Permission = "gameorder.create"
	GameOrderRead   Permission = "gameorder.read"
	GameOrderUpdate Permission = "gameorder.update"
	GameOrderDelete Permission = "gameorder.delete"
)

var Users []User = []User{
	{
		Team: "凸凹ジャイアンツ",
		Name: "なBつね",
		Roles: []Role{
			RoleAdmin,
		},
	},
	{
		Team: "凸凹ジャイアンツ",
		Name: "長妻すげお",
		Roles: []Role{
			RoleDirector,
		},
	},
	{
		Team: "凸凹ジャイアンツ",
		Name: "とんこつ由伸",
		Roles: []Role{
			RoleMember,
		},
	},
	{
		Team: "凸凹ジャイアンツ",
		Name: "暑い秀喜",
		Roles: []Role{
			RoleMember,
		},
	},
	{
		Team: "薬流党スワローズ",
		Name: "2ばくろう",
		Roles: []Role{
			RoleAdmin,
		},
	},
	{
		Team: "薬流党スワローズ",
		Name: "新田敦也",
		Roles: []Role{
			RoleDirector,
			RoleMember,
		},
	},
	{
		Team: "薬流党スワローズ",
		Name: "栗山",
		Roles: []Role{
			RoleMember,
		},
	},
}

func main() {
	env, err := cel.NewEnv(
		cel.Declarations(
			decls.NewVar(
				"user",
				// decls.NewObjectType(".User"),
				decls.Dyn,
			),
		),
	)
	if err != nil {
		glog.Exit(err)
	}
	ast, iss := env.Parse(`user.Team == '凸凹ジャイアンツ'`)
	if iss.Err() != nil {
		glog.Exit(iss.Err())
	}
	checked, iss := env.Check(ast)
	if iss.Err() != nil {
		glog.Exit(iss.Err())
	}
	if !proto.Equal(checked.ResultType(), decls.Bool) {
		glog.Exitf(
			"Got %v, wanted %v result type",
			checked.ResultType(), decls.String,
		)
	}
	program, err := env.Program(checked)
	if err != nil {
		glog.Exitf("program error: %v", err)
	}

	result, _, err := program.Eval(&Users[0])
	if err != nil {
		glog.Exitf("eval error: %v", err)
	}
	fmt.Println(result)
}
