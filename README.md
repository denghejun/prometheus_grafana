# prometheus_grafana

1. Update your local machine ip in `./prometheus.yaml`
2. Make sure you already have a local application startup with prometheus metrics exporter(usually done with official
   prometheus client sdk and coding embed in your application, and the http exported port is 2112)

3. To startup locally

```
docker-compose up
```

4. Visit prometheus: http://localhost:9090

5. Visit grafana: [http://localhost:3000](http://localhost:3000): username/password = admin/admin, you also need to config the prometheus server url(e.g. in this
   case: http://{your-host-ip-not-localhost}:9090) as a datasource at the first time you startup grafana.
