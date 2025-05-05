# Usar la imagen oficial de PostgreSQL versión 17.2
FROM postgres:17.2

# Copiar el archivo de inicialización SQL al contenedor
COPY ddl.sql /docker-entrypoint-initdb.d/

# Exponer el puerto predeterminado de PostgreSQL
EXPOSE 5432