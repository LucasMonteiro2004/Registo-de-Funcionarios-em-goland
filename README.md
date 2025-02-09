# Gerenciamento de Equipes e Funcionários em Go

Este projeto implementa um sistema simples de gerenciamento de equipes e funcionários utilizando a linguagem Go.

## 📌 Funcionalidades

- Criar funcionários com atributos como idade, nome, cargo, salário, etc.
- Criar equipes e associar funcionários a elas.
- Listar funcionários de uma equipe.
- Adicionar e remover funcionários de equipes.

## 📂 Estrutura do Código

### 1. Estruturas de Dados
- **Funcionario**: Representa um funcionário com atributos como idade, NIF, nome, cargo, etc.
- **Equipa**: Representa uma equipe contendo um grupo de funcionários e informações como nome e função.

### 2. Funções Principais
- `cria_Funcionario(...)` → Cria um funcionário.
- `cria_equipe(...)` → Cria uma equipe e adiciona à lista de equipes.
- `print_Funcionarios(...)` → Exibe a lista de funcionários.
- `adiciona_funcionario_equipa(...)` → Adiciona um funcionário a uma equipe existente.
- `removePessoa(...)` → Remove um funcionário de uma equipe específica.
- `print_all_equipes()` → Lista todos os funcionários de todas as equipes.

## 🚀 Como Executar

1. Clone o repositório:
   ```bash
   git clone <URL_DO_REPOSITORIO>
   cd <NOME_DO_PROJETO>
   ```
2. Compile e execute o programa:
   ```bash
   go run main.go
   ```

## 📝 Exemplo de Uso

```go
var funcionario_1 = Funcionario{idade: 23, nome: "Lucas", horario_trabalho: "09:00 - 18:00", salario: 13000, cargo: "Estagiário", data_contratacao: "21_03_2023", id: 2, nif: 91834526}
adiciona_funcionario_equipa(funcionario_1, 1)
removePessoa(1, 2)
```

## 📜 Licença

---
📌 Desenvolvido para fins educacionais.

