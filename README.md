# Servicios matemáticos

Microservicios escritos en Go que permiten hacer varios calculos.

## Contenido

- [Funcionalidad](#funcionalidad)
- [Requisitos previos](#requisitos-previos)
- [Cómo ejecutar el proyecto](#cómo-ejecutar-el-proyecto)
- [Ejecución de pruebas](#ejecución-de-pruebas)
- [Reglas de PR y merge](#reglas-de-pr-y-merge)
- [Licencia](#licencia)

## Funcionalidad

### Calcular Pi: Los usuarios pueden especificar cuántos dígitos desean calcular y el microservicio devolverá una aproximación del número Pi con la precisión solicitada.

- Endpoint que permite calcular dígitos específicos del número Pi.
- Precisión configurable de la aproximación.
- Salida formateada para fácil lectura.

## Requisitos previos

- Go 1.21
- Docker (opcional para despliegue con contenedores)

## Cómo ejecutar el proyecto

1. Clona el repositorio:

```bash
git clone https://github.com/lowlander1982/math-services
cd math-services
```

2. Ejecuta el servicio:

```bash
go run main.go
```

Esto iniciará el microservicio en el puerto `8080`.

## Ejecución de pruebas

Para ejecutar las pruebas unitarias, navega al directorio del proyecto y ejecuta:

```bash
go test
```

Esto ejecutará todas las pruebas unitarias del proyecto y mostrará los resultados.

## Reglas de PR y merge

1. **Pruebas unitarias**: Todos los PR deben pasar las pruebas unitarias antes de ser considerados para el merge.
2. **Revisiones de código**: Cada PR debe ser revisado por al menos otro miembro del equipo.
3. **GitHub Actions**: Se deben pasar todos los flujos de trabajo de GitHub Actions, incluidas las pruebas unitarias y la construcción de Docker (si corresponde), antes de que un PR sea aprobado.
4. **Descripción del PR**: Cada PR debe tener una descripción clara de los cambios realizados y el motivo del PR.
5. **Commits pequeños y significativos**: Trata de hacer commits pequeños que describan un cambio particular. Esto facilita las revisiones y la detección de problemas.
6. **No se permiten merges directos a `main`**: Todo cambio en `main` debe pasar por un PR.

## Licencia

Este proyecto está bajo la licencia MIT. Ver el archivo `LICENSE` para más detalles.
