## Telegram Echo Bot
Link to actual bot: [@JustAnotherEchoBot](https://t.me/yc_go_function_echo_bot)  
Main guide used:   [How to create a Telegram bot](https://cloud.yandex.com/en/docs/functions/tutorials/telegram-bot-serverless)

### Components needed:
- Telegram:
  - Bot token from [@BotFather](https://t.me/BotFather)
  - Setup WebHook
- Yandex Cloud:
  - API-Gateway
  - Cloud Function
  - Managed YBD (or any other)

### Notes:
#### Yandex Cloud roles:
- `serverless.functions.admin` for GitHub Actions to deploy  cloud functions 
- `serverless.functions.invoker`, `ydb.editor` for executing function, read/write to YDB

#### Setup Telegram WebHook:
```
curl `
    --request POST `
    --url https://api.telegram.org/bot<bot-token>/setWebhook `
    --header '"content-type: application/json"' `
    --data '"{ \"url\": \"<url-for-webhook>" }"' `
    --ssl-no-revoke 
```
`<bot-token>` = bot token [@BotFather](https://t.me/BotFather)  
`<url-for-webhook>` = (API-Gateway -> Default domain ) + (API-Gateway -> Default domain -> Specification -> `paths`)  
