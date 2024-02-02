# lzw-testdata

## Build

```bash
make build
```

## Usage

Using pipes
```bash
cat <<EOF | bin/lzw-testdata -
{"some":"test","data":1}
EOF
```

Using file param
```bash
bin/lzw-testdata -f some-data.json
```
