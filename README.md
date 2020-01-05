# Estructura Api Base Golang
Proyecto que usa el framework echo para lograr una estructura de proyecto básica y base para comenzar cualquier api.
Se inspiró en el framework Laravel para PHP.

## Requisitos

- GO 1.13

## Instalación

- $ mv env-ejemplo.json env.json
- $ go mod download

## Run
- $ go run main.go 

## Opcional - Certificados para auth
- $ openssl genrsa -out private.rsa 1024
- $ openssl rsa -in private.rsa -pubout > public.rsa.pub