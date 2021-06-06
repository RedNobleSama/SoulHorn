/**
* @Author: oreki
* @Date: 2021/6/5 11:08
* @Email: a912550157@gmail.com
*/

package main

import (
	"SoulHorn/model"
	"SoulHorn/routes"
	"SoulHorn/utils"
	"strings"
)


func main() {
	model.InitDb()
	r := routes.InitRouter()
	//address := strings.Join([]string{utils.Address, utils.HttpPort}, "")
	_ = r.Run(strings.Join([]string{utils.Address, utils.HttpPort}, ""))
}
