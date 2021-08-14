# 🇺🇸 Ports

Projeto com diversos _Ports_ para se integrar com _Cache_, _PubSub_, servidores e clientes **HTTP**, camada de _log_, etc.,
no intuito de permitir que um projeto que faça uso deste possa conter apenas códigos de regra de negócio, particulares a ele.

## Como utilizar

Ao definir a `main` de um projeto é necessário instanciar as dependências do mesmo, e ao fazer isso nos deparamos sempre
com a necessidade de usar as variáveis de ambiente. Para obter essas variáveis pode-se utilizar um `Repository`, contido
no _package_ `config` e obtido por meio da função `config.NewRepository`. Com um `Repository` é possível obter variáveis de
ambiente com o seu método `Get`, e essas variáveis podem ser carregadas por diversas fontes. Cada uma dessas fontes pode ser
carregada a qualquer momento durante a inicialização de acordo com a necessidade de uso delas. Por exemplo:

```go
r := config.NewRepository()
// É válido sempre ao inicializar um repository definir sua versão
r.Add("VERSION", "1.0.0")

// Define as variáveis do sistema como fonte
r.Source(context.Background(), osenvsadapter.New("APP_"))
```

Isso permite remover do seu projeto toda a camada de integração com variáveis do sistema para obter o que precisa para instanciar as suas dependências.

Além de auxiliar durante a inicialização, este projeto possui um pacote com uma especificação que nos permite definir métodos
que serão executados ao finalizar o projeto, o chamado `shutdown`. Seja ao usar numa **API** ou num _consumer_, basta criar
um `shutdown` com o método `stoppage.NewShutdown` e ir adicionando cada _adapter_ que implementar a _interface_ `Stopper`.
Para ativar o pacote e fazê-lo executar os métodos ao fim do projeto basta chamar o método `shutdown.GracefulSignal` que
fornece um `context` que pode ser usado na lógica da sua **API**/_Consumer_. Após instanciar o seu _server_, basta chamar a
função `shutdown.Graceful`, que será executada quando o serviço finalizar.

```go
// Para instanciá-lo é necessário utilizar algo que implemente um Logger, pois o mesmo faz uso de logs
shutdown := stoppage.NewShutdown(logger)

// Adiciona cada uma das implementações de Stopper que se deseja utilizar no seu projeto
shutdown.Add(server)
shutdown.Add(cacher)
shutdown.Add(publisher)

// Obtém uma nova instância de context que deve ser utilizada ao longo do projeto,
// utilizando o mesmo context usado para instanciar suas dependências anteriormente
ctx = shutdown.GracefulSignal(ctx)

router := http.NewRouter()
_ = server.Listen(ctx, router)
```

A _interface_ `Stopper` contém apenas dois métodos, `Stop(ctx context.Context) error` e	`StopError() error`, e permite-lhe
centralizar toda a lógica de finalização do seu _adapter_ neles.

## Pacotes em destaque

### Gerenciamento de erros

Erros de aplicações podem ser gerados por diversas causas e cada erro pode ser tratado de uma maneira específica. Para padronizar
a forma como esses erros são gerados e tratados foi criado o pacote `errorhandler`. Com ele é possível criar erros e definir
como os mesmos serão logados e/ou enviados de volta numa requisição **HTTP**.

O método `Handle` permite que a sua aplicação trate os erros que implementam as especificações desse pacote num único lugar.
Abaixo um exemplo de como utilizar este método, supondo que `Open` retorne variável do tipo `error`, mesmo que ele implemente
ou não as especificações de `errorhandler`. Caso ele não implemente é retornado apenas erro interno na aplicação.
```go
err := i.Open(ctx)
if err != nil {
    h.Handle(ctx, res, err)
    return
}
```