# 🔋🚲 e-bike rental service
Uma aplicação para gerenciar o aluguel de bicicletas elétricas, onde usuários podem registrar-se, localizar bicicletas disponíveis, fazer reservas, iniciar e finalizar aluguéis, além de avaliar as bicicletas após o uso. Administradores podem gerenciar a frota de bicicletas, visualizar relatórios de uso e lidar com manutenção.

## Funcionalidades principais
#### Autenticação e autorização:
- Registro de usuários.
- Login e logout.
- Autorização baseada em roles (usuário comum e administrador).
- Gerenciamento de Usuários:

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

#### Reservas e aluguéis:
- Localização de bicicletas disponíveis.
- Reserva de bicicletas.
- Início e término de aluguéis.
- Cálculo automático de tarifas.
  
#### Avaliações e feedback:
- Avaliação das bicicletas após o uso.
- Comentários sobre o serviço.

#### Relatórios e Estatísticas (para Administradores):
- Relatórios de uso das bicicletas.
- Estatísticas de aluguéis.
- Relatórios de manutenção.

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

#### Avaliações e feedback:
- Usuários podem avaliar bicicletas apenas após um aluguel.
- Avaliações podem incluir uma nota de 1 a 5 estrelas e um comentário opcional.

#### Manutenção de bicicletas:
- Bicicletas com avaliações negativas ou relatórios de problemas são marcadas automaticamente para manutenção.
- Administradores podem alterar o status da bicicleta para "em manutenção" e vice-versa.

## Requisitos técnicos
### Endpoints da API:
#### Autenticação e autorização:
- `POST /v1/auth/register`: Registrar um novo usuário.
- `POST /v1/auth/login`: Autenticar um usuário.
- `POST /v1/auth/logout`: Deslogar um usuário.

#### Gerenciamento de Usuários:
- `GET /v1/users/{id}`: Obter detalhes de um usuário.
- `PUT /v1/users/{id}`: Atualizar detalhes de um usuário.
- `GET /v1/users/{id}/rentals`: Obter histórico de aluguéis de um usuário.

#### Gerenciamento de Bicicletas:
- `POST /v1/admin/bikes`: Adicionar uma nova bicicleta.
- `GET /v1/bikes`: Listar todas as bicicletas.
- `GET /v1/bikes/{id}`: Obter detalhes de uma bicicleta.
- `PUT /v1/bikes/{id}`: Atualizar detalhes de uma bicicleta.
- `DELETE /v1/bikes/{id}`: Remover uma bicicleta.

#### Reservas e Aluguéis:
- `POST /v1/bikes/{id}/reserve`: Reservar uma bicicleta.
- `POST /v1/bikes/{id}/rent`: Iniciar o aluguel de uma bicicleta.
- `POST /v1/bikes/{id}/return`: Finalizar o aluguel de uma bicicleta.

#### Avaliações e Feedback:
- `POST /v1/bikes/{id}/review`: Avaliar uma bicicleta.
- `GET /v1/bikes/{id}/reviews`: Obter avaliações de uma bicicleta.

#### Relatórios e Estatísticas:
- `GET /v1/admin/reports`: Obter relatórios de uso.
- `GET /v1/admin/statistics`: Obter estatísticas de aluguéis.

### Persistência de Dados:
- Banco de dados relacional: `PostgreSQL`.
- Tabelas sugeridas: `users`, `bikes`, `rentals`, `reviews`, `maintenance_logs`.

### Autenticação e segurança:
- JWT (JSON Web Tokens) para autenticação.
- Senhas armazenadas com hashing seguro (e.g., bcrypt).


## Instalação

## Rodando

### Rodando com Air

### Rodando com Docker

## O que foi usado

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
