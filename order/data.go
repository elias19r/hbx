package order

var data []*Order
var nextID = 1

func init() {
	// Initialize data skipping the zero index.
	data = append(data, nil)

	// Load data from Redis if any.
	// TODO
}

func NextID() int {
	return nextID
}

func Insert(o *Order) error {
	// Validations.
	if err := o.Validation(); err != nil {
		return err
	}
	if o.state != StateNew {
		return ErrMustHaveStateNew
	}

	data = append(data, o)
	id = len(data) - 1

	o.id = id
	return nil
}

func FindByID(id int) (*Order, error) {
	if id < 0 || id >= len(data) {
		return nil, ErrOrderNotFound
	}
	return data[id], nil
}
