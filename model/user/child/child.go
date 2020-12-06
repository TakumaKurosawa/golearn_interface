package child

import (
	"fmt"
	"golearninterface/interface/user"
	"golearninterface/model/user/entity"
)

type GenderType = int

const (
	Man     GenderType = 1
	Woman   GenderType = 2
	Unknown GenderType = 0
)

type Child struct {
	entity.UserData
	Gender GenderType `json:"gender"`
}

func New(userData entity.UserData, gender GenderType) user.UserInterface {
	return &Child{
		UserData: userData,
		Gender:   gender,
	}
}

func (c *Child) Print() {
	if c.Age > 18 {
		fmt.Println("19歳以上の方はご利用になれません．")
		return
	}

	switch c.Gender {
	case Man:
		fmt.Printf("ようこそ！%sさん（♂）．あなたの誕生日は%sですね！\n", c.Name, c.Birthday)
	case Woman:
		fmt.Printf("ようこそ！%sさん（♀）．あなたの誕生日は%sですね！\n", c.Name, c.Birthday)
	case Unknown:
		fmt.Printf("ようこそ！%sさん．あなたの誕生日は%sですね！\n", c.Name, c.Birthday)
	}
}

func (c *Child) UpdateBaseData(data entity.UserData) {
	if c.Age > 18 {
		fmt.Println("19歳以上の方はご利用になれません．")
		return
	}
	c.UserData = data
}
