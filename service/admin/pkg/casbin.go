package pkg

import (
	"errors"
	"strings"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

// 默认api
var DefaultCasbin = []CasbinInfo{
	{Path: "/menu/getMenu", Method: "POST"},
	{Path: "/jwt/jsonInBlacklist", Method: "POST"},
	{Path: "/base/login", Method: "POST"},
	{Path: "/user/register", Method: "POST"},
	{Path: "/user/changePassword", Method: "POST"},
	{Path: "/user/setUserAuthority", Method: "POST"},
	{Path: "/user/setUserInfo", Method: "PUT"},
	{Path: "/user/getUserInfo", Method: "GET"},
}

type CasbinInfo struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

type CasbinModel struct {
	Ptype       string `json:"ptype" gorm:"column:ptype"`
	AuthorityId string `json:"rolename" gorm:"column:v0"`
	Path        string `json:"path" gorm:"column:v1"`
	Method      string `json:"method" gorm:"column:v2"`
}

type CasbinService struct {
	// 数据库
	DB *gorm.DB
	// 配置文件路径
	ModelPath string
}

// 实例
var Casbin = new(CasbinService)

// 更新规则
func (l *CasbinService) UpdateCasbin(authorityId string, casbinInfos []CasbinInfo) error {
	l.ClearCasbin(0, authorityId)
	rules := [][]string{}
	for _, v := range casbinInfos {
		cm := CasbinModel{
			Ptype:       "p",
			AuthorityId: authorityId,
			Path:        v.Path,
			Method:      v.Method,
		}
		rules = append(rules, []string{cm.AuthorityId, cm.Path, cm.Method})
	}
	e := l.Casbin()
	success, _ := e.AddPolicies(rules)
	if !success {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	return nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateCasbinApi
//@description: API更新随动
//@param: oldPath string, newPath string, oldMethod string, newMethod string
//@return: error
func (l *CasbinService) UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := l.DB.Table("casbin_rule").Model(&CasbinModel{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPolicyPathByAuthorityId
//@description: 获取权限列表
//@param: authorityId string
//@return: pathMaps []request.CasbinInfo
func (l *CasbinService) GetPolicyPathByAuthorityId(authorityId string) (pathMaps []CasbinInfo) {
	e := l.Casbin()
	list := e.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ClearCasbin
//@description: 清除匹配的权限
//@param: v int, p ...string
//@return: bool
func (l *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := l.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success

}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Casbin
//@description: 持久化到数据库  引入自定义规则
//@return: *casbin.Enforcer

var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func (l *CasbinService) Casbin() *casbin.SyncedEnforcer {
	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(l.DB)
		syncedEnforcer, _ = casbin.NewSyncedEnforcer(l.ModelPath, a)
		syncedEnforcer.AddFunction("ParamsMatch", l.ParamsMatchFunc)
	})
	return syncedEnforcer
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ParamsMatch
//@description: 自定义规则函数
//@param: fullNameKey1 string, key2 string
//@return: bool

func (l *CasbinService) ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ParamsMatchFunc
//@description: 自定义规则函数
//@param: args ...interface{}
//@return: interface{}, error

func (l *CasbinService) ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return l.ParamsMatch(name1, name2), nil
}
