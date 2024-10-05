## О проекте

Этот проект представляет собой REST сервер для мессенджера

## Для запуска приложения:

### Конфигурация

Прежде чем запустить проект, убедитесь, что в вашей директории проекта создан файл `config.yaml` с необходимыми переменными окружения. Далее представлен шаблон содержимого файла `config.yaml`:

```yaml
env: "local"
graceful_shutdown_timeout: 10

server:
  address: 0.0.0.0:8080
  read_timeout: 10
  write_timeout: 10
  idle_timeout: 10

logger:
  path: "logs/logger.log"
```

### Запуск сервера

   ``` bash
   make serve
   ```

## Разработчики
Разработкой сервера занимались:
- [Mutalibov Alaudin](https://github.com/KrizzMU)
- [Paradeev Nikolay](https://github.com/Cr4z1k)
- [Fadeev Vsevolod](https://github.com/fFH255)
