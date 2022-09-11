package main

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/oauth2/jws"
)

const (
	one_day_expire = 3600 * 24
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, errorResponse ErrorResponse) {
	respondWithJSON(w, code, errorResponse)
}

func generateJWT(saKeyfile string, saEmail string, audience string, expireTime int64, userPayload UserPayload) (string, error) {

	now := time.Now().Unix()

	jwt := &jws.ClaimSet{
		Iat:           now,
		Exp:           now + expireTime,
		Iss:           saEmail,
		Aud:           audience,
		Sub:           saEmail,
		PrivateClaims: map[string]interface{}{"email": saEmail, "userInfo": userPayload},
	}

	jwsHeader := &jws.Header{
		Algorithm: "RS256",
		Typ:       "JWT",
	}

	saKeyfileReplaced := strings.Replace(saKeyfile, "\\n", "\n", -1)

	saKeyfileAsByteArr := []byte(saKeyfileReplaced)

	block, _ := pem.Decode(saKeyfileAsByteArr)

	if block == nil {
		return "", errors.New("pem not found")
	}

	parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)

	if err != nil {
		return "", fmt.Errorf("private key parse error: %v", err)
	}

	rsaKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return "", errors.New("private key failed rsa.PrivateKey type assertion")
	}

	return jws.Encode(jwsHeader, jwt, rsaKey)
}

func generate(w http.ResponseWriter, r *http.Request) {

	var userPayload UserPayload
	var audience string
	var generateResponse GenerateResponse
	var service_account_email string
	var errorResponse ErrorResponse

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userPayload)

	if err != nil {
		errorResponse.Message = "Invalid request payload"
		errorResponse.Status = http.StatusBadRequest
		respondWithError(w, http.StatusBadRequest, errorResponse)
		return
	}

	defer r.Body.Close()

	audience = os.Getenv("JWT_AUDIENCE")

	service_account_email = os.Getenv("SERVICE_ACCOUNT_CLIENT_EMAIL")

	service_account_private_key := os.Getenv("SERVICE_ACCOUNT_PRIVATE_KEY")

	has_audience := audience != ""
	has_service_account_email := service_account_email != ""
	has_service_accourt_private_key := service_account_private_key != ""

	if !has_audience || !has_service_account_email || !has_service_accourt_private_key {
		errorResponse.Message = "Enviroment variables is missing"
		errorResponse.Status = http.StatusInternalServerError
		respondWithError(w, http.StatusInternalServerError, errorResponse)
		return
	}

	jwt, err := generateJWT(service_account_private_key, service_account_email, audience, int64(one_day_expire), userPayload)

	if err != nil {
		errorResponse.Message = err.Error()
		errorResponse.Status = http.StatusInternalServerError
		respondWithError(w, http.StatusInternalServerError, errorResponse)
		return
	}

	generateResponse.Token = jwt

	respondWithJSON(w, http.StatusOK, generateResponse)

}

func Decode(payload string) (*TokenDecoded, error) {

	split_payload := strings.Split(payload, ".")

	has_valid_token := len(split_payload) > 2

	if !has_valid_token {
		return nil, errors.New("jws: invalid token received")
	}

	decoded, err := base64.RawURLEncoding.DecodeString(split_payload[1])
	if err != nil {
		return nil, err
	}

	tokenDecoded := &TokenDecoded{}

	err = json.NewDecoder(bytes.NewBuffer(decoded)).Decode(tokenDecoded)

	return tokenDecoded, err
}

func refresh(w http.ResponseWriter, r *http.Request) {

	var userPayload UserPayload
	var audience string
	var generateResponse GenerateResponse
	var service_account_email string
	var errorResponse ErrorResponse

	bearer_token := r.Header.Get("Authorization")

	has_token := bearer_token != ""

	if !has_token {
		errorResponse.Message = "Header Authorization is missing."
		errorResponse.Status = http.StatusBadRequest
		respondWithError(w, http.StatusBadRequest, errorResponse)
	}

	token := strings.TrimPrefix(bearer_token, "Bearer")
	tokenDecoded, err := Decode(token)

	if err != nil {
		errorResponse.Message = err.Error()
		errorResponse.Status = http.StatusBadRequest
		respondWithError(w, http.StatusBadRequest, errorResponse)
	}

	audience = os.Getenv("JWT_AUDIENCE")

	service_account_email = tokenDecoded.Email

	service_account_private_key := os.Getenv("SERVICE_ACCOUNT_PRIVATE_KEY")

	has_audience := audience != ""
	has_service_account_email := service_account_email != ""
	has_service_accourt_private_key := service_account_private_key != ""

	if !has_audience || !has_service_account_email || !has_service_accourt_private_key {
		errorResponse.Message = "Enviroment variables is missing"
		errorResponse.Status = http.StatusInternalServerError
		respondWithError(w, http.StatusInternalServerError, errorResponse)
		return
	}

	userPayload = tokenDecoded.UserInfo

	jwt, err := generateJWT(service_account_private_key, service_account_email, audience, int64(one_day_expire), userPayload)

	if err != nil {
		errorResponse.Message = err.Error()
		errorResponse.Status = http.StatusInternalServerError
		respondWithError(w, http.StatusInternalServerError, errorResponse)
		return
	}

	generateResponse.Token = jwt

	respondWithJSON(w, http.StatusOK, generateResponse)

}

func main() {

	test := os.Getenv("SECRET_ENV_TEST")
	fmt.Println(test)

	router := mux.NewRouter()
	port := os.Getenv("PORT")
	router.HandleFunc("/api/v2/generate", generate).Methods("POST")
	router.HandleFunc("/api/v2/refresh", refresh).Methods("POST")
	fmt.Println("GO REST server running on " + port)
	http.ListenAndServe(":"+port, router)

}
