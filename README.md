<div align="center">
  <br>
  <h1>Currency conversion API</h1>
</div>






## Summary

An API that provides a small but well-thought service converting Euro to US Dollar and vice-versa. That API should only be accessible if you are in the possession of an API-KEY.



##  Running

Just type in the terminal

```bash
docker-compose build
```

```bash
docker-compose up
```

Then you can test the API from [postman collection](postman%20collection) by mporting the collection in postman application, or you can test it from browser:

- To get an API key the route will be:

  ```
  http://localhost:8080/apikey
  ```

- To get the currency rate conversion to conversion (price of the EUR in USD)

  ```
  http://localhost:8080/currency_rate?api_key=<VALID API KEY>
  ```

- To convert from USD to EUR:

  ```
  http://localhost:8080/convert_usd?api_key=<VALID API KEY>&usd=<FLOAT AMOUNT OF MONEY>
  ```

- To convert from EUR to USD:

  ```
  http://localhost:8080/convert_eur?api_key=<VALID API KEY>&eur=<FLOAT AMOUNT OF MONEY>
  ```



## Tasks which done

- End point to generate an API key.
- Endpoint to get the currency rate conversion to conversion (price of the EUR in USD)
- Enpoint to convert from USD to EUR
- Endpoint to convert from EUR to USD



## Extra pieces of information

I isolated packages to avoid import cycles
