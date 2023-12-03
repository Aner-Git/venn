:8080
encode gzip
log stdout
root * /webapp
header / {
	Strict-Transport-Security max-age=31536000
	Content-Security-Policy "frame-ancestors 'self'"
	X-Xss-Protection "1; mode=block"
	Referrer-Policy strict-origin
	X-Content-Type-Options nosniff
}

try_files  {path}  /index.html
file_server
