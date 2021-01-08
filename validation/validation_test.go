package validation

import (
	"crypto/sha256"
	"encoding/base64"
	"testing"
)

type ValidateLoginTest struct {
	Username string
	Password string
	Token    string
	Expected bool
}

func runValidateLogin(username string, password string, token string, expected bool) bool {
	hash := getHash(username, password)
	result, _ := ValidateLogin(hash, token)
	return result == expected
}

func getHash(username string, password string) string {
	h := sha256.New()
	h.Write([]byte(username + password))
	b := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(b)
}

func runValidateLoginTests(tests []ValidateLoginTest, testName string, t *testing.T) {
	for _, test := range tests {
		t.Run(testName, func(t *testing.T) {
			result := runValidateLogin(test.Username, test.Password, test.Token, test.Expected)
			if !result {
				t.Errorf("%s, %s, %s returned %t. Expected %t", test.Username, test.Password, test.Token, test.Expected, result)
			}
		})
	}
}

func TestValidateLoginValidCredentials(t *testing.T) {
	var tests = []ValidateLoginTest{
		{"c137@onecause.com", "#th@nH@rm#y#r!$100%D0p#", "1201", true},
		{"c137@onecause.com", "#th@nH@rm#y#r!$100%D0p#", "1202", true},
	}

	runValidateLoginTests(tests, "TestValidateLoginValidCredentials", t)
}

func TestValidateLoginInvalidCredentials(t *testing.T) {
	var tests = []ValidateLoginTest{
		{"", "#th@nH@rm#y#r!$100%D0p#", "1201", false},                    //empty email
		{"c137@onecause.com", "", "1201", false},                          //empty password
		{"bademail@google.com", "#th@nH@rm#y#r!$100%D0p#", "1201", false}, //invalid email, valid password
		{"c137@onecause.com", "badpassword", "1201", false},               //valid email, invalid password
		{"bademail@google.com", "badpassword", "1201", false},             //invalid both
		{"c137@onecause.com", "#th@nH@rm#y#r!$100%D0p#", "", false},       //empty token
		{"c137@onecause.com", "#th@nH@rm#y#r!$100%D0p#", "aaa", false},    //bad token format
		{"c137@onecause.com", "#th@nH@rm#y#r!$100%D0p#", "123456", false}, //token too long
	}

	runValidateLoginTests(tests, "TestValidateLoginInvalidCredentials", t)
}

func TestValidateLoginInvalidReuseToken(t *testing.T) {
	username := "c137@onecause.com"
	password := "#th@nH@rm#y#r!$100%D0p#"
	token := "1201"

	runValidateLogin(username, password, token, true)
	// should be invalid now with same token
	if !runValidateLogin(username, password, token, false) {
		t.Errorf("reused token, still valid login")
	}
}
