
function initEvents() {
	console.log("initEvents");
	const {remote} = require('electron');
	const {BrowserWindow} = remote;
	const win = BrowserWindow.getFocusedWindow();

	document.getElementById("close-win").onclick = function() {
	  win.close();
	}
	document.getElementById("min-win").onclick = function() {
	  win.minimize();
	}
	document.getElementById("max-win").onclick = function() {
	  win.isMaximized() ? win.unmaximize() : win.maximize();
	}
}

jQuery(document).ready(function() {
	//alert("ready");
	//initEvents();
});

window.onload = function() {
	console.log("onload");
	initEvents();
}
