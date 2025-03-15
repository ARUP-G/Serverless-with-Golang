package handlers

import (
	"fmt"
	"net/http"

	"github.com/ARUP-G/Serverless-with-Golang/pkg/user"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var ErrorMethodNotAllowed = "method not allowed"

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func GetUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	email := req.QueryStringParameters["email"]

	// Prompt user to input email if not provided
	if email == "" {
		return apiResponse(http.StatusBadRequest, "Error: Please provide an email address to fetch user details.")
	}
	// Fetch user by email
	result, err := user.FetchUser(email, tableName, dynaClient)
	if len(email) > 0 {
		if err != nil {
			return apiResponse(http.StatusBadRequest, fmt.Sprintf("Error fetching user: %s", result))
		}
	}
	return apiResponse(http.StatusOK, fmt.Sprintf("User found: %v", result))

	// Fetch all users
	// result, err := user.FetchUsers(tableName, dynaClient)
	// if err != nil {
	// 	return apiResponse(http.StatusBadRequest, ErrorBody{
	// 		aws.String(err.Error()),
	// 	})
	// }
	// return apiResponse(http.StatusOK, result)

}

// create user
func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	result, err := user.CreateUser(req, tableName, dynaClient)
	if err != nil {
		return apiResponse(http.StatusBadRequest, fmt.Sprintf("Error creating user: %s", err.Error()))
	}

	if result == nil {
		resopnseMessage := "User already exists with this email"
		return apiResponse(http.StatusCreated, resopnseMessage)
	}
	resopnseMessage := fmt.Sprintf("✅ Success! User created with email: %s", result.Email)
	return apiResponse(http.StatusCreated, resopnseMessage)
}

// Update user
func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error,
) {
	result, err := user.UpdateUser(req, tableName, dynaClient)
	if err != nil {
		resopnseMessage := fmt.Sprintf("Error updating user: %s", err.Error())
		return apiResponse(http.StatusBadRequest, resopnseMessage)
	}
	resopnseMessage := fmt.Sprintf("User updated successfully👍: %v", result)
	return apiResponse(http.StatusOK, resopnseMessage)
}

// Delete user
func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error,
) {
	err := user.DeleteUser(req, tableName, dynaClient)
	if err != nil {
		resopnseMessage := fmt.Sprintf("Error deleting user: %s", err.Error())
		return apiResponse(http.StatusBadRequest, resopnseMessage)
	}
	resopnseMessage := ("🚮 User deleted successfully")
	return apiResponse(http.StatusOK, resopnseMessage)
}

func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {
	resopnseMessage := "Error: Method Not Allowed"
	return apiResponse(http.StatusMethodNotAllowed, resopnseMessage)
}
