package os_test

import (
	"os/user"
	"testing"
)

// 获取系统当前的用户对象
func TestUserCurrent(t *testing.T) {
	user, err := user.Current()
	if err != nil {
		t.Fatalf("Get current user failed, err = %v", err)
	}
	t.Logf("user = %#v", user)
}

// 通过用户名称获取系统当前的用户对象
func TestUserLookup(t *testing.T) {
	user, err := user.Lookup("DESKTOP-I67M0ID\\\\stridedot")
	if err != nil {
		t.Fatalf("Lookup failed, err = %v", user)
	}
	t.Logf("user = %#v", user)
}

// 通过用户id获取系统当前的用户对象
func TestUserLookupId(t *testing.T) {
	user, err := user.LookupId("S-1-5-21-139621390-717672982-1645291921-1001")
	if err != nil {
		t.Fatalf("Lookup by id faild, err = %v", err)
	}
	t.Logf("user = %#v", user)
}

// 获取系统当前用户所属的组 IDs
func TestUserGroupIds(t *testing.T) {
	user, err := user.Current()
	if err != nil {
		t.Fatalf("Get current user failed, err = %v", err)
	}
	gIds, err := user.GroupIds()
	if err != nil {
		t.Fatalf("Get groupids failed, err =%v", err)
	}
	t.Logf("gids = %v", gIds)
}

// 获取系统当前用户所属组的组名称
func TestUserLookupGroupId(t *testing.T) {
	group, err := user.LookupGroupId("S-1-5-21-139621390-717672982-1645291921-513")
	if err != nil {
		t.Fatalf("Lookup by group id failed, err = %v", err)
	}
	t.Logf("group = %#v", group)
}

// 根据组名获取系统用户所属的组信息
func TestUserLookupGroup(t *testing.T) {
	group, err := user.LookupGroup("None")
	if err != nil {
		t.Fatalf("Look up by group name failed, err = %v", err)
	}
	t.Logf("group = %v", group)
}
