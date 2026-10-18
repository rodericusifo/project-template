# 🏗️ Infrastructure

This directory contains optional local dependencies for the included example.
Infrastructure is kept separate from [`workload/`](../workload/) so that
dependencies can be selected, replaced, or removed without changing the
workload layout.

The included definitions are development-oriented Docker Compose stacks. They
are useful as a reference and local environment, but should be reviewed
before being used for production or real data.

## 🗂️ Directory Layout

```text
infrastructure/
├── cache/
│   ├── Makefile
│   └── redis/
│       ├── docker-compose.yaml
│       └── env/.env.example
├── database/
│   └── sql/
│       ├── Makefile
│       └── postgres/
│           ├── docker-compose.yaml
│           ├── env/.env.example
│           ├── primary/
│           │   ├── config/
│           │   └── initdb/
│           └── replica/config/
└── messaging/
    ├── Makefile
    ├── redpanda/
    │   ├── docker-compose.yaml
    │   ├── config/
    │   └── env/.env.example
    └── lavinmq/
        ├── docker-compose.yaml
        └── env/.env.example
```

## 🧱 Included Resources

| Category | Selection value | Local topology |
| --- | --- | --- |
| SQL database | `postgres` | PostgreSQL primary and a disposable streaming replica |
| Cache | `redis` | Redis primary and replica |
| Messaging | `redpanda` | Three Redpanda brokers, Console, and RustFS |
| Messaging | `lavinmq` | Three LavinMQ nodes and etcd |

These technologies are sample choices, not template requirements. A project
can replace them with another database, cache, broker, hosted dependency, or
orchestration system.

## 🎛️ Selecting Resources

The root [stack.mk](../stack.mk) controls what the root `make start` and
`make stop` commands operate:

```makefile
MESSAGING := redpanda lavinmq
CACHE := redis
DATABASE_SQL := postgres
```

Use one or more supported values in each variable. The category Makefiles
validate selections before invoking Docker Compose:

- `infrastructure/database/sql/Makefile` currently supports `postgres`.
- `infrastructure/cache/Makefile` currently supports `redis`.
- `infrastructure/messaging/Makefile` currently supports `redpanda` and
  `lavinmq`.

An unsupported value fails fast instead of starting an unknown Compose stack.

## 🔐 Environment Configuration

Every resource has an `.env.example`. Generate environment-specific files
from the repository root:

```bash
make setup-env ENV=dev
```

The command creates files such as
`infrastructure/cache/redis/env/dev.env` and
`infrastructure/database/sql/postgres/env/dev.env`. Each category loads its
selected file with Docker Compose.

The main configuration groups are:

```dotenv
# PostgreSQL
DATABASE_SQL_POSTGRES_USERNAME=
DATABASE_SQL_POSTGRES_PASSWORD=
DATABASE_SQL_POSTGRES_REPLICATOR_PASSWORD=
DATABASE_SQL_POSTGRES_TIME_ZONE=UTC

# Redis
CACHE_REDIS_USERNAME=
CACHE_REDIS_PASSWORD=

# LavinMQ
MESSAGING_LAVINMQ_MAIN_DEFAULT_USER=
MESSAGING_LAVINMQ_MAIN_DEFAULT_PASSWORD_HASH=
```

Redpanda also requires broker credentials, Console credentials and JWT
configuration. See
[`redpanda/env/.env.example`](messaging/redpanda/env/.env.example) for the
complete list.

Review every generated file before starting the stack. Do not commit generated
environment files, passwords, tokens, or cloud-storage credentials.

## 🌐 Network and Naming

The root Makefile creates one external Docker network per environment:

```text
<project-slug>_<environment>
```

For example, the default project may use `api_template_dev`. All infrastructure
and workloads join this network and should use Docker service/container names
for internal connections. Host-published ports are intended for local tools
running outside Docker.

The stack also derives resource volume names from `SLUG` and `ENV`, which keeps
different environments isolated:

```text
<slug>_<resource>_<environment>
```

