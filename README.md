# 📺 Series Tracker API

Backend en Go para una aplicación de seguimiento de series. Permite gestionar series, ratings y archivos de imagen asociados.

---

## 🚀 Tecnologías usadas

- Go (net/http)
- SQLite (modernc.org/sqlite)
- API REST
- JSON
- Manejo manual de uploads
- CORS middleware

---

## 📁 Estructura del proyecto

```
backend/
├── internal/
│   ├── database/
│   ├── handlers/
│   │   ├── series/
│   │   └── ratings/
│   ├── middleware/
│   ├── models/
│   ├── routes/
│   └── utils/
├── uploads/
├── main.go
```

---

## ⚙️ Instalación y ejecución

### 1. Clonar repositorio

```bash
git clone https://github.com/jdivass/Proyecto1-Web-Backend.git
cd Proyecto1-Web-Backend
```

### 2. Ejecutar el servidor

```bash
go run main.go
```

El servidor corre en:

```
http://localhost:8080
```

---

## 🗄️ Base de datos

Se usa SQLite y se inicializa automáticamente al iniciar el proyecto.

### Tabla: series

- id
- title
- genre
- description
- platform
- status (0,1,2)
- total_seasons
- total_episodes
- current_season
- current_episode
- image_path
- created_at
- updated_at

### Tabla: ratings

- id
- series_id (único, 1 rating por serie)
- content
- stars_quantity (1 a 5)
- created_at

---

## 📌 Endpoints

### 🎬 SERIES

- **GET /series** → Obtener todas las series  
- **GET /series/{id}** → Obtener una serie con su rating  
- **POST /series** → Crear serie (multipart/form-data con imagen)  
- **PUT /series/{id}** → Actualizar serie  
- **DELETE /series/{id}** → Eliminar serie  

---

### ⭐ RATINGS

- **POST /series/{id}/rating** → Crear rating  
- **PUT /series/{id}/rating** → Crear o actualizar rating (upsert)  
- **DELETE /series/{id}/rating** → Eliminar rating  

---

## 🖼️ IMÁGENES

Las imágenes se almacenan en:

```
/uploads
```

Y se acceden mediante:

```
http://localhost:8080/uploads/nombre.jpg
```

En producción:

```
https://proyecto1-web-backend-production.up.railway.app/uploads/nombre.jpg
```

---

## 🌐 DEPLOY

Backend desplegado en:

```
https://proyecto1-web-backend-production.up.railway.app
```

---

## 🔐 NOTAS IMPORTANTES

- Cada serie tiene máximo 1 rating
- Ratings están ligados directamente a series
- SQLite se genera automáticamente
- CORS está habilitado

---



## 👨‍💻 AUTOR
Julián Divas


Backend desarrollado en Go para proyecto de Series Tracker

Repositorio:
https://github.com/jdivass/Proyecto1-Web-Backend.git