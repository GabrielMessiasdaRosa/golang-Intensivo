package entity // o pacote recebe o mesmo nome da pasta semelhante ao python, C# e Java.
import ("errors")


type Order struct { // structs em go são semelhantes a classes em outras linguagens.
	ID string
	Price float64
	Tax float64
	FinalPrice float64
}

// pra criar um metodo em go, basta criar uma função com o nome da struct e o nome do metodo.
// o * significa que o metodo é um ponteiro, ou seja, ele vai alterar o valor da struct.
// se não colocar o * o metodo não vai alterar o valor da struct. criando uma copia dela.

// o metodo CalculateFinalPrice() recebe um ponteiro para Order e não retorna nada. Por enquanto.


func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := &Order{
		ID: id,
		Price: price,
		Tax: tax,
	}
	order.CalculateFinalPrice()
	err := order.Validate()
	if err != nil {
 		return nil, err
	}
	return order, nil
}

func (order *Order) Validate() error {
	if order.ID == "" {
		return errors.New("ID is empty")
	}
	if order.Price <= 0 {
		return errors.New("Price is invalid")
	}
	if order.Tax < 0 {
		return errors.New("Tax is invalid")
	}
	return nil
}
func (order *Order) CalculateFinalPrice() error {
	if err := order.Validate(); err != nil {
		return errors.New("Final price is invalid")
	}
	order.FinalPrice = order.Price + order.Tax
	return nil
}