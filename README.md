# gofer
gofer is a golang script designed to output the cert information for various websites

## Example run
You can supply multiple sites with port (ie. :443) as command line arguments

```
nick$ go run gofer.go brave.com:443 amazon.com:443 namecheap.com:443
--  brave.com  --
Certificate:
        DNS Name: [brave.com]
        Serial Number: 337965366830398027248761021639441482885235
        Issuer:  CN=R3,O=Let's Encrypt,C=US
        Validity:
                Not Before:      2021-10-28 20:45:08 +0000 UTC
                Not After:       2022-01-26 20:45:07 +0000 UTC
        Subject: CN=brave.com
        Authority Information Access:
                 OCSP - URI:     [http://r3.o.lencr.org]
                 CA Issuers - URI:       [http://r3.i.lencr.org/]
        SignatureAlgorithm: SHA256-RSA
        PublicKeyAlgorithm: RSA
--

--  amazon.com  --
Certificate:
        DNS Name: [amazon.co.uk uedata.amazon.co.uk www.amazon.co.uk origin-www.amazon.co.uk *.peg.a2z.com amazon.com amzn.com uedata.amazon.com us.amazon.com www.amazon.com www.amzn.com corporate.amazon.com buybox.amazon.com iphone.amazon.com yp.amazon.com home.amazon.com origin-www.amazon.com origin2-www.amazon.com buckeye-retail-website.amazon.com huddles.amazon.com amazon.de www.amazon.de origin-www.amazon.de amazon.co.jp amazon.jp www.amazon.jp www.amazon.co.jp origin-www.amazon.co.jp *.aa.peg.a2z.com *.ab.peg.a2z.com *.ac.peg.a2z.com origin-www.amazon.com.au www.amazon.com.au *.bz.peg.a2z.com amazon.com.au origin2-www.amazon.co.jp]
        Serial Number: 18953053220451131611646046652406462681
        Issuer:  CN=DigiCert Global CA G2,O=DigiCert Inc,C=US
        Validity:
                Not Before:      2021-10-06 00:00:00 +0000 UTC
                Not After:       2022-09-19 23:59:59 +0000 UTC
        Subject: CN=*.peg.a2z.com
        Authority Information Access:
                 OCSP - URI:     [http://ocsp.digicert.com]
                 CA Issuers - URI:       [http://cacerts.digicert.com/DigiCertGlobalCAG2.crt]
        SignatureAlgorithm: SHA256-RSA
        PublicKeyAlgorithm: RSA
--

--  namecheap.com  --
Certificate:
        DNS Name: [www.namecheap.com namecheap.com]
        Serial Number: 151405733054384844453721065555891153903
        Issuer:  CN=Sectigo RSA Extended Validation Secure Server CA,O=Sectigo Limited,L=Salford,ST=Greater Manchester,C=GB
        Validity:
                Not Before:      2021-11-23 00:00:00 +0000 UTC
                Not After:       2022-12-17 23:59:59 +0000 UTC
        Subject: SERIALNUMBER=4077265,CN=www.namecheap.com,O=Namecheap\, Inc.,ST=Arizona,C=US,2.5.4.15=#131450726976617465204f7267616e697a6174696f6e,1.3.6.1.4.1.311.60.2.1.2=#130844656c6177617265,1.3.6.1.4.1.311.60.2.1.3=#13025553
        Authority Information Access:
                 OCSP - URI:     [http://ocsp.sectigo.com]
                 CA Issuers - URI:       [http://crt.sectigo.com/SectigoRSAExtendedValidationSecureServerCA.crt]
        SignatureAlgorithm: SHA256-RSA
        PublicKeyAlgorithm: RSA
--
```