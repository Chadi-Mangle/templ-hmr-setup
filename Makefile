dev: 
	parallel -j 2 --line-buffer ::: 'npx tailwindcss -i input.css -o static/css/tw.css --watch' 'templ generate -watch -proxy="http://localhost:3000" -proxyport="8080" -cmd="go run ."'