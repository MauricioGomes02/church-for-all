# Church For All

Church For All trata-se de um projeto de estudos, com foco em demonstrar a evolução de uma aplicação para gestão de igrejas, buscando alcançar as melhores práticas. A evolução do projeto será de forma progressiva e poderá ser acompanhada através de posts no [Linkedin](https://www.linkedin.com/in/mauricioandradegomes/).

## Índice

- [Sobre o Projeto](#sobre-o-projeto)
- [Como rodar a aplicação](#como-rodar-a-aplicação)
- [Histórico de Versões](#histórico-de-versões)
- [Licença](#licença)
- [Contato](#contato)

## Como rodar a aplicação

1. Certifique-se de ter o Docker e o Go instalados.
2. No terminal, navegue até a pasta `backend` do projeto.
3. Suba o banco de dados MySQL com:

```sh
docker-compose up -d
```

4. Execute a aplicação Go:

```sh
go run cmd/api/main.go
```

5. Acesse a rota da API em [http://localhost:8080/churchs](http://localhost:8080/churchs)
   - Para cadastrar uma igreja, utilize o método **POST**
   - Para listar igrejas, utilize o método **GET**.

## Histórico de Versões

### v1.0.0
-Reestruturação completa do projeto, seguindo a estrutura recomendada pela comunidade Go, com ajustes para o contexto da aplicação.

- Adoção do framework Gin para controle e manipulação mais eficiente da API.

- Criação de testes iniciais, promovendo maior confiabilidade no código e facilitando a evolução do projeto.

- Abstração de dependências por meio de interfaces, permitindo testes de unidade mais eficazes e maior flexibilidade na implementação.

- Segregação de interfaces conforme o princípio da Segregação de Interfaces (ISP) do SOLID, garantindo que clientes dependam apenas do que realmente utilizam.

- Melhoria geral na organização do código, visando facilitar a manutenção e evolução da aplicação.

### v0.1.1
- Refatoração do código seguindo o princípio da Responsabilidade Única (SRP) do SOLID, separando as responsabilidades em diferentes funções, arquivos ou camadas para melhorar a organização e manutenção do projeto.

### v0.1.0
- Integração da API com banco de dados MySQL.
- Adição de um arquivo docker-compose para facilitar o setup do banco.
- Agora os dados de igrejas são persistidos no banco de dados.
- Scripts de inicialização SQL podem ser executados automaticamente ao subir o banco.

### v0.0.0
- Estruturação inicial do projeto e documentação básica.
- Implementação da rota POST em `/churchs` para cadastro de igrejas.
- Implementação da rota GET em `/churchs` para listar todas as igrejas cadastradas.
- Melhoria nos comentários do código para facilitar o entendimento.

Veja todas as versões e downloads em [Releases do GitHub](https://github.com/MauricioGomes02/church-for-all/releases)

## Licença

Este projeto está licenciado sob a licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes. 

## Contato

Maurício Andrade Gomes - mauricioandradegomes@gmail.com

[LinkedIn](https://www.linkedin.com/in/mauricioandradegomes/)

 