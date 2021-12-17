# gofer
gofer is a golang script designed to output the cert information for various websites

## Examples
You can supply multiple sites with port (ie. :443) as command line arguments

### Successful Examples

```
go run gofer.go brave.com:443 namecheap.com:443
--  brave.com:443  --
Certificate:
        DNS Name: [brave.com]
        Serial Number: 337965366830398027248761021639441482885235
        Issuer:  CN=R3,O=Let's Encrypt,C=US
        Validity:
                Not Before:              2021-10-28 20:45:08 +0000 UTC
                Not After:               2022-01-26 20:45:07 +0000 UTC
        Subject: CN=brave.com
        Authority Information Access:
                 OCSP - URI:             [http://r3.o.lencr.org]
                 CA Issuers - URI:       [http://r3.i.lencr.org/]
        SignatureAlgorithm: SHA256-RSA
        PublicKeyAlgorithm: RSA
--

--  namecheap.com:443  --
Certificate:
        DNS Name: [www.namecheap.com namecheap.com]
        Serial Number: 151405733054384844453721065555891153903
        Issuer:  CN=Sectigo RSA Extended Validation Secure Server CA,O=Sectigo Limited,L=Salford,ST=Greater Manchester,C=GB
        Validity:
                Not Before:              2021-11-23 00:00:00 +0000 UTC
                Not After:               2022-12-17 23:59:59 +0000 UTC
        Subject: SERIALNUMBER=4077265,CN=www.namecheap.com,O=Namecheap\, Inc.,ST=Arizona,C=US,2.5.4.15=#131450726976617465204f7267616e697a6174696f6e,1.3.6.1.4.1.311.60.2.1.2=#130844656c6177617265,1.3.6.1.4.1.311.60.2.1.3=#13025553
        Authority Information Access:
                 OCSP - URI:             [http://ocsp.sectigo.com]
                 CA Issuers - URI:       [http://crt.sectigo.com/SectigoRSAExtendedValidationSecureServerCA.crt]
        SignatureAlgorithm: SHA256-RSA
        PublicKeyAlgorithm: RSA
--

```

### Failure examples

```
go run gofer.go expired.badssl.com:443 wrong.host.badssl.com:443 self-signed.badssl.com:443 incomplete-chain.badssl.com:443
--  expired.badssl.com:443  --
x509: certificate has expired or is not yet valid: current time 2021-12-13T16:00:35-06:00 is after 2015-04-12T23:59:59Z
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
* There is a 30 second tcp timeout
* Script does not handle revoked certs correctly (i.e. no feedback for cert revoked).
