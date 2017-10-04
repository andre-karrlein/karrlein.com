FROM php:7.1.8-fpm

RUN \
    apt-get update && \
    rm -rf /var/lib/apt/lists/*

RUN docker-php-ext-install -j$(nproc) mysqli pdo_mysql

COPY . /code