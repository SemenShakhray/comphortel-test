
## Установка

1. Клонируйте репозиторий:

```bash
git clone https://github.com/SemenShakhray/comphortel-test.git
cd comphortel-test
```

2. Запустите миграции

```bash
make install-deps //при необходиомсти
make migration-up
```

3. Запустите приложение в контейнере

```bash
docker-compose up
```

