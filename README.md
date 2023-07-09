# Uniswap monitor

![API](https://github.com/marcelosf89/uniswap-monitor/actions/workflows/api.yml/badge.svg)
![Monitor](https://github.com/marcelosf89/uniswap-monitor/actions/workflows/monitor.yml/badge.svg)

## Challeger Description

The objective of the assignment is to create a monitoring service for Uniswap V3 pools that continuously tracks and logs essential data points, stores them in a persistent datastore, and provides access to the data through a REST endpoint

The service should be built using golang and uses the go-ethereum package to interact with the blockchain, can use the datastore and the web framework of your choice

### Functionality

- The addresses of the pools to track are given as config and the service
- The service fetches and calculates the following datapoints from the configured pool contracts every 12 blocks and saves them into a persisted datastore of choice
- The service then exposes this data over the REST endpoints which have the following schema
    - Get the balances saved in the datastore with a block query filter
- Get the historical values of the balances in the datastore, along with the token deltas which represent the change since the previous blockNumber

- This is a monitoring system, hence nice to have recovery mechanisms in place to make the resilient on RPC provider failures

## Prerequisites

- Golang: 1.20
- Docker
- Make

## Running the Application

```make up-rebuild```

## Default Configuration

### Monitor

- RPC_ENDPOINT=wss://mainnet.infura.io/ws/v3/59c4fc44196840e99911c9c43617a803
- DATABASE_CONNECTION_STRING=mongodb://localhost:27017

By default the monitor will create a dummy Monitor state on MongoDB using this configurations 

```json
{
  "id":"1",
  "address":"0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640",
  "lastblock":17651601,
  "token0balance":"4766794417568",
  "token1balance":"-2558144422968827284321",
  "enabled": true
}
```

but you can create a new configuration saving a new `MonitorState` on `monitor` database on Mongo

> Note: if you set the `enabled` as `false` the monitor on restart will not trigger more this state

### Api
- DATABASE_CONNECTION_STRING=mongodb://localhost:27017
- PORT=3000
 

### API Mapping

- /health

Get a api health check, this health is getting the DB communication status

- /v1/api/pool/:pool_id
  - Query options:
    - `block` : list of the block that you would like to get the data, 

Sample: /v1/api/pool/0x88e6A0c2dDD26FEEb64F039a2c41296FcB3f5640?block=latest,123456,654024

This endpoint get all the data point date to theses specifics blocks if exists

- /v1/api/pool/:pool_id/historic

Sample: /v1/api/pool/0x88e6A0c2dDD26FEEb64F039a2c41296FcB3f5640/historic

Get all pool id data point


> Note: you can play if all endpoints using these http files on [Http Files](.http)