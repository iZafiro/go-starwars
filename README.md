# go-starwars

Grupo 45

Fabián Levicán 201603012-3

Paula Pérez 201603022-0

Felipe Vicencio 201603028-k

# Instalación

Debería estar todo instalado en las máquinas.

# Uso

Para cada máquina se debe usar el comando "make [programa]" con los programas correspondientes desde la carpeta go-starwars.

En dist177 ejecutar "make broker" desde en una consola.

En dist178 ejecutar "make informant" desde en una consola, y en otra consola ejecutar "make fulcrum1".

En dist179 ejecutar "make informant" desde en una consola, y en otra consola ejecutar "make fulcrum2".

En dist180 ejecutar "make leia" desde en una consola, y en otra consola ejecutar "make fulcrum3".

# Comandos

Por favor, ingresar comandos en el formato especificado **separados por comas en vez de espacios**. Por ejemplo, 
```
AddCity,Tattooine,Mos_Eisley,5
```
Además, supusimos que el comando GetNumberRebelds tenía un typo, así que el comando que implementamos es GetNumberRebels.

# Merge

Cada dos minutos llamamos a una subrutina en fulcrum1.go que coordina el merge. Obtiene los logs de registro de los otros dos fulcrum en el siguiente formato
```
[
    "Comando 1\nVector 1",
    "Comando 2\nVector 2",
    ...
]
```
donde Vector X hace referencia al reloj de vector del planeta en el Comando X más actualizado que tiene el fulcrum consultado. Posteriormente, ejecuta los comandos del fulcrum 2, y luego del fulcrum3. Es decir, los cambios con más prioridad son los del fulcrum 3, seguidos de los del fulcrum 2, seguidos de los del fulcrum 1, en el sentido previamente descrito. Los relojes de vector se actualizan varias veces, manteniendo los valores más altos en cada coordenada. Finalmente, envía los archivos y los relojes de vector que fueron modificados en el proceso anterior en el siguiente formato
```
[
    "Planeta 1\nVector 1\nLínea1\nLínea 2 ..."
    "Planeta 2\nVector 2\nLínea1\nLínea 2 ..."
]
```
para que se propaguen los cambios.

# Read Your Writes

Garantizamos Read Your Writes en los informantes pues en la primera operación de un informante X a un planeta redirigimos a un fulcrum al azar, y en las siguientes, a un fulcrum al azar con un reloj de vector del planeta más actualizado que el que tiene X.

# Monotonic Reads

Garantizamos Monotonic Reads en Leia pues en la primera consulta de Leia a un planeta llamamos a un fulcrum al azar que tiene información del planeta (si existe), y en las siguientes, a un fulcrum al azar con un reloj de vector del planeta más actualizado que el que tiene Leia.