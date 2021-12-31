# gofer
gofer is a golang script designed to output the cert information for supplied sites

## Examples
You can supply multiple sites as command line arguments.  If no port is specified, :443 is implied.

### Successful Examples

```
go run gofer.go brave.com central.crashplan.com:4287
--  brave.com:443  --
Certificate:
        DNS Name: [brave.com]
        Serial Number: 369484983675259537284761682550997610411307
        Issuer:  CN=R3,O=Let's Encrypt,C=US
        Validity:
                Not Before:              2021-12-27 20:45:25 +0000 UTC
                Not After:               2022-03-27 20:45:24 +0000 UTC
        Subject: CN=brave.com
        Authority Information Access:
                 OCSP - URI:             [http://r3.o.lencr.org]
                 CA Issuers - URI:       [http://r3.i.lencr.org/]
        SignatureAlgorithm: SHA256-RSA
        PublicKeyAlgorithm: RSA
--

--  central.crashplan.com:4287  --
Certificate:
        DNS Name: [*.us2.code42.com *.us2.crashplan.com *.crashplan.com]
        Serial Number: 18083294273383594768584993584401356329
        Issuer:  CN=DigiCert TLS RSA SHA256 2020 CA1,O=DigiCert Inc,C=US
        Validity:
                Not Before:              2021-05-20 00:00:00 +0000 UTC
                Not After:               2022-05-25 23:59:59 +0000 UTC
        Subject: CN=*.us2.code42.com,O=Code42 Software Inc,L=Minneapolis,ST=Minnesota,C=US
        Authority Information Access:
                 OCSP - URI:             [http://ocsp.digicert.com]
                 CA Issuers - URI:       [http://cacerts.digicert.com/DigiCertTLSRSASHA2562020CA1.crt]
        SignatureAlgorithm: SHA256-RSA
        PublicKeyAlgorithm: RSA
--

```

### Failure examples

```
➜  gofer git:(main) go run gofer.go expired.badssl.com wrong.host.badssl.com self-signed.badssl.com incomplete-chain.badssl.com    
--  expired.badssl.com:443  --
x509: certificate has expired or is not yet valid: current time 2021-12-31T12:47:58-06:00 is after 2015-04-12T23:59:59Z
--

--  wrong.host.badssl.com:443  --
x509: certificate is valid for *.badssl.com, badssl.com, not wrong.host.badssl.com
--

--  self-signed.badssl.com:443  --
x509: certificate signed by unknown authority
--

--  incomplete-chain.badssl.com:443  --
x509: certificate signed by unknown authority
--
```

## Caveats
* If no TCP port is supplied, :443 is checked.
* There is a 30 second tcp timeout
* Script does not handle revoked certs correctly (i.e. no feedback for cert revoked).
