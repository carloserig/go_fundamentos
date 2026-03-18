# Principais comandos de terminal (sistema operacional)

Referência rápida dos comandos mais usados no terminal (macOS/Linux), úteis ao trabalhar com Go e no dia a dia.

---

## Navegação e diretórios

| Comando | Descrição |
|---------|-----------|
| `pwd` | Mostra o diretório atual (caminho completo). |
| `cd pasta` | Entra na pasta (ex: `cd fundamentos`). |
| `cd ..` | Volta um nível (pasta pai). |
| `cd ~` | Vai para o diretório home do usuário. |
| `cd /` | Vai para a raiz do sistema. |
| `cd -` | Volta ao diretório anterior. |
| `ls` | Lista arquivos e pastas do diretório atual. |
| `ls -l` | Lista em formato longo (permissões, tamanho, data). |
| `ls -la` | Lista tudo, incluindo arquivos ocultos (que começam com `.`). |
| `ls -la pasta` | Lista o conteúdo de uma pasta específica. |

---

## Tela e histórico

| Comando | Descrição |
|---------|-----------|
| `clear` | Limpa a tela do terminal (ou Ctrl+L). |
| `history` | Mostra o histórico de comandos. |
| `!!` | Repete o último comando. |
| `!comando` | Repete o último comando que começava com "comando". |

---

## Arquivos e pastas

| Comando | Descrição |
|---------|-----------|
| `mkdir nome` | Cria uma pasta. |
| `mkdir -p a/b/c` | Cria a árvore de pastas (cria pais se não existirem). |
| `touch arquivo.txt` | Cria um arquivo vazio ou atualiza a data de modificação. |
| `cp origem destino` | Copia arquivo ou pasta. |
| `cp -r pasta destino` | Copia pasta recursivamente (com todo o conteúdo). |
| `mv origem destino` | Move ou renomeia arquivo/pasta. |
| `rm arquivo` | Remove arquivo (não vai para lixeira). |
| `rm -r pasta` | Remove pasta e todo o conteúdo. |
| `rm -rf pasta` | Remove pasta forçando, sem pedir confirmação (use com cuidado). |
| `cat arquivo` | Mostra o conteúdo do arquivo no terminal. |
| `head arquivo` | Mostra as primeiras linhas (padrão: 10). |
| `tail arquivo` | Mostra as últimas linhas (padrão: 10). |
| `tail -f arquivo` | Acompanha o arquivo em tempo real (útil para logs). |

---

## Busca e conteúdo

| Comando | Descrição |
|---------|-----------|
| `find . -name "*.go"` | Busca arquivos pelo nome (ex: todos os .go a partir da pasta atual). |
| `grep "texto" arquivo` | Busca "texto" dentro do arquivo. |
| `grep -r "texto" .` | Busca "texto" em todos os arquivos do diretório atual (recursivo). |
| `grep -i "texto" arquivo` | Busca ignorando maiúsculas/minúsculas. |

---

## Permissões

| Comando | Descrição |
|---------|-----------|
| `chmod +x arquivo` | Torna o arquivo executável. |
| `chmod 755 arquivo` | Define permissões (leitura/execução para todos, escrita só dono). |
| `chown usuario:grupo arquivo` | Muda dono e grupo do arquivo (geralmente precisa de `sudo`). |

---

## Processos e sistema

| Comando | Descrição |
|---------|-----------|
| `ps` | Lista processos do terminal atual. |
| `ps aux` | Lista processos de todos os usuários. |
| `kill PID` | Encerra processo pelo número (PID). |
| `kill -9 PID` | Força o encerramento do processo. |
| `top` | Mostra processos em tempo real (sair: `q`). |
| `which comando` | Mostra o caminho do executável (ex: `which go`). |
| `echo $VARIAVEL` | Mostra o valor de uma variável de ambiente. |

---

## Atalhos úteis no terminal

| Atalho | Descrição |
|--------|-----------|
| Ctrl+C | Interrompe o comando em execução. |
| Ctrl+L | Limpa a tela (como `clear`). |
| Ctrl+D | Encerra a entrada (sair do terminal ou de um programa que lê input). |
| Ctrl+Z | Pausa o processo (depois use `fg` para continuar ou `bg` para rodar em segundo plano). |
| Tab | Autocompleta nomes de arquivos e comandos. |
| Setas ↑ e ↓ | Navega no histórico de comandos. |

---

## Exemplos no contexto Go

```bash
# Ir até a pasta do projeto e listar
cd ~/Documents/ISEPE/2026/Backend/go/fundamentos
ls -la

# Limpar tela, rodar o programa e ver o código
clear
go run exercicios/exercicio01/exercicio01.go
cat exercicios/exercicio01/exercicio01.go

# Buscar onde está o executável do Go
which go
```
