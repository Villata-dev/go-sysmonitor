# Go-SysMonitor: Monitor de recursos del sistema en tiempo real para la terminal.

Go-SysMonitor es una herramienta sencilla y eficiente escrita en Go para monitorear los recursos básicos de tu sistema directamente desde la línea de comandos.

## 🚀 Características (Features)

*   🖥️ **Monitoreo de RAM:** Visualiza el total de memoria y la memoria usada/libre.
*   💾 **Estado del Disco Principal:** Información sobre el espacio total y libre en tu disco principal.
*   🎨 **Alertas visuales:** El uso de recursos se muestra con colores (Verde/Amarillo/Rojo) dependiendo del nivel de carga para una identificación rápida.
*   🔄 **Refresco automático:** Los datos se actualizan automáticamente cada 2 segundos.

## 🛠️ Instalación y Uso

### Prerrequisitos

*   Tener [Go](https://golang.org/dl/) instalado en tu sistema.

### Clonar el repositorio

```bash
git clone <URL_DEL_REPOSITORIO>
cd go-sysmonitor
```

### Ejecución directa

Para ejecutar el monitor sin compilar:

```bash
go run main.go
```

### Compilar y Ejecutar

Si prefieres tener un binario ejecutable:

1.  Compilar el binario:
    ```bash
    go build -o sysmonitor
    ```
2.  Ejecutar el monitor:
    ```bash
    ./sysmonitor
    ```

## 💻 Compatibilidad

Gracias a que está desarrollado utilizando únicamente la librería estándar de Go, Go-SysMonitor es compatible con:

*   🐧 Linux
*   🍎 macOS
*   🪟 Windows

---
Desarrollado con ❤️ usando Go.
