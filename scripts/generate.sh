#!/bin/bash
set -e

moovio_xsd2go convert \
   xsd/xmldsig-core-schema.xsd \
   github.com/moov-io/rtp20022 \
   gen \
   --template-name=internal/templates/rtp20022/signature.tmpl \
   --xmlns-override="http://www.w3.org/2000/09/xmldsig#=xmldsig"

moovio_xsd2go convert \
   xsd/messages.xsd \
   github.com/moov-io/rtp20022 \
   gen \
   --template-name=internal/templates/rtp20022/messages.tmpl \
   --xmlns-override="urn:iso:std:ma:20022:tech:xsd:admn.007.001.01=admn_007_001_01" \
   --xmlns-override="urn:iso:std:ma:20022:tech:xsd:admn.008.001.01=admn_008_001_01" \
   --xmlns-override="urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02=acmt_022_001_02" \
   --xmlns-override="urn:iso:std:iso:20022:tech:xsd:admi.002.001.01=admi_002_001_01" \
   --xmlns-override="urn:iso:std:iso:20022:tech:xsd:admi.004.001.02=admi_004_001_02" \
   --xmlns-override="urn:iso:std:iso:20022:tech:xsd:admn.001.001.01=admn_001_001_01" \
   --xmlns-override="urn:iso:std:iso:20022:tech:xsd:admn.002.001.01=admn_002_001_01" \
   --xmlns-override="urn:iso:std:iso:20022:tech:xsd:admn.003.001.01=admn_003_001_01" \
   --xmlns-override="urn:iso:std:iso:20022:tech:xsd:admn.004.001.01=admn_004_001_01" \
   --xmlns-override="urn:iso:std:iso:20022:tech:xsd:admn.005.001.01=admn_005_001_01" \
   --xmlns-override="urn:iso:std:iso:20022:tech:xsd:admn.006.001.01=admn_006_001_01" \
   --xmlns-override="urn:iso:std:iso:20022:tech:xsd:camt.026.001.07=camt_026_001_07" \
   --xmlns-override="urn:iso:std:iso:20022:tech:xsd:camt.028.001.09=camt_028_001_09" \
   --xmlns-override="urn:iso:std:iso:20022:tech:xsd:camt.029.001.09=camt_029_001_09" \
   --xmlns-override="urn:iso:std:iso:20022:tech:xsd:camt.035.001.05=camt_035_001_05" \
   --xmlns-override="urn:iso:std:iso:20022:tech:xsd:camt.056.001.08=camt_056_001_08" \
   --xmlns-override="urn:iso:std:iso:20022:tech:xsd:head.001.001.01=head_001_001_01" \
   --xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10=pacs_002_001_10" \
   --xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.008.001.08=pacs_008_001_08" \
   --xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.009.001.08=pacs_009_001_08" \
   --xmlns-override="urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03=pacs_028_001_03" \
   --xmlns-override="urn:iso:std:iso:20022:tech:xsd:pain.013.001.07=pain_013_001_07" \
   --xmlns-override="urn:iso:std:iso:20022:tech:xsd:pain.014.001.07=pain_014_001_07" \
   --xmlns-override="urn:iso:std:iso:20022:tech:xsd:remt.001.001.04=remt_001_001_04"

# run go fmt and goimports for every generated file
files=($(find ./gen -name '*.go'))
for file in "${files[@]}"
do
    gofmt -w $file
    goimports -w $file
done
