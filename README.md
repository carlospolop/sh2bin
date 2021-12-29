# sh2bin

Use this program to **embed sh scripts in binaries**.
Feel free to **fork this project** and **improve the obfuscation of the script** inside the new binaries (currently just B64), this could be very useful to **bypass AVs**.

## Automatic steps
```bash
git clone https://github.com/carlospolop/sh2bin
cd sh2bin
bash build.sh </path/to/yourscript.sh> # This will build binaries for several architectures
```

The new binaries will be located in the `builds/` directory


## Manual steps
1. Clone this repository
2. Base64 encode the script you want to embed: `base64 -w0 [/path/to/yourscript.sh] > script.sh.b64` (*The final file must be called "script.sh.b64"*)
3. Build the binary: `env GOOS=linux GOARCH=amd64 go build -o builds/sh2bin_$arch`
