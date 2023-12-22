package handler

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	f := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
	
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Go developer | Ilyas Amantaev</title>
	</head>
	<body style="padding: 0 15% 0 15%;">
		<style>
			h1 {
				box-sizing: border-box;
				break-after: avoid;
				color: rgb(51, 51, 51);
				font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
				font-size: 32px;
				font-weight: 700;
				letter-spacing: 0.2px;
				line-height: 54.4px;
				margin: 0 0 27.2px 0;
				overflow-wrap: break-word;
				position: relative;
				--rem: 16
			}
			h5 {
				box-sizing: border-box;
				break-after: avoid;
				color: rgb(51, 51, 51);
				font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
				font-size: 16px;
				font-weight: 700;
				letter-spacing: 0.2px;
				line-height: 27.2px;
				margin: 20.4px 0 13.6px 0;
				overflow-wrap: break-word;
				position: relative;
				text-rendering: optimizelegibility;
				-moz-text-size-adjust: none;
				--rem: 16
			}
			p {
				box-sizing: border-box;
				color: rgb(51, 51, 51);
				font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
				font-size: 16px;
				letter-spacing: 0.2px;
				line-height: 27.2px;
				margin: 0 0 13.6px 0;
				overflow-wrap: break-word;
				text-rendering: optimizelegibility;
				-moz-text-size-adjust: none;
				--rem: 16
			}
			blockquote {
				border-left-color: rgb(229, 229, 229);
				border-left-style: solid;
				border-left-width: 4px;
				box-sizing: border-box;
				break-inside: avoid;
				color: rgb(133, 133, 133);
				font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
				font-size: 16px;
				letter-spacing: 0.2px;
				line-height: 27.2px;
				margin: 0 0 0 13.6px;
				overflow-wrap: break-word;
				padding: 0 15px 0 15px;
				text-rendering: optimizelegibility;
				-moz-text-size-adjust: none;
				--rem: 16;
			}
	
		</style>
		<h1>Ilyas Amantaev's Homepage</h1>
		<h5>About me</h5>
		<p>I'm a Go developer with 2 years of experience in programming. I love cats and computers. I would be happy to talk to you!</p> 
		<br>
		
		<section id="portfolio">
			<h5>Portfolio</h5>
			<p>You can view my projects on <a href="https://github.com/Kartochnik010">Github</a>!</p>
			<blockquote><p>I'm happy to hear your constructive feedback!</p></blockquote>
		</section>
		<br>
	
		<section id="contact">
			<h5>Contact info</h5>
			<p>Email: <a href="mailto:ilyas.amantaev@gmail.com">ilyas.amantaev@gmail.com</a></p>
			<p>LinkedIn: <a href="https://www.linkedin.com/in/ilyas-amantaev/">linkedin.com/in/ilyas-amantaev</a></p>
			<p>Telegram: <a href="https://t.me/ilyas_amantaev">@ilyas_amantaev</a></p>
			<p>Github: <a href="https://github.com/Kartochnik010">github.com/Kartochnik010</a></p>
		</section>
	</body>
	</html>`
	fmt.Fprint(w, f)
}
