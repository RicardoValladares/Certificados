openssl req -x509 -newkey rsa:4096 -sha256 -days 3650 -nodes -keyout example.com.key -out example.com.crt -subj "/CN=example.com"  -addext "subjectAltName=DNS:example.com,DNS:*.example.com,IP:127.0.0.1"




Cada uno de los comandos anteriores crea un certificado que es

válido para el dominio example.com (SAN), también válido para el dominio comodín *.example.com (SAN), también válido para la dirección IP 10.0.0.1 (SAN), relativamente fuerte (a partir de 2023) y válido para 3650 días (~10 años). Se generan los siguientes archivos:

Clave privada: example.com.key Certificado: example.com.crt Toda la información se proporciona en la línea de comando. No hay ninguna entrada interactiva que te moleste. No hay archivos de configuración con los que tengas que jugar. Todos los pasos necesarios se ejecutan mediante una única invocación de OpenSSL: desde la generación de la clave privada hasta el certificado autofirmado.

Observación #1: Parámetros criptográficos Dado que el certificado está autofirmado y los usuarios deben aceptarlo manualmente, no tiene sentido utilizar una caducidad corta o una criptografía débil.

En el futuro, es posible que desee utilizar más de 4096 bits para la clave RSA y un algoritmo hash más potente que sha256, pero a partir de 2023 estos son valores sensatos. Son lo suficientemente potentes y, al mismo tiempo, compatibles con todos los navegadores modernos.

Observación #2: Parámetro "-nodes" Teóricamente podría omitir el parámetro -nodes (que significa "sin cifrado DES"), en cuyo caso ejemplo.key se cifraría con una contraseña. Sin embargo, esto casi nunca es útil para la instalación de un servidor, porque también tendría que almacenar la contraseña en el servidor o ingresarla manualmente en cada reinicio.
