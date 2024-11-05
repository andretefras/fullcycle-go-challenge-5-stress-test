Para executar o stress test via imagem Docker execute:

```docker run andretefras/go-stress-test --url=https://google.com --requests=10 --concurrency=2```

Caso queira realizar o build da imagem Docker localmente execute:

```docker build -t go-stress-test .```

E para executar o stress test via imagem Docker execute:

```docker run go-stress-test --url=https://google.com --requests=10 --concurrency=2```