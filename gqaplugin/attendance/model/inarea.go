package model

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"time"
)

type Inarea struct {
	UserGUID    string    `json:"userGUID" gorm:"column:UserGUID"`
	InAreaTime  time.Time `json:"inAreaTime" gorm:"column:InAreaTime"`
	DeviceSN    string    `json:"deviceSN" gorm:"column:DeviceSN"`
	InAreaImage string    `json:"in_area_image" gorm:"column:InAreaImage"`
	CardNo      string    `json:"card_no" gorm:"column:CardNo"`
	Type        int       `json:"type" gorm:"column:Type"`
}

type Userinfo struct {
	GUID               string    `json:"GUID" gorm:"column:GUID"`
	WorkNumber         string    `json:"workNumber" gorm:"column:WorkNumber"`
	CardNumber         string    `json:"cardNumber" gorm:"column:CardNumber"`
	DepID              int       `json:"depID" gorm:"column:DepID"`
	TeamID             int       `json:"teamID" gorm:"column:TeamID"`
	PassageTimeGroupID int       `json:"passageTimeGroupID" gorm:"column:PassageTimeGroupID"`
	PassageAreaID      int       `json:"passageAreaID" gorm:"column:PassageAreaID"`
	State              int       `json:"state" gorm:"column:State"`
	UserName           string    `json:"userName" gorm:"column:UserName"`
	UserSex            string    `json:"userSex" gorm:"column:UserSex"`
	UserNation         string    `json:"userNation" gorm:"column:UserNation"`
	UserBrithday       string    `json:"userBrithday" gorm:"column:UserBrithday"`
	UserAddress        string    `json:"userAddress" gorm:"column:UserAddress"`
	UserIDCardNumber   string    `json:"userIDCardNumber" gorm:"column:UserIDCardNumber"`
	EffectiveStart     string    `json:"effectiveStart" gorm:"column:EffectiveStart"`
	EffectiveEnd       string    `json:"effectiveEnd" gorm:"column:EffectiveEnd"`
	UserImage          string    `json:"userImage" gorm:"column:UserImage"`
	Organ              string    `json:"organ" gorm:"column:Organ"`
	InOrOut            int       `json:"inOrOut" gorm:"column:InOrOut"`
	IsUpdate           int       `json:"isUpdate" gorm:"column:IsUpdate"`
	UpdateTime         time.Time `json:"updateTime" gorm:"column:UpdateTime"`
	UpowerID           int       `json:"upowerID" gorm:"column:UpowerID"`
	CreateTime         time.Time `json:"createTime" gorm:"column:CreateTime"`
	UserRole           int       `json:"userRole" gorm:"column:UserRole"`
}

type InareaWithUserinfo struct {
	Inarea
	Userinfo
}

type InareaWithUserinfoDateRequest struct {
	global.RequestPageAndSort
	WorkNumber string `json:"workNumber"`
	InAreaTime string `json:"inAreaTime"`
	UserName   string `json:"userName"`
}

type InareaWithUserinfoYearRequest struct {
	WorkNumber string `json:"workNumber"`
}
