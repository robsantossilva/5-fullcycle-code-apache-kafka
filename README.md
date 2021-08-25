# Apache Kafka
---

O **Apache Kafka** é uma plataforma distribuída open-source de **streaming de eventos** que é utilizada por milhares de empresas para uma alta performance em pipe line de dados, stream de analytics, integração de dados e aplicações de missão crítica.

https://kafka.apache.org

### O Mundo dos eventos
A cada dia que passa precisamos processar mais e mais eventos em diversos tipos de plataforma. Desde sistemas que precisam se comunicar, devices para IOT, monitoramento de aplicações, sistemas de alarmes, etc.

Perguntas:
- Onde salvar esses eventos?
- Como recuperar de forma rápida e simples, fazendo com que o feedback entre um processo e outro, ou mesmo entre um sistema e outro possa acontecer de maneira fluida e em tempo real?
- Como escalar?
- Como ter resiliência e alta disponibilidade?

### Kafka e seus super poderes!
- Altíssimo throughput. (Aguenta pancada)
- Latência extremamente baixa (2ms)
- Escalável
- Armazenamento
- Alta disponibilidade
- Se conecta com quase tudo
- Bibliotecas prontas para as mais diversas tecnologias
- Ferramenta open-source

### Empresas usando Kafka
- Linkedin
- Netflix
- Uber
- Twitter
- Dropbox
- Spotify
- Paypal
- Bancos...

### Conceitos e dinâmica básica de funcionamento
![Conceitos e dinâmica básica de funcionamento](./.github/kafka-dinamica.png)

- Kafka é um cluster (conjunto de maquinas), e esse cluster é formado de nós, e esses nós são chamados de Broker.
- Cada broker é uma máquina com banco de dados próprio.
- Producer: Enviar mensagem para o Kafka
- Consumer: Acessa o Kafka para ler a mensagem
- A recomendação para a utilização do kafka é de no minimo 3 Brokers

### Tópicos
É o canal de comunicação responsável por receber e disponibilizar os dados enviados para o Kafka.

