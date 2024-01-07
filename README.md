# Projeto do Zero ao Avançado com Go Lang

Este repositório contém um projeto desenvolvido em Go, abrangendo desde a configuração do ambiente até tópicos avançados.

Minha motivação surgiu após responder a um questionário de uma plataforma. Esse questionário tinha como objetivo avaliar meu conhecimento em algumas tecnologias, como Java, Node, Back-end e Front-end.

Como comecei meus estudos em Go, decidi adaptar a lista que a plataforma me apresentou para este projeto.

Espero que, até o final, eu tenha adquirido de forma **Hardcode** bastante habilidade com essa tecnologia, que até então meu conhecimento não passa de um simples "Hello, World!" :).

## 🏆 Conquistas da Jornada

> 🧑‍💻 EM DESENVOLVIMENTO

Listei com ajuda do ChatGpt algumas coisas que não seram seguidas a riscas, pois não faço ideia de como trabalhar com o Go,
foi mais coisas que é possivel fazer com Java e Javascript.


✅ Montar e configurar ambientes em Golang: 

- Golang: Configurar o ambiente de desenvolvimento Go usando GOPATH ou módulos Go.

🏅 Análise da API REST com monitoramento de saúde:

- Golang: Use bibliotecas como "github.com/gin-gonic/gin" para criar APIs REST e explore ferramentas de monitoramento e análise específicas do ecossistema Go.

🏅 Segurança de sessões via token:

- Golang: Use bibliotecas como "github.com/dgrijalva/jwt-go" para lidar com tokens JWT.

🏅 Implementar processo de memória cache:

- Golang: Use bibliotecas como "github.com/patrickmn/go-cache" para implementar cache em memória.

🏅 Trabalhar com filas e eventos:

- Golang: Use bibliotecas como "github.com/nsqio/go-nsq" para filas e "github.com/nats-io/nats.go" para eventos.

🏅 Mapeamento de entidades (se for possivel):

- Golang: Use tags de estrutura para mapear entidades.

🏅 Mapeamento de entidades (se possivel):

- Golang: Use bibliotecas como "github.com/jinzhu/copier" ou atribuição direta de campos.

🏅 Construir relatórios (Se for possivel):

- Golang: Similar ao item anterior.

🏅 Visualização orientada a eventos:

- Golang: Use bibliotecas como "github.com/gorilla/websocket" para comunicação em tempo real.

🏅 Desenvolvimento web:

- Golang: Use frameworks web como "github.com/gin-gonic/gin" ou "github.com/labstack/echo".

🏅 Acesso a dados

- Golang: Use bibliotecas como "github.com/jinzhu/gorm" para acessar dados de bancos.

- Golang: Use "database/sql" com drivers específicos para mapear consultas SQL.


## Tecnologias

- [Ubuntu 22.04.3 LTS](https://ubuntu.com/download/desktop): Sistema Operacional.
- [Git](https://git-scm.com/): Git é um sistema de controle de versão distribuído gratuito e de código aberto.
- [Github](https://github.com/): GitHub é uma plataforma de hospedagem de código-fonte e arquivos com controle de versão usando o Git.
- [Go](https://golang.org/): Linguagem de programação principal do projeto.
- [VScode](#): O Visual Studio Code é um editor de código-fonte desenvolvido pela Microsoft para Windows, Linux e macOS.

### Pacotes
- [fmt](https://pkg.go.dev/fmt): O pacote `fmt` em Go é utilizado para formatação de entrada e saída. Ele fornece funcionalidades para imprimir e formatar texto no console, além de manipular strings de maneira eficiente.

**Funcionalidades fmt**:

`Println` é usada para imprimir valores no console, adicionando uma nova linha ao final.

`Printf` permite formatar e imprimir texto com base em um formato específico, semelhante à função `printf` em outras linguagens.

`Sprintf` é usado para formatar strings e retornar o resultado formatado como uma nova string, sem imprimir no console.

`Scanf` é usado para ler valores do console com base em um formato específico.

O pacote fmt é muito útil para operações de entrada e saída formatadas, além de ser amplamente utilizado para depuração e exibição de informações durante o desenvolvimento.

- [Ferramenta ABC](#): Em construçao.




## Configuração do Ambiente em Linux 🛫

As configurações asseguir apenas o processo de instalação do Ubuntu e criação do repositorio no Github foram desconsiderados.

Mas se preciso for, tem vários conteúdos na internet de como fazer.

### _GIT

> Passo 1: Instalação do Git

Se o Git ainda não estiver instalado, você pode instalá-lo usando o seguinte comando no terminal:

```bash
sudo apt update
sudo apt install git

```
> Passo 2: Configuração Inicial

Depois de instalar o Git, configure seu nome de usuário e endereço de e-mail. Isso é usado para identificar suas contribuições nos repositórios.

```bash
git config --global user.name "Seu Nome"
git config --global user.email "seu@email.com"

```

> Passo 3: Configuração do GitHub

**I. Geração de Chaves SSH (se ainda não tiver uma)**

As chaves SSH permitem a comunicação segura entre seu computador e o GitHub. Se você já tiver uma, pode pular esta etapa.

```bash
ssh-keygen -t rsa -b 4096 -C "seu@email.com"

```
Pressione Enter para aceitar o local padrão para salvar as chaves (geralmente ~/.ssh/id_rsa), e crie uma senha se desejar.

**II. Adicionando a Chave SSH ao Agente SSH**

```bash
eval "$(ssh-agent -s)"
ssh-add ~/.ssh/id_rsa

```

**III. Adicionando a Chave SSH ao GitHub**

```bash
cat ~/.ssh/id_rsa.pub
```
*Copie a saída do comando e adicione a chave SSH ao GitHub. No **GitHub**, vá para Settings -> SSH and GPG keys -> New SSH key, cole a chave e salve.*

### Passo 4: Verificação da Conexão

Para verificar se a configuração está correta, execute o seguinte comando:

```bash
ssh -T git@github.com
```

Você deve receber uma mensagem de confirmação.

```bash
The authenticity of host 'github.com (##.###.##.###)' can't be established.
ED25519 key fingerprint is SHA256:****************************************.
This key is not known by any other names
Are you sure you want to continue connecting (yes/no/[fingerprint])? yes
Warning: Permanently added 'github.com' (#######) to the list of known hosts.
Hi neresfabio! You've successfully authenticated, but GitHub does not provide shell access.

```

**Obs: Com minha conta criada no Github criei meu repositorio e em seguida fiz o processo acima referente ao Git e conexão com o Github, **"NA MINHA MAQUINA FUNCIONA"** se estiver tendo dificuldade recomendo pesquisar nos repositórios, foruns ou Youtube.**

### _Go

Para configurar o ambiente Go no Linux, siga os seguintes passos:

**I. Instale o Go usando o gerenciador de pacotes da sua distribuição. Exemplo para o Ubuntu:**

```bash
   sudo apt-get update
   sudo apt-get install golang
```

**II. Configure as variáveis de ambiente no seu arquivo de perfil (ex: ~/.bashrc):**

```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

```
**III. Reinicie o terminal ou execute source ~/.bashrc para aplicar as alterações.**
Execute o seguinte comando para viasualizar a versão do Go: ```go version```

