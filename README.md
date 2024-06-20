# store-restful
Restful project. Port by default 8090. 

> **_NOTE:_**  .env contains database's password

# commands to test!^^
Sign Up
``` shell
$curl -d '{"name":"rmarsu","username":"rmarsu","password":"1234"}' -H "Content-Type: application/json" -X POST http://localhost:8090/auth/sign-up
```

Sign In
``` shell
$curl -d '{"username":"rmarsu","password":"1234"}' -H "Content-Type: application/json" -X POST http://localhost:8090/auth/sign-in
```
Buy 
``` shell
$curl -XGET -H 'Authorization: Bearer (your auth token)' 'http://localhost:8090/api/products/buy/1'
```
Upload product
``` shell
$curl -XPOST -H 'Authorization: Bearer (your auth token)' -H "Content-type: application/json" -d '{"id": "1" ,"image":"some image", "name":"plushie", "price":12.5}' 'http://localhost:8090/api/products/createProduct'
```
Show all products
``` shell
$curl http://localhost:8090
```
