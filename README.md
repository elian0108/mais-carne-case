# mais-carne-case

## Instruções de Execução

Este projeto foi construído em Go. Para facilitar a execução da aplicação sem precisar instalar o Go no seu computador, foi disponibilizado um arquivo executável para Windows.

### Como Rodar o Executável

1. Baixe o repositório ou faça um clone em sua máquina:
   ```powershell
   git clone https://github.com/elian0108/mais-carne-case.git
   ```
2. Acesse a pasta do projeto pelo terminal ou duplo clique na pasta:
   ```powershell
   cd mais-carne-case
   ```
3. Execute o arquivo gerado:
   - **Pelo Explorador de Arquivos:** Dê um duplo clique no arquivo `mais-carne.exe`.
   - **Pelo Terminal (Powershell/CMD):**
     ```powershell
     .\mais-carne.exe
     ```

Caso deseje compilar novamente o executável a partir do código fonte (necessário ter o Go instalado):
```powershell
go build -o mais-carne.exe main.go
```