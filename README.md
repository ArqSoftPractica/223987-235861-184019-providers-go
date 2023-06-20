## Arquitectura de software en la práctica - Microservicio Providers - Obligatorio 2
Docentes: Nicolás Fornaro y Guillermo Areosa

Evelyn Jodus (223987), Hernán Reyes (235861), Michael Ellis (184019)

4 de Mayo de 2023
_____________________________________________________________________________________________

## Contenido de este Repositorio
- [Introducción y funcionalidades](#IntroYFunc)
- [Desarrollo y tecnologías utilziadas](#tecnologias)
- [Requerimientos funcionales](#RF)
- [Diagramas de vistas](#vistas)
- [Guía de despliegue](#despliegue)



## 1. Introducción y funcionalidades <a name="IntroYFunc"></a>

A lo largo de este repositorio se detallarán distintos aspectos importantes con lo que respecta al microservicio de Providers que forma parte del backend de la aplicación que se nos solicitó desarrollar

La aplicación que se nos solicitó desarrollar tiene como base el manejo de inventario para empresas. 
El principal objetivo por el cual se crea es para lograr reducir errores humanos y mejorar la eficiencia en la gestión del inventario. 
Basándonos, tanto en los requerimientos funcionales como en los no funcionales es que surge esta aplicación basada en roles. Más adelante en la documentación, detallaremos las distintas acciones capaces de realizar cada uno de los roles que se encuentran en nuestro sistema, pero adelantaremos que contaremos tanto con Administradores como con Empleados.

Este fue el microservicio que se eligió para desarrollar en otro lenguaje. El mismo fue implementado en Golang. Se encarga del manejo de los proveedores.


## 2. Desarrollo y tecnologías utilziadas  <a name="tecnologias"></a>
### Desarrollo
En esta segunda versión. el equipo determinó tener una arquitectura basada en microservicios cloud native por ser la más apropiada. 
A continuación detallaremos las decisiones de tecnologías utilizadas para este microservicio: 
#### Backend
Para el requerimiento que solicitaba desarrollar un microservicio en otro lenguaje decidimos utilizar Golang, el mismo por haber sido mencionado en muchos ejemplos en clase. En base a esto, decidimos investigar los beneficios de utilizar Go para el desarrollo del servicio requerido en otro lenguaje y nos encontramos con características de ser un lenguaje con un muy buen rendimiento en comparación con otros. Por ejemplo, a nivel de performance, es más rápido incluso que NodeJs, lo que concluimos incluso en nuestra propia experiencia, luego de tener el mismo servicio desarrollado en NodeJs vs el que desarrollamos en Go. También ofrece características excelentes en cuanto al manejo y performance de concurrencia. A su vez nos encontramos con que es un lenguaje open-source, con una comunidad muy amplia, lo que lo hace ser muy confiable.


#### Despliegue 
Para el despliegue de la aplicación Backend utilizamos AWS. La misma fue elegida por ser la ejemplificada en el curso. 
Más adelante en el Readme se tendrá acceso a la guia de despliegue utilizada

## 3. Requerimientos funcionales <a name="RF"></a>
En la siguiente tabla se verán reflejados los distintos requerimientos funcionales solicitados, en conjunto con los actores correspondientes: 
Tengamos en cuenta que todos estos requerimientos no se ven implementados en este microservicio, sino que en el conjunto de todos los microservicios de backend

| **Requerimiento** | **Descripción** | **Actor** |
|:---:|:---:|:---:|
| RF1 | Permite el registro de un usuario de tipo Administrador | Usuario Administrador |
| RF2 | Permite el registro tanto de usuarios administradores como usuarios empleados vía un link por mail | Usuario Administrador / Usuario Empleado |
| RF3 | Permite tanto el login como el logout del sistema de los distintos usuarios | Usuario Administrador / Usuario Empleado |
| RF4.1 | Permite dar de alta productos | Usuario Administrador |
| RF4.2 | Permite la modificación de productos | Usuario Administrador |
| RF4.3 | Permite la eliminación lógica de productos | Usuario Administrador |
| RF5.1 | Permite dar de alta proveedores | Usuario Administrador |
| RF5.2 | Permite la modificación de proveedores | Usuario Administrador |
| RF5.3 | Permite la eliminación lógica de proveedores | Usuario Administrador |
| RF6 | Permite el registro de compras de productos  | Usuario Administrador |
| RF7 | Permite el registro de ventas de productos | Usuario Administrador / Usuario Empleado |
| RF8 | Permite llevar un control del inventario de la tienda, actualizando cada vez que se registre una compra o una venta | Usuario Administrador / Usuario Empleado |
| RF9 | Se permite visualizar una pantalla con las ventas realizadas para un cierto período | Usuario Administrador / Usuario Empleado |
| RF10 | Disponibilización de un endpoint con los 3 productos más vendidos de forma histórica para una empresa | Endpoint público |
| RF11 | Disponibilización de un endpoint que dado un periodo de tiempo, te devuelva todas las compras a un proveedor X en ese tiempo | Endpoint público |
| RF12 | Permite la subscripción a producto desde un usuario | U. Administrador |
| RF13 | Permite la de-subscripción a producto desde un usuario | U. Administrador |
| RF14 | Permite realizar una venta programada  | U. Administrador / U. Empleado |
| RF15 | Permite recibir reportes de la empresa a la que pertenece | U. Administrador / U. Empleado |
| RF16 | Permite recibir notificaciones de productos que no tienen stock | U. Administrador |

## 4. Diagramas de vistas <a name="vistas"></a>
Describiremos la arquitectura de nuestro sistema representados con las distintas vistas de arquitectura. Cada una de ellas nos provee de distintos detalles del sistema, así pudiendo llegar a lo que es el mismo en su completitud.

### Vista de módulos 
Las vistas de módulos se usan para presupuestar, estimar y asignar tareas y seguimiento de los proyectos. Además, a nivel de la construcción, nos proveen una especie de plano que refleja el código fuente. También nos ayudan a analizar que impacto de las modificaciones. Por último, podemos decir que son muy útiles para poder instruir a nuevos desarrolladores.

#### Vista de descomposición 
En la vista de descomposición podemos describir la estructura jerárquica del sistema, partiendo del elemento de más alto nivel y documentando así de forma recursiva los elementos de la jerarquía.
![image](https://github.com/ArqSoftPractica/223987-235861-184019-providers/assets/44271850/7a1d6629-d0a1-4a46-bc33-6732666ff7ef)






#### Vista de usos
La siguiente vista describe las dependencias de usos entre los módulos del sistema
![image](https://github.com/ArqSoftPractica/223987-235861-184019-providers/assets/44271850/113ca16e-7138-4537-a931-413c563f6831)





#### Vista de asignación y despliegue
En esta sección describiremos los estilos que consideramos fundamentales para comunicar la forma como los elementos del software, principalmente cómo se relacionan los componentes con su entorno. En particular se describe la forma como los componentes se despliegan en los diferentes nodos físicos (servidores). Además de describir cómo se relacionan los módulos con el ambiente de desarrollo. Representamos los servicios como nodos, ya que son elementos físicos que existen en tiempo de ejecución y representan un recurso computacional, que poseen memoria y capacidad de procesamiento de los datos

![Diagramas de vistas-Despliegue OBL2 drawio (1)](https://github.com/ArqSoftPractica/223987-235861-184019-providers-go/assets/44271850/5a064416-ef79-4bc9-abf9-04e398ae7ea8)


En este diagrama podemos ver los distintos microservicios y la forma en la cual estos se comunican con el cliente. 
Para poder comunicar el frontend con los distintos microservicios hacemos uso de una API Gateway. En este caso, el frontend le hace la request a la API Gateway y luego la misma la redirige al microservicio que corresponda.


## 5. Guia de despliegue <a name= "despliegue"></a>
### Guia para crear un deploy desde cero 

**Paso 1:** Se debe comprimir todo el contenido de la carpeta del proyecto a un archivo .zip. 

**Paso 2:** Iniciar laboratorio en AWS.

**Paso 3:** Dentro de la consola ir a Elastic Beanstalk y crear una nueva aplicación: Solo modificar los campos definidos en las capturas de pantalla. El resto de campos dejarlos por defecto. 

**Paso 4:** Seleccionar el tipo de environment: web server environment.

![image](https://user-images.githubusercontent.com/44271850/236257662-75535a92-6edb-45ae-95a7-d7018db423f3.png)

**Paso 5:** Ingresar un nombre de aplicación. En nuestro caso fue: backend-ellis-jodus-reyes

![image](https://user-images.githubusercontent.com/44271850/236257715-589b23fc-7eec-4c56-9462-c099f57799cc.png)

**Paso 6:** Seleccionar como plataforma a "Docker". (Dentro de los elementos que contenía la carpeta que comprimimos se encuentra el Docekerfile.)

![image](https://user-images.githubusercontent.com/44271850/236257775-cd1473dc-c7f3-4dde-9ccd-dda9f412c10b.png)

**Paso 7:** Seleccionar "Upload your code"

![image](https://user-images.githubusercontent.com/44271850/236257839-8eb52d4d-aac6-4aba-ba8d-fadbc7469ae5.png)

**Paso 8:** Seleccionar "Local file" y hacer click en "Choose file" (Verificar que el Dockerfile esta en la raíz de lo que comprimimos)

![image](https://user-images.githubusercontent.com/44271850/236257927-b13d3232-3d5f-43c5-b722-df8cda9d816a.png)

**Paso 9:** Click en "Configure More Options"

![image](https://user-images.githubusercontent.com/44271850/236257968-f67cee38-a382-447b-93b5-1318deff6154.png)

**Paso 10:** Para evitar un error que dice que no tenemos un cierto permiso, tenemos que darle a EB un rol que nos trae AWS Academy que tiene más permisos (LabRole). Para eso:

![image](https://user-images.githubusercontent.com/44271850/236258137-d6cbb75b-975d-42f6-8d14-837f38bc3be2.png)

**Paso 11:** Hacer click en "Create application"
El proceso de creación toma unos minutos. Podemos observar los procesos que se van generando:

![image](https://user-images.githubusercontent.com/44271850/236258201-39846707-7a8a-48b5-9bbc-f49e396d5e5d.png)

Finalmente hacer click en la URL disponible para acceder a la aplicación
Desplegada. En nuestro caso: backend-ellis-jodus-reyes.us-east-1.elasticbeanstalk.com
En el ejemplo: TestNodeBeanstalk-env.eba-mhvdih5x.us-east-1.elasticbeanstalk.com

![image](https://user-images.githubusercontent.com/44271850/236258273-61a93591-2bd2-4b5f-ad73-51e7ba85f503.png)

### Guia para crear una nueva versión de deploy 

**Paso 1:** Dirigirse a Elastic Beanstalk en aws, ir al environment que ya se había creado. Deberá ver una pantalla como la siguiente:

![image](https://user-images.githubusercontent.com/44271850/236259566-937a3d7a-dba5-4c2a-a683-287ab2bb7d8a.png)

**Paso 2:** fijarse el número de la última versión deployada y generar el nuevo número, ingresandolo en package.json

Cambiar el número de versión en el archivo package.json del proyecto.

![image](https://user-images.githubusercontent.com/44271850/236259658-aea42d20-d7b1-4f3b-8f94-0236a0339da1.png)

Si la versión previa era la 0.0.4, ingrese 0.0.5 por ejemplo, o dependiendo del versionado que se utilice, cambie los números que correspondan.

**Paso 3:** seleccione todos los archivos de la raíz del proyecto y comprimirlos, generando un .zip

**Paso 4:** hacer click en Upload and deploy.

**Paso 5:** haga click en choose file y seleccione el archivo .zip que comprimió en el paso 3.

![image](https://user-images.githubusercontent.com/44271850/236259759-a9f34d0c-e01b-4f1c-a631-412c58d76a4f.png)

**Paso 6:** Ingrese la versión correcta, que fue la que ingresó en el archivo package.json del proyecto.

![image](https://user-images.githubusercontent.com/44271850/236259841-3998d824-d736-4086-8616-3a934e07c937.png)

**Paso 7:** Seleccione “Deploy”.

![image](https://user-images.githubusercontent.com/44271850/236259913-8bfd3d17-0e8c-4896-9994-757eaf67b3ec.png)

**Paso 8:** AWS hará las modificaciones  y los updates correspondientes, poniendo en marcha el servicio de la nueva versión. Podrá ver los eventos que aws realiza en la pestaña de eventos.

**Paso 9:** Una vez finalizado este proceso, podrá confirmarlo viendo el estatus del proyecto:

![image](https://user-images.githubusercontent.com/44271850/236259988-620dbe5d-5e39-4d7d-b354-4af7d1c0003c.png)

En la siguiente pantalla:

![image](https://user-images.githubusercontent.com/44271850/236260257-41f729bb-76db-4186-9a9b-53cd7953dca7.png)
