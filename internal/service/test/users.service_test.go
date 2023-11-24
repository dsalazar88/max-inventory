package test

import (
	"context"
	"courses/golang/inventory-project/internal/service"
	"testing"
)

// The TestRegisterUser function tests the RegisterUser method by running multiple test cases with
// different inputs and expected outputs.
func TestRegisterUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		UserName      string
		Password      string
		ExpectedError error
	}{
		{
			Name:          "RegisterUser_Success",
			Email:         "test@test.com",
			UserName:      "test",
			Password:      "validPassword",
			ExpectedError: nil,
		},
		{
			Name:          "RegisterUser_UserAlreadyExists",
			Email:         "test@exists.com",
			UserName:      "test",
			Password:      "validPassword",
			ExpectedError: service.ErrUserAlreadyExists,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			err := s.RegisterUser(ctx, tc.Email, tc.Name, tc.Password)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError.Error(), err)
			}
		})
	}
}

// The TestLoginUser function tests the login functionality for a user by providing different test
// cases with expected errors.
func TestLoginUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		Password      string
		ExpectedError error
	}{
		{
			Name:          "loginUser_success",
			Email:         "test@exists.com",
			Password:      "validPassword",
			ExpectedError: nil,
		},
		{
			Name:          "LoginUser_invalidPassword",
			Email:         "test@exists.com",
			Password:      "invalidPassword",
			ExpectedError: service.ErrInvalidCredentials,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)
		})

		_, err := s.LoginUser(ctx, tc.Email, tc.Password)
		if err != tc.ExpectedError {
			t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
		}

	}
}

// The TestSaveUserRole function tests the SaveUserRole function by providing different test cases and
// checking if the expected error matches the actual error.
func TestSaveUserRole(t *testing.T) {
	testCases := []struct {
		Name          string
		UserId        int64
		RoleId        int64
		ExpectedError error
	}{
		{
			Name:          "AddUserRole_Success",
			UserId:        1,
			RoleId:        2,
			ExpectedError: nil,
		},
		{
			Name:          "AddUserRole_UserAlreadyHasRole",
			UserId:        1,
			RoleId:        2,
			ExpectedError: service.ErrRoleAlreadyAdded,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			repo.Mock.Test(t)

			err := s.SaveUserRole(ctx, tc.UserId, tc.RoleId)
			if err != nil {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}

// The TestRemoveUserRole function tests the RemoveUserRole method by providing test cases with
// different user IDs and role IDs and checking if the expected error matches the actual error.
func TestRemoveUserRole(t *testing.T) {
	testCases := []struct {
		Name          string
		UserId        int64
		RoleId        int64
		ExpectedError error
	}{
		{
			Name:          "RemoveUserRole_Success",
			UserId:        1,
			RoleId:        1,
			ExpectedError: nil,
		},
		{
			Name:          "RemoveUserRole_UserDoesNotHaveRole",
			UserId:        1,
			RoleId:        3,
			ExpectedError: service.ErrRoleNotFound,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			repo.Mock.Test(t)

			err := s.RemoveUserRole(ctx, tc.UserId, tc.RoleId)
			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}
}
