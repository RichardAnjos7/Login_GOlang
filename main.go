package main

import (
	"fmt"
	"net/http"
)

// Usuário e senha para efetuar o login!
const (
	validUsername = "admin"
	validPassword = "123"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {

			// Obtém os valores do formulário que retorna o valor associado a variável username e password.
			username := r.FormValue("username")
			password := r.FormValue("password")

			// Verifica se as credenciais são válidas
			if username == validUsername && password == validPassword {
				fmt.Fprintln(w, `
			<!DOCTYPE html>
				<html>

				<head>
					<title>Autenticação</title>

					<style>

        				body{
            				font-family: Arial, Helvetica, sans-serif;
            				background-image: linear-gradient(45deg, cyan, yellow);
        				}

        				form{
            				background-color: rgba(0, 0, 0, 0.9);
            				position: absolute;
            				top: 50%;
            				left: 50%;
            				transform: translate(-50%,-50%);
            				padding: 80px;
            				border-radius: 15px;
            				color: #fff;
        				}
     				
    				</style>
				</head>

				<body>
					<form method="post">
						<h1>Login bem-sucedido! Você está autenticado.</h1>
					</form>
				</body>

				</html>
			`)

			} else {
				fmt.Fprintln(w, `
			  <!DOCTYPE html>
			  	<html>

				<head>
					<title>Erro</title>

					<script>
        					function redirectToPage() {
            					window.location.href = "/";
        					}
    				</script>

					<style>

        				body{
            				font-family: Arial, Helvetica, sans-serif;
            				background-image: linear-gradient(45deg, cyan, yellow);
        				}

        				form{
            				background-color: rgba(0, 0, 0, 0.9);
            				position: absolute;
            				top: 50%;
            				left: 50%;
            				transform: translate(-50%,-50%);
            				padding: 80px;
            				border-radius: 15px;
            				color: #fff;
        				}
     				
    				</style>
				</head>

				<body>
					<form method="post">
						<h1>Credenciais inválidas! Tente novamente.</h1>
						<br><br>
						<button type="button" onclick="redirectToPage()" style="background-color: white; border: none; color: black;">Voltar</button>


					</form>
				</body>

				</html>
			`)

			}

		} else {
			fmt.Fprintln(w, `
			<!DOCTYPE html>
				<html>

				<head>
					<title>Login</title>

					<style>

        				body{
            				font-family: Arial, Helvetica, sans-serif;
            				background-image: linear-gradient(45deg, cyan, yellow);
        				}

        				form{
            				background-color: rgba(0, 0, 0, 0.9);
            				position: absolute;
            				top: 50%;
            				left: 50%;
            				transform: translate(-50%,-50%);
            				padding: 80px;
            				border-radius: 15px;
            				color: #fff;
        				}

        				input{
            				padding: 15px;
            				border: none;
            				outline: none;
            				font-size: 15px;
        				}

        				button{
            				background-color: dodgerblue;
            				border: none;
            				padding: 15px;
            				width: 100%;
            				border-radius: 10px;
            				color: white;
            				font-size: 15px;
                    	}

        				button:hover{
            				background-color: deepskyblue;
            				cursor: pointer;
        				}
    				</style>
				</head>

				<body>
					
					<form method="post">
						<h1>Login</h1>
						<input type="text" name="username" placeholder="Usuário" required>
						<br><br>
						<input type="password" name="password" placeholder="Senha" required>
						<br><br>
						<button type="submit">Entrar</button>

					</form>
				</body>

				</html>
			`)
		}
	})

	// Essa linha imprime a mensagem “Servidor rodando em http://localhost:8080” no console.
	fmt.Println("Servidor rodando em http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
