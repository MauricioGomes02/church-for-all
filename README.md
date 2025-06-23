# Church For All

Church For All trata-se de um projeto de estudos, com foco em demonstrar a evolução de uma aplicação para gestão de igrejas, buscando alcançar as melhores práticas. A evolução do projeto será de forma progressiva e poderá ser acompanhada através de posts no [Linkedin](https://www.linkedin.com/in/mauricioandradegomes/).

## Índice

- [Sobre o Projeto](#church-for-all)
- [Evolução do Projeto](#evolução-do-projeto)
- [Como rodar a aplicação](#como-rodar-a-aplicação)
- [Licença](#licença)
- [Contato](#contato)

## Evolução do Projeto

Acompanhe aqui as principais etapas e avanços deste projeto:

- Estruturação inicial do repositório e documentação.
- Criação de uma API em Go utilizando apenas bibliotecas internas.
- Implementação da rota `/churchs` para cadastro (POST) de igrejas.
- Definição do payload para cadastro de igreja:

```json
{
  "id": 1,
  "name": "Igreja Central",
  "address": "Rua das Flores, 123",
  "phones": ["11999999999", "11888888888"],
  "email": "contato@igreja.com"
}
```

Novas funcionalidades e melhorias serão registradas nesta seção conforme o projeto evolui.

## Como rodar a aplicação

1. Certifique-se de ter o Go instalado ([download aqui](https://go.dev/dl/)).
2. No terminal, navegue até a pasta `backend` do projeto.
3. Execute o comando:

```sh
go run main.go
```

4. Acesse a rota da API em [http://localhost:8080/churchs](http://localhost:8080/churchs)
   - Para cadastrar uma igreja, utilize o método **POST** com um payload JSON conforme o exemplo na seção de evolução do projeto.

## Licença

Este projeto está licenciado sob a licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes. 

## Contato

Maurício Andrade Gomes - mauricioandradegomes@gmail.com

[LinkedIn](https://www.linkedin.com/in/mauricioandradegomes/)

 