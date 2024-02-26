### How to run

Execute in command prompt

```
go run main.go
```

or

```
app.exe
```

### How to test

Postman for API call

```http
  POST /email/classify
```

| Parameter   | Type     | Description               |
| :---------- | :------- | :------------------------ |
| `keywords`  | `array`  | Array of words or phrases |
| `emailText` | `string` | Your email text           |

#### Example

```
{
  "keywords": [" EmaIL ", "CONfideNtial   ", "top secret  "],
  "emailText": "This email contains confidential and classified information. Therefore, this email is top secret. Again. This email is TOP SECRET."
}
```
