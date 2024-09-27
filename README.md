# Go, Prometheus, Grafana

This is an example repo of how you would setup Prometheus and Grafana to monitor your Go application.

You can launch the application using `docker compose up -d`

Navigate to `localhost:9090` for the Prometheus dashboard, and to `localhost:3000` for the Grafana dashboard.

Make sure to add Prometheus as a data source in Grafana. You can impoort a [Go Metrics](https://grafana.com/grafana/dashboards/10826-go-metrics/)
dashboard to quicky setup your metrics.
