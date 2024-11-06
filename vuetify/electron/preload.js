// const { contextBridge } = require('electron');

// contextBridge.exposeInMainWorld('electron', {
//   // Define functions or variables you want to expose to your Vite app
// });

// This will run when the preload script is loaded.
window.addEventListener('DOMContentLoaded', () => {
    console.log('Preload script loaded successfully');
    
    // Example: Replace text on the page (optional)
    const replaceText = (selector, text) => {
      const element = document.getElementById(selector);
      if (element) element.innerText = text;
    };
  
    // Replace these IDs with any elements in your index.html
    for (const type of ['chrome', 'node', 'electron']) {
      replaceText(`${type}-version`, process.versions[type]);
    }
  });
  
  // Securely expose any APIs here that you want to use in the renderer.
  const { contextBridge, ipcRenderer } = require('electron');
  
  // Expose `ipcRenderer` methods to the renderer process safely
  contextBridge.exposeInMainWorld('electronAPI', {
    send: (channel, data) => ipcRenderer.send(channel, data),
    receive: (channel, func) => ipcRenderer.on(channel, (event, ...args) => func(...args)),
  });