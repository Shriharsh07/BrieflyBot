const { contextBridge, ipcRenderer } = require("electron");

contextBridge.exposeInMainWorld("windowControls", {
  minimize: () => ipcRenderer.send("window:minimize"),
  close: () => ipcRenderer.send("window:close")
});

contextBridge.exposeInMainWorld("wsClient", {
  connect: (url, onMessage, onStatus) => {
    const socket = new WebSocket(url);

    socket.onopen = () => onStatus("connected");
    socket.onclose = () => onStatus("disconnected");
    socket.onerror = () => onStatus("error");

    socket.onmessage = (event) => {
      try {
        onMessage(JSON.parse(event.data));
      } catch {
        console.error("Invalid WS payload");
      }
    };

    return socket;
  }
});
