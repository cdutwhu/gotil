#!/bin/bash

rm -rf ./embres/cache
rm -f ./io/*.txt

# delete all binary files
find . -type f -executable -exec sh -c "file -i '{}' | grep -q 'x-executable; charset=binary'" \; -print | xargs rm -f
for f in $(find ./ -name '*.log' -or -name '*.csv' -or -name '*.txt'); do rm $f; done
find ./ -type d -empty -delete