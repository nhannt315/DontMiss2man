package controllers_test

import (
	"context"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/nhannt315/real_estate_api/internal/model"
	openapiv1 "github.com/nhannt315/real_estate_api/internal/openapi/v1"
	openapiv1_controllers "github.com/nhannt315/real_estate_api/internal/openapi/v1/controllers"
	jwt_generator_mock "github.com/nhannt315/real_estate_api/internal/services/jwt/mock"
	"github.com/nhannt315/real_estate_api/internal/test"
	openapitest "github.com/nhannt315/real_estate_api/internal/test/openapi"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegistrationController_Register(t *testing.T) {
	ctx := context.Background()
	testHelper := test.NewHelper(t)

	tests := []struct {
		name                    string
		prepareTestData         func(userID uint64) error
		prepareMockJWTGenerator func(mockedGenerator *jwt_generator_mock.MockGenerator, userID uint64)
		userID                  uint64
		requestBody             *openapiv1.RegistrationRequest
		want                    *openapiv1.AuthenticationResponse
		wantErr                 *openapiv1.Error
	}{
		{
			name: "Normal case, all parameters are valid",
			prepareTestData: func(userID uint64) error {
				return nil
			},
			prepareMockJWTGenerator: func(mockedGenerator *jwt_generator_mock.MockGenerator, userID uint64) {
				mockedGenerator.EXPECT().GenerateJWTString(gomock.Any()).Return("token", nil)
			},
			userID: 321,
			requestBody: &openapiv1.RegistrationRequest{
				Email:                "test@mail.com",
				Password:             "new_password",
				PasswordConfirmation: "new_password",
			},
			want: &openapiv1.AuthenticationResponse{
				Email: "test@mail.com",
				Token: "token",
			},
		},
		{
			name: "password confirmation is not match",
			prepareTestData: func(userID uint64) error {
				return nil
			},
			prepareMockJWTGenerator: func(mockedGenerator *jwt_generator_mock.MockGenerator, userID uint64) {},
			userID: 321,
			requestBody: &openapiv1.RegistrationRequest{
				Email:                "test@mail.com",
				Password:             "new_password",
				PasswordConfirmation: "new_password_not_match",
			},
			wantErr: &openapiv1.Error{
				Detail: "Password not match",
				Title:  "Bad Request",
				Type:   "bad_request",
			},
		},
		{
			name: "email is already used",
			prepareTestData: func(userID uint64) error {
				newUser := &model.User{ID: userID, Email: "test@mail.com"}
				err := testHelper.Registry().UserRepository().WithContext(ctx).Create(newUser)
				if err != nil {
					return err
				}
				return nil
			},
			prepareMockJWTGenerator: func(mockedGenerator *jwt_generator_mock.MockGenerator, userID uint64) {},
			userID: 321,
			requestBody: &openapiv1.RegistrationRequest{
				Email:                "test@mail.com",
				Password:             "new_password",
				PasswordConfirmation: "new_password",
			},
			wantErr: &openapiv1.Error{
				Detail: "Email can't be used for registration",
				Title:  "Bad Request",
				Type:   "bad_request",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			err := tt.prepareTestData(tt.userID)
			if err != nil {
				t.Error(err)
				return
			}

			jwtGenerator := jwt_generator_mock.NewMockGenerator(ctrl)
			tt.prepareMockJWTGenerator(jwtGenerator, tt.userID)

			testServer, err := openapitest.NewTestServer(ctx, &openapiv1_controllers.InitializeContext{
				Logger:       testHelper.Logger(),
				Reg:          testHelper.Registry(),
				JWTGenerator: jwtGenerator,
			})
			if err != nil {
				t.Error(err)
				return
			}

			reqBody, err := json.Marshal(tt.requestBody)
			if err != nil {
				t.Error(err)
				return
			}

			rawResponse := testServer.PostJSON("/auth/register", string(reqBody))

			resBody := &openapiv1.AuthenticationResponse{}
			resError := &openapiv1.Error{}

			err = openapitest.ParseResponse(rawResponse, resBody, resError)
			if err != nil {
				t.Error(err)
				return
			}

			if tt.want != nil {
				assert.Equal(t, tt.want, resBody)
			}
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr, resError)
			}

			if err = testHelper.ClearDB(ctx); err != nil {
				t.Error(err)
			}
		})
	}
}
