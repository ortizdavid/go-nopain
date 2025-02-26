package encryption

import (
    "testing"
    "strings"
)

func TestGenerateUUID(t *testing.T) {
    uuid := GenerateUUID()
    if len(uuid) != 36 {
        t.Errorf("GenerateUUID() returned an invalid UUID: %s", uuid)
    }
}

func TestGenerateRandomToken(t *testing.T) {
    token := GenerateRandomToken(100)
    if len(token) != 136 { // 100 bytes when base64 encoded
        t.Errorf("GenerateRandomToken() returned an invalid token: %s", token)
    }
}

func TestGenerateCode(t *testing.T) {
    prefix := "PREFIX"
    code := GenerateCode(prefix)
    if !strings.HasPrefix(code, prefix) {
        t.Errorf("GenerateCode() did not prepend the prefix: %s", code)
    }
    if len(code) != len(prefix)+14 { // 14 characters for the timestamp
        t.Errorf("GenerateCode() generated an invalid code: %s", code)
    }
}
