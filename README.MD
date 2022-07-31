# Ports and Adapters with [![Go](https://img.shields.io/badge/--00ADD8?logo=go&logoColor=ffffff)](https://golang.org/)

Exemplo de projeto [golang](https://go.dev) utilizando o conceito de [arquitetura hexagonal (ports and adapters)](https://alistair.cockburn.us/hexagonal-architecture/)

O app se trata de um gerenciador de produtos, onde é possivel inserir, consultar, habilitar e desabilitar. 

O mesmo é bem simples e não deve ser levado como exemplo de CRUD, mas sim de como é possivel reaproveitar o codigo e isolar o dominio da aplicacao das complexidades tecnicas impostas pelos frameworks e bibliotecas externas. 

Nesse exemplo o programa pode rodar utilizando tanto a linha de comando quanto via requisicoes http, e a unica diferença entre os codigos é a implementação do adapter de cada uma da formas de input, sendo desnecessaria qualquer alteração nas regras das entidades da aplicação. 


# Rodando via CLI 
``` 
  # adicionar um produto
  go run ./main.go cli -a=create -n=teste -p=10

  # ativar um produto
  go run ./main.go cli -a=enable -i=id_produto
  
  # desabilitar um produto
  go run ./main.go cli -a=enable -i=id_produto

  # buscar um produto pelo id
  go run ./main.go cli -a=enable -i=id_produto
``` 

# Utilizando via web 
``` 
  #Subir o servidor http
  go run./main.go http


  # adicionar um produto
  POST http://localhost:8080/products 
  body: {
    "name": string,
    "price": number
  }

  # ativar um produto
  GET http://localhost:8080/products/{id_produto}/enable
  
  # desabilitar um produto
  GET http://localhost:8080/products/{id_produto}/disable

  # buscar um produto pelo id
  GET http://localhost:8080/products/{id_produto}
``` 
