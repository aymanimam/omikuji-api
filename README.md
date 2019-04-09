# omikuji-api

`omikuji-api` is the last task in [Gopher dojo](https://docs.google.com/presentation/d/1Ri0sN-jnQTDEhV0Oiet4T_voSg7Sr9qnRDSOIWQsdW8/edit#slide=id.g395fe3d780_122_756)
It provide the following endpoint:
```$xslt
GET /omikuji
```
The response will be a random Omikuji from a predefined Omikujis list. The format of the response will be.
```$xslt
Successful response:
200 OK
{
    "omikuji":"中吉"
}

Error response:
500 Internal server error
{
    "message":"..."
    "code":"100"
}
```
It returns Daikichi Omikuji only from January 1st to January 3rd. 