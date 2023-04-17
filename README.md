# Bistat

## Reporte de avances

### Avance 1

- Se escribió la gramática del lenguaje en un archivo de ANTLR para que el lexer y el parser fueran generados en Go. 
- Se escribieron listeners en Go que detectan cuando ciertas reglas son llamadas para comprobar el correcto funcionamiento del lenguaje
- Se comprobó que se puede parsear un programa de ejemplo (hello.bs) exitosamente excepto por el reconocimiento de los tokens PARAM_TYPE y RETURN_TYPE

Resumen de commits:
- Versión inicial del grammar: se agregó el grammar presente en la propuesta a un archivo de antlr
- Agregar listeners: se agregaron listeners para comprobar que el parser esté entrando a las reglas correctas
- Cambiar convención de acciones a camel case: se empezó a a usar camel case para las acciones de antlr
- Agregar más listeners: se agregaron más listeners para comprobar reglas adicionales y se descubrió que algunas reglas no están funcionando correctamente