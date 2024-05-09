package valueobject

import (
	"errors"
	"strconv"
	"strings"
)

type CPF struct {
	Value string
}

func NewCPF(value string) (*CPF, error) {
	cpf := &CPF{
		Value: value,
	}

	err := cpf.Validate()
	if err != nil {
		return nil, err
	}

	return cpf, nil
}

func (c *CPF) GetValue() string {
	return c.Value
}

func (c *CPF) Validate() error {
	cpf := c.Value
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	if len(cpf) != 11 {
		return errors.New("o CPF deve conter 11 dígitos")
	}

	// Verifica se todos os dígitos são iguais
	if allSameDigits(cpf) {
		return errors.New("CPF inválido [mesmo dígito repetido]")
	}

	d1, d2 := 0, 0
	for i, r := range cpf[:9] {
		num, _ := strconv.Atoi(string(r))
		d1 += num * (10 - i)
		d2 += num * (11 - i)
	}

	d1 = (d1 * 10) % 11
	if d1 == 10 {
		d1 = 0
	}

	if d1 != int(cpf[9]-'0') {
		return errors.New("CPF inválido [dígito verificador 1]")
	}

	d2 += d1 * 2
	d2 = (d2 * 10) % 11
	if d2 == 10 {
		d2 = 0
	}
	if d2 != int(cpf[10]-'0') {
		return errors.New("CPF inválido [dígito verificador 2]")
	}

	return nil
}

func allSameDigits(cpf string) bool {
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			return false
		}
	}
	return true
}
