const { contextBridge, ipcRenderer } = require("electron");

contextBridge.exposeInMainWorld("controls", {
  minimize: () => ipcRenderer.send("window:minimize"),
  close: () => ipcRenderer.send("window:close")
});
