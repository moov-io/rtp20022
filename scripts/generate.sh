#!/bin/bash
set -e

#../xsd2go/gocomply_xsd2go convert \
#   --xsd-file=xsd/xmldsig-core-schema.xsd \
#   --output-dir=gen \
#   --output-file=signature.go \
#   --go-package=head_001_001_01 \
#   --template-package=rtp

../xsd2go/gocomply_xsd2go convert \
   xsd/messages.xsd \
   github.com/moov-io/rtp20022 \
   gen \
   --template-package=internal/templates/rtp20022 \
   --template-name=messages.tmpl
