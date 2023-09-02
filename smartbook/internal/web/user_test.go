/*
该程序就是一个单元测试文件
如果想要：
1.测试该文件，用go test user_test.go
2.测试所有文件，用go test .
*/
package web

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestEncrypt(t *testing.T) {
	password := "helloworld123"
	//对密码进行加密,encrypted就是加密后的字符
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
