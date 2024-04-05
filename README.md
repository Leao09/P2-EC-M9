# Prova 2 EC-M9

## Como rodar o projeto 
- É necessário primeiramente clonar o repositório utilizando o seguinte comando:
<pre><code>
 git clone  https://github.com/Leao09/P2-EC-M9.git
</code></pre>
- Depois entrar no diretório do projeto
<pre><code>
 cd P2-EC-M9
</code></pre> 
- Em sequida é necessário ter instalado a linguagem  go  e a configuração do conductor feita
- Esse links podem ajudar [Go](https://go.dev/dl/) e [Conductor](https://rmnicola.github.io/m9-ec-encontros/kafka-setup)
- Assim, inicie um módulo para o diretório e depois baixe as depencias para o projeto
<pre><code>
 go mod init seu-no
 go mod tidy
</code></pre>
- Agora é necessário rodar este código para subir um docker compose com o seu cluster kafka
<pre><code>
 docker compose up 
</code></pre> 
- Com o kafka funcionando voce pode acessar a sua interface do conductor para conectar o seu cluster com os comando explicandos na configuração assim para conectar com a porta do seu kafka e após visualizar as informações do seu cluster no conductor basta entrar nos diretórios de produce e consumer
<pre><code>
cd produce
 go run producer.go
</code></pre> 
<pre><code>
cd consumer
 go run consumer.go
</code></pre> 
- Por fim, podemos rodar os tests que irão fazer validações sobre o nosso producer
<pre><code>
go test -v -cover 
</code></pre> 
## Video 
[Link](https://youtu.be/iRLPVaaczQk)