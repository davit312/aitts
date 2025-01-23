package main

import "fmt"

func makeStartpage(port int) string {
	css := `
			.btn {
			display: inline-block;
			font-weight: 400;
			color: #212529;
			text-align: center;
			vertical-align: middle;
			cursor: pointer;
			-webkit-user-select: none;
			user-select: none;
			border: 1px solid transparent;
			padding: 0.375rem 0.75rem;
			font-size: 1rem;
			line-height: 1.5;
			border-radius: 0.25rem;
			min-width: 100px;
		}
		.btn:disabled,
		.btnton[disabled]{
			opacity: 0.56;
			cursor: default;
		}
		.btn-primary{
			color: #fff;
			background-color: #007bff;
			border-color: #007bff;
		}
		
		.btn-primary:hover:enabled {
			color: #fff;
			background-color: #0069d9;
			border-color: #0062cc;
		}
		.footer{
			display: flex;
			justify-content: space-between;
		}
		.form-check {
			position: relative;
			display: block;
			padding-top: 8px;
		}

		.form-check-input {
			width: 1rem;
			height: 1rem;
			margin-top: 0;
			vertical-align: middle;
			background-color: #fff;
			background-repeat: no-repeat;
			background-position: 50% 50%;
			background-size: 50% 50%;
			border: 1px solid rgba(0, 0, 0, 0.25);
			transition: background-color 0.15s ease-in-out, border-color 0.15s ease-in-out, box-shadow 0.15s ease-in-out;
			float: left;
			left: -1.25rem;
		}

		.form-check-input:checked {
			background-color: #007bff;
			border-color: #007bff;
			background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' viewBox='-4 -4 8 8'%3e%3cpath stroke='%23fff' stroke-linecap='round' stroke-linejoin='round' stroke-width='2' d='M0 0h24v24H0z fill-none'/%3e%3cpath d='M20 6L9 17l-5-5'/%3e%3c/svg%3e");
		}

		.form-check-label {
			margin-bottom: 0;
		}

		body{
			height: 100vh;
			font-size: large;
			margin: 0;
			padding: 0;
		}


		.container {
			display: flex;
			height: 100%;
			max-height: 100%;
			max-height: calc(100%); /* Full height viewport */
		}
		
		.left-part {
			flex: 5;
			padding: 15px;
			display: flex;
			flex-direction: column;
		}
		
		.right-part {
			flex: 1; /* Smaller part */
			display: flex;
			flex-direction: column;
			justify-content: flex-start;
			align-items: center;
			padding: 15px;
		}
	`

	js := `
	function getText(){
		return document.querySelector('#text').value
	}
	function play(file){
		let player = document.querySelector('#speech')
		let address = (atob(file)).split('\\')
		alert(address[address.length - 1])
		player.src = "/audio/"+address[address.length - 1]
		player.play()
	}
	`
	return `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<base href="http://localhost:` + fmt.Sprintf("%d", port) + `">
	
	<style>` + css + `
	</style>
</head>
<body>
	<div class="container">
		<div class="left-part">
			<audio id='speech' hidden></audio>
			<textarea id='text' style="width: 100%; height: 100%; resize: none; padding: 10px;"></textarea>
			<div class="footer">
				<div class="form-check">
					<input class="form-check-input" type="checkbox" value="" id="flexCheckDefault">
					<label class="form-check-label" for="flexCheckDefault">
						Read clipboard
					</label>
				</div>
				<select>
					<option value="aaa">yyy</option>
					<option value="zzz">yyy</option>
				</select>
				<button>Settings</button>
			</div>
		</div>
		<div class="right-part">
			<button class="btn btn-primary" onclick="read(getText())" style="margin-bottom: 10px;">Speak</button>
			<button class="btn btn-primary">Get Audio</button>
		</div>
	</div>
	<script>` + js + `</script>
</body>
</html>
`
}
