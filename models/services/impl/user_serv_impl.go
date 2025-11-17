package impl

import (
	"be/helper"
	"be/models/domains"
	"be/models/repositories"
	"be/models/requests/user"
	user2 "be/models/resources/user"
	"fmt"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserServImpl struct {
	UserRepo  repositories.UserRepo
	Db        *gorm.DB
	Validator *validator.Validate
	JwtKey    string
}

func NewUserServImpl(userRepo repositories.UserRepo, db *gorm.DB, validator *validator.Validate, jwtKey string) *UserServImpl {
	return &UserServImpl{UserRepo: userRepo, Db: db, Validator: validator, JwtKey: jwtKey}
}

func (serv *UserServImpl) RegisterMember(request user.RegisterRequest) error {
	errValidator := helper.ErrValidator(request, serv.Validator)
	if errValidator != nil {
		return errValidator
	}

	// Find user by email
	findUser, _ := serv.UserRepo.FindByEmail(serv.Db, request.Email)
	if findUser != nil {
		return fmt.Errorf("user already exists")
	}

	// Define model
	model := user.RegisterRequestToDomain(request)
	model.Role = domains.RoleMember

	// Call repo
	errCreate := serv.UserRepo.Create(serv.Db, model)
	if errCreate != nil {
		log.Printf("[UserRepo.Create] error: %v", errCreate)
		return fmt.Errorf("failed to register, please try again later")
	}

	return nil
}

func (serv *UserServImpl) LoginMember(request user.LoginRequest) (*string, error) {
	errValidator := helper.ErrValidator(request, serv.Validator)
	if errValidator != nil {
		return nil, errValidator
	}

	// Find user by email
	findUser, errFindUser := serv.UserRepo.FindByEmail(serv.Db, request.Email)
	if findUser == nil || errFindUser != nil {
		log.Printf("[UserRepo.FindByEmail] error: %v", errFindUser)
		return nil, fmt.Errorf("user not found")
	}

	// Define model
	model := user.LoginRequestToDomains(request)

	// Check Password
	valid, errPw := serv.UserRepo.CheckPasswordValid(serv.Db, model)
	if errPw != nil {
		log.Printf("[UserRepo.CheckPasswordValid] error: %v", errPw)
		return nil, fmt.Errorf("failed to login, please try again later")
	}

	if !valid {
		return nil, fmt.Errorf("password invalid")
	}

	// User Resource
	resource := user2.ToSingleResource(findUser)

	// Create JWT Token
	duration := time.Hour * 24
	jwt, errJwt := helper.GenerateJWT(serv.JwtKey, duration, &resource)
	if errJwt != nil {
		log.Printf("[GenerateJWT] error: %v", errJwt)
		return nil, fmt.Errorf("failed to login, please try again later")
	}

	return &jwt, nil
}

func (serv *UserServImpl) Logout(userId int) error {
	// Call repo
	err := serv.UserRepo.ResetToken(serv.Db, userId)
	if err != nil {
		log.Printf("[UserRepo.ResetToken] error: %v", err)
		return fmt.Errorf("failed to logout, please try again later")
	}

	return nil
}

func (serv *UserServImpl) UpdateData(request user.UpdateDataRequest) (*string, error) {
	errValidator := helper.ErrValidator(request, serv.Validator)
	if errValidator != nil {
		return nil, errValidator
	}

	// Define model
	model := user.UpdateDataRequestToDomain(request)

	// Call repo
	result, err := serv.UserRepo.Update(serv.Db, model)
	if err != nil {
		log.Printf("[UserRepo.Update] error: %v", err)
		return nil, fmt.Errorf("failed to update, please try again later")
	}

	// User Resource
	resource := user2.ToSingleResource(result)

	// Create JWT Token
	duration := time.Hour * 24
	jwt, errJwt := helper.GenerateJWT(serv.JwtKey, duration, &resource)
	if errJwt != nil {
		log.Printf("[GenerateJWT] error: %v", errJwt)
		return nil, fmt.Errorf("failed to login, please try again later")
	}

	return &jwt, nil
}
