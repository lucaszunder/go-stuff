package main

type Dictionary map[string]string

const (
	NotFoundError = ErrDictionary("Not Found")
)

type ErrDictionary string

func (e ErrDictionary) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	value, found := d[key]

	if found {
		return value, nil
	}
	return "", NotFoundError
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case NotFoundError:
		d[key] = value
	case nil:
		d[key] = value
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) error {
	delete(d, key)
	return nil
}
