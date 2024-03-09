NYT Best Sellers and Go

Exploratory for now
Trying to brush up on my Go skills

Next steps

- pagination on booklists, by 10 to start
- prompting with promptui or other
- clustering
- remove the build files (BestSellers and playground)

Future...

- connect to a DB
- something AI
- data analytics
- tests ;p

To run

1. download go
2. clone repo
3. Follow the instructions at [NYT API Documents](https://developer.nytimes.com/get-started) to allow access for each API endpoint you want to use. Then save the API key in your `.env` file like `APIKEY=youruniqueapikey`.
4. `cd bestsellers`
5. `go run main.go`

or

1. download go
2. clone repo
3. `go build .`
4. `./bestsellers CMD`
