# GoTaskManager

<div id="portuguese"></div>

<p align="right">(<a href="#english">English version</a>)</p>

GoTaskManager é um utilitário de linha de comando escrito em Go.
Ele fornece um sistema de gerenciamento de tarefas onde os usuários podem criar, atualizar, excluir e listar tarefas.
Cada tarefa consiste em uma descrição, um status, uma data de vencimento e uma data de conclusão.

As tarefas são armazenadas em um arquivo JSON local chamado 'tasks.json'. O arquivo é carregado quando o programa é iniciado,
e quaisquer alterações feitas durante a execução do programa são salvas de volta no arquivo quando o programa é encerrado.
Isso permite o armazenamento persistente de tarefas entre diferentes execuções do aplicativo.

## Baixando o Projeto

Para baixar o projeto, clique no botão "Code" no GitHub e, em seguida, clique na opção "Download ZIP". Depois de baixar o arquivo ZIP, você poderá descompactá-lo e executar o programa.
O usuário interage com o GoTaskManager através de um sistema de menu:

## Como usar

MENU
- C - CREATE
- U - UPDATE
- D - DELETE
- L - LIST
- O - COMPLETE
- Q - QUIT

Por exemplo:
- Escolher 'C' solicita ao usuário que insira a descrição da tarefa e sua data de vencimento, criando uma nova tarefa.
- Escolher 'U' solicita ao usuário que insira o ID da tarefa a ser atualizada e a nova descrição para essa tarefa.
- Escolher 'D' solicita ao usuário que insira o ID da tarefa a ser excluída.
- Escolher 'L' lista todas as tarefas existentes.
- Escolher 'O' completa uma tarefa.
- Escolher 'Q' salva as tarefas atuais no arquivo JSON e sai do programa.

## Como compilar

Para compilar este programa para Linux, use o seguinte comando:

    go build -o GoTaskManager

Para compilar este programa para Windows, use o seguinte comando:

    GOOS=windows GOARCH=amd64 go build -o GoTaskManager.exe

<div id="english"></div>

# GoTaskManager

<p align="right">(<a href="#portuguese">Versao em Portugues</a>)</p>

GoTaskManager is a command-line utility written in Go.
It provides a task management system where users can create, update, delete, and list tasks.
Each task consists of a description, a status, a due date, and a completion date.

Tasks are stored in a local JSON file described as 'tasks.json'. The file is loaded when the program starts,
and any changes made during the program's execution are saved back to the file when the program is quit.
This allows for persistent storage of tasks between different executions of the application.

## Downloading the Project

To download the project, click on the "Code" button on GitHub and then select the "Download ZIP" option. Once the ZIP file is downloaded, you can unzip it and run the executable.

## How to use

The user interacts with GoTaskManager through a menu system:

MENU
- C - CREATE
- U - UPDATE
- D - DELETE
- L - LIST
- O - COMPLETE
- Q - QUIT

For example:
- Choosing 'C' prompts the user to enter the description of the task and its due date, creating a new task.
- Choosing 'U' prompts the user to enter the ID of the task to update and the new description for that task.
- Choosing 'D' prompts the user to enter the ID of the task to delete.
- Choosing 'L' lists all the existing tasks.
- Choosing 'O' completes a task.
- Choosing 'Q' saves the current tasks to the JSON file and quits the program.

## How to compile

To compile this program for Linux, use the following command:

```bash
go build -o GoTaskManager
```

To compile this program for Windows, use the following command:

```bash
GOOS=windows GOARCH=amd64 go build -o GoTaskManager.exe
```