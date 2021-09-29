<h1 align="center">ASCII-ART-WEB</h1>

## About The Project
Ascii-art-web is a program which consists in running HTTP server, which is written in GoLang. In this server possible to use a web GUI of [ascii-art](https://github.com/SpectreH/ascii-art) project.

## Installation
```
git clone https://github.com/SpectreH/ascii-art-web.git
cd ascii-art-web
```
## Usage
```
go run . main.go
```
Then open your browser and go to the IP address <code>localhost:8000</code>, and you should see the ASCII-ART-GENERATOR page.

Available banners:
* <code>standard</code>
* <code>shadow</code>
* <code>thinkertoy</code>

Available options:
* <code>Generate</code> - Generate art from the inputted string
* <code>Export</code> - Save generated ascii-art into txt file 

## Additional information

When you firstly get to the web page - program is going to use GET method to request generate default ascii-art <code>Hello World</code> and receive it to display at the page. With submitting button - program is going to use POST method to send data (text and a banner) to Go server to generate ascii-art and return it as a string value.

Prepared <code>Dockerfile</code> for using program at the Docker platform.

Only standard go packages were in use.  

## Author

* SpectreH (https://github.com/SpectreH)
