# 🔋🚲 e-bike rental service
Uma aplicação para gerenciar o aluguel de bicicletas elétricas, onde usuários podem registrar-se, localizar bicicletas disponíveis, fazer reservas, iniciar e finalizar aluguéis, além de avaliar as bicicletas após o uso. Administradores podem gerenciar a frota de bicicletas, visualizar relatórios de uso e lidar com manutenção.

## Funcionalidades principais
#### Autenticação e autorização:
- Registro de usuários.
- Login e logout.
- Autorização baseada em roles (usuário comum e administrador).
- Validar conta do usuário.

#### Gerenciamento de usuário:
- Editar dados pessoais.
  - Foto de perfil
  - Nome e e-mail
  - Informações de pagamento
  - Informações de contato
- Alterar senha.
- Deletar conta.
- Histórico de aluguéis.

#### Gerenciamento de bicicletas:
- Registro e remoção de bicicletas.
- Atualização de status das bicicletas (disponível, em manutenção, alugado).

#### Aluguéis:
- Localização de bicicletas disponíveis.
- Início e término de aluguéis.
- Cálculo automático de tarifas.

## Regras de negócio
#### Cadastro e login:
- O e-mail do usuário deve ser único.
- Senhas devem ter no mínimo 8 caracteres, incluindo uma letra maiúscula, uma letra minúscula e um número.
- Usuários não podem alugar bicicletas se houverem pendências financeiras.

#### Reserva de bicicletas:
- Bicicletas podem ser reservadas por até 15 minutos antes do início do aluguel.
- Se a bicicleta não for alugada dentro de 15 minutos, a reserva é cancelada automaticamente.

#### Início e término de aluguéis:
- O aluguel começa quando o usuário desbloqueia a bicicleta.
- O aluguel termina quando a bicicleta é devolvida a uma estação de devolução.
- O custo do aluguel é calculado com base no tempo de uso.

## Requisitos técnicos
### Endpoints da API:
#### Autenticação e autorização:
- `POST /v1/auth/register`: Registrar um novo usuário. ✅
- `POST /v1/auth/login`: Autenticar um usuário. ✅
- `POST /v1/auth/logout`: Deslogar um usuário. ✅
- `POST /v1/auth/verify-account/{token}`: Verificar conta de usuário. ✅
- `POST /v1/auth/forgot-password`: Gera um token para o usuário recuperar a senha. ✅
- `POST /v1/auth/reset-password/{token}`: Cria uma nova senha para conta do usuário baseada no token de recuperação. ✅

#### Gerenciamento de usuários:
- `GET /v1/admin/users`: Obter detalhes de todos os usuários. ✅
- `GET /v1/users/{id}`: Obter detalhes de um usuário. ✅
- `PUT /v1/users/{id}`: Atualizar detalhes de um usuário. ✅
- `PUT /v1/users/{id}/password`: Atualizar a senha de um usuário. ✅
- `DELETE /v1/users/{id}/delete`: Deletar a conta de um usuário. ✅

#### Gerenciamento de bicicletas:
- `POST /v1/admin/bikes/`: Adicionar uma nova bicicleta. ✅
- `PUT /v1/admin/bikes/{id}`: Atualizar detalhes de uma bicicleta. ✅
- `DELETE /v1/admin/bikes/{id}`: Remover uma bicicleta. ✅
- `GET /v1/bikes`: Listar todas as bicicletas. ✅
- `GET /v1/bikes/{id}`: Obter detalhes de uma bicicleta. ✅

#### Reservas e aluguéis:
- `POST /v1/rentals/rent/{bikeId}`: Iniciar o aluguel de uma bicicleta. ✅
- `POST /v1/rentals/return/{rentalId}`: Finalizar o aluguel de uma bicicleta. ✅
- `GET /v1/rentals/{userId}`: Lista os aluguéis do usuário. ✅
- `GET /v1/admin/rentals`: Lista todos os aluguéis da plataforma. ✅

### Persistência de dados:
- Banco de dados relacional: `PostgreSQL`.
- Tabelas sugeridas: `users`, `bikes`, `rentals`, `reviews`, `maintenance_logs`.

### Autenticação e segurança:
- JWT (JSON Web Tokens) para autenticação.
- Senhas armazenadas com hashing seguro (e.g., bcrypt).


## Instalação
Para usar este projeto, você precisa seguir estas etapas:

1. Clone o repositório: `git clone https://github.com/vinniciusgomes/mybooks-api`
2. Instale as dependências: `go mod download`
3. Construa o aplicativo: `go build`
4. Execute o aplicativo: `go run ./cmd/api/main.go`

## Rodando
Existe algumas formas de rodar um projeto em Go, aqui estão duas delas configuradas nesse repositório:

### Rodando com Air
Para executar o serviço localmente, você pode usar [Air](https://github.com/cosmtrek/air) para usar o hot reload. Execute o seguinte comando:
```
air init
```

Altere essa linha no arquivo gerado `.air.toml`:
```
cmd = "go build -o ./tmp/main ."
``` 

Para:
```
  cmd = "go build -o ./tmp/main ./cmd/api/main.go"
```

Rode para iniciar:
```
air
```

### Rodando com Docker
Este projeto inclui um arquivo `Dockerfile` e `docker-compose.yml` para fácil conteinerização e implantação. Aqui estão os comandos Docker e Docker Compose mais comuns que você pode querer usar:

- `docker build -t your-image-name .`: Construa uma imagem Docker para o projeto. Substitua `your-image-name` por um nome para sua imagem.
- `docker run -p 8080:8080 -e PORT=8080 your-image-name`: Execute um contêiner com base na imagem construída. Substitua `your-image-name` pelo nome que você usou ao construir a imagem. Você pode alterar o número da porta, se necessário.

Se você quiser usar o Docker Compose, siga estes comandos:

- `docker compose build`: Construa os serviços definidos no arquivo `docker-compose.yml`.
- `docker compose up`: Execute os serviços definidos no arquivo `docker-compose.yml`.

Para parar e remover contêineres, redes e volumes definidos no arquivo `docker-compose.yml`, execute:

```sh
docker-compose down
```

Para obter mais informações sobre Docker e Docker Compose, consulte a documentação oficial:

- [Docker](https://docs.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## O que foi usado
Este projeto utiliza as seguintes ferramentas:

- [Golang](https://golang.org/) para desenvolvimento back-end
- [Go-Gin](https://github.com/gin-gonic/gin) para gerenciamento de rotas
- [GoORM](https://gorm.io/) para comunicação com banco de dados

## Contribuição

Para contribuir com este projeto, siga estas diretrizes:

1. Clone o repositório
2. Crie uma nova branch: `git checkout -b feature/your-feature-name`
3. Faça suas alterações e confirme-as usando Conventional Commits
4. Faça um push da sua branch: `git push origin feature/your-feature-name`
5. Envie uma pull request

---

## Licença

Este projeto está licenciado sob a licença MIT - consulte o arquivo LICENSE.md para obter detalhes.
