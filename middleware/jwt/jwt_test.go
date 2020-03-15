package jwt

import (
	"reflect"
	"testing"
	"time"
)

func TestJWT(t *testing.T) {
	testKey := "test-key"
	testIssuer := "test-issuer"

	expectedClaims := &Claims{
		UUID:  "uuid",
		Email: "email",
	}

	token, _ := GenerateJWT(expectedClaims, testKey, testIssuer, time.Hour*2)

	actualClaims, err := ParseJWT(token, testKey)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(actualClaims, expectedClaims) {
		t.Errorf("want : %v", expectedClaims)
		t.Errorf("got  : %v", actualClaims)
	}
}
