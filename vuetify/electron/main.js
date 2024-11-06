const { app, BrowserWindow } = require('electron');
const path = require('path');

// function createWindow() {
//     try {
//         const win = new BrowserWindow({
//             width: 800,
//             height: 600,
//             webPreferences: {
//                 preload: path.join(__dirname, 'preload.js'),
//                 contextIsolation: true,
//                 enableRemoteModule: false,
//             },
//         });

//         if (process.env.VITE_DEV_SERVER_URL) {
//             win.loadURL(process.env.VITE_DEV_SERVER_URL);
//         } else {
//             win.loadFile(path.join(__dirname, '../index.html'));
//         }
//     } catch (error) {
//         console.error("Error creating Electron window:", error);
//     }
// }

// app.whenReady().then(createWindow);

// app.on('window-all-closed', () => {
//     if (process.platform !== 'darwin') app.quit();
// });

// app.on('activate', () => {
//     if (BrowserWindow.getAllWindows().length === 0) createWindow();
// });

function createWindow() {
    const win = new BrowserWindow({
      width: 800,
      height: 600,
      webPreferences: {
        preload: path.join(__dirname, 'preload.js'),
        contextIsolation: true,
      },
    });
  
    // Development: load the Vite server URL
    if (process.env.VITE_DEV_SERVER_URL) {
      console.log("Loading Vite dev server URL:", process.env.VITE_DEV_SERVER_URL); // For debugging
      win.loadURL(process.env.VITE_DEV_SERVER_URL).catch((err) => {
        console.error("Failed to load Vite server:", err);
      });
    } 
    // Production: load the `dist` folder’s `index.html`
    else {
      win.loadFile(path.join(__dirname, '../index.html')).catch((err) => {
        console.error("Failed to load index.html from dist:", err);
      });
    }
  }

app.on('ready', createWindow);
app.on('window-all-closed', () => {
  if (process.platform !== 'darwin') app.quit();
});
app.on('activate', () => {
  if (BrowserWindow.getAllWindows().length === 0) createWindow();
});