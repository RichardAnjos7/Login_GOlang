# README: Autenticação Simples com Go

Este projeto implementa um servidor HTTP básico em Go que oferece autenticação simples por meio de um formulário de login. Os usuários podem inserir um nome de usuário e senha para acessar o servidor.

## Funcionalidades

- **Autenticação de Usuário**:
  - O servidor verifica se as credenciais inseridas no formulário são válidas (usuário e senha).
  - Se as credenciais forem corretas, o usuário é autenticado e pode acessar o servidor.
  - Caso contrário, o acesso é negado.

## Como Usar

1. **Instalação**:
   - Certifique-se de ter o Go instalado em sua máquina.
   - Copie e cole este repositório para o seu diretório local.

2. **Execução**:
   - Navegue até o diretório do projeto.
   - Execute o seguinte comando para iniciar o servidor:
     ```
     go run main.go
     ```

3. **Acesso ao Servidor**:
   - Abra o navegador e acesse `http://localhost:8080`.
   - Você verá um formulário de login.
   - Insira o nome de usuário e senha corretos (usuário: "admin", senha: "123").
   - Se as credenciais forem válidas, você verá a mensagem "Servidor rodando em http://localhost:8080" no console.

## Créditos

Este projeto foi criado por Richard dos Anjos Oliveira.
