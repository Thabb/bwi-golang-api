# Aufgabe
Zu entwickeln ist eine REST API zum Berechnen von (einfachen) mathematischen Formeln. Diese Formeln sollen  analysiert (Lexer, Parser) werden und das Ergebnis soll zurück geliefert werden. Formeln können die Grundrechenarten, Zahlen und Klammern enthalten. Beispiele:
- 5*3/2
- 4 * (3+7)

Wird eine ungültige Formel oder eine nicht lösbare Rechnung mitgegeben, wird ein Fehler zurückgegeben.
Die REST API bietet folgenden API Call:
- POST /calculate
  Parameter: Formel
  Format: JSON 
  Beispiel: { 'form': '(10+7)*5' }
  Response: { 'result': 85 }

Bedingungen:
- Die Formel soll analysiert werden mit einem eigenen Lexer und Parser
- 'eval' und andere Möglichkeiten, die die Formel mathematisch analysiert oder auswertet dürfen nicht genutzt werden.