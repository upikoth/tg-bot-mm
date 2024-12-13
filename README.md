# Telegram bot

Бот для управления дежурствами

```shell
TELEGRAM_TOKEN="---"
CLOUD_FUNCTION_URL="---"

curl --data "url=$CLOUD_FUNCTION_URL" https://api.telegram.org/bot$TELEGRAM_TOKEN/SetWebhook
```

Нужно закомментировать строки установки команд после того как они будут установлены
```
err = srv.setCommands()

if err != nil {
    log.Println("Error: ", err)
    return
}
```

Уведомления нужно включать только после того как setCommands установлены, они не работают вместе
