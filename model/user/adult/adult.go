package adult

import (
	"encoding/json"
	"fmt"
	"golearninterface/interface/user"
	"golearninterface/model/user/entity"
	"io/ioutil"
	"log"
)

const filename = "user.json"

type Adult struct {
	entity.UserData
	Phone   string `json:"phone"`
	Married bool   `json:"married"`
}

func New(userData entity.UserData, phone string, married bool) user.UserInterface {
	return &Adult{
		UserData: userData,
		Phone:    phone,
		Married:  married,
	}
}

func (u *Adult) Print() {
	if u.Age <= 18 {
		fmt.Println("18歳以下の方はご利用になれません．")
		return
	}
	file, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	if err := ioutil.WriteFile(filename, file, 0644); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s様のデータを %s に保存しました．\n", u.Name, filename)
}

func (u *Adult) UpdateBaseData(data entity.UserData) {
	if u.Age <= 18 {
		fmt.Println("18歳以下の方はご利用になれません．")
		return
	}
	u.UserData = data
}
