## Cadastro Básico em Go: API RESTful para gerenciamento de clientes

**Descrição**

Este projeto é uma API RESTful simples desenvolvida em Go para gerenciar cadastros de clientes. Ele serve como um ponto de partida para projetos mais complexos, demonstrando a implementação de funcionalidades CRUD (Create, Read, Update, Delete) utilizando as tecnologias Go, MySQL e Gorilla Mux.

**Funcionalidades:**

* **Criar um novo cliente:** Envia uma requisição POST para `/cadastro` com os dados do cliente no corpo da requisição (formato JSON).
* **Listar todos os clientes:** Envia uma requisição GET para `/clientes` para obter uma lista de todos os clientes cadastrados.
* **Atualizar um cliente:** Envia uma requisição PUT para `/atualizar` com os dados atualizados do cliente no corpo da requisição.
* **Deletar um cliente:** Envia uma requisição DELETE para `/deletar` com o ID do cliente no corpo da requisição.

**Tecnologias Utilizadas:**

* **Go:** Linguagem de programação utilizada para o desenvolvimento da API.
* **MySQL:** Banco de dados relacional para armazenar as informações dos clientes.
* **Gorilla Mux:** Roteador HTTP para Go, utilizado para definir as rotas da API.
* **JSON:** Formato de troca de dados utilizado para a comunicação entre o cliente e o servidor.

**Estrutura do Projeto:**

* **`src`:** Contém o código fonte da aplicação, organizado em pacotes para facilitar a manutenção.
  * `controller`: Contém os controladores que manipulam as requisições HTTP e interagem com o modelo de dados.
  * `db`: Contém a lógica de acesso ao banco de dados MySQL.
* **`main.go`:** Ponto de entrada da aplicação, onde o servidor HTTP é iniciado e as rotas são configuradas.


