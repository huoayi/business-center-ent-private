package enum

type OrderStatus string

const (
	Canceled OrderStatus = "canceled"
	Created  OrderStatus = "created"
	Payed    OrderStatus = "Payed"
	Succeed  OrderStatus = "succeed"
)

func (obj OrderStatus) Values() []string {
	return []string{
		string(Canceled),
		string(Created),
		string(Payed),
		string(Succeed),
	}
}
func (obj OrderStatus) Ptr() *OrderStatus {
	if obj != "" {
		return &obj
	} else {
		return nil
	}
}
