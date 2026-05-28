const { app, BrowserWindow } = require('electron');
const path = require('path');

function createWindow() {
    const win = new BrowserWindow({
        width: 1000,
        height: 700,
        title: "Gestión de Empleados",
        webPreferences: {
            nodeIntegration: true,
            contextIsolation: false 
        }
    });

    win.loadFile(path.join(__dirname, 'src', 'index.html'));
    
    // Abre las herramientas de desarrollo (consola) automáticamente
    // win.webContents.openDevTools();
}

app.whenReady().then(() => {
    createWindow();

    app.on('activate', () => {
        if (BrowserWindow.getAllWindows().length === 0) createWindow();
    });
});

app.on('window-all-closed', () => {
    if (process.platform !== 'darwin') app.quit();
});