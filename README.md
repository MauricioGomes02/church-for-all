# Church For All

Church For All trata-se de um projeto de estudos, com foco em demonstrar a evolução de uma aplicação para gestão de igrejas, buscando alcançar as melhores práticas. A evolução do projeto será de forma progressiva e poderá ser acompanhada através de posts no [Linkedin](https://www.linkedin.com/in/mauricioandradegomes/).

## Índice

- [Sobre o Projeto](#sobre-o-projeto)
- [Como rodar a aplicação](#como-rodar-a-aplicação)
- [Histórico de Versões](#histórico-de-versões)
- [Licença](#licença)
- [Contato](#contato)

## Como rodar a aplicação

1. Certifique-se de ter o Go instalado ([download aqui](https://go.dev/dl/)).
2. No terminal, navegue até a pasta `backend` do projeto.
3. Execute o comando:

```sh
go run main.go
```

4. Acesse a rota da API em [http://localhost:8080/churchs](http://localhost:8080/churchs)
   - Para cadastrar uma igreja, utilize o método **POST** com um payload JSON conforme o exemplo na seção de evolução do projeto.

## Histórico de Versões

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

 