Changing `SLUG` or `ENV` points Compose at a different set of external volumes.

## ▶️ Lifecycle

The normal root workflow creates the network, starts database/cache/messaging
resources, and then starts the workloads:

```bash
make start ENV=dev
make stop ENV=dev
```

`make stop` stops Compose projects but does not remove named volumes. This
allows the next start to reuse local data. To inspect resources:

```bash
docker compose ls
docker ps
docker volume ls
docker network ls
```

To operate one category directly, pass the variables normally supplied by the
root Makefile:

```bash
make -C infrastructure/cache \
  CACHE=redis ENV=dev SLUG=api_template NAME=api-template run-cache

make -C infrastructure/cache \
  CACHE=redis ENV=dev SLUG=api_template NAME=api-template stop-cache
```

Direct category commands assume the target environment file and external
network already exist. Prefer the root commands for a complete local stack.

## 🔌 Ports

The current Compose files publish these host ports:

| Resource | Host ports | Purpose |
| --- | --- | --- |
| PostgreSQL primary | `5433` | Primary database |
| PostgreSQL replica | `5434` | Replica database |
| Redis primary | `6380` | Primary cache |
| Redis replica | `6381` | Replica cache |
| Redpanda broker 0 | `19092`, `18081`, `18082`, `19644` | Kafka, Schema Registry, HTTP Proxy, Admin API |
| Redpanda brokers 1–2 | `29092`/`39092` and related API ports | Additional broker endpoints |
| Redpanda Console | `8080` | Web console |
| RustFS | `9000`, `9001` | S3 API and console |
| LavinMQ nodes | `5672`–`5674`, `15672`, `25672`, `35672` | AMQP and management endpoints |

Internal workload connections should use the service names and container ports
defined in Compose, not the host-published ports.

## 💾 Persistence and Safety

The Makefiles create external Docker volumes before starting a resource.
`make stop` preserves those volumes. Remove them only when intentionally
resetting local state, and inspect the exact names first:

```bash
docker volume ls
docker volume inspect <volume-name>
docker volume rm <volume-name>
```

The PostgreSQL replica is intentionally disposable for local development. Its
Compose command clears the replica data directory and obtains a fresh base
backup from the primary whenever the replica is started. Do not use this
definition unchanged for a promoted or production database.

## 🔑 LavinMQ Password Hash

LavinMQ expects `MESSAGING_LAVINMQ_MAIN_DEFAULT_PASSWORD_HASH`, not a plain-text
password. Generate a hash with a temporary container:

```bash
docker run -d --name lavinmq_temp \
  -p 5672:5672 -p 15672:15672 cloudamqp/lavinmq
docker exec -it lavinmq_temp lavinmqctl hash_password yourdesiredpassword
docker rm -f lavinmq_temp
```

Copy the command output into the selected environment file, then remove the
temporary container as shown.

## 🩺 Troubleshooting

### A resource is rejected as unavailable

Check the selection values in `stack.mk` against the supported values listed
above. The Makefiles reject unknown values before Compose starts.

### Containers cannot reach one another

Confirm that the environment-specific network exists and that all stacks use
the same `SLUG` and `ENV`:

```bash
docker network ls
docker network inspect <slug>_<env>
```

### Configuration changes are not visible

Compose reads the environment file for the selected resource. Confirm that
the file exists at the expected path and restart that category after changing
it:

```bash
make -C infrastructure/messaging \
  MESSAGING=lavinmq ENV=dev SLUG=api_template NAME=api-template stop-messaging
make -C infrastructure/messaging \
  MESSAGING=lavinmq ENV=dev SLUG=api_template NAME=api-template run-messaging
```

### A replica does not start

Inspect the dependency logs and health status:

```bash
docker ps
docker logs primary_postgres_dev
docker logs replica_postgres_dev
```

For the PostgreSQL replica, verify the primary credentials and replication
password in the environment file. Remember that restarting the replica
recreates its local data from the primary.
