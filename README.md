reflectext
==========

Stub package for reflect extensions implemented at https://code.google.com/r/carlchatfield-go-arrayof-structof

Install
=======

To enable the extensions FuncOf/ArrayOf/StructOf/InterfaceOf/Name:
 1. Clone and build https://code.google.com/r/carlchatfield-go-arrayof-structof
 2. Set the clone's root directory as $GOROOT
 3. `cd $GOPATH/src/github.com/0xfaded/reflectext`
 4. `go install -tags reflectext`

go install (or go get) without `-tags reflectext` will install panicking stubs.
