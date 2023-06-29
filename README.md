# mini-go-service

## Contactos

- Developer: pablo.halamaj@sendati.com
- Administrador: pablo.halamaj@sendati.com

## Inicio

Esta es una aplicación que escucha en un puerto HTTP y devuelve un mensaje.

## Archivos Provistos

Se proveen los siguientes archivos

- Dockerfile : archivo para la construcción de un container utilizando el codigo fuente de la app.
- README.md : este archivo, pequeña guia del repostiorio
- app.go : codigo fuente de la aplicacion.
- go.mod : metadata de la aplicación necesaria para la compilación de la misma


## Modo de uso de la aplicación

La aplicación se puede usar de dos maneras:

- Binario: se compila el codigo fuente y se ejecuta el binario final, las lineas para compilarlo se pueden tomar del Dockerfile.
- Contenedor: utilizando el dockerfile provisto se genera un contenedor y se ejecuta el mismo como cualquier otro contenedor.

## Arquitectura de la aplicación.

Ésta aplicación presenta las siguientes funcionalidades:

1. Puerto HTTP: Escucha en un puerto http que se defina con la variable de entorno PUERTO, en caso de que la misma no exista escuchará en un puerto aleatorio entre el 1001 y el 10000.
2. Log de inicios: La aplicación registra cada inicio de la misma en un archivo definido con la variable de entorno LOGFILE, en caso de que la misma no esté definida lo registrará en un archivo aleatorio dentro de la ruta /tmp
3. Respuesta a pedidos: A cada pedido HTTP la aplicación responde con el texto del archivo /config/respuesta.txt, en caso de que dicho archivo no exista responderá el mensaje "Mensaje de texto generico"
4. Información de Conexiones: La aplicación informa por consola cada vez que alguien se conecta a la misma informando Fecha e IP del origen.

## Container Registry

Para alojar las imagenes de la app este repositorio cuenta con una registry en la siguiente URL:

- registry.gitlab.com/sendati-training/mini-go-service

Para alojar sus imagenes usar un tag propio siguiendo los siguientes pasos.

1. Pedir un personal Token al administrador del Repositorio.
2. Loguearse a la registry con el Token

~~~
$ docker login registry.gitlab.com
Username (gitops-docker-push): xxxxxxx
Password: 
WARNING! Your password will be stored unencrypted in /home/xxxxxxxx/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store

Login Succeeded
$
~~~

3. Buildear el contenedor y taguearlo de forma apropiada

~~~
$ docker build -t registry.gitlab.com/sendati-training/mini-go-service/myuser:vX.Y .
~~~

4. Subir la imagen a la registry con el tag asignado.

~~~
$ docker push registry.gitlab.com/sendati-training/mini-go-service/myuser:vX.Y
~~~
