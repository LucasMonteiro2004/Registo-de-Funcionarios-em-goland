package estruturas

import "fmt"

type Funcionario struct {
	idade            int
	nome             string
	horario_trabalho string
	salario          int
	cargo            string
	data_contratacao string
	id               int
}

func print_Funcionarios(funcionarios []Funcionario) {
    for _, funcionario := range funcionarios {
        fmt.Printf("Nome: %s, Cargo: %s\n", funcionario.nome, funcionario.cargo)
    }
}

func cria_Funcionario(idade int, nome string, horario_trabalho string, salario int, cargo string, data_contratacao string, id int) Funcionario {
	funcionario := Funcionario{
		idade:            idade,
		nome:             nome,
		horario_trabalho: horario_trabalho,
		salario:          salario,
		cargo:            cargo,
		data_contratacao: data_contratacao,
		id:               id,
	}
	return funcionario
}

//(23, "Lucas", "09:00 - 18:00", 15000, "Estagiario", "21_03_2015", 3)