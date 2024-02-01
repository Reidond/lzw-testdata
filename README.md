# charidy_testdata

## Build

```bash
make build
```

## Usage

Using pipes
```bash
cat <<EOF | bin/testdata
{"some":"test","data":1}
EOF
```

Using file param
```bash
bin/testdata -f some-data.json
```
