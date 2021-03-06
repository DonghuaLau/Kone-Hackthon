const {app, BrowserWindow} = require('electron')
const path = require('path')
const url = require('url')
const jQuery = require('jquery')

let win

function createWindow () {
  // Create the browser window.
  win = new BrowserWindow({
			 width: 800
			,height: 560
			,minWidth: 800
			,minHeight: 560
			,resizable: true 
			,frame: false
		})

  // and load the index.html of the app.
  win.loadURL(url.format({
    pathname: path.join(__dirname, 'index.html'),
    protocol: 'file:',
    slashes: true
  }))


  //win = BrowserWindow.getFocusedWindow()


  //win.webContents.openDevTools()



  // Emitted when the window is closed.
  win.on('closed', () => {
    // Dereference the window object, usually you would store windows
    // in an array if your app supports multi windows, this is the time
    // when you should delete the corresponding element.
    win = null
  })

}

function initEvents() {

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

// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
// Some APIs can only be used after this event occurs.
app.on('ready', createWindow)

// Quit when all windows are closed.
app.on('window-all-closed', () => {
  // On macOS it is common for applications and their menu bar
  // to stay active until the user quits explicitly with Cmd + Q
  if (process.platform !== 'darwin') {
    app.quit()
  }
})

app.on('activate', () => {
  // On macOS it's common to re-create a window in the app when the
  // dock icon is clicked and there are no other windows open.
  if (win === null) {
    createWindow()
  }
})


