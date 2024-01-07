# Roteiro de Fundamentos Básicos em Go (Golang):

1. Instalação do Go:

Baixe e instale o Go a partir do site oficial (https://golang.org/dl/).
Configure corretamente as variáveis de ambiente, como GOPATH e GOROOT.

2. Hello, World!:

Crie um programa "Hello, World!" em Go.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}

```
Depois de instalar e configurar, criar o repositorio na minha area de trabalho, tentei criar meu primeiro "Hello, World!". Usando o script acima, a propria IDE me informa a seguinte informação:

![Imagem IDE com informação de erro](/_img/01.png)

```
gopls was not able to find modules in your workspace.
When outside of GOPATH, gopls needs to know which modules you are working on.
You can fix this by opening your workspace to a folder inside a Go module, or
by using a go.work file to specify multiple modules.
See the documentation for more information on setting up your workspace:
https://github.com/golang/tools/blob/master/gopls/doc/workspace.md.go list
```
>gopls não conseguiu encontrar módulos em seu espaço de trabalho.
>Quando estiver fora do GOPATH, o gopls precisa saber em quais módulos você está trabalhando.
>Você pode corrigir isso abrindo seu espaço de trabalho em uma pasta dentro de um módulo Go ou
>usando um arquivo go.work para especificar vários módulos.
>Consulte a documentação para obter mais informações sobre como configurar seu espaço de trabalho:
>https://github.com/golang/tools/blob/master/gopls/doc/workspace.md.go lista

![Imagem IDE com informação de erro](/_img/02.png)

## Chatgpt

### _gopls

Go Language Server, é uma ferramenta usada para fornecer recursos como autocompletar, análise de código e navegação em projetos Go.

O erro que você compartilhou indica que o gopls não conseguiu encontrar módulos em seu espaço de trabalho Go. 

### _go.mod

O arquivo go.mod é um componente essencial no ecossistema de desenvolvimento Go (Golang). Ele é utilizado para definir e gerenciar os módulos Go em um projeto.

A introdução de módulos Go foi uma mudança significativa introduzida no Go 1.11 para melhorar o gerenciamento de dependências e versionamento de pacotes. Antes dos módulos, o Go usava o conceito de GOPATH para organizar os pacotes, mas isso trouxe algumas limitações e desafios.

Com o arquivo go.mod, você pode especificar as dependências do seu projeto, incluindo a versão de cada pacote. O arquivo também contém informações sobre o próprio módulo, como o nome do módulo e a versão do Go compatível.

Aqui está um exemplo básico de um arquivo go.mod:

```go
module example.com/meuprojeto

go 1.16

require (
    github.com/pacote/exemplo v1.2.3
    outra/biblioteca v0.4.1
)

```

- `module example.com/meuprojeto`: Especifica o caminho do módulo.
- `go 1.16`: Indica a versão mínima do Go necessária para compilar o módulo.
- `require`: Lista as dependências do módulo, incluindo os caminhos dos pacotes e as versões específicas.

Ao usar o comando go get ou ao importar novas dependências em seu código, o arquivo go.mod é atualizado automaticamente para refletir essas alterações.

Os módulos Go fornecem uma maneira mais robusta e flexível de gerenciar dependências em comparação com o GOPATH, tornando o desenvolvimento em Go mais escalável e controlado.

Mesmo que tente executar o programa `go run .` ele não executa, isso é caracteristas do Go. Ele espera que tenha o modulo na raiz do projeto.

```bash
$go run .
go: cannot find main module, but found .git/config in /home/neres/develop/_github/_go
        to create a module there, run:
        cd .. && go mod init
```

Para criar o arguivo go.mod no terminal `go mod init localhost/neres/hello`, esse exemplo segue o pardrão recomendado para criar o arquivo, para quem usa Java e Springboot, esse go.mod se comporta como **Maven** que é um gerenciador de dependencias do Java. `localhost/` represento o local do repositorio, `neres/` o desenvolvedor e `hello` o projeto. Isso pelo que entendi é como se fosse o ponteiro do Go para o resto do programa. 

Executando o comando:

```bash
$go mod init localhost/neres/hello
go: creating new go.mod: module localhost/neres/hello
go: to add module requirements and sums:
    go mod tidy
