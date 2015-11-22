go-id-cipher
=========

use RC4 cipher a integer to another integer

usage
------

in console use this to install lib

>
    go get github.com/delongw/go-int-cipher
    
and then, import it from source

>
    import "github.com/delongw/go-int-cipher"
    
use it with a private key, ** 1 < len(key) < 256 ** 

    a := int_cipher.Encrypt(1, "this is key") // a is 2554174457
    b := int_cipher.Decrypt(2554174457, "this is key") // b is 1

    c := int_cipher.Encrypt8(1, "this is key") // c is 153
    d := int_cipher.Encrypt8(153, "this is key") // d is 1
