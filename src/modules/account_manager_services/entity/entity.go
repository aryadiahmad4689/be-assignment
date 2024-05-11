package entity

type User struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"-"`
	Gender     string `json:"-"`
	Email      string `json:"email"`
	Password   string `json:"-"`
	IsVerified int    `json:"-"`
}

type AccountPayment struct {
	Id     string
	Name   string
	Type   string
	UserId string `json:"-"`
	Amount float64
}

type SingUpReq struct {
	Id          string
	Name        string
	Age         int
	Gender      string
	Email       string
	Password    string
	IsVerified  int
	NameAccount string
	Type        string
	UserId      string
	Amount      float64
}

type UserWithPayments struct {
	User
	Payments []AccountPayment
}
