# prometheus_grafana

1. Update your local machine ip in `prometheus.yaml`
2. Make sure you already have a local application startup with prometheus metrics exporter(usually done with official
   prometheus client sdk and coding embed in your application)

3. To startup locally

```
docker-compose up
```

4. Visit prometheus: http://localhost:9090

5. Visit grafana: http://localhost:3000, you also need to config the prometheus server url(e.g. in this
   case: http://localhost:9090) as a datasource.
