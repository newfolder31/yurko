package profession


type ClientRepository interface {
	Store(user *Client) error
	FindById(id int) (*Client, error)
}

type Client struct {
	Id int
}
