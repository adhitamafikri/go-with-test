package dictionary

import "testing"

func TestDictionarySearch(t *testing.T) {
	t.Run("should return correct value if the given key exists in dictionary", func(t *testing.T) {
		dictionary := InitDictionary()
		result, err := dictionary.Search("protege")
		expected := "The student"

		assertNoError(t, err)
		assertValue(t, result, expected)
	})

	t.Run("error should be raised if key is not found in the dictionary", func(t *testing.T) {
		dictionary := InitDictionary()
		_, err := dictionary.Search("skuba")

		assertError(t, err, ErrKeyNotFound)
	})
}

func TestDictionaryAdd(t *testing.T) {
	t.Run("should be able to add a new dictionary key", func(t *testing.T) {
		dictionary := InitDictionary()
		errAdd := dictionary.Add("newgen", "This shi is so easy LMFAO")
		result, errSearch := dictionary.Search("newgen")

		assertNoError(t, errAdd)
		assertNoError(t, errSearch)
		assertValue(t, result, "This shi is so easy LMFAO")
	})

	t.Run("should NOT be able to add a value with the key that already exists", func(t *testing.T) {
		dictionary := InitDictionary()
		err := dictionary.Add("propane", "This is the new value for propane")

		assertError(t, err, ErrDuplicateKey)
	})
}

func TestDictionaryUpdate(t *testing.T) {
	t.Run("should be able to update a dictionary value", func(t *testing.T) {
		dictionary := InitDictionary()
		errUpdate := dictionary.Update("propane", "New value for propane")
		result, errSearch := dictionary.Search("propane")

		assertNoError(t, errUpdate)
		assertNoError(t, errSearch)
		assertValue(t, result, "New value for propane")
	})

	t.Run("should NOT be able to update a nonexistant dictionary value", func(t *testing.T) {
		dictionary := InitDictionary()
		err := dictionary.Update("tm88", "This is the producer for 808 Mafia")

		assertError(t, err, ErrUpdateKeyNotExist)
	})
}

func TestDictionaryRemove(t *testing.T) {
	t.Run("should be able to remove a value from dictionary", func(t *testing.T) {
		dictionary := InitDictionary()
		errRemove := dictionary.Remove("propane")
		_, errSearch := dictionary.Search("propane")

		assertNoError(t, errRemove)
		assertError(t, errSearch, ErrKeyNotFound)
	})

	t.Run("should NOT be able to remove a nonexistant dictionary value", func(t *testing.T) {
		dictionary := InitDictionary()
		err := dictionary.Remove("tm88")

		assertError(t, err, ErrRemoveKeyNotExist)
	})
}

func assertValue(t testing.TB, result string, expected string) {
	t.Helper()
	if result != expected {
		t.Errorf("Expected: %q, got: %q", result, expected)
	}
}

func assertError(t testing.TB, err error, expected error) {
	t.Helper()
	if err == nil {
		t.Fatal("Expected to get error")
	}

	if err != expected {
		t.Errorf("Got error: %s, Expected: %s", err, expected)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("Expected NOT to get error")
	}
}
