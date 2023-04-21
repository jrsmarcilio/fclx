# 1. Docker e Containers e Microserviçoes do ChatGTP
<hr>


## Agenda

  - Entender o projeto prático
  - Tecnologias que serão utilizadas
  - Docker
  - Início do Microserviço


## Projeto prático

  - Desenvolver duas UI para utilizar o ChatGPT
  - Integração com o WhatsApp


## Arquitetura do projeto

```txt
                            <- gRPC <- [ Next.js (backend) ] <- [ Next.js (frontend) ]
[ OpenAI ] <- [  Chat MS ]
                            <- HTTP <- [ Twilio ] <- [ WhatsApp ] 
```                    


## OpenAI API e ChatGPT
  
  ### Tokens do ChatGPT

  Para o ChatGPT, "tokens" são sequências de caracteres ou símbolos que representam unidades semânticas, como palavras, pontuação e caracteres especiais, que são usados ​​como entrada para o modelo de linguagem GPT (Generative Pre-trained Transformer).

  O ChatGPT usa uma técnica conhecida como tokenização para quebrar o texto em "tokens", que são então codificados e usados ​​como entrada para o modelo. A tokenização é um processo importante para o modelo GPT, pois ajuda a transformar o texto em uma representação numérica que pode ser compreendida e processada pelo modelo.

  Os "tokens" são usados ​​para treinar o modelo e também para gerar respostas para as perguntas dos usuários durante as interações do chatbot. O modelo GPT usa uma grande variedade de "tokens" para entender a linguagem natural e gerar respostas coerentes e relevantes.


  ### Modelos do ChatGPT

  #### GPT-3.5

  Os modelos GPT-3.5 podem entender e gerar linguagem natural ou código. O modelo mais capaz e econômico na família GPT-3.5 é o gpt-3.5-turbo, que foi otimizado para chat, mas também funciona bem para tarefas tradicionais de preenchimento.

  | Modelo mais recente  | Descrição  | Máximo de Tokens | Dados de treinamento |
  |---|---|---|---|
  | gpt-3.5-turbo | Modelo GPT-3.5 mais capaz e otimizado para chat, custando 1/10 do valor do text-davinci-003. Será atualizado com última iteração de modelo. | 4.096 | Até setembro de 2021 |
  | gpt-3.5-turbo-0301 | Instantâneo do gpt-3.5-turbo em 1º de março de 2023. Ao contrário do gpt-3.5-turbo, este modelo não receberá atualizações e será descontinuado 3 meses após o lançamento de uma nova versão. | 4.096 | Até setembro de 2021 |
  | text-davinci-003 | Pode realizar qualquer tarefa de linguagem com melhor qualidade, saída mais longa e seguimento de instruções consistente do que os modelos curie, babbage ou ada. Também suporta inserção de conclusões dentro do texto. | 4.097 | Até junho de 2021 |
  | text-davinci-002 | Possui capacidades semelhantes às do text-davinci-003, mas é treinado com ajuste fino supervisionado em vez de aprendizado por reforço. | 4.097 | Até junho de 2021 |
  | code-davinci-002 | Otimizado para tarefas de completar códigos. | 8.001 | Até junho de 2021 |

  Recomendamos o uso do ```gpt-3.5-turbo``` sobre os outros modelos GPT-3.5 devido ao seu custo mais baixo.

  [OpenAI - ChatGPT-3.5](https://platform.openai.com/docs/models/gpt-3-5)

  <hr><br>


  #### GPT-4 - ```Limited beta```

  O GPT-4 é um grande modelo multimodal (aceitando entradas de texto e emitindo saídas de texto, aceitará entradas de imagem no futuro) que pode resolver problemas difíceis com maior precisão do que qualquer um de modelos anteriores, graças ao seu amplo conhecimento geral e capacidades avançadas de raciocínio.Assim como o ```gpt-3.5-turbo```, o GPT-4 é otimizado para bate-papo, mas funciona bem para tarefas de conclusão tradicionais usando a [```API de conclusão de Chat```](https://platform.openai.com/docs/api-reference/chat). Aprenda a usar o GPT-4 na [```Documentação do chat```](https://platform.openai.com/docs/guides/chat).

  | Modelo mais recente  | Descrição  | Máximo de Tokens | Dados de treinamento |
  |---|---|---|---|
  | gpt-4  | Mais capaz do que qualquer modelo GPT-3.5, capaz de realizar tarefas mais complexas e otimizado para chat. Será atualizado com a última iteração do modelo. | 8.192 tokens | Até setembro de 2021 |
  | gpt-4-0314 | Instantâneo do gpt-4 de 14 de março de 2023. Ao contrário do gpt-4, este modelo não receberá atualizações e será depreciado 3 meses após o lançamento de uma nova versão. | 8.192 tokens | Até setembro de 2021 |
  | gpt-4-32k | Mesmas capacidades do modelo base gpt-4, mas com 4x o comprimento do contexto. Será atualizado com a última iteração do modelo. | 32.768 tokens | Até setembro de 2021 |
  | gpt-4-32k-0314 | Instantâneo do gpt-4-32k de 14 de março de 2023. Ao contrário do gpt-4-32k, este modelo não receberá atualizações e será depreciado 3 meses após o lançamento de uma nova versão. | 32.768 tokens | Até setembro de 2021 |

  Para muitas tarefas básicas, a diferença entre os modelos GPT-4 e GPT-3.5 não é significativa. No entanto, em situações de raciocínio mais complexas, o GPT-4 é muito mais capaz do que qualquer um dos modelos anteriores.

  [OpenAI - ChatGPT-4](https://platform.openai.com/docs/models/gpt-4)

  <hr><br>


## ChatGPT / API 0 Funcionamento básico

  - Mensagem inicial do sistema ```(system)```
  - Pergunta do usuário ```(user)```
  - Resposta do chatgpt ```(assistant)```
  - Pergunta do usuário ```(user)```
  - Resposta do chatgpt ```(assitant)```

  - Mensagens vão se acumulando para armazenar o contexto
  - Quando não couber mais tokens, precisamos remover mensagens para a nova poder entrar


## Tokens e Contexto

  - Segredo de tudo é fazer a contagem dos tokens
  - Sabendo quantos tokens estamos utilizando e a quantidade máxima do modelo, podemos acumular mensagens
  - Quanto mais mensagens, melhor a resposta por conta do contexto das anteriores


## Microsserviço de Chat

  - [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

  ![Arquitetura Limpa](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)


## Pontos importantes

  - Coração da aplicação deve ter suas regras de negócio consolidadas
  - Coração da aplicação não sabe que existe a API da OpenAl
  - Armazenar todas as conversações em um banco de dados
  - Usuário poderá informar seu "user_id" como referência para ter acesso as conversas de um determinado usuário
  - Servidor Web e RPC para realizar as conversas Precisaremos gerar um Token no site da OpenAl para termos acesso a API
  - A autenticação de nosso microsserviço também será realizada via um token fixo em um arquivo de configuração


## Tecnologias

  ### Docker

  ### Containers:

  Containers no Docker são unidades de software que encapsulam um aplicativo e todas as suas dependências em um pacote portátil, que pode ser executado em qualquer sistema operacional compatível com o Docker. Eles permitem que o aplicativo seja executado de maneira isolada, sem afetar o ambiente do host ou de outros aplicativos em execução no mesmo host.

  Os containers são criados a partir de imagens do Docker, que contêm todo o software necessário para executar um aplicativo, incluindo bibliotecas, dependências e arquivos de configuração. Ao executar um container, o Docker cria uma instância da imagem e adiciona uma camada de gravação de sistema de arquivos em cima dela, que permite que o aplicativo modifique arquivos e dados dentro do container.

  Os containers do Docker são leves e rápidos de inicializar e parar, o que os torna ideais para implantação em escala e em ambientes de nuvem. Eles também fornecem uma maneira conveniente de empacotar e distribuir aplicativos, garantindo que eles funcionem de maneira consistente em qualquer ambiente onde o Docker esteja instalado. ```by ChatGPT```

  ### Containers vs VMs:

  Os contêineres do Docker e as máquinas virtuais são maneiras de implantar aplicativos em ambientes isolados do hardware subjacente. A principal diferença é o nível de isolamento.

  <img src="https://images.contentstack.io/v3/assets/blt300387d93dabf50e/bltb6200bc085503718/5e1f209a63d1b6503160c6d5/containers-vs-virtual-machines.jpg" alt="Containers vs virtual machines" width="800">

  1. [Docker vs Virtual Machines (VMs) : A Practical Guide to Docker Containers and VMs](https://www.weave.works/blog/a-practical-guide-to-choosing-between-docker-containers-and-vms)
  2. [dockerlabs](https://dockerlabs.collabnix.com)
      

  ### Dockerfile:
    
  O Dockerfile é um arquivo de texto usado para definir a configuração e o ambiente de um contêiner Docker. Ele contém uma série de instruções que o Docker usa para construir uma imagem personalizada. As instruções podem incluir coisas como a imagem base a ser usada, os comandos necessários para instalar e configurar o software no contêiner e a definição de variáveis de ambiente. Uma vez que um Dockerfile é criado, ele pode ser usado para construir uma imagem repetidamente, permitindo a criação de ambientes de contêiner consistentes e reprodutíveis.
      
  1. [Docs Dockerfile](https://docs.docker.com/engine/reference/builder/)


  ### Mapeamento de portas:

  Para realizar o mapeamento de portas basta saber qual porta será mapeada no host e qual deve receber essa conexão dentro do container.

  ```txt
  docker container run -it --rm -p "<host>:<container>" python
  ```
      
  Um exemplo com a porta 80 do host para uma porta 8080 dentro do container tem o seguinte comando:
  
  ```docker
  docker container run -it --rm -p 80:8080 python
  ```
  Com o comando acima temos a porta 80 acessível no Docker host que repassa todas as conexões para a porta 8080 dentro do container. Ou seja, não é possível acessar a porta 8080 no endereço IP do Docker host, pois essa porta está acessível apenas dentro do container que é isolada a nível de rede, como já dito anteriormente.

  1. [Stack Expert](https://stack.desenvolvedor.expert/appendix/docker/comandos.html)


  ### Docker Compose:

  O Docker Compose é uma ferramenta que permite definir e executar aplicativos multi-container Docker de maneira fácil e eficiente. Ele permite que você descreva as dependências de seus aplicativos em um arquivo YAML e, em seguida, execute e gerencie todos os contêineres necessários com um único comando. Com o Docker Compose, é possível configurar ambientes de desenvolvimento complexos com vários serviços, como bancos de dados, servidores web, cache, entre outros, e facilmente escalá-los ou destruí-los com apenas alguns comandos. O Docker Compose é uma extensão do Docker CLI e funciona em conjunto com o Docker Engine para simplificar o processo de gerenciamento de aplicativos em contêineres.

  Exemplo de Dockerfile para MySQL na versão 8

  ```yaml
  version: '3'
  services:
    mysql:
      image: mysql:8
      container_name: mysql
      restart: always
      environment:
        MYSOL_ROOT_PASSWORD: root
        MYSOL_DATABASE: chat
        MYSOL_USER: root
      port:
        - 3306:3306
      volumes:
        - ./.docker/mysql:/var/lib/mysql
  ```
  
  Comandos comuns

  ```bash
  # Inicia contêineres existentes para um ou mais serviços.
  docker-compose start

  # Interrompe a execução de contêineres sem removê-los.
  docker-compose stop

  # Pausa a execução de contêineres de um ou mais serviços.
  docker-compose pause

  # Retoma contêineres pausados ​​de um ou mais serviços.
  docker-compose unpause

  # Lista contêineres..
  docker-compose ps

  # Builds, recria, inicia e anexa a contêineres para um ou mais serviços.
  docker-compose up

  # Interrompe contêineres e remove contêineres, redes, volumes e imagens criadas.
  docker-compose down
  ```
  
  Para executar corretamente os comandos é necessário está no mesmo ```path``` do arquivo do ```Dockerfile```

  1. [Docs Docker Compose](https://docs.docker.com/compose/)
  2. [Docker do Zero a Produção](https://www.youtube.com/watch?v=62jN36WLBLk)
  3. [Docker Compose CheatSheet](https://devhints.io/docker-compose)


  ### Go
  
  Go é uma linguagem de programação de código aberto criada pela Google em 2007. Ela é projetada para ser eficiente em termos de tempo de compilação, velocidade de execução e gerenciamento de memória. Go é uma linguagem compilada, estaticamente tipada e fortemente tipada, o que significa que os tipos de variáveis ​​precisam ser declarados antes da sua utilização e não podem ser alterados durante a execução do programa.

  Go é uma linguagem multiplataforma e tem sido amplamente adotada para desenvolvimento de software em sistemas operacionais como Linux, macOS, Windows, entre outros. Ela é frequentemente usada para construir aplicativos de alto desempenho, como servidores web, ferramentas de linha de comando e sistemas distribuídos. Go tem uma sintaxe concisa e simples, tornando-a relativamente fácil de aprender e usar. Além disso, a linguagem é suportada por uma grande comunidade de desenvolvedores e tem uma ampla gama de bibliotecas e frameworks disponíveis.

  ### Baixar/Instalar Go
  1. [Instalar o Go no Host](https://go.dev/doc/install)
  2. [Instalar Extensão no VSCode](https://code.visualstudio.com/docs/languages/go)
    Inicie o VS Code Quick Open (Ctrl+P), cole o seguinte comando e pressione enter.
    
    ```txt
    ext install golang.Go
    ```
  3. Inicie o VS Code Quick Open (Ctrl+P), cole o seguinte comando e pressione enter.
    
    ```txt
    go: Install/Update Tools
    ```
    Selecione todas opções e clique em ```OK```

  4. Se seu SO é Linux e teve problemas no último passo, faça:
    - Adicione essas linhas a seu arquino ```~/.bashrc``` ou ```~/.zshrc```:
      
    ```shell
    echo export GOROOT=/usr/local/go
    echo export GOPATH=\$HOME/go
    echo export PATH=\$PATH:\$GOROOT/bin:\$GOPATH/bin
    ```
    - A extensão depende de ```go, gopls, dlv e outras ferramentas opcionais```. Se alguma das dependências estiver ausente, o aviso Faltando ferramentas de análise será exibido. Clique no aviso para baixar as dependências.

      ![baixar as dependências](https://github.com/golang/vscode-go/raw/master/docs/images/installtools.gif)

    - Depois disso, volte para o passo 3

  ### MySQL