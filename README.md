# Projeto do Zero ao Avançado com Go Lang

Este repositório contém um projeto desenvolvido em Go, abrangendo desde configuração de ambiente até tópicos avançados.

## Tecnologias

- [Ubuntu 22.04.3 LTS](https://ubuntu.com/download/desktop): Sistema Operacional.
- [Git](https://git-scm.com/): Git é um sistema de controle de versão distribuído gratuito e de código aberto.
- [Github](https://github.com/): GitHub é uma plataforma de hospedagem de código-fonte e arquivos com controle de versão usando o Git.
- [Go](https://golang.org/): Linguagem de programação principal do projeto.
- [Biblioteca XYZ](#): Utilizada para [descreva a funcionalidade].
- [Ferramenta ABC](#): Usada para [descreva a finalidade].


## Configuração do Ambiente em Linux

As configurações asseguir apenas o processo de instalação do Ubuntu e criação do repositorio no Github nforam desconsiderados.

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

**Obs: Com minha conta criada no Github criei meu repositorio e em seguida fiz o processo acima referente ao Git e conexão com o Gihub, "NA MINHA MAQUINA FUNCIONA" se estiver tendo dificuldade recomendo pesquisar nos repositórios, foruns ou Youtube.**