NotifyHub Go skeleton

This folder contains a Go project skeleton mapped from the Java NotifyHub implementation.

Layout mapping:
- cmd/notifyhub -> NotifyHubApplication.java (entry)
- internal/data -> data/Data.java + global/EventEnum.java
- internal/core -> core/Observer.java, core/Publisher.java
- internal/service -> service/EventHub.java, service/Graph.java
- internal/component -> component/HubManager.java
- internal/model -> model/SinkNode.java
- internal/config -> config/Sinks.java
- internal/factory -> service/factory/Factory.java
- internal/api -> api controllers

To implement: fill each file with the Go code described in the mapping and run `go run ./cmd/notifyhub` from this directory.
