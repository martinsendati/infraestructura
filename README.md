mini-go-svc

Para el uso de esta app hay que clonar este repositorio dentro de donde se vaya a trabajar:

git clone https://github.com/martinsendati/infraestructura.git

Esto va a crear un directorio con el nombre del repositorio dentro del sector en el que se haya 
ejecutado el comando.

Una vez el repositorio clona, dentro del directorio generado, hacer un:

kubectl apply -f . 

El punto lo que hace es indicar que se quieren crear todos los recursos del directorio, caso contrario, indicar cual es el que se quiere emplear.

Este repositorio consta de los recursos necesarios para poder acceder a la app mini-go (deployment) un configMap para la edición del puerto por el cual escucha el pod, un servicio que tiene que ser editado en paralelo al configMap ya que hay que cambiar el puerto al cual se tieen que conectar (puerto del POD), y un ingress. 
Para el funcionamiento del ingress hay que hacer las siguientes ediciones con permisos de root:
Dentro del /etc/hosts/.

IP DEL CLUSTER <nombre-de-la-app>

Listo esto, el ingress queda utilizable.

No tan solo traen los recursos mencionados sino que tambien los que van a ayudar a la persistencia de la información, como un Storage Class para darle una identidad a la clase de datos y el funcionamiento que tiene que dar de sí para con los datos. Un PV que indica el almacenamiento y donde será guardada la data y por ultimo el PVC, el encargado de escuchar y enlazar el espacio necesario para el pod que solicita o reclama, todo esto por el volumen que se declaró dentro del deployment de dicho pod.

