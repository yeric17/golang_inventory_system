package responses

import (
	"encoding/json"
	"fmt"
	"io"
)

//IEncodeJSON es una interface que tiene el metodo ToJSON
type IEncodeJSON interface {
	ToJSON(w io.Writer, i interface{}) error
}

//IDecodeJSON es una interface que tiene el metodo FromJSON
type IDecodeJSON interface {
	FromJSON(r io.Reader, i interface{}) error
}

//EncodeJSON estructura que implementa la interface IEncodeJSON
type EncodeJSON struct{}

//DecodeJSON estructura que implementa l interface IDecodeJSON
type DecodeJSON struct{}

//ToJSON imprime al usuario información de structuras en JSON
func (ej EncodeJSON) ToJSON(w io.Writer, i interface{}) error {
	var err error
	encode := json.NewEncoder(w)
	err = encode.Encode(i)
	if err != nil {
		fmt.Printf("[Error] Al generar encode de un error JSON: %s", err.Error())
		return err
	}
	return nil
}

//FromJSON transforma a JSON
func (ej *DecodeJSON) FromJSON(r io.Reader, i interface{}) error {
	var err error
	decode := json.NewDecoder(r)
	err = decode.Decode(i)
	if err != nil {
		fmt.Printf("[Error] Al generar dencode de un error JSON: %s", err.Error())
		return err
	}
	return nil
}

func ToJSON(w io.Writer, ej IEncodeJSON) {
	ej.ToJSON(w, ej)
}

func FromJSON(w io.Reader, dj IDecodeJSON) {
	dj.FromJSON(w, dj)
}
