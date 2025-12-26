# Guía de Pruebas Unitarias - Backend Siniestros

## 📋 Resumen

Se han creado pruebas unitarias completas para el proyecto Backend Siniestros. Todas las pruebas están organizadas en la carpeta `tests/` con una estructura clara y mantenible.

## 🗂️ Estructura de Pruebas

```
tests/
├── README.md                              # Documentación detallada de las pruebas
├── services/                              # Pruebas de servicios de negocio
│   ├── sinister_test.go                  # 5 pruebas para SinisterPaymentFinder
│   └── collaborator_finder_test.go       # 6 pruebas para CollaboratorFinder
├── mappers/                               # Pruebas de transformación de datos
│   ├── sinister_payment_mappers_test.go  # 7 pruebas para mappers de pagos
│   └── collaborator_mappers_test.go      # 5 pruebas para mappers de colaboradores
├── storage/                               # Pruebas de repositorios
│   ├── sinister_payment_test.go          # 2 pruebas para repositorio de pagos
│   ├── collaborators_test.go             # 2 pruebas para repositorio de colaboradores
│   └── soat_return_test.go               # 2 pruebas para repositorio de SOAT
├── controllers/                           # Pruebas de controladores HTTP
│   ├── collaborator_test.go              # 3 pruebas para CollaboratorHandler
│   └── sinister_save_test.go             # 5 pruebas para SinisterHandler
├── apihelpers/                            # Pruebas de helpers de API
│   └── apihelpers_test.go                # 8 pruebas para ResponseWrapper y helpers
└── helpers/                               # Pruebas de utilidades generales
    ├── utils_test.go                      # 5 pruebas para utilidades de archivos
    └── config_loader_test.go              # 8 pruebas para configuración
```

**Total: 58+ pruebas unitarias creadas**

## 🚀 Comandos para Ejecutar las Pruebas

### Ejecutar todas las pruebas
```bash
go test ./tests/... -v
```

### Ejecutar pruebas por módulo
```bash
# Servicios
go test ./tests/services/... -v

# Mappers
go test ./tests/mappers/... -v

# Storage
go test ./tests/storage/... -v

# Controllers
go test ./tests/controllers/... -v

# API Helpers
go test ./tests/apihelpers/... -v

# Helpers generales
go test ./tests/helpers/... -v
```

### Ejecutar con reporte de cobertura
```bash
go test ./tests/... -cover
```

### Generar reporte HTML de cobertura
```bash
go test ./tests/... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

### Ejecutar pruebas en modo corto (skip long tests)
```bash
go test ./tests/... -short
```

### Ejecutar con salida detallada y sin caché
```bash
go test ./tests/... -v -count=1
```

## 📊 Cobertura de Pruebas

### Services
- ✅ **SinisterPaymentFinder**: Casos exitosos, errores de repositorio, documentos no encontrados, validación de mapeo
- ✅ **CollaboratorFinder**: Decodificación base64, casos exitosos, errores, códigos inválidos

### Mappers
- ✅ **SinisterPaymentMappers**: Valores vacíos, caracteres especiales, diferentes formatos de montos, sobrescritura
- ✅ **CollaboratorMappers**: Estados activo/inactivo, tipos de documento (DNI/CE), estados en mayúsculas/minúsculas

### Storage
- ✅ **Repositorios**: Verificación de constructores, firmas de métodos, implementación de interfaces

### Controllers
- ✅ **CollaboratorHandler**: Creación de handlers, manejo de errores del finder
- ✅ **SinisterHandler**: Operaciones CRUD, búsqueda por documento, guardado, búsqueda por caso

### Helpers
- ✅ **ApiHelpers**: ResponseWrapper, códigos de respuesta, headers CORS, diferentes tipos de datos
- ✅ **Utils**: Detección de content-type (texto, PDF, PNG, JPEG), lectura de archivos (pequeños, grandes, vacíos)
- ✅ **ConfigLoader**: Valores válidos, valores por defecto, tipos de datos (int, string, bool, nil)

## 🔧 Características de las Pruebas

1. **Mocks**: Se utilizan mocks para todas las dependencias externas (BD, servicios)
2. **Casos de Prueba**: Se cubren casos exitosos, errores y edge cases
3. **Aislamiento**: Cada prueba es independiente y no depende del estado de otras
4. **Nomenclatura Clara**: Nombres descriptivos siguiendo convención `Test<Component>_<Scenario>`
5. **Documentación**: Comentarios claros en cada función de prueba

## 📝 Convenciones Utilizadas

- Los archivos de prueba terminan con `_test.go`
- Los packages de prueba usan el sufijo `_test` (ej: `services_test`)
- Las funciones de prueba comienzan con `Test`
- Se usan mocks que implementan las interfaces del código de producción
- Se verifican múltiples escenarios por cada función

## 🎯 Próximos Pasos

1. **Ejecutar las pruebas** para verificar que todo compila correctamente
2. **Revisar la cobertura** y agregar pruebas adicionales si es necesario
3. **Integrar con CI/CD** para ejecutar pruebas automáticamente
4. **Agregar pruebas de integración** complementarias
5. **Mantener actualizado** al agregar nuevas funcionalidades

## 📚 Recursos Adicionales

- [Documentación oficial de Go Testing](https://golang.org/pkg/testing/)
- [Mejores prácticas de testing en Go](https://golang.org/doc/code#Testing)
- Ver `tests/README.md` para información detallada de cada componente

## ⚠️ Notas Importantes

- Las pruebas de **storage** solo verifican estructura y métodos, no funcionalidad con BD real
- Los **mocks** simulan respuestas sin conectar a servicios externos
- Para pruebas de integración con BD, crear carpeta separada `integration_tests/`
- Asegúrate de tener Go 1.18 o superior instalado

## 🤝 Contribuir

Al agregar nuevas funcionalidades:
1. Crea las pruebas correspondientes en la carpeta `tests/`
2. Sigue las convenciones de nomenclatura
3. Incluye casos exitosos, errores y edge cases
4. Actualiza esta documentación si es necesario

---

**Creado**: Diciembre 2025  
**Versión**: 1.0  
**Estado**: ✅ Completado

