package goconfluence

import (
	"testing"
)

func Test_TestExtenderAddCategoryResponseType(t *testing.T) {
	prepareTest(t, []int{TestExtenderAddCategoryResponseType})

	ok, err2 := testClient.AddSpaceCategory("ds", "test")
	//	defer CleanupH(resp)
	if err2 == nil {
		if ok == nil {
			t.Error("Expected Spaces. Spaces is nil")
		} else {
			if ok.Status != "category 'test' added to 'Demonstration Space (ds)' space" {
				t.Errorf("Expected Success, received: %v Spaces \n", ok.Status)
			}
		}
	} else {
		t.Error("Received nil response.")
	}
}

func Test_ExtenderSpacePermissionTypes(t *testing.T) {
	prepareTest(t, []int{TestExtenderSpacePermissionTypes})

	permissionTypes, err2 := testClient.GetPermissionTypes()
	//	defer CleanupH(resp)
	if err2 == nil {
		if permissionTypes == nil {
			t.Error("Expected Spaces. Spaces is nil")
		} else {
			if len(*permissionTypes) == 0 {
				t.Errorf("Expected Success, received: %v Spaces \n", len(*permissionTypes))
			}
		}
	} else {
		t.Error("Received nil response.")
	}
}

func Test_TestExtenderSpacePermissionTypes(t *testing.T) {
	prepareTest(t, []int{TestExtenderSpaceUserPermission})

	usersWithAnyPermission, err2 := testClient.GetAllUsersWithAnyPermission("~admin", &PaginationOptions{}) // StartAt: 0, MaxResults: 50
	//	defer CleanupH(resp)
	if err2 == nil {
		if usersWithAnyPermission == nil {
			t.Error("Expected Spaces. Spaces is nil")
		} else {
			if len(usersWithAnyPermission.Users) == 0 {
				t.Errorf("Expected Success, received: %v Spaces \n", len(usersWithAnyPermission.Users))
			}
		}
	} else {
		t.Error("Received nil response.")
	}
}

func Test_ExtenderSpaceAnyUserPermission(t *testing.T) {
	prepareTest(t, []int{TestExtenderSpaceAnyUserPermission})

	userPermissionsForSpace, err2 := testClient.GetUserPermissionsForSpace("~admin", "admin")
	//	defer CleanupH(resp)
	if err2 == nil {
		if userPermissionsForSpace == nil {
			t.Error("Expected Spaces. Spaces is nil")
		} else {
			if len(userPermissionsForSpace.Permissions) == 0 {
				t.Errorf("Expected Success, received: %v Spaces \n", len(userPermissionsForSpace.Permissions))
			}
		}
	} else {
		t.Error("Received nil response.")
	}
}

func Test_ExtenderGetGroups(t *testing.T) {
	prepareTest(t, []int{TestExtenderGetGroups})

	getGroups, err2 := testClient.GetGroups(nil)
	//	defer CleanupH(resp)
	if err2 == nil {
		if getGroups == nil {
			t.Error("Expected Groups. Groups is nil")
		} else {
			if len(getGroups.Groups) == 0 {
				t.Errorf("Expected Success, received: %v Groups \n", len(getGroups.Groups))
			}
		}
	} else {
		t.Error("Received nil response.")
	}
}

func Test_ExtenderGetUsers(t *testing.T) {
	prepareTest(t, []int{TestExtenderGetUsers})

	getUsers, err2 := testClient.GetUsers("confluence-users", nil)
	//	defer CleanupH(resp)
	if err2 == nil {
		if getUsers == nil {
			t.Error("Expected Users. Users is nil")
		} else {
			if len(getUsers.Users) == 0 {
				t.Errorf("Expected Success, received: %v Users \n", len(getUsers.Users))
			}
		}
	} else {
		t.Error("Received nil response.")
	}
}

func Test_AllGroupsWithAnyPermission(t *testing.T) {
	prepareTest(t, []int{TestAllGroupsWithAnyPermission})
	opt := PaginationOptions{}
	opt.StartAt = 0
	opt.MaxResults = 10
	getGroups, err2 := testClient.GetAllGroupsWithAnyPermission("ds", &opt)
	//	defer CleanupH(resp)
	if err2 == nil {
		if getGroups == nil {
			t.Error("Expected Groups. Groups is nil")
		} else {
			if len(getGroups.Groups) == 0 {
				t.Errorf("Expected Success, received: %v Groups \n", len(getGroups.Groups))
			}
		}
	} else {
		t.Error("Received nil response.")
	}
}

func Test_GetGroupPermissionsForSpace(t *testing.T) {
	prepareTest(t, []int{TestGetGroupPermissionsForSpace})
	opt := PaginationOptions{}
	opt.StartAt = 0
	opt.MaxResults = 10
	getPermissions, err2 := testClient.GetGroupPermissionsForSpace("ds", "confluence-users")
	//	defer CleanupH(resp)
	if err2 == nil {
		if getPermissions == nil {
			t.Error("Expected Permissions. Permissions is nil")
		} else {
			if len(getPermissions.Permissions) == 0 {
				t.Errorf("Expected Success, received: %v Permissions \n", len(getPermissions.Permissions))
			}
		}
	} else {
		t.Error("Received nil response.")
	}
}
