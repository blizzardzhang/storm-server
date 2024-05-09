package tool

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	pwd := HashPwd("code-storm")
	fmt.Println(pwd)
}

func TestCheckPwd(t *testing.T) {
	//$2a$04$rQzCu6O84S0YE/aNIyR0xOVo1UZgkUlS21mGnulBZMkzLv3uuZAEu
	pwd := CheckPwd("$2a$04$rQzCu6O84S0YE/aNIyR0xOVo1UZgkUlS21mGnulBZMkzLv3uuZAEu", "12356")
	fmt.Println(pwd)
}
