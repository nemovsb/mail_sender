package app

type Recipient struct {
	MailAddr string `form:"mailaddr"`
	Name     string `form:"name"`
	Surname  string `form:"surname"`
	Birthday string `form:"birthday"`
}
