## Semicolon: is optional, but go complier will import semicolon in the following cases:
- An identifier
- A literal value for string, Boolean, numeric, or complex
- A control directive such as brak, continue, return
- A closing parenthesis or bracket such as ), }, ]
- The increasement ++  or decreasement -- operator
-> Therefore, golang need } at the end of statement or , at the end of last element in struct, ... to identify this is the end of statement, otherwise, golang will add semicolor lead to error

