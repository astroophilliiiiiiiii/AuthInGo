package controllers

// in this JSON reponses will be sent -- by utils -- json.go responses
import (
	"AuthInGo/dto"
	"AuthInGo/services"
	"AuthInGo/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// no interface -- as it only does one work of passing the logic to the service layer --- passer !!

// it will connect to the service layer interface
type UserController struct {
	UserService services.UserService
}

// constructor in which the service layer will be passed
func NewUserController(__userService services.UserService) *UserController {
	return &UserController{
		UserService: __userService,
	}
}

// member function
func (uc *UserController) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateUser called in UserController")

	payload := r.Context().Value("payload").(*dto.CreateUserRequestDto)

	user, err1 := uc.UserService.CreateUserService(*payload)

	if err1 != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Creation failed", err1)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusCreated, user, "user created successfully")
	fmt.Println("User created successfullyy !! ")
}

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetUserByID called in UserController")

	// call the service layer
	var id int
	strid := chi.URLParam(r, "id") // returns the string

	// CONDITION 1: Agar URL me ID di hai (jaise /profile/1)
	if strid != "" {
		parsedId, err := strconv.Atoi(strid)
		if err != nil {
			utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid User ID in URL", err)
			return // 🛑 Return zaroori hai!
		}
		id = parsedId
	} else {
		// CONDITION 2: Agar URL me ID NAHI hai, toh Context se utha (Middleware se)
		ctxVal := r.Context().Value("userID")
		if ctxVal == nil {
			utils.WriteJsonErrorResponse(w, http.StatusUnauthorized, "User ID missing in request", nil)
			return
		}

		// Piche humne middleware me int64 banaya tha, usko int me convert kar rahe hain
		ctxID, ok := ctxVal.(int64) // type assertion -- ki isme int64 dala tha vo nikalke dedo
		if !ok {
			utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Context ID is not int64", nil)
			return
		}
		id = int(ctxID) // finally converting in the int of this id
	}

	response, nerr := uc.UserService.GetUserByIdService(&id)

	// sending JSON response to the frontend
	if nerr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "User not found with the given id", nerr)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, response, "fetched user with given id successfully")
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login user called in UserController")

	//.(dto.LoginUserRequestDTO) = Context se nikli any value ko LoginUserRequestDTO type me convert (type assert) karna. ✅
	// taki fir mein payload.Email payload.password ye sb krr saku
	payload := r.Context().Value("payload").(*dto.LoginUserRequestDTO)

	token, err := uc.UserService.LogInUser(payload) // user service -- contains logic of hwt and all --validating

	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Login failed", err)
		return
	}

	// send the response -- JSON format
	utils.WriteJsonSuccessResponse(w, http.StatusOK, token, "User logged in sucessfully ! ")
}
