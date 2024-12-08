# Bubble Sort API com Gin

Este projeto é uma API simples construída com o framework Gin em Go, que implementa o algoritmo de ordenação Bubble Sort. A API permite que você envie um conjunto de números inteiros através de uma query string e recebe o resultado da ordenação juntamente com os passos intermediários do algoritmo.

## Funcionalidades

- **Bubble Sort**: Ordena um array de inteiros.
- **Passos Intermediários**: A API retorna o array após cada iteração externa do algoritmo, permitindo acompanhar o progresso da ordenação.
- **Validação de Entrada**: A API valida se os números fornecidos são inteiros válidos.

## Endpoints

### `GET /bubblesort`

Este endpoint recebe um conjunto de números inteiros separados por vírgulas na query string e retorna o array ordenado, juntamente com os passos intermediários do algoritmo.


ex: /bubblesort?numbers=1,70,20,54,269


#### Parâmetros

- `numbers` (obrigatório): Uma string contendo números inteiros separados por vírgulas. Exemplo: `numbers=5,3,8,2,1`

#### Resposta

- **Status 200**: Se a ordenação for realizada com sucesso, a resposta será um objeto JSON contendo:
  - `array_ordenado`: O array ordenado.
  - `passos`: Uma lista dos estados do array após cada iteração externa.

- **Status 400**: Se os parâmetros forem inválidos ou ausentes, a resposta será um erro com a descrição do problema.

## Como Rodar

1. **Instale o Go**: Se você ainda não tem o Go instalado, baixe e instale a partir do [site oficial do Go](https://golang.org/dl/).

2. **Clone o repositório**:

    ```bash
    git clone https://github.com/thalesz/apibubblesortgo
    cd bubble-sort-api
    ```

3. **Inicialize o módulo Go**:

    ```bash
    go mod init api-go
    ```

4. **Instale as dependências**:

    ```bash
    go get github.com/gin-gonic/gin
    ```

5. **Execute o servidor**:

    ```bash
    go run main.go
    ```

    O servidor estará disponível em `http://localhost:8080`.
