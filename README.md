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
![Tópicos](./.github/topico-1.png)

### Tópico ~= Log
![Tópico ~= Log](./.github/topico-log.png)

### Anatomia de um registro
![Tópico ~= Log](./.github/anatomia-de-um-registro.png)

### Partições
Cada tópico pode ter uma ou mais partições para conseguir garantir a distribuição e resiliência de seus dados.
![Tópico ~= Log](./.github/particoes-em-um-topico.png)

### Sobre as "Keys"
##### Garantindo ordem de entrega
- Só é possivel garantir a ordem dentro da mesma partição

Risco: Existe uma grande chace de o Estorno ser executado antes da efetivação da transferencia, pelo fato do consumirdor 1 estar lento; Isso é um problema.
![](./.github/keys-1.png)

**Transferencia e estorno precisam estar na mesma partição**, e podemos garantir isso através das "Keys"

![](./.github/keys-2.png)

### Partições distribuidas

![](./.github/particoes-distribuidas-1.png)

**Replication Factor**
![](./.github/particoes-distribuidas-2.png)

### Partition Lidership - Partições Liderança
![](./.github/particoes-distribuidas-3.png)

### Producer: Garantia de entrega

A mensagem é enviada ao Leader sem a necessidade de confirmação
![](./.github/garantia-ack-0.png)

A mensagem é enviada e o Producer aguarda a confirmação do Leader
![](./.github/garantia-ack-1.png)

A mensagem é enviada e o Producer aguarda a confirmação e duplicação para os Followers ser realizada pelo Leader
![](./.github/garantia-ack-menos-1.png)

##### Performance VS Garantia
**At most once:** Melhor performance. Pode perder algumas mensagens
![](./.github/at-most-once.png)

**At least once:** Performance moderada. Pode duplicar mensagens
![](./.github/at-least-once.png)

**Exacly once:** Pior performance. Exatamente uma vez
1,2,3,4,5 -> [Kafka Process] -> 1,2,3,4,5
![](./.github/exacly-once.png)

### Producer: Indepotência
Producer Indepotente **OFF** : Mensagem duplicada
Producer Indepotente **ON** : Descarta mensagem duplicada
É um recurso que gera lentidão mas garante que não havera duplicidade
![](./.github/producer-indepotente.png)

### Consumer e Consumer Group

Grupo com 2 consumers lendo 3 partições.
Um consumer irá ler 2 partições
![](./.github/consumer-group-1.png)

Grupo com 3 consumers lendo 3 partições, cada um lê uma partição
![](./.github/consumer-group-2.png)

Grupo com 4 consumers lendo 3 partições.
Um consumer irá ler ficar parado, sem ler nenhuma partição.
**Se existir apenas 1 consumer no grupo ele iá ler todas as partições**
![](./.github/consumer-group-3.png)



Docker Compose:
- Zookeeper
- Kafka
- Control Center Confluent

### Commands

Create a new topic
``` bash
kafka-topics --create --topic=teste --bootstrap-server=localhost:9092 --partitions=3
kafka-topics --list --bootstrap-server=localhost:9092
```
``` bash
#Detalhamento do topic
kafka-topics --bootstrap-server=localhost:9092 --topic=teste --describe
Topic: teste    TopicId: PWQDtuVVSk2sfriWQYRZmg PartitionCount: 3       ReplicationFactor: 1    Configs: 
        Topic: teste    Partition: 0    Leader: 1       Replicas: 1     Isr: 1
        Topic: teste    Partition: 1    Leader: 1       Replicas: 1     Isr: 1
        Topic: teste    Partition: 2    Leader: 1       Replicas: 1     Isr: 1
```

``` bash
#start consumer
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste

#producer a message
kafka-console-producer --bootstrap-server=localhost:9092 --topic=teste

#--from-beginning
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste --from-beginning
```