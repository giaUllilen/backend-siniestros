# Pruebas Unitarias - Backend Siniestros

Este directorio contiene todas las pruebas unitarias del proyecto Backend Siniestros.

## Estructura

```
tests/
├── services/           # Pruebas de servicios de negocio
│   ├── sinister_test.go
│   └── collaborator_finder_test.go
├── mappers/            # Pruebas de mappers (transformación de datos)
│   ├── sinister_payment_mappers_test.go
│   └── collaborator_mappers_test.go
├── storage/            # Pruebas de repositorios y acceso a datos
│   ├── sinister_payment_test.go
│   ├── collaborators_test.go
│   └── soat_return_test.go
├── controllers/        # Pruebas de controladores HTTP
│   ├── collaborator_test.go
│   └── sinister_save_test.go
├── apihelpers/         # Pruebas de helpers de API
│   └── apihelpers_test.go
└── helpers/            # Pruebas de utilidades y helpers
    ├── utils_test.go
    └── config_loader_test.go
```

## Ejecutar las pruebas

### Ejecutar todas las pruebas
```bash
go test ./tests/... -v
```

### Ejecutar pruebas por componente
```bash
# Pruebas de services
go test ./tests/services/... -v

# Pruebas de mappers
go test ./tests/mappers/... -v

# Pruebas de storage
go test ./tests/storage/... -v

# Pruebas de controllers
go test ./tests/controllers/... -v

# Pruebas de apihelpers
go test ./tests/apihelpers/... -v

# Pruebas de helpers
go test ./tests/helpers/... -v
```

### Ejecutar con cobertura
```bash
go test ./tests/... -cover
```

### Generar reporte de cobertura HTML
```bash
go test ./tests/... -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

## Convenciones

1. **Nomenclatura**: Los archivos de prueba deben terminar con `_test.go`
2. **Package**: Los packages de prueba usan el sufijo `_test` (ej: `services_test`)
3. **Funciones de prueba**: Comienzan con `Test` seguido del nombre descriptivo
4. **Mocks**: Se crean mocks para simular dependencias externas
5. **Casos de prueba**: Se cubren casos exitosos, errores y edge cases

## Componentes Testeados

### Services
- **SinisterPaymentFinder**: Búsqueda de pagos de siniestros por número de documento
- **CollaboratorFinder**: Búsqueda de colaboradores con decodificación base64

### Mappers
- **SinisterPaymentMappers**: Transformación de modelos de pago de siniestros a respuestas
- **CollaboratorMappers**: Transformación de modelos de colaboradores con lógica condicional por estado

### Storage
- **SinisterPaymentRepository**: Operaciones de acceso a datos de pagos
- **CollaboratorRepository**: Operaciones de acceso a datos de colaboradores
- **SoatReturnRepository**: Operaciones de acceso a datos de devoluciones SOAT

### Controllers
- **CollaboratorHandler**: Manejo de requests HTTP para colaboradores
- **SinisterHandler**: Manejo de requests HTTP para siniestros

### Helpers
- **ApiHelpers**: Utilidades para respuestas HTTP y wrappers
- **Utils**: Utilidades para detección de content-type y lectura de archivos
- **ConfigLoader**: Utilidades para cargar y obtener valores de configuración

## Notas

- Las pruebas utilizan mocks para evitar dependencias de bases de datos y servicios externos
- Los mocks implementan las interfaces correspondientes del código de producción
- Se verifica el comportamiento tanto en casos exitosos como en casos de error
- Las pruebas de storage solo verifican la estructura y métodos, no la funcionalidad de BD real

## Contribuir

Al agregar nuevas funcionalidades al código, asegúrate de:
1. Crear las pruebas unitarias correspondientes
2. Mantener una cobertura de código adecuada
3. Seguir las convenciones de nomenclatura
4. Documentar casos especiales o edge cases

