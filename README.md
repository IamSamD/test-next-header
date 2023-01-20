# test-next-header

During an investigation on 431 responses being returned by the TGG NextJS proxy, we also saw 502 reponses returned for what we beleve to be the same call containing oversized headers

This repository is a tool for testing the NextJS API with headers of different sizes

The repository contains two applications:
- nextjs-app
- test-next-header

<br>

## nextjs-app
A template Next JS application with a POST api endpoint

<br>

### Run Locally
**You must have [NodeJS](https://nodejs.org) installed**

<br>

From the root directory
```
cd nextjs-app
npm install
npm run dev
```

You wil now be able to access the Post endpoint at
```
http://localhost:3000/api/post
```
```
http://127.0.0.1:3000/api/post
```

If the request is successfull the endpoint will return the cookie in the response body

You can use whichever tool you like to send requests to this endpoint

For example, `test-next-headers`, `Postman`, `curl`, `Powershell` or your own script

<br>

---

<br>

## test-next-headers
A simple GO CLI that creates a random base64 string of a given size in bytes 

if the `--request` flag is used, it sends a post requst to the nextjs-app enpoint with the generated string in the cookie header

if the `--request` flag is used, the HTTP response and response body will be output to the terminal, otherwise the generated string will be output. 

<br>

## Builds

Built executables for Windows, Mac and Linux can be found in `./test-next-header/dist`

<br>

## Usage

Generate a random base64 string of a given size
```
cd test-next-header\dist\tnh_windows_amd64_v1
.\tnh.exe --stringSize <DESIRED_STRING_SIZE_IN_BYTES>
```
Example
```
.\tnh.exe --stringSize 16384
```

Generate a random string of a given size and send it in the cookie header to the nextjs-app
```
.\tnh.exe --stringSize <DESIRED_STRING_SIZE_IN_BYTES> --request
```

Example
```
.\tnh.exe --stringSize 16384 --request
```

If using Mac or Linux, the same usage applies but you should use the binary for your OS

**Mac** - `\dist\tnh_darwin_amd64_v1`

**Linux** - `\dist\tnh_linux_amd64_v1`

Execute with `./tnh`