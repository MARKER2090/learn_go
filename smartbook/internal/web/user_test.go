package web

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestEncrypt(t *testing.T) {
	password := "helloworld123"
	//对密码进行加密
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}
	/*
		可以理解为将加密后的密码A进行盐值的提取，然后利用盐值对现有的密码B进行加密，加密后的B和密码A进行对比
	*/
	err = bcrypt.CompareHashAndPassword(encrypted, []byte(password))
	assert.NoError(t, err)
}
