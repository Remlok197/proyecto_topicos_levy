# Proyecto Final para Topicos de Despliegue de Aplicaciones
Este es el repo para el proyecto final de Levy, consiste en hacer un server RestAPI con un cliente con software de escritorio asi es viva Linux
## Stack Tecnologico
* **Backend:** Para el servidor usaremos Go, especificamente su framework Gin
* **Database:** Para la base de datos usaremos MySQL, se hara sobre la base de datos de prueba de este, especificamente sobre la tabla employees
* **Frontend:** Para el cliente usaremos Electron y ya jeje.
* **Infraestructura:** Se usara Dockersito y de preferencia Docker-Compose (tu padre)
* **CI/CD:** Se usara GitHub Actions jeje

## Requisitos
1. Git
2. Docker Desktop (o Docker Engine + Docker Compose)
3. Go (Versión 1.20 o superior)
4. Node.js y npm (Para el cliente de Electron)

## API Endpoints
GET ---> /api/employees
GET ---> /api/employees/id
POST ---> /api/employees
PUT ---> /api/employees/id
DELETE ---> /api/employees/id
