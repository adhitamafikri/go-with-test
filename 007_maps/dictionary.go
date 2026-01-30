package dictionary

const (
	ErrKeyNotFound       = DictionaryErr("Key not found in the given dictionary")
	ErrDuplicateKey      = DictionaryErr("Key already exists in the dictionary")
	ErrUpdateKeyNotExist = DictionaryErr("Cannot update value that does not exist in dictionary")
	ErrRemoveKeyNotExist = DictionaryErr("Cannot update value that does not exist in dictionary")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func InitDictionary() Dictionary {
	return Dictionary{
		"protege": "The student",
		"propane": "Some kind of chemical gas substance",
		"hockey":  "Some kind of sport",
	}
}

func (d Dictionary) Search(key string) (string, error) {
	value, found := d[key]

	if !found {
		return "", ErrKeyNotFound
	}

	return value, nil
}

func (d Dictionary) Add(key string, value string) error {
	_, found := d[key]

	if found {
		return ErrDuplicateKey
	}

	d[key] = value
	return nil
}

func (d Dictionary) Update(key string, value string) error {
	_, found := d[key]

	if !found {
		return ErrUpdateKeyNotExist
	}

	d[key] = value
	return nil
}

func (d Dictionary) Remove(key string) error {
	_, found := d[key]

	if !found {
		return ErrRemoveKeyNotExist
	}

	delete(d, key)
	return nil
}
