package main

type Dictionary map[Word]string

type Definition string

type Word string

type DictionaryErrs string

const (
	ErrNotFound         = DictionaryErrs("could not find the word you were looking for")
	ErrDuplicateWord    = DictionaryErrs("this word already exists")
	ErrWordDoesNotExist = DictionaryErrs("can't update because this word does not exist")
)

func (d Dictionary) Define(w Word) (string, DictionaryErrs) {
	definition, ok := d[w]

	if !ok {
		return "", ErrNotFound
	}

	return definition, ""
}

func (d Dictionary) Add(word Word, def string) (string, DictionaryErrs) {
	if d[word] != "" {
		return "", ErrDuplicateWord
	}

	d[word] = def

	return d[word], ""
}

func (d Dictionary) Update(word Word, def string) (string, DictionaryErrs) {
	_, err := d.Define(word)

	switch err {
	case ErrNotFound:
		return "", ErrWordDoesNotExist
	case "":
		d[word] = def
	}

	return d[word], err
}

func (d Dictionary) Delete(word Word) {
	delete(d, word)
}
