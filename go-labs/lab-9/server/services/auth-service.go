package services

type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

var jwtKey = []byte("tut_mogla-bit__w-a-s-h-a__REKLAMA")

func GenerateTokens(password []byte) {

}
