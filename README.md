openssl req -x509 -newkey rsa:4096 -sha256 -days 3650 -nodes -keyout example.com.key -out example.com.crt -subj "/CN=example.com"  -addext "subjectAltName=DNS:example.com,DNS:*.example.com,IP:127.0.0.1"

Each of the above commands creates a certificate that is

valid for the domain example.com (SAN),
also valid for the wildcard domain *.example.com (SAN),
also valid for the IP address 10.0.0.1 (SAN),
relatively strong (as of 2023) and
valid for 3650 days (~10 years).
The following files are generated:

Private key: example.com.key
Certificate: example.com.crt
All information is provided at the command line. There is no interactive input that annoys you. There are no config files you have to mess around with. All necessary steps are executed by a single OpenSSL invocation: from private key generation up to the self-signed certificate.

Remark #1: Crypto parameters
Since the certificate is self-signed and needs to be accepted by users manually, it doesn't make sense to use a short expiration or weak cryptography.

In the future, you might want to use more than 4096 bits for the RSA key and a hash algorithm stronger than sha256, but as of 2023 these are sane values. They are sufficiently strong while being supported by all modern browsers.

Remark #2: Parameter "-nodes"
Theoretically you could leave out the -nodes parameter (which means "no DES encryption"), in which case example.key would be encrypted with a password. However, this is almost never useful for a server installation, because you would either have to store the password on the server as well, or you'd have to enter it manually on each reboot.
