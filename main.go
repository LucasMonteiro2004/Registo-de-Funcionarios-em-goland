package main

import "fmt"

type Funcionario struct {
	idade            int
	nif              int
	nome             string
	horario_trabalho string
	salario          int
	cargo            string
	data_contratacao string
	id               int
}

type Equipa struct {
	funcionarios           []Funcionario
	numero_de_funcionarios int
	id                     int
	nome                   string
	funcao                 string
}

var equipes []Equipa

func print_Funcionarios(funcionarios []Funcionario) {
	for _, funcionario := range funcionarios {
		fmt.Printf("Idade %d, Nome %s, Horario %s, salario %d, Cargo %s, Inicio de contrato %s, id %d, nif %d\n", funcionario.idade, funcionario.nome, funcionario.horario_trabalho, funcionario.salario, funcionario.cargo, funcionario.data_contratacao, funcionario.id, funcionario.nif)
	}
}

func cria_Funcionario(idade int, nome string, horario_trabalho string, salario int, cargo string, data_contratacao string, id int, nif int) Funcionario {
	funcionario := Funcionario{
		idade:            idade,
		nome:             nome,
		horario_trabalho: horario_trabalho,
		salario:          salario,
		cargo:            cargo,
		data_contratacao: data_contratacao,
		id:               id,
		nif:              nif,
	}
	return funcionario
}

func adiciona_funcionario_equipe(pessoa Funcionario, id int) {
	// Verificar se a equipe já existe
	var equipeExistente bool
	for i := range equipes {
		if id == equipes[i].id {
			equipes[i].funcionarios = append(equipes[i].funcionarios, pessoa)
			equipeExistente = true
			break
		}
	}

	// Se a equipe não existir, criar uma nova equipe
	if !equipeExistente {
		novaEquipe := Equipa{id: id}
		novaEquipe.funcionarios = append(novaEquipe.funcionarios, pessoa)
		equipes = append(equipes, novaEquipe)
	}
}

func removePessoa(idEquipe, idPessoa int) {
	// Verificar se a equipe existe
	var equipeIndex int
	var equipeExistente bool
	for i, equipe := range equipes {
		if idEquipe == equipe.id {
			equipeExistente = true
			equipeIndex = i
			break
		}
	}

	if !equipeExistente {
		fmt.Println("Equipe nao encontrada para o ID:", idEquipe)
		return
	}

	// Remover a pessoa da fatia
	for j, funcionario := range equipes[equipeIndex].funcionarios {
		if idPessoa == funcionario.id {
			equipes[equipeIndex].funcionarios = append(equipes[equipeIndex].funcionarios[:j], equipes[equipeIndex].funcionarios[j+1:]...)
			fmt.Println("Pessoa removida com sucesso.")
			return
		}
	}

	fmt.Println("Pessoa nao encontrada para o ID:", idPessoa)
}

func main() {
	// Criação de uma equipe e adição de funcionários
	var vec []Funcionario
	equipe1 := Equipa{funcionarios: vec, numero_de_funcionarios: 0, id: 1, nome: "Equipe A", funcao: "Des"}
	funcionario1 := cria_Funcionario(25, "João", "9h-17h", 5000, "Desenvolvedor", "01/01/2022", 1, 123456789)
	funcionario2 := cria_Funcionario(30, "Maria", "10h-18h", 6000, "Analista", "01/02/2022", 2, 987654321)

	adiciona_funcionario_equipe(funcionario1, 1)
	adiciona_funcionario_equipe(funcionario2, 1)

	// Impressão dos funcionários antes da remoção
	fmt.Println("Funcionários da Equipe A antes da remoção:")
	print_Funcionarios(equipe1.funcionarios)

	// Remoção de um funcionário pelo ID da equipe e da pessoa
	removePessoa(1, 2)

	// Impressão dos funcionários após a remoção
	fmt.Println("Funcionários da Equipe A após a remoção:")
	print_Funcionarios(equipe1.funcionarios)
}
