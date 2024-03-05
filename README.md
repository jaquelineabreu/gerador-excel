# Como usar pprof no projeto 🤔
### Documentação: https://pkg.go.dev/net/http/pprof


Importe o pacote

```go
import _ "net/http/pprof"
```

Inicie o servidor HTTP de perfilização:

```go
go func() {
    log.Println(http.ListenAndServe("localhost:6060", nil))
}()
```

Execute a aplicação e um dos comandos a seguir paralelamente:

### Para coletar perfil da CPU:
```go
go tool pprof http://localhost:6060/debug/pprof/profile
```

### Para alloc_space:
```go
go tool pprof -alloc_space http://localhost:6060/debug/pprof/heap
```

### Para inuse_space:
```go
go tool pprof -inuse_space http://localhost:6060/debug/pprof/heap
```

Após executar o comando go tool pprof -alloc_space ou inuse_space aqui estão alguns comandos úteis que você pode usar:

- top: Mostra as funções que alocam a maior quantidade de memória.
- list <function>: Lista o código-fonte para uma função específica.
- web: Abre um visualizador web para visualizar o perfil.
- svg: Exporta o perfil para um arquivo SVG para visualização - gráfica.

Por exemplo, você pode usar o comando top para ver as funções que alocam mais memória:

```go
(pprof) top
```


**alloc_space:** O perfil gerado com a opção alloc_space mostra a quantidade de memória alocada no heap durante a execução do programa, incluindo a memória que já foi liberada. Ele inclui todos os blocos de memória que foram alocados, mesmo que alguns deles já tenham sido liberados. Isso pode ajudá-lo a entender quanto espaço de memória seu programa está alocando em total, incluindo a quantidade que ainda está sendo usada e a quantidade que foi liberada.


**inuse_space:** Por outro lado, o perfil gerado com a opção inuse_space mostra a quantidade de memória atualmente em uso no heap durante a execução do programa. Ele inclui apenas a memória que está atualmente em uso e não inclui a memória que já foi alocada, mas posteriormente liberada. Isso pode ajudá-lo a entender a quantidade de memória que seu programa está realmente usando ativamente em um determinado momento.