```
![Imagem IDE do go.mod criado](/_img/03.png)

O comando `go mod init` é usado para inicializar um novo módulo Go em um projeto. Quando você executa este comando, ele cria um arquivo `go.mod` na raiz do seu projeto. Parece que você já executou o comando e obteve a mensagem indicando que o módulo foi criado com sucesso.

Isso criará um módulo com o caminho `localhost/neres/hello`. O arquivo `go.mod` gerado pelo comando conterá essa informação.

A mensagem "go: to add module requirements and sums: go mod tidy" está sugerindo que você execute `go mod tidy` para garantir que as dependências e somas do módulo estejam em conformidade com o que está no seu código. O comando `go mod tidy` remove qualquer dependência que não é necessária e ajusta as versões das dependências conforme necessário, com base no código atual.

Se você precisar adicionar ou remover dependências manualmente, você pode fazer isso diretamente no arquivo `go.mod` e, em seguida, executar `go mod tidy` para garantir a consistência.

### Agora executando o programa no terminal.

```bash
$go run .

Oi, Mundo!
```
![Imagem IDE do go.mod criado](/_img/04.png)

## PACOTE `fmt`

Esse pacote será bastante utilizado nesse primeiro contato.

**I. Println**: 
```go
package main

import "fmt"

func main() {
    nome := "Alice"
    idade := 30

    fmt.Println("Nome:", nome, "Idade:", idade)
}

```

**II. Printf**: 
```go
package main

import "fmt"

func main() {
    nome := "Bob"
    idade := 25

    fmt.Printf("Nome: %s, Idade: %d\n", nome, idade)
}

```

**III. Sprintf**: 
```go
package main

import "fmt"

func main() {
    nome := "Charlie"
    idade := 35

    resultadoFormatado := fmt.Sprintf("Nome: %s, Idade: %d", nome, idade)
    fmt.Println(resultadoFormatado)
}

```

**IV. Scanf**: 
```go
package main

import "fmt"

func main() {
    var nome string
    var idade int

    fmt.Print("Digite o nome: ")
    fmt.Scanln(&nome)

    fmt.Print("Digite a idade: ")
    fmt.Scanln(&idade)

    fmt.Printf("Nome: %s, Idade: %d\n", nome, idade)
}

```
## Variaveis e tipos de dados

![Imagem IDE do go.mod criado](/_img/tiposDB.png)

**1. Declaração e Atribuição de Valores as Variaveis**

```go
package main

import "fmt"

func main() {
    // Declaração e atribuição de variáveis
    var idade int
    idade = 25

    nome := "João" // Inferência de tipo
    altura := 1.75 // Inferência de tipo

    // Impressão dos valores
    fmt.Println("Nome:", nome)
    fmt.Println("Idade:", idade)
    fmt.Println("Altura:", altura)
}

```
**2. Tipos de Dados Básicos:**

```go
package main

import "fmt"

func main() {
    // Tipos de dados básicos
    var numeroInteiro int
    numeroInteiro = 42

    var numeroDecimal float64
    numeroDecimal = 3.14

    texto := "Golang"
    
    verdadeiro := true

    // Impressão dos valores
    fmt.Println("Número Inteiro:", numeroInteiro)
    fmt.Println("Número Decimal:", numeroDecimal)
    fmt.Println("Texto:", texto)
    fmt.Println("Booleano:", verdadeiro)
}

```
Esses exemplos demonstram a declaração e atribuição de variáveis em Go, bem como os tipos de dados básicos como int, float64, string e boolean. A inferência de tipo (:=) é usada quando o tipo pode ser deduzido automaticamente pelo compilador.

