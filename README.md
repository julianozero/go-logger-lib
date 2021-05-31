# Go-Logger-Lib

Composto por dois sub pacotes: **gin_middlewares** e **nlog**.

## nlog

Este pacote nasceu em decorrência dos seguintes requisitos:

- Abstrair a lib de log para facilitar manutenções futuras, a Zerolog.
- Evitar confusões de tipos ao enviar informações básicas para o log centralizado.
- Padronizar a escrita de logs.

#### Como usar?

1 - Crie uma instância única de NLogger:
```
import "github.com/julianozero/go-logger-lib/nlog"

func main() {
    logger := nlog.NewLogger("<service_name>", "<service_version>", "<level>")
}
```

2 - Logando uma mensagem:
```
import "github.com/julianozero/go-logger-lib/nlog"

func main() {
    logger := nlog.NewLogger("export", "1.0.0", "info")
    logger.Info().Send("exportaçao processada com sucesso")
}
```

3 - Você pode passar argumentos para a mensagem com `Sendf(string, ...interface{})`:
```
import "github.com/julianozero/go-logger-lib/nlog"

func main() {
    logger := nlog.NewLogger("export", "1.0.0", "info")
    logger.Info().Sendf("exportação [%s] processada com sucesso", exportID)
}
```

4 - Você pode utilizar os métodos auxiliares para um log mais completo:
```
import "github.com/julianozero/go-logger-lib/nlog"

func main() {
    logger := nlog.NewLogger("export", "1.0.0", "info")
    logger.Info()
        .TraceID(requestID)
        .Org(clientID, userID)
        .Req(url, method)
        .ElapsedTime(elapsedTime)
        .Res(http.StatusNoContent)
        .Sendf("exportação [%s] processada com sucesso", exportID)
}
```

5 - Você pode logar erro:
```
import "github.com/julianozero/go-logger-lib/nlog"

func main() {
    logger := nlog.NewLogger("export", "1.0.0", "info")
    logger.Error()
        .TraceID(requestID)
        .Org(clientID, userID)
        .Req(url, method)
        .ElapsedTime(elapsedTime)
        .Res(http.StatusInternalServerError)
        .Err(err)
        .Send("falha na conexão com o MongoDB")
}
```

6 - Se quiser logar mais informações, use `Dict()`:
```
import "github.com/julianozero/go-logger-lib/nlog"

func main() {
    logger := nlog.NewLogger("export", "1.0.0", "info")
    logger.Error()
        .TraceID(requestID)
        .Org(clientID, userID)
        .Req(url, method)
        .ElapsedTime(elapsedTime)
        .Res(http.StatusInternalServerError)
        .Err(err)
        .Dict("<subobject_name>", map[string]interface{}{"<key>", "<value>"})
        .Send("falha na conexão com o MongoDB")
}
```

## gin_middlewares

