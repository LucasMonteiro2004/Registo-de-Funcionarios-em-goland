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

func cria_equipe(funcionarios []Funcionario, n_funcionarios int, id int, nome string, funcao string) {
	// Verificar se a equipe já existe com base no ID e nome
	for _, equipe := range equipes {
		if id == equipe.id || nome == equipe.nome {
			fmt.Println("Equipe já existe. Não será adicionada.")
			return
		}
	}

	// Criar a nova equipe e adicioná-la à lista de equipes
	newEquipa := Equipa{
		funcionarios:           funcionarios,
		numero_de_funcionarios: n_funcionarios,
		id:                     id,
		nome:                   nome,
		funcao:                 funcao,
	}

	equipes = append(equipes, newEquipa)

	fmt.Println("Equipe adicionada com sucesso.")
}

func print_all_equipes(){
	for _, equipa := range equipes{
		print_Funcionarios(equipa.funcionarios)
	}
}

func adiciona_funcionario_equipa(pessoa Funcionario, id int) {
	// Verificar se a equipe já existe
	var equipeExistente bool
	for _, equipe := range equipes{
		if id == equipe.id {
			equipeExistente = true
			equipe.funcionarios = append(equipe.funcionarios, pessoa)
			fmt.Println("Funcionario adicionado a equipe com sucesso!")
			print_Funcionarios(equipe.funcionarios)
		}
	}

	// Se a equipe não existir, criar uma nova equipe
	if !equipeExistente {
		fmt.Println("Equipe não encontrada")
	}
}

func removePessoa(idEquipe, idFunc int) {
	for i, equipe := range equipes {
		if equipe.id == idEquipe {
			for j, funcionario := range equipe.funcionarios {
				if funcionario.id == idFunc {
					// Remover o funcionário do slice usando técnicas de slicing
					equipes[i].funcionarios = append(equipes[i].funcionarios[:j], equipes[i].funcionarios[j+1:]...)
					fmt.Println("Funcionário removido da equipe com sucesso!")
					print_Funcionarios(equipes[i].funcionarios)
					return
				}
			}
			fmt.Println("Funcionário não encontrado na equipe.")
			return
		}
	}
	fmt.Println("Equipe não encontrada.")
}

func procurar_equipe_para_fun(fun Funcionario, id int){
	for _, equipa := range equipes{
		if equipa.id == id{
			adiciona_funcionario_equipa(fun, id)
		}
	}
}

func main() {

	var eq []Funcionario
	cria_equipe(eq, 0, 1, "Equipe A", "Design")

	var funcionario_1 = Funcionario{idade: 23, nome: "Lucas", horario_trabalho: "09:00 - 18:00", salario: 13000, cargo: "Estagiario", data_contratacao: "21_03_2023", id: 2, nif: 91834526}

	var funcionario_2 = Funcionario{idade: 20, nome: "Miguel", horario_trabalho: "09:00 - 18:00", salario: 16000, cargo: "junior", data_contratacao: "27_05_2022", id: 1, nif: 97165347}
	
	//eq = append(eq, funcionario_1)
	adiciona_funcionario_equipa(funcionario_1, 1)
	adiciona_funcionario_equipa(funcionario_2, 1)

	removePessoa(1, 1)
}
