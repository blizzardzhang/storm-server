package uniqueid

import (
	"fmt"
	"github.com/sony/sonyflake"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

var flake *sonyflake.Sonyflake

func init() {
	flake = sonyflake.NewSonyflake(sonyflake.Settings{})
}

func GenId() int64 {

	id, err := flake.NextID()
	if err != nil {
		logx.Severef("flake NextID failed with %s \n", err)
		panic(err)
	}
	strId := strconv.FormatInt(int64(id), 10)
	fmt.Printf("%T-%v", strId, strId)
	return int64(id)
}

func GenIdStr() string {

	id, err := flake.NextID()
	if err != nil {
		logx.Severef("flake NextID failed with %s \n", err)
		panic(err)
	}
	return strconv.FormatInt(int64(id), 10)
}
