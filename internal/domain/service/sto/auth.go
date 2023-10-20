package sto

type LoginUserInput struct {
	Login    string `json:"login" binding:"required,max=25,min=5"`
	Password string `json:"password" binding:"required,max=50,min=8"`
}

type RegisterUserInput struct {
	Login       string `json:"login" binding:"required,max=25,min=5"`
	Password    string `json:"password" binding:"required,max=50,min=8"`
	FullName    string `json:"fullName" binding:"required,max=100"`
	Position    string `json:"position" binding:"required,max=50"`
	Email       string `json:"email" binding:"required,email,max=200"`
	PhoneNumber string `json:"phoneNumber" binding:"required,e164"`
	UserRole    string `json:"role" binding:"required,max=50"`
	AirlineCode string `json:"airlineCode" binding:"required,max=25"`
}

func NewRegisterUserInput(login string, password string, fullName string, position string, email string, phoneNumber string, userRole string, airlineCode string) *RegisterUserInput {
	return &RegisterUserInput{Login: login, Password: password, FullName: fullName, Position: position, Email: email, PhoneNumber: phoneNumber, UserRole: userRole, AirlineCode: airlineCode}
}

type RefreshTokenInput struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

type TokenPairOutput struct {
	Access  string `json:"accessToken" binding:"required"`
	Refresh string `json:"refreshToken" binding:"required"`
}

func NewTokenPairOutput(access string, refresh string) *TokenPairOutput {
	return &TokenPairOutput{Access: access, Refresh: refresh}
}
