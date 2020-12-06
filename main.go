package main

import (
	"fmt"
	"golearninterface/model/user/adult"
	"golearninterface/model/user/child"
	"golearninterface/model/user/entity"
)

func main() {
	fmt.Println("-------- Child --------")
	childUser := child.New(
		entity.UserData{
			Name:     "kuro",
			Age:      18,
			Birthday: "2001/12/07",
		},
		child.Man,
	)
	childUser.Print()
	childUser.UpdateBaseData(
		entity.UserData{
			Name:     "黒澤",
			Age:      17,
			Birthday: "1998/03/07",
		},
	)
	childUser.Print()

	fmt.Println("-------- Adult --------")
	adultUser := adult.New(
		entity.UserData{
			Name:     "Shiro",
			Age:      22,
			Birthday: "1998/12/07",
		},
		"080-1111-2222",
		true,
	)
	adultUser.Print()
	adultUser.UpdateBaseData(entity.UserData{
		Name:     "白澤",
		Age:      23,
		Birthday: "1997/12/07",
	})
	adultUser.Print()
}
