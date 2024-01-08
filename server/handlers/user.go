package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"server/config/db"
	"server/internal/auth"
    "server/internal/struct"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
	UserID string `bson:"_id"`
	Email string `bson:"email"`
	Password string `bson:"password"`
}

type ValidationError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
}





var userCollection *mongo.Collection = database.UserCollection()
var wg sync.WaitGroup


func SignUp(ctx *gin.Context) {

    startTime := time.Now()

	fmt.Println("sign new new err")

    var newUser SignUpCredentials

    errChan := make(chan error, 4) 
    hashedPasswordCh := make(chan string, 1) 
    invalidInputCh := make(chan ValidationError, 4)
    emailExistChan := make(chan emailExist, 1) 


    wg.Add(4)

    go func() {
        defer wg.Done()
        bindUserDataJson(ctx, &newUser, errChan)
    }()

	
    go func() {
        defer wg.Done()
        checkEmailExist(&newUser.Email, emailExistChan, errChan)
    }()


    go func() {
        defer wg.Done()
        hashPassword([]byte(newUser.Password), hashedPasswordCh, errChan)
    }()


    go func() {
        defer wg.Done()
        validate(&newUser, invalidInputCh)
    }()


    
	go func() {
		wg.Wait()
		close(errChan)
	}()



    for err := range errChan {

        if err != mongo.ErrNoDocuments {
            log.Fatalln("SIGNING UP ERRORS:", err.Error())
        }
        
    }


    for invalid := range invalidInputCh {
        ctx.JSON(http.StatusInternalServerError, gin.H{invalid.Field: invalid.Tag})
        return
    }


	user := <-emailExistChan

    if user.Email != "" {
        ctx.JSON(http.StatusInternalServerError, gin.H{"emailExist": "Email Address taken"})
        return
    }


    newUser.Password = <-hashedPasswordCh

    signUpResult, signupErr := userCollection.InsertOne(context.Background(), newUser)
    if signupErr != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"INSERTING USER TO DB ERROR": signupErr.Error()})
        return
    }


    executionTime := time.Since(startTime)

    ctx.JSON(http.StatusOK, gin.H{"New User": signUpResult.InsertedID, "ExecutionTime": executionTime.String()})
}



// ------------------------------------- LOGINNN HANDLER ----------------------------------------


func Login(ctx *gin.Context){


    invalidCredChan := make(chan error, 3)
    emailExistChan := make(chan emailExist, 1) 

	var loginCredentials LoginCredentials

	wg.Add(3)

	go func ()  {
		defer wg.Done()
		bindUserDataJson(ctx, &loginCredentials, invalidCredChan)
	}()


	go func ()  {
		defer wg.Done()
		checkEmailExist(&loginCredentials.Email, emailExistChan, invalidCredChan)
	}()

    
	user := <- emailExistChan

	go func ()  {
		defer wg.Done()
		if user.Password != "" {
			isValidPassword([]byte(loginCredentials.Password), []byte(user.Password), invalidCredChan)
		}
	}()


    go func ()  {
        wg.Wait()
        close(invalidCredChan)
    }()


    for err := range invalidCredChan {
        
        if err != nil{
            ctx.JSON(http.StatusUnauthorized, gin.H{"Invalid Credentials": "Invalid Username or Password"})
            return
        }
    }


	tokenString, tokenErr := auth.CreateJwtToken(&user.Email, &user.UserID)
	if tokenErr != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"Creating Token ERROR": tokenErr.Error()})
		ctx.Abort()
		return
    }

    fmt.Println("tokenString", tokenString)

			   
	if tokenString != "" {

		cookieExp3Days := 259200

        fmt.Println("Cookie bag nan")

        ctx.SetCookie("auth_token", tokenString, cookieExp3Days, "/", "" , false, true)
	}

}


// ------------------------------------- AUTH HELPERS ----------------------------------------


func bindUserDataJson[CredType SignUpCredentials | LoginCredentials](ctx *gin.Context, userCredentials *CredType, errChan chan error) {

    jsonBindErr := ctx.ShouldBindJSON(userCredentials)

    if jsonBindErr != nil {
        errChan <- jsonBindErr
    }
}


func checkEmailExist(email *string, emailExistChan chan emailExist, errChan chan error) {

	defer close(emailExistChan)

	fmt.Println("inputted email new:", email)

    var result emailExist 
    field := bson.M{"email": email}
    
    findEmailErr := userCollection.FindOne(context.Background(), field).Decode(&result)
	fmt.Println("result", result)
	
    if findEmailErr != nil {
        errChan <- findEmailErr 
        return
    }

   
	emailExistChan <- result

}


func hashPassword(password []byte, hashedPasswordChan chan string, errChan chan error){
    
	defer close(hashedPasswordChan)

	hashedPassword, hashErr := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	
	if hashErr != nil {
		errChan <- hashErr
	}

	hashedPasswordChan <- string(hashedPassword)
	
}



func validate[CredType SignUpCredentials | LoginCredentials](userData *CredType, invalidInputCh chan<- ValidationError) {
	defer close(invalidInputCh)

    validate := validator.New()
    credValidationErr := validate.Struct(userData)

    if credValidationErr != nil {
        for _, err := range credValidationErr.(validator.ValidationErrors) {
            invalidInputCh <- ValidationError{Field: err.Field(), Tag: err.Tag()}
        }
    }

	
}


func isValidPassword(password []byte, hashedPassword []byte, invalidCredChan chan error) {

	compareHashErr := bcrypt.CompareHashAndPassword(hashedPassword, password)

	if compareHashErr != nil{
		invalidCredChan <- compareHashErr
	}

    fmt.Println("compareHashErr", compareHashErr)

	
}

// ------------------------------------- END OF AUTH HELPERS ----------------------------------------




// ------------------------------------- JWT AUTHENTICATION HANDLER ----------------------------------------

func UserData(ctx *gin.Context){

    startTime := time.Now()

    userData, exists := ctx.Get("UserData")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"ERROR": "Internal server error"})
		return
	}

    user, ok := userData.(*user_struct.UserDataClaims)
    if !ok {
        ctx.JSON(http.StatusInternalServerError, gin.H{"ERROR": "Claiming JWT Data"})
		return
    }

    fmt.Println("user data in handler", user)


    executionTime := time.Since(startTime)

    response := gin.H{
        "user_id": user.ID,
        "email": user.Email,
        "executiontime": executionTime.String(),
    }

    ctx.JSON(http.StatusOK, response)
    ctx.Abort()

}