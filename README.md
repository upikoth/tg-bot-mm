# Telegram bot

Бот для управления дежурствами

```shell
TELEGRAM_TOKEN="---"
CLOUD_FUNCTION_URL="---"

curl --data "url=$CLOUD_FUNCTION_URL" https://api.telegram.org/bot$TELEGRAM_TOKEN/SetWebhook
```
