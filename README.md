# telegram-echo-bot

API-шлюз -> Cпецификация  
```
paths:
  /telegram-echo-bot:
    post:
      x-yc-apigateway-integration:
        type: cloud_functions
        function_id: <cloud-function-id>
      operationId: <cloud-function-name>
```

Настройка WebHook 
```
curl `
    --request POST `
    --url https://api.telegram.org/bot<bot-token>/setWebhook `
    --header '"content-type: application/json"' `
    --data '"{ \"url\": \"<url-for-webhook>" }"' `
    --ssl-no-revoke 
```
(bot-token) = токен для бота от https://t.me/BotFather  
(url-for-webhook) = (API-шлюз -> Служебный домен) + (cloud-function-name)  
