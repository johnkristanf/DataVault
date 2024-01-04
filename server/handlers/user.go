package handler

import (
	"context"
	"log"
	"net/http"
	"server/config/db"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)


type SignUpCredentials struct {
	Firstname string `json:"firstname" validate:"required"` 
	Lastname string `json:"lastname" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginCredentials struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type emailExist struct{
	Email string `bson:"email"`
	Password string `bson:"password"`
}

type ValidationError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
}



var userCollection *mongo.Collection = database.UserCollection()
var wg sync.WaitGroup


func SignUp(ctx *gin.Context){

	startTime := time.Now()

	var newUser SignUpCredentials

	hashedPasswordCh := make(chan string, 1)
	invalidInputCh := make(chan ValidationError) 

	emailExistCh := make(chan emailExist)


	wg.Add(4)

	go func (){
		defer wg.Done()

		if jsonbindErr := bindUserDataJson(ctx, &newUser); jsonbindErr != nil{
			ctx.JSON(http.StatusInternalServerError, gin.H{"JSON BINDING ERROR": jsonbindErr.Error()})
			ctx.Abort()
			return
		}
	}()


	go func ()  {
		defer wg.Done()

		emailExist, emailExistErr := checkEmailExist(&newUser.Email)
		log.Println("emailExist", emailExist)

		if emailExistErr != nil{
			ctx.JSON(http.StatusInternalServerError, gin.H{"CHECKING EMAIL ERROR": emailExistErr.Error()})
			close(emailExistCh)
			ctx.Abort()
			return
		}

		emailExistCh <- emailExist
		close(emailExistCh)
	}()

	    user := <-emailExistCh

	    if user.Email != "" {
		    ctx.JSON(http.StatusInternalServerError, gin.H{"emailExist": "Email Address taken"})
		    ctx.Abort()
		    return
	    }

	
	go func ()  {
		defer wg.Done()
		hashedPassword, hashErr := hashPassword([]byte(newUser.Password))

		if(hashErr != nil){
			ctx.JSON(http.StatusInternalServerError, gin.H{"HASHING PASSWORD ERROR": hashErr.Error()})
			ctx.Abort()
			close(hashedPasswordCh)
		    return
		}

		hashedPasswordCh <- hashedPassword
		close(hashedPasswordCh)
	}()
	

	newUser.Password = <- hashedPasswordCh
	

	go func ()  {
		defer wg.Done()	
		validate(ctx, &newUser, invalidInputCh)
		close(invalidInputCh)
	}()


	for invalid := range invalidInputCh{
		ctx.JSON(http.StatusInternalServerError, gin.H{invalid.Field: invalid.Tag})
		ctx.Abort()
		return
	}


	wg.Wait()


	signUpResult, signupErr := userCollection.InsertOne(context.Background(), newUser)
	if signupErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"SIGNING UP ERROR": signupErr.Error()})
		ctx.Abort()
		return
	}


	executionTime := time.Since(startTime)
	ctx.JSON(http.StatusOK, gin.H{"bag o": signUpResult, "ExecutionTime": executionTime.String()})
	ctx.Abort()

}


func Login(ctx *gin.Context){

	loginCredExistCh := make(chan emailExist)
	var loginCredentials LoginCredentials

	wg.Add(3)

	go func ()  {
		defer wg.Done()

		jsonbindErr := bindUserDataJson(ctx, &loginCredentials)

		if jsonbindErr != nil{
			ctx.JSON(http.StatusInternalServerError, gin.H{"JSON BINDING ERROR": jsonbindErr.Error()})
			ctx.Abort()
			return
		}
	}()


	go func ()  {

		defer wg.Done()

		userCred, emailExistErr := checkEmailExist(&loginCredentials.Email)
		log.Println("emailExist", userCred)

		if emailExistErr != nil{
			ctx.JSON(http.StatusInternalServerError, gin.H{"Invalid Credentials": "Invalid Username "})
			close(loginCredExistCh)
			ctx.Abort()
			return
		}

		loginCredExistCh <- userCred
		close(loginCredExistCh)

	}()

	    user := <- loginCredExistCh


		go func ()  {
			defer wg.Done()

			if user.Password != "" {
				checkvalidPassErr := isValidPassword([]byte(loginCredentials.Password), []byte(user.Password))
				if checkvalidPassErr != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"Invalid Credentials": "Invalid  Password"})
					ctx.Abort()
					return
				}
			}
			
		}()
		    


	wg.Wait()
}



func bindUserDataJson[CredType SignUpCredentials | LoginCredentials](ctx *gin.Context, userCredentials *CredType) error {

    jsonBindErr := ctx.ShouldBindJSON(userCredentials)
    if jsonBindErr != nil {
        return jsonBindErr
    }

	return nil
}


func checkEmailExist(email *string) (emailExist, error) {

    var result emailExist 
	
	log.Println("emaildfdf:",  email)
    field := bson.M{"email": email}
    
    findEmailErr := userCollection.FindOne(context.Background(), field).Decode(&result)
	log.Println("result:",  result)

	
    if findEmailErr != nil {

        if findEmailErr == mongo.ErrNoDocuments {
            log.Println("Email not found")
            return result, findEmailErr
        }
    }

    if result.Email != "" {
        return result, nil 
    }

    return result, findEmailErr
}


func hashPassword(password []byte) (string, error) {

	hashedPassword, hashErr := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	
	if hashErr != nil {
		return "", hashErr
	}

	return string(hashedPassword), nil
}



func validate[CredType SignUpCredentials | LoginCredentials](ctx *gin.Context, userData *CredType, invalidInputCh chan<- ValidationError) {

    validate := validator.New()
    credValidationErr := validate.Struct(userData)

    if credValidationErr != nil {
        for _, err := range credValidationErr.(validator.ValidationErrors) {
            invalidInputCh <- ValidationError{Field: err.Field(), Tag: err.Tag()}
        }
    }
}


func isValidPassword(password []byte, hashedPassword []byte) error {
	compareHashErr := bcrypt.CompareHashAndPassword(hashedPassword, password)

	if compareHashErr != nil{
		return compareHashErr
	}

	return nil
}
