package controllers_test

import (
	"context"
	"encoding/json"
	"github.com/nhannt315/real_estate_api/internal/model"
	"github.com/nhannt315/real_estate_api/pkg/password"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/golang/mock/gomock"
	openapiv1 "github.com/nhannt315/real_estate_api/internal/openapi/v1"
	openapiv1_controllers "github.com/nhannt315/real_estate_api/internal/openapi/v1/controllers"
	jwt_generator_mock "github.com/nhannt315/real_estate_api/internal/services/jwt/mock"
	"github.com/nhannt315/real_estate_api/internal/test"
	openapitest "github.com/nhannt315/real_estate_api/internal/test/openapi"
)

func Test_AuthenticationController_Login(t *testing.T) {
	ctx := context.Background()
		testHelper := test.NewHelper(t)

	tests := []struct {
		name                    string
		prepareMockJWTGenerator func(mockedGenerator *jwt_generator_mock.MockGenerator, userID uint64)
		prepareTestData         func(userID uint64) error
		userID                  uint64
		requestBody             *openapiv1.LoginRequest
		want                    *openapiv1.AuthenticationResponse
		wantErr                 *openapiv1.Error
	}{
		{
			name: "Email and password are valid",
			prepareTestData: func(userID uint64) error {
				hashedPassword, err := password.HashPassword("123456")
				if err != nil {
					return err
				}
				return testHelper.Registry().UserRepository().Create(&model.User{
					ID:             userID,
					Email:          "test@gmail.com",
					PasswordDigest: hashedPassword,
				})
			},
			prepareMockJWTGenerator: func(mockedGenerator *jwt_generator_mock.MockGenerator, userID uint64) {
				mockedGenerator.EXPECT().GenerateJWTString(userID).Return("token", nil)
			},
			userID: 23,
			requestBody: &openapiv1.LoginRequest{
				Email:    "test@gmail.com",
				Password: "123456",
			},
			want: &openapiv1.AuthenticationResponse{
				Email: "test@gmail.com",
				Token: "token",
			},
		},
		{
			name: "Email or Password are not valid",
			prepareTestData: func(userID uint64) error {
				hashedPassword, err := password.HashPassword("123456")
				if err != nil {
					return err
				}
				return testHelper.Registry().UserRepository().Create(&model.User{
					ID:             userID,
					Email:          "test@gmail.com",
					PasswordDigest: hashedPassword,
				})
			},
			prepareMockJWTGenerator: func(mockedGenerator *jwt_generator_mock.MockGenerator, userID uint64) {},
			userID:                  23,
			requestBody: &openapiv1.LoginRequest{
				Email:    "test@gmail.com",
				Password: "wrongpass",
			},
			wantErr: &openapiv1.Error{
				Detail: "Login failed",
				Title:  "Invalid Credential",
				Type:   "invalid_credential",
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

			rawResponse := testServer.PostJSON("/auth/login", string(reqBody))

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
