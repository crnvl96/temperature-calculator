## Temperature calculator

https://temperature-calculator-rypv67tq2a-uc.a.run.app/calculate?zipcode=<CEP>

## Exemplo

https://temperature-calculator-rypv67tq2a-uc.a.run.app/calculate?zipcode=71625045

```json
{"temp_C":22.3,"temp_F":72.1,"temp_K":295.3}
```

## Execução em dev
[weather api key](https://www.weatherapi.com/)

#### Requisitor

```bash
git clone https://github.com/crnvl96/temperature-calculator.git && cd temperature-calculator
```

```bash
go mod download
```

```bash
cp .env.example .env
```

```env
WEATHER_API_KEY=<your_weather_api_key>
PORT=8080
ENVIRONMENT=development
```

```bash
docker compose up -d
```

```bash
curl http://localhost:8080/calculate?zipcode=<CEP>
```

```bash
go test ./...
```
