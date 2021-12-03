package models

type Credential struct {
	Host     string `gorm:"default:nil"`
	App      string `gorm:"default:nil"`
	User     string `gorm:"default:nil"`
	Password string `gorm:"default:nil"`
}

func (c *Credential) Get() {
}

func (c *Credential) Save() {
	/* Example SQL Query:
	INSERT INTO vault(host, app , user, passwd)
	VALUES ("10.11.1.2", "ssh", "developer", "T6xGy%2z<'c>k");
	*/
}

func (c *Credential) Delete() {
}

func (c *Credential) Update(id uint) {
}

func (c *Credential) Purge(id uint) {
}
