const { app, BrowserWindow, ipcMain } = require("electron");
const path = require("path");

/* Disable overlay scrollbars (keep this) */
app.commandLine.appendSwitch("disable-features", "OverlayScrollbar");

let mainWindow;

function createWindow() {
  mainWindow = new BrowserWindow({
    width: 380,
    height: 520,
    frame: false,
    alwaysOnTop: true,
    transparent: true,
    resizable: true,
    webPreferences: {
      preload: path.join(__dirname, "preload.js")
    }
  });

  mainWindow.loadFile("index.html");

  /* ✅ IMPORTANT: allow normal close */
  mainWindow.on("closed", () => {
    mainWindow = null;
  });
}

/* IPC actions */
ipcMain.on("window:minimize", () => {
  if (mainWindow) mainWindow.minimize();
});

ipcMain.on("window:close", () => {
  app.quit(); // 🔥 FULL TERMINATION
});

/* App lifecycle */
app.whenReady().then(createWindow);

app.on("window-all-closed", () => {
  app.quit(); // 🔥 exit app completely
});
