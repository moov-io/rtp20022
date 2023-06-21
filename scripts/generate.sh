#!/bin/bash
set -e

files=(
  "xsd/head.001.001.01.xsd":"head"
  "xsd/pacs.008.001.08.xsd":"ct"
  "xsd/pacs.002.001.10.xsd":"ps"
  "xsd/camt.035.001.05.xsd":"ac"
  "xsd/camt.056.001.08.xsd":"rt"
  "xsd/pain.013.001.07.xsd":"pr"
  "xsd/camt.029.001.09.xsd":"tr"
  "xsd/admn.005.001.01.xsd":"er"
  "xsd/admn.006.001.01.xsd":"re"
  "xsd/admn.003.001.01.xsd":"fr"
  "xsd/admn.004.001.01.xsd":"rf"
  "xsd/admn.001.001.01.xsd":"sr"
  "xsd/admn.002.001.01.xsd":"rs"
  "xsd/admi.004.001.02.xsd":"ne"
  "xsd/admi.002.001.01.xsd":"mr"
)

for file in "${files[@]}"
do
    xsdFile=${file%%:*}
    nsPrefix=${file#*:}
    name=$(basename ${xsdFile})
    name=${name%".xsd"}
    name=$(echo "${name}" | tr '.' '_')
    echo "creating ${name} from ${xsdFile}"

    mkdir -p gen/${name}

    ../xsd2go/gocomply_xsd2go convert \
       --xsd-file=${xsdFile} \
       --output-dir=gen/${name} \
       --output-file=${name}.go \
       --go-package=${name} \
       --namespace-prefix=${nsPrefix} \
       --template-package=rtp
